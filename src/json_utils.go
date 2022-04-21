package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	readCli()
}

func readCli() {
	fmt.Print("Enter path: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")

	readJson(input)
}

func readJson(input string) {
	fmt.Println(input)
	data := PersonData{}
	files, err := ioutil.ReadDir(input)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		jsonFile, _ := ioutil.ReadFile(input + "/" + file.Name())
		fmt.Println(file.Name())
		json.Unmarshal([]byte(jsonFile), &data)
		fmt.Println(data)
		data.Names[0].Name = "John-new"
		res, _ := json.MarshalIndent(data, "", "    ")
		ioutil.WriteFile(input+"/"+file.Name(), res, os.ModePerm)
	}
}
