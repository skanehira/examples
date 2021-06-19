package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

func newDB() *xorm.Engine {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=UTC",
		"test", "test", "localhost", "3307", "test")

	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

type User struct {
	Created int64 `xorm:"created"`
}

func main() {
	db := newDB()
	_, err := db.Exec("truncate table user")
	if err != nil {
		log.Fatal(err)
	}

	input := "2021/06/18 13:00:10.123"
	tz := "America/New_York"
	format := "2006/01/02 15:04:05.999"
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal(err)
	}

	t, err := time.ParseInLocation(format, input, loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	msec := t.UTC().UnixNano() / int64(time.Millisecond)
	fmt.Println("original time", input)
	fmt.Println("original tz", tz)
	fmt.Println("save time", t.UTC())
	fmt.Println("save msec", msec)

	data := User{Created: msec}

	if _, err := db.NoAutoTime().Table("user").Insert(data); err != nil {
		log.Fatal(err)
	}
	users := []User{}
	if err := db.Find(&users); err != nil {
		log.Fatal(err)
	}

	c := users[0].Created
	fmt.Println("got msec", c)
	got := time.Unix(0, c*int64(time.Millisecond)).UTC().In(loc)
	fmt.Println("got time", got.Format(format))
}
