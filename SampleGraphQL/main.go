package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//DB
var DB *sqlx.DB

//user
type user struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

//userType
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

//queryType
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						var u = user{}
						err := DB.Get(&u, "SELECT * FROM user WHERE id=$1", idQuery)
						if err != nil {
							return nil, nil
						}
						return &u, nil
					}

					return nil, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result := executeQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

func main() {
	connectDB()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func connectDB() {
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=OncoMessage password=root sslmode=disable")

	db, err := sqlx.Connect("postgres","host=localhost port=5432 user=postgres dbname=OncoMessage password=root sslmode=disable")
	if err != nil {
		fmt.Print("Error:", err)
	}
	DB = db
}
// If
// user = root
// And pass = root
// And database = rasaui

// postgres://root:root@localhost:5432/rasaui