package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// http://go-database-sql.org/errors.html

func main() {
	db, err := sql.Open("mysql", "root:!azxsw2@tcp(127.0.0.1:3306)/wuji")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(3)
	for i := 0; i < 5; i++ {
		go func() {

			rows, err := db.Query("select id, title from activity_invite")
			if err != nil {
				panic(err)
			}
			for rows.Next() {
				var id int64
				var title string
				err = rows.Scan(&id, &title)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(id, title)
			}
			if rows.Err() != nil {
				fmt.Println(rows.Err())
			}
		}()
	}
	for {

	}

}
