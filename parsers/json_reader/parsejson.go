package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
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
	defer duration(track("Parsin File: " + filename))

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
	defer duration(track("Parsin File: " + filename))
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

	// Print Memory Usage after reading employees
	PrintMemUsage()

	// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
	data = nil
	PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()

	fmt.Printf("emp Struct: %#v\n", emp)
}

func ParseJsonlFile(filename string) {
	defer duration(track("Parsin File: " + filename))
	fmt.Println("Parsin File: " + filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	type dataentry map[string]interface{}
	var data dataentry

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		content := scanner.Bytes()
		err = json.Unmarshal(content, &data)
		layout := "2020-10-30T12:24:10+0700"
		datetext := data["date"].(string)
		fmt.Println(datetext)
		t, err := time.Parse(layout, datetext)
		fmt.Println(data["data"], t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data["date"].(string))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
