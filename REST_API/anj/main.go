package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

type Address struct {
	gorm.Model
	State   string `json:"state"`
	Country string `json:"country"`
	City    string `json:"city"`
	ZipCode int    `json:"zipcode"`
	UserID uint
}

// ,omitempty

type Person struct {
	gorm.Model
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Email string  `gorm:"unique" json:"email"`
	Add   Address `gorm:"foreignkey:UserID;association_foreignkey:ID" json:"address"`
}


var persons []Person
var people  Person
var addresses Address



func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=sharma sslmode=disable")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("success")

	//db.Model(&people).Association("ID")
	db.AutoMigrate(&Person{}, &Address{})
	
	// db.Model(&people).Related(&addresses, "Id")
	e := echo.New()

	e.POST("/createPerson", CreatePerson)
	e.POST("/updatePerson", UpdatePerson)
	e.DELETE("/deletePerson", DeletePerson)
	e.GET("/getAll", GetALL)

	e.Logger.Fatal(e.Start(":1323"))
}

func CreatePerson(c echo.Context) error {
	var p Person

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=sharma sslmode=disable")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	// fmt.Println("sucess")

	if err := c.Bind(&p); err != nil {
		fmt.Println(err)
		return err
	}

	db.Create(&p)
	// db.Create(&p.Add)

	c.JSON(http.StatusCreated, p)

	fmt.Println(p)

	return c.JSON(http.StatusOK, "Added")
}

func GetALL(c echo.Context) error {

	//var personRes []PersonResponse
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=sharma sslmode=disable")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var list []Person

	db.Preload("Add").Find(&list)

	//JSON.NewEncoder(w).Encoder(persons)
	
	return c.JSON(http.StatusOK, list)
}


func UpdatePerson(c echo.Context) error {
		
	
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=sharma sslmode=disable")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	
	type RAddress struct {
		ID uint `json:"id"`
		State   string `json:"state"`
		Country string `json:"country"`
		City    string `json:"city"`
		ZipCode int    `json:"zipcode"`
	}
	type Request struct{
		ID uint `json:"id"`
		Name string `json:"name"`
		Age  int `json:"age"`
		Email string `json:"email"`
		Add RAddress `json:"Add"`
	}
	r :=new(Request)
 	if err := c.Bind(r); err != nil {
		fmt.Println(err) 
		return err
	 }
	 
	//  db.Model(&Person{}).Where("id=?",r.ID).Update(r)


	var p Person
	p.ID = r.ID
	db.Preload("Add").Find(&p)

	p.Name = r.Name
	p.Age = r.Age
	p.Email = r.Email
	p.Add.ID = r.Add.ID
	p.Add.State = r.Add.State
	p.Add.City = r.Add.City
	p.Add.Country = r.Add.Country
	p.Add.ZipCode = r.Add.ZipCode

	db.Save(&p)

	//  db.Where("ID= ?",r.ID).Find(&people)
	// db.Model(&people).Updates(&Person{Name : r.Name, Email : r.Email, Age : r.Age, Add: Address{ State : r.Add.State, City : r.Add.City, Country: r.Add.Country, ZipCode: r.Add.ZipCode }} )


return GetALL(c)
 }


 func DeletePerson(c echo.Context) error {



	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=sharma sslmode=disable")
	defer db.Close()
	if err != nil {
		panic(err)
	}


	type Request struct{
		ID int `json:"id"`
	}
	
	r :=new(Request)
	
		
	if err := c.Bind(r); err != nil {
	   fmt.Println(err) 
	   return err
	}
	fmt.Println(r.ID)
	
		   
		db.Where("Id=?", r.ID).Delete(&people)

		return  c.JSON(http.StatusOK, "Deleted")
	 }