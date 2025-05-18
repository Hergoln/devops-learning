package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

// use golang container to run this code before comiting

func main() {
	callRandomAPI()
}

func callRandomAPI() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(responseData))
}