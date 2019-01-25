package main

//> go build
//> mypgoject // the exe file and then hit the API form postman

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

// type Model struct {
// 	ID        uint `gorm:"primary_key"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt *time.Time
// }

type Person struct {
	gorm.Model
	Id        int     `gorm:"id,omitempty"`
	Firstname string  `gorm:"firstname,omitempty"`
	Lastname  string  `gorm:"lastname,omitempty"`
	Age       uint    `gorm:"age"`
	Address   Address `gorm:"address,omitempty"`
	// GenderID  uint
	Gender Gender `gorm:"gender"`
	// ContactID uint
	Contact Contact `gorm"contact"`
}

// type People struct {
// 	People []Person `gorm:"people"`
// }
// func (Person) TableName() string {
// 	return "people"
// }

type Address struct {
	gorm.Model
	UserID uint
	City   string `gorm:"city,omitempty`
	State  string `gorm:state,omitempty`
}

type Gender struct {
	gorm.Model
	UserID     uint
	GenderType string `gorm:"gendertype"`
}

type Contact struct {
	gorm.Model
	MobileNo string `gorm:"mobileno"`
	Email    string `gorm:"email"`
}

var people []Person

func PostPerson(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testdb password=root sslmode=disable")

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Problem while connecting database[Postperson] ")
		defer db.Close()
		return err
	}
	// db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testdb password=root sslmode=disable")
	// if err != nil {
	// 	fmt.Println("Problem while conncting database")
	// 	defer db.Close()
	// }
	db.Create(&Person{})
	db.Create(&Address{})
	db.Create(&Contact{})
	db.Create(&Gender{})

	u := new(Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	// sqlStatement := "INSERT INTO Person (firstname, lastname, age)VALUES ($1, $2, $3)"
	// err :=
	db.Create(u)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	// fmt.Println("-->", res)
	// 	return c.JSON(http.StatusCreated, u)
	// }
	return c.JSON(http.StatusOK, u)
}
func Getall(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testdb password=root sslmode=disable")

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Problem while connecting database[GEtall] ")
		defer db.Close()
		return err
	}
	u := new(Person)
	if err := c.Bind(u); err != nil {
		return err
	}
	var list []Person

	db.Find(&list)
	return c.JSON(http.StatusOK, list)
}

func Getbyid(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testdb password=root sslmode=disable")

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Problem while connecting database[Getbyid] ")
		defer db.Close()
		return err
	}

	e := []Person{}
	// id := c.Param("id")
	// db.Where(&Person{"Id": &id}).First(&e)
	// db.Where("id = ?", id).Find(&e)
	db.Select([]string{"firstname", "age"}).Find(&e)
	return c.JSON(http.StatusOK, e)

}

// func DeletePersonbyId(c echo.Context) error {
// 	db, err := gorm.Open("Postgres", "host=localhost port=5433 user=postgres dbname=testdb password=root sslmode=disable")
// 	if err != nil {
// 		fmt.Println("Problem while connecting database")
// 		defer db.Close()
// 	}

// }

// // PUT methode to update data
// func Updateperson(c echo.Context) error {
// 	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testdb password=root sslmode=disable")
// 	if err != nil {
// 		fmt.Println("Problem while conncting database")
// 		defer db.Close()
// 	}
// 	fmt.Println(c)
// 	u := new(Person)
// 	if err := c.Bind(&u); err != nil {
// 		return err
// 	}
// 	paramId := c.Param("id")
// 	attrMap := map[string]interface{}{"firstname": Person.Firstname, "lastname": Person.Lastname, "age": Person.Age}
// 	db.Model(&Person{}).Where("id= ?", paramId).Updates(attrMap)
// 	return c.NoContent(http.StatusOK)

// 	// sqlStatement := "UPDATE Person SET Firstname=$1,Lastname=$2,Age=$3 WHERE ID=ID"
// 	// res, err := db.Update(sqlStatement, u.Firstname, u.Lastname, u.Age, u.ID)
// 	// db.Update(u)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return c.JSON(http.StatusCreated, u)
// 	// } else {
// 	// 	fmt.Println(u)
// 	// 	return c.JSON(http.StatusCreated, u)
// 	// }
// 	// return c.JSON(http.StatusOK, u)
// }

func initMigrate() {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=testdb password=root sslmode=disable")
	if err != nil {
		fmt.Println("Problem while conncting database")
		defer db.Close()
	}
	db.AutoMigrate(
		&Person{},
		&Address{},
		&Gender{},
		&Contact{},
	)
}

func handlerequest() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/POSTperson", PostPerson)
	// e.PUT("/PUTperson/:ID", Updateperson)
	e.GET("/GETperson", Getall)
	e.GET("/GETperson/:ID", Getbyid)
	e.Logger.Fatal(e.Start(":12345"))

}

func main() {
	handlerequest()
	initMigrate()
}
