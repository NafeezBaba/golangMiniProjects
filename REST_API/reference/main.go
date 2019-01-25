package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}

func main() {
	initMigrate()
	run()
}

func initMigrate() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=root sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Employee{})
}

func run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/employees", showAllEmployees)
	e.GET("/employee/:id", showEmployee)
	e.PUT("/employee/:id", updateEmployees)
	e.POST("/employee", newEmployees)
	e.DELETE("/employee/:id", deleteEmployee)

	log.Fatal(e.Start(":8080"))
}
func showAllEmployees(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=root sslmode=disable")
	defer db.Close()
	checkError(err)
	var employees []Employee
	db.Find(&employees)
	return c.JSON(http.StatusOK, employees)
}

func showEmployee(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=root sslmode=disable")
	checkError(err)
	defer db.Close()
	var e Employee
	id := c.Param("id")
	db.Where("id=?", id).Find(&e)
	return c.JSON(http.StatusOK, e)
}

func newEmployees(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=root sslmode=disable")
	defer db.Close()
	checkError(err)

	employee := new(Employee)
	if err := c.Bind(employee); err != nil {
		return err
	}

	db.Create(&employee)
	return c.String(http.StatusOK, "OK")
}

func updateEmployees(c echo.Context) error {
	var employee Employee
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=root sslmode=disable")
	checkError(err)
	defer db.Close()
	if err := c.Bind(&employee); err != nil {
		return err
	}
	paramId := c.Param("id")
	attrMap := map[string]interface{}{"name": employee.Name, "salary": employee.Salary, "age": employee.Age}
	db.Model(&employee).Where("id= ?", paramId).Updates(attrMap)
	return c.NoContent(http.StatusOK)
}

func deleteEmployee(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=root sslmode=disable")
	checkError(err)
	defer db.Close()
	var e Employee
	id := c.Param("id")
	db.Where("id=?", id).Find(&e).Delete(&e)
	return c.JSON(http.StatusOK, e)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// $ curl http://localhost:8080/employees      // (1) showAllEmployees()
// $ curl http://localhost:8080/employee/{id}  // (2) showEmployee()
// $ curl -X POST -H 'Content-Type: application/json' -d '{ "name":"fuga","salary":"fuga",age:"fuga" }' localhost:8080/employee      // (3)newEmployees()
// $ curl -X PUT  -H 'Content-Type: application/json' -d '{ "name":"hoge","salary":"hoge",age:"hoge" }' localhost:8080/employee/{id} // (4)updateEmployees()
// $ curl -X DELETE http://localhost:8080/employee/{id} // (5)deleteEmployee()
