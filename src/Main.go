package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type PersonData struct {
	Id          string   `json:"id"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Names       []Person `json:"names"`
}

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	data := PersonData{}
	resource := "resources/"
	files, err := ioutil.ReadDir(resource)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		jsonFile, _ := ioutil.ReadFile(resource + file.Name())
		fmt.Println(file.Name())
		json.Unmarshal([]byte(jsonFile), &data)
		fmt.Println(data)
		data.Names[0].Name = "John-new"
		res, _ := json.MarshalIndent(data, "", "    ")
		ioutil.WriteFile(resource+file.Name(), res, os.ModePerm)
	}
}
