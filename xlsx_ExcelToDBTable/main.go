package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tealeg/xlsx"
)

//MstNotification struct
type MstNotification struct {
	gorm.Model
	Name             string `json:"name"`
	Question         string `json:"question"`
	RedirectURL      string `json:"redirect_url"`
	NotificationBody string `json:"notification_body"`
	Type             string `json:"type"`
	Role             string `json:"role"`
}

// host=localhost port=5432 user=postgres dbname=personDB password=root sslmode=disable
func main() {
	// connect to database
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			"oncopower-staging.cfirvgcpmj9o.us-east-1.rds.amazonaws.com",
			"5432",
			"oncopowerProd",
			"oncopower_test",
			"Onco_power-2018",
		),
	)
	// db, err := gorm.Open("postgres",
	// 	`host=localhost
	// 	port=5432
	// 	user=postgres
	// 	dbname=personDB
	// 	password=root
	// 	sslmode=disable`,
	// )
	if err != nil {
		fmt.Println("Error:", err)
	}
	// auto migrate table
	db.AutoMigrate(
		&MstNotification{},
	)
	// open xlxs file
	excelFileName := "condition.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// read file row by row and add data to database
	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex != 0 {
				var mstNot MstNotification
				for index, item := range row.Cells {
					switch index {
					case 0:
						mstNot.Name = item.String()
						break
					case 1:
						mstNot.Question = item.String()
						break
					case 2:
						mstNot.RedirectURL = item.String()
						break
					case 3:
						mstNot.NotificationBody = item.String()
						break
					case 4:
						mstNot.Type = item.String()
						break
					case 5:
						mstNot.Role = item.String()
						break
					}
				}
				// fmt.Println("-->", mstNot.Role)
				db.Create(&mstNot)
			}
		}
	}
}
