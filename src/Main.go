package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	file, err := ioutil.ReadFile("struct.json")
	if err != nil {
		fmt.Println("file does not exist")
	} else {
		json.Unmarshal([]byte(file), &data)
		fmt.Println(data)
		data.Names[0].Name = "John-new"
		res, _ := json.MarshalIndent(data, "", "    ")
		ioutil.WriteFile("struct.json", res, os.ModePerm)
	}
}
