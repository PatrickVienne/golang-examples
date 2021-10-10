package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type employee struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

// type mytype []map[string]string
type mytype []employee

func main() {
	var filename string = os.Args[1]
	if filename == "users.json" {
		ParseUsers(filename)
	} else if strings.HasSuffix(filename, "jsonl") {
		ParseJsonlFile(filename)
	} else if strings.HasSuffix(filename, "json") {
		ParseJsonFile(filename)
	} else {
		fmt.Println("Unknown Type")
	}
}

func ParseUsers(filename string) {
	// read our opened jsonFile as a byte array.
	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}

func ParseJsonFile(filename string) {
	fmt.Println("Parsin File: " + filename)

	var emp employee
	var data mytype

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Some error occured while reading file. Error: %s", err)
	}
	// err = json.Unmarshal([]byte(file), &emp)
	// if err != nil {
	// 	log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	// }

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	fmt.Printf("emp Struct: %#v\n", emp)

}

func ParseJsonlFile(filename string) {
	fmt.Println("Parsin File: " + filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var emp employee
	var data mytype

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		content := scanner.Bytes()
		err = json.Unmarshal(content, &data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)

		fmt.Printf("emp Struct: %#v\n", emp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
