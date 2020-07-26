## List of well-known struct tags
Tag       | Documentation
----------|---------------
xml       | https://godoc.org/encoding/xml
json      | https://godoc.org/encoding/json
asn1      | https://godoc.org/encoding/asn1
reform    | https://godoc.org/gopkg.in/reform.v1
dynamodb  | https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/dynamodbattribute/#Marshal
bigquery  | https://godoc.org/cloud.google.com/go/bigquery
datastore | https://godoc.org/cloud.google.com/go/datastore
spanner   | https://godoc.org/cloud.google.com/go/spanner
bson      | https://godoc.org/labix.org/v2/mgo/bson, https://godoc.org/go.mongodb.org/mongo-driver/bson/bsoncodec
gorm      | http://gorm.io/docs/models.html
yaml      | https://godoc.org/gopkg.in/yaml.v2
validate  | https://github.com/go-playground/validator
mapstructure | https://godoc.org/github.com/mitchellh/mapstructure
protobuf  | https://github.com/golang/protobuf
db        | https://github.com/jmoiron/sqlx

### Json Tags

Field     |Json tags | Comment
----------|----------|-------------
Field int | `json:"myName"` | // JSON as key "myName"
Field int | `json:"myName,omitempty"` | // omit if empty
Field int | `json:",omitempty"` | // JSON as key "Field"
Field int | `json:"-"`  | // Field is ignored by this package.
Field int | `json:"-,"` | // 


### Gorm Tags

```go
type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // set field size to 255
  MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
  Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
  Address      string  `gorm:"index:addr"` // create index with name `addr` for address
  IgnoreMe     int     `gorm:"-"` // ignore this field
}
```

Tag	| Description
-----------|--------------
Column   |	Specifies column name
Type   |	Specifies column data type
Size   |	Specifies column size, default 255
PRIMARY_KEY   |	Specifies column as primary key
UNIQUE   |	Specifies column as unique
DEFAULT   |	Specifies column default value
PRECISION   |	Specifies column precision
NOT   | NULL	Specifies column as NOT NULL
AUTO_INCREMENT   |	Specifies column auto incrementable or not
INDEX   |	Create index with or without name, same name creates composite indexes
UNIQUE_INDEX   |	Like INDEX, create unique index
EMBEDDED   |	Set struct as embedded
EMBEDDED_PREFIX   |	Set embedded struct’s prefix name
-   |	Ignore this fields

###  Protobuf Tags
 