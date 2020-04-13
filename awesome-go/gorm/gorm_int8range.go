package main

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
)

func appendArrayQuotedBytes(b, v []byte) []byte {
	b = append(b, '"')
	for {
		i := bytes.IndexAny(v, `"\`)
		if i < 0 {
			b = append(b, v...)
			break
		}
		if i > 0 {
			b = append(b, v[:i]...)
		}
		b = append(b, '\\', v[i])
		v = v[i+1:]
	}
	return append(b, '"')
}

type Int8rangeArray [][]int64

func parseArray(src, del []byte) (dims []int, elems [][]byte, err error) {
	var depth, i int

	if len(src) < 1 || src[0] != '{' {
		return nil, nil, fmt.Errorf("pq: unable to parse array; expected %q at offset %d", '{', 0)
	}

Open:
	for i < len(src) {
		switch src[i] {
		case '{':
			depth++
			i++
		case '}':
			elems = make([][]byte, 0)
			goto Close
		default:
			break Open
		}
	}
	dims = make([]int, i)

Element:
	for i < len(src) {
		switch src[i] {
		case '{':
			if depth == len(dims) {
				break Element
			}
			depth++
			dims[depth-1] = 0
			i++
		case '"':
			var elem = []byte{}
			var escape bool
			for i++; i < len(src); i++ {
				if escape {
					elem = append(elem, src[i])
					escape = false
				} else {
					switch src[i] {
					default:
						elem = append(elem, src[i])
					case '\\':
						escape = true
					case '"':
						elems = append(elems, elem)
						i++
						break Element
					}
				}
			}
		default:
			for start := i; i < len(src); i++ {
				if bytes.HasPrefix(src[i:], del) || src[i] == '}' {
					elem := src[start:i]
					if len(elem) == 0 {
						return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
					}
					if bytes.Equal(elem, []byte("NULL")) {
						elem = nil
					}
					elems = append(elems, elem)
					break Element
				}
			}
		}
	}

	for i < len(src) {
		if bytes.HasPrefix(src[i:], del) && depth > 0 {
			dims[depth-1]++
			i += len(del)
			goto Element
		} else if src[i] == '}' && depth > 0 {
			dims[depth-1]++
			depth--
			i++
		} else {
			return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
		}
	}

Close:
	for i < len(src) {
		if src[i] == '}' && depth > 0 {
			depth--
			i++
		} else {
			return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
		}
	}
	if depth > 0 {
		err = fmt.Errorf("pq: unable to parse array; expected %q at offset %d", '}', i)
	}
	if err == nil {
		for _, d := range dims {
			if (len(elems) % d) != 0 {
				err = fmt.Errorf("pq: multidimensional arrays must have elements with matching dimensions")
			}
		}
	}
	return
}
func scanLinearArray(src, del []byte, typ string) (elems [][]byte, err error) {
	dims, elems, err := parseArray(src, del)
	if err != nil {
		return nil, err
	}
	if len(dims) > 1 {
		return nil, fmt.Errorf("pq: cannot convert ARRAY%s to %s", strings.Replace(fmt.Sprint(dims), " ", "][", -1), typ)
	}
	return elems, err
}

func (d *Int8rangeArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		elems, err := scanLinearArray(src, []byte{','}, "ArrayArray")
		if err != nil {
			return err
		}
		b := make(Int8rangeArray, len(elems))
		for i, elem := range elems {
			subElems := strings.Split(strings.Replace(strings.Replace(string(elem), "[", "", -1), ")", "", -1), ",")
			for i2, i3 := range subElems {
				pv, err := strconv.ParseInt(i3, 10, 64)
				if err != nil {
					return fmt.Errorf("pq: parsing array element index %d: %v", i, err)
				}
				b[i] = append(b[i], pv)
				fmt.Println(b[i][i2], err)

			}

		}
		*d = b
	case nil:
		*d = nil
	}

	return nil
}

func (d Int8rangeArray) Value() (driver.Value, error) {
	if d == nil {
		return nil, nil
	}

	if n := len(d); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+3*n)
		b[0] = '{'

		b = appendArrayQuotedBytes(b, []byte(fmt.Sprintf(
			"[%d, %d)",
			d[0][0],
			d[0][1],
		)))
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = appendArrayQuotedBytes(b, []byte(fmt.Sprintf(
				"[%d, %d)",
				d[i][0],
				d[i][1],
			)))
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

type Act struct {
	Id     int32          `gorm:"not null; column:id" json:"id"`
	Ranges Int8rangeArray `gorm:"not null; type:int8range[]; column:range" json:"range"`
}

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", "postgres", "postgres", "gobyexample"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Act{})

	/*
		Create
	*/
	r := make(Int8rangeArray, 2)
	r[0] = []int64{1, 3}
	r[1] = []int64{4, 7}
	db.Debug()
	db.LogMode(true)
	db.Create(&Act{
		Id:     0,
		Ranges: r,
	})

	res := new([]*Act)
	// find
	db.Model(&Act{}).Where(`lower(r) >= 3::bigint
  and upper(r) <= 20`).Find(res)
	fmt.Println(res)
}
