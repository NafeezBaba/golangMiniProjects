package main

import (
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("oncotest-28bcc-firebase-adminsdk-z48jn-bf684f0185.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	// claims := map[string]interface{}{
	// 	"primiumAccount": true,
	// }
	token, err := client.CustomToken(context.Background(), "some-uid")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Got Token:- %v\n", token)

}

// func verifyToken(token string, app *App) {

// 	client, err := app.Auth(context.Background())
// 	if err != nil {
// 		log.Fatalf("error getting Auth client: %v\n", err)
// 	}

// 	token1, err := client.VerifyIDToken(ctx, idToken)
// 	if err != nil {
// 		log.Fatalf("error verifying ID token: %v\n", err)
// 	}

// 	log.Printf("Verified ID token: %v\n", token1)
// }
