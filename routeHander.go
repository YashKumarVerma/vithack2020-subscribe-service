package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
)

// declare a schema to accept from client
type ParticipantStruct struct {
	Email string `json:"email"`
	Ip    string `json:"ip"`
	Time  int16  `json:"timestamp"`
}

func addListener(res http.ResponseWriter, req *http.Request) {
	statement, err := db.Prepare("INSERT INTO participants(email, ip, timestamp) VALUES (?,?,?) ")
	if err != nil {
		// panic(err.Error())
		log.Fatalln("Error Generating Prepared Statement")
		returnError(res, "Internal Server Error", 500)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// panic(err.Error)
		log.Fatalln("Error Reading  request body")
		returnError(res, "Error Reading Request Body", 400)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	// reading data from user request
	email := keyVal["email"]
	ip := getIPAdress(req)
	time := time.Now().UnixNano()

	_, err = statement.Exec(email, ip, time)
	if err != nil {
		// panic(err.Error())
		log.Fatalln("Error Inserting Record")
		returnError(res, "Error Inserting records", 500)
	}

	log.Printf("New User Added to list : %s", email)
	returnSuccess(res, "new subscriber added")
}

func returnSuccess(res http.ResponseWriter, message string) {
	json := simplejson.New()
	json.Set("seccess", true)
	json.Set("message", message)

	payload, err := json.MarshalJSON()
	if err != nil {
		returnError(res, "Internal Server Error", 500)
		log.Fatalln("Error Creating Success Response")
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(200)
	res.Write(payload)
}

func returnError(res http.ResponseWriter, message string, statusCode int) {
	json := simplejson.New()
	json.Set("seccess", false)
	json.Set("message", message)

	payload, err := json.MarshalJSON()
	if err != nil {
		returnError(res, "Internal Server Error", 500)
		log.Fatalln("Error Creating Success Response")
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(statusCode)
	res.Write(payload)
}
