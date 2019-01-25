localhost:12345/person          <-- To #POST the Data in body 
 {
	 	"Firstname": "TestFisrtname",
	 	"Lastname": "TestLastname",
	 	"age": 21,
 	"Gender": "male",
 	"Address": {
 		"City": "CityName",
 		"State": "StateName",
 		"pin": "000000"
 	},
 	"Contact": {
 		"Mobile": "96161",
 		"Email": "test@test.com"
 	}
 }
------------------------------------------------
 localhost:12345/               <-- To #GET the  Hello User!
 localhost:12345/person         <-- To #GET all records from Database 
 localhost:12345/person/10      <-- To #GET 10th number person's Data
 localhost:12345/person/f/10    <-- To #GET the  10th number person's Fullname
 localhost:12345/person/a/10    <-- To #GET the  10th number person's Address
 localhost:12345/person/c/10    <-- To #GET the  10th number person's Contact
 localhost:12345/person/10      <-- To #DELETE the  10th number person's Data
 localhost:12345/person         <-- To #PUT update the  10th number person's Address
  {
	 	"Firstname": "TestFisrtnameChanged",
	 	"Lastname": "TestLastnameChanged",
	 	"age": 21, , <-- change te age
 	"Gender": "male", 
 	"Address": {
 		"City": "CityNameChanged",
 		"State": "StateNameChanged",
 		"pin": "000000" <-- change the pin 
 	},
 	"Contact": {
 		"Mobile": "96161", <-- change the mobile 
 		"Email": "test@test.com" <-- change the email
 	}
 }
 ------------------------------------------------