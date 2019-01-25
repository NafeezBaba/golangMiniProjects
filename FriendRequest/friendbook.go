package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

type Mydata struct {
	// gorm.Model
	ID   int    `gorm:"primary_key:true;"`
	Name string `json:"name;"`
	// Friend []Friend `gorm:"Friendlist;association_foreigenkey:FriendID;foreignkey:ID"`
}
type Friend struct {
	// gorm.Model
	ID       int `gorm:"primary_key:true;"`
	MyID     int
	FriendID int
	Name     string `json:"name;"`
	Status   string
}

func initiatMigrate() {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(
		&Mydata{},
		&Friend{},
	)
}

func checkError(err error) {
	if err != nil {
		log.Panic("Error detected-->", err)
	}
}

func handlerequest() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloWorld)
	e.POST("/friend", FriendData)
	e.POST("/friendRequest", FriendRequest)
	e.POST("/friendRequest/id", ResponseRequest)
	e.GET("/friend", GetList)
	e.Logger.Fatal(e.Start(":12345"))

}

// Handler
func helloWorld(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello, World!")
}

//DBConnection
func DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres",
		`host=localhost 
							port=5432 
							user=postgres 
							dbname=personDB 
							password=root 
			sslmode=disable`)
	return db, err
}

//FriendData
func FriendData(c echo.Context) (err error) {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()
	var m Mydata
	if err := c.Bind(&m); err != nil {
		return err
	}
	if err := db.Create(&m).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &m)
}

//GetList
func GetList(c echo.Context) (err error) {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()
	var m []Mydata
	db.Preload("Friend").Preload("Mydata").Find(&m)
	return c.JSON(http.StatusOK, &m)
}

//FriendRequest
func FriendRequest(c echo.Context) (err error) {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()

	var f Friend

	if err := c.Bind(&f); err != nil {
		return err
	}
	if err := db.Create(&f).Error; err != nil {
		return err
	}
	fmt.Println("G->", f)
	return c.JSON(http.StatusOK, &f)

}

//ResponseRequest
func ResponseRequest(c echo.Context) (err error) {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()
	type ResponseRequestStatus struct {
		// gorm.Model
		// ID       int `gorm:"primary_key:true;"`
		// MyID     int
		// FriendID int
		// Response     string
		Status int
	}
	r := new(ResponseRequestStatus)
	if err := c.Bind(r); err != nil {
		fmt.Println(err)
		return err
	}

	f := new(Friend)
	id, err := strconv.Atoi(c.Param("id"))
	checkError(err)
	f.ID = id
	switch r.Status {
	case 0:
		{
			f.Status = "Pending"
		}
	case 1:
		{
			f.Status = "Aproved"
		}
	case 2:
		{
			f.Status = "Blocked"
		}
	}
	db.Where("id = ?", id).Update(&f)
	db.Find(&f)
	fmt.Println("-->", f)
	return c.JSON(http.StatusOK, &f)

}
func main() {
	fmt.Println("Friend Book")
	initiatMigrate()
	handlerequest()

}
