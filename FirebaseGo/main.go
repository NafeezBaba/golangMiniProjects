package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	fcm "github.com/NaySoftware/go-fcm"
)

const (
	serverKey = "AAAAHWT_RUg:APA91bEMZsS60mlW8c349Sn591wKueShHssAWZFEDL2EQtI0ulsKlEjmBigj8_CmtSRSUDlkkQgFF_MO42e1WHuBnRaklYWHB1ycuCbew7iOnL5YZbDL7bIATz8sftC-S8iBpa1o2XfU"
	topic     = "/topics/someTopic"
)

func main() {

	conf := &firebase.Config{
		DatabaseURL: "gs://oncotest-28bcc.appspot.com/",
	}

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(context.Background(), conf, serverKey)
	if err != nil {
		log.Fatalln("Error initializing app:", serverKey)
	}

	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatalln("Error initializing database client:", serverKey)
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("restricted_access/secret_document")
	var data map[string]interface{}
	if err := ref.Get(context.Background(), &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)
	// data := map[string]string{
	// 	"msg": "Hello World1",
	// 	"sum": "Happy Day",
	// }

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmMsgTo(topic, data)

	status, err := c.Send()
	fmt.Println("status", status)

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println("err", err)
	}

}
