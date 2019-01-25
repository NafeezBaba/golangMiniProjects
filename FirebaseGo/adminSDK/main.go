package main


import (
	"google.golang.org/api/option"
	"firebase.google.com/go"

)

func main() {
	sa := "AAAAHWT_RUg:APA91bEMZsS60mlW8c349Sn591wKueShHssAWZFEDL2EQtI0ulsKlEjmBigj8_CmtSRSUDlkkQgFF_MO42e1WHuBnRaklYWHB1ycuCbew7iOnL5YZbDL7bIATz8sftC-S8iBpa1o2XfU"
	app, err := firebae.NewApp(context.Background(), nil, sa)
	if err := nil {
		log.Fatalln(errr)
	}
	defer client.Close()
}

func getQuote() {

	"quote": "play it sam"
	"author": "casablnca"
}