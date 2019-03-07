package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

//Person
type Person struct {
	// gorm.Model
	ID        int     `gorm:"primary_key:true"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       uint    `json:"age"`
	Gender    string  `json:"gender"`
	Address   Address `json:"foreignkey:ID;association_foreignkey:ID" gorm:"address"`
	Contact   Contact `json:"foreignkey:ID;association_foreignkey:ID" gorm:"contact"`
}

//Address
type Address struct {
	// gorm.Model
	ID    uint
	City  string `json:"city"`
	State string `json:"state"`
}

//Contact
type Contact struct {
	// gorm.Model
	ID     uint
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

//IndexCreateRespnse
type IndexCreateRespnse struct {
	took     uint
	timedOut bool
	_shards  Shards
	hits     Hits
}

//Shards
type Shards struct {
	total      uint
	successful uint
	skipped    uint
	failed     uint
}

//Hits
type Hits struct {
	total     uint
	max_score *uint
	hits      []_source
}

type _source struct {
	_index string
	_type  string
	_id    uint
	_score float64
	Person Person
}

// {
// 	"_index": "person",
// 	"_type": "_doc",
// 	"_id": "2",
// 	"_score": 1.0,
// 	"_source": {
// 		"ID": 2,
// 		"first_name": "rohan",
// 		"last_name": "bagul",
// 		"age": 25,
// 		"gender": "male",
// 		"Address": {
// 			"ID": 2,
// 			"city": "nashik",
// 			"state": "maharashtra"
// 		},
// 		"Contact": {
// 			"ID": 2,
// 			"mobile": "920",
// 			"email": "r@g.com"
// 		}
// 	}
// },

//DBConnection function returns the database object
func DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(`postgres`, `host=localhost port=5433 user=postgres dbname=testone password=root sslmode=disable`)
	checkError(err)
	fmt.Println(`Database connected...`)
	return db, err
}

func handlerequest() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloWorld)

	e.POST("/person", PostPerson)
	e.GET("/person", GetAll)
	e.GET("/person/:ID", GetByID)

	e.PUT("/createElasticSearchIndex", CreateIndex)
	e.GET("/getByIDElastic/:ID", GetByIDElastic)
	e.GET("/getAllElastic", GetAllElastic)

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
	db, err := DBConnection()
	checkError(err)
	defer db.Close()

	db.AutoMigrate(
		&Person{},
		&Address{},
		&Contact{},
		&Address{},
	)
}

var people []Person

// Handler
func helloWorld(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello, World!")
}

//PostPerson new record to people
func PostPerson(c echo.Context) (err error) {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()

	u := new(Person)

	if err := c.Bind(u); err != nil {
		return err
	}
	db.Create(&u)

	dt, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
	}
	body := bytes.NewBuffer(dt)
	fmt.Println("body->", body)

	PUTuri := `http://localhost:9200/person/_doc/` + strconv.Itoa(u.ID) + ``

	req, err := http.NewRequest(http.MethodPut, PUTuri, body)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("DATA", string(data))
	return c.JSON(http.StatusOK, &u)
}

//GetAll records
func GetAll(c echo.Context) error {
	db, err := DBConnection()
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

//GetAllElastic records
func GetAllElastic(c echo.Context) error {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()

	GETIndex := `http://localhost:9200/person/_search?pretty=true&q=*:*`
	response, err := http.Get(GETIndex)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	d := string(data)
	return c.JSON(http.StatusOK, &d)
}

//GetByID one record by id
func GetByID(c echo.Context) error {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()

	var e []Person
	id := c.Param("ID")
	db.Where("id = ?", id).Find(&e)
	return c.JSON(http.StatusOK, &e)
}

//GetByIDElastic one record by id
func GetByIDElastic(c echo.Context) error {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()
	id := c.Param("ID")
	GETIndex := `http://localhost:9200/person/_doc/` + id + ``
	response, err := http.Get(GETIndex)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	d := string(data)
	return c.JSON(http.StatusOK, &d)
}

//DeletePerson record
func DeletePerson(c echo.Context) error {
	db, err := DBConnection()
	checkError(err)
	defer db.Close()

	var e []Person
	id := c.Param("ID")
	db.Model(&e).Where("id=?", id).Delete(&e)
	return c.JSON(http.StatusOK, &e)
}

//UpdatePerson the person by id
func UpdatePerson(c echo.Context) error {
	var e Person
	db, err := DBConnection()
	checkError(err)
	defer db.Close()

	if err := c.Bind(&e); err != nil {
		return err
	}
	ID := c.Param("id")
	up := map[string]interface{}{"firstname": e.FirstName, "lastname": e.LastName, "age": e.Age}
	db.Model(&e).Where("id= ?", ID).Updates(&up)
	return c.JSON(http.StatusOK, &up)
}

//CreateIndex to create elstic search index
func CreateIndex(c echo.Context) error {

	PUTIndex := `http://localhost:9200/person`
	reqIndex, err := http.NewRequest(http.MethodPut, PUTIndex, nil)
	reqIndex.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(reqIndex)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	d := string(data)
	return c.JSON(http.StatusOK, &d)
}

//main func
func main() {
	initiatMigrate()
	handlerequest()
}

// {
//     "ID": 1,
//     "first_name": "one",
//     "last_name": "one",
//     "age": 25,
//     "gender": "male",
//     "Address": {
//         "ID": 1,
//         "city": "nashik",
//         "state": "maharashtra"
//     },
//     "Contact": {
//         "ID": 1,
//         "mobile": "11",
//         "email": "1@g.com"
//     }
// }
