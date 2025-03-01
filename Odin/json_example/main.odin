package main

import "core:encoding/json"
import "core:fmt"
import "core:os"

Jtest :: struct {
	Name:       string `json:"name"`,
	Generation: int `json:"generation"`,
	Ability:    string `json:"ability"`,
}

main :: proc() {
	fmt.println("JSON manipulation")
	data, ok := os.read_entire_file_from_filename("data/example.json")
	if !ok {
		fmt.println("Data read incorrectly.")
		return
	}
	defer delete(data)

	jt := Jtest{}
	err := json.unmarshal(data, &jt)
	if err != nil {
		fmt.println(err)
		return
	}

	fmt.println(jt)

	// saving to json
	jotaro := Jtest {
		Name       = "Jotaro Kujo",
		Generation = 3,
		Ability    = "Star Platinum",
	}
	giorno := Jtest {
		Name       = "Giorno Giovanna",
		Generation = 5,
		Ability    = "Golden Experience",
	}
	res := []Jtest{jotaro, giorno}

	fmt.println(res)

	res_byte, err2 := json.marshal(res)
	ok = os.write_entire_file("data/example_save.json", res_byte)
	if !ok {
		fmt.println("Data write incorrectly.")
		return
	}
}
