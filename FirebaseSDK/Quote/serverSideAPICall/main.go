package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Starting the application...")
	// response, err := http.Get("https://httpbin.org/ip")
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(data))
	// }

	jsonData := map[string]interface{}{
		"notification": map[string]interface{}{
			"title":        "OncoPower",
			"body":         "Hello! this is test notification on your browser. have a nice day",
			"click_action": "http://oncopower.org/",
			"icon":         "https://oncopower-test-photo-user.s3.amazonaws.com/2edf5475-8344-4ec6-ad99-6c0f90f81eaa.png",
		},
		"collapse_key": "type_a",
		"to":           "fDQwEDN61QE:APA91bH-Q1SPpkpAw8TftMVR0hE3L1UCHAwRCqQ8qsI4d9enBFQZZZTOftz9y_f3Y500-ed_jR8D8p1OJjYxh66PCryzgobGPs17pmwwNjHbI8g9lwxhddaF69ukjjwNnm2Q0MTJCZel",
	}

    // jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
//    / Headers := {"Authorization": "Key=AAAAZsIAg7g:APA91bHXECuDFO7k095_fJA1fJP9hlXhdi8aknnzjQf2rUQUSwQhG-nQe_4xh8xm-pwcuTa6Wg0s21OjzKXIxSoxDqexKaeoAOOdpzDBoYnmlcAm49QWO33blP9WNCFOuoUKPzRXgdHm",
//         		"Content-Type":  "application/json",
//     			}
	fmt.Println("jsonData->", jsonData)
	jsonValue, _ := json.Marshal(jsonData)
	fmt.Println("jsonValue->", jsonValue)
	response, err := http.Post(
		"https://fcm.googleapis.com/fcm/send",
		"Authorization": "Key=AAAAZsIAg7g:APA91bHXECuDFO7k095_fJA1fJP9hlXhdi8aknnzjQf2rUQUSwQhG-nQe_4xh8xm-pwcuTa6Wg0s21OjzKXIxSoxDqexKaeoAOOdpzDBoYnmlcAm49QWO33blP9WNCFOuoUKPzRXgdHm",
        		"Content-Type":  "application/json",
		jsonValue,
	)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("string(data)->", string(data))
	}
	fmt.Println("Terminating the application...")
}
