package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

type Person struct {
	// gorm.Model
	ID        int     `gorm:"primary_key:true;"`
	Firstname string  `json:"firstname;"`
	Lastname  string  `json:"lastname;"`
	Age       uint    `json:"age"`
	Gender    string  `json:"gendertype;"`
	Address   Address `gorm:"foreignkey:ID;association_foreignkey:ID" json:"address"`
	Contact   Contact `grom:"foreignkey:ID;association_foreignkey:ID" json: "contact"`
}

type Address struct {
	// gorm.Model
	ID    uint
	City  string `json:"city;"`
	State string `json:"state;"`
}
type Contact struct {
	// gorm.Model
	ID       uint
	Mobile string `json:"mobile;"`
	Email    string `json:"email;"`
}

// type Gender struct {
// 	gorm.Model
// 	UserID     uint
// 	GenderType string
// }

func handlerequest() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloWorld)
	e.POST("/person", PostPerson)
	e.GET("/person", GetAll)
	e.GET("/person/:ID", GetById)
	e.PUT("/person/:ID", UpdatePerson)
	e.DELETE("/person/:ID", DeletePerson)
	e.Logger.Fatal(e.Start(":12345"))

}
func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initiatMigrate() {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testone password=root sslmode=disable")
	checkError(err)
	defer db.Close()

	db.AutoMigrate(
		&Person{},
		&Address{},
		&Contact{},
		&Address{},
		// &Gender{},
	)
}

var people []Person

// var adresses []Address
// var contacts []Contact
// var genders []Gender

// Handler
func helloWorld(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello, World!")
}

//PostPerson new record to people
func PostPerson(c echo.Context) (err error) {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testone password=root sslmode=disable")
	checkError(err)
	defer db.Close()

	u := new(Person)

	// u := &Person{}
	if err := c.Bind(u); err != nil {
		return err
	}
	db.Create(&u)
	fmt.Println("u->", u)
	return c.JSON(http.StatusOK, &u)
}

//to Getall records
func GetAll(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testone password=root sslmode=disable")
	checkError(err)
	defer db.Close()

	u := new(Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	var all []Person
	db.Find(&all)
	return c.JSON(http.StatusOK, &all)
}

// to Getbyid one record by id
func GetById(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testone password=root sslmode=disable")
	checkError(err)
	defer db.Close()

	var e []Person
	id := c.Param("ID")
	db.Where("id = ?", id).Find(&e)
	return c.JSON(http.StatusOK, &e)
}

//to delete person record
func DeletePerson(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testone password=root sslmode=disable")
	checkError(err)
	defer db.Close()

	var e []Person
	id := c.Param("ID")
	db.Model(&e).Where("id=?", id).Delete(&e)
	// db.Where("id=?", id).Find(&e).Delete(&e)
	return c.JSON(http.StatusOK, &e)
}

//to update the person by id
func UpdatePerson(c echo.Context) error {
	var e Person
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testone password=root sslmode=disable")
	checkError(err)
	defer db.Close()

	if err := c.Bind(&e); err != nil {
		return err
	}
	ID := c.Param("id")
	up := map[string]interface{}{"firstname": e.Firstname, "lastname": e.Lastname, "age": e.Age}
	// up := &db.Person{"firstname": e.Firstname, "lastname": e.Lastname, "age": e.Age}
	db.Model(&e).Where("id= ?", ID).Updates(&up)
	return c.JSON(http.StatusOK, &up)
}

//main func
func main() {
	initiatMigrate()
	handlerequest()
}
