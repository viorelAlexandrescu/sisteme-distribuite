package main

import (
    "fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
)

func parsePayload(arr []string) string {
	name := arr[0]
	gender := arr[1]
	birthdate := strings.Split(arr[2], "-")

	birthyear, err := strconv.ParseInt(birthdate[0], 10, 64)

	if err != nil {
        panic(err)
    }

	age := 2017 - birthyear

	result := "Hello "

	if (gender == "Male") {
		result += "dude!"
	} else {
		result += "sister!"
	}

	result += "\nI am glad to meet you " + name

	result += "\nYou are " + strconv.FormatInt(age, 10) + " years old! Phew..."

	return result
}

//payload is csv -> Name,Gender,Birthdate
func handler(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, err := ioutil.ReadAll(r.Body)
	body := ""
	for _, char := range bodyBuffer {
        body += string(char)
	}
	
	if err != nil {
        panic(err)
    }

	payload := strings.Split(body, ",")
	responseMessage := parsePayload(payload)

	fmt.Fprintln(w, responseMessage)
}

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)
	fmt.Println("Go Web server listenting on " + port)
	http.ListenAndServe(port, nil)
}