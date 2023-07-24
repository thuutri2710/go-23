package main

import (
	"fmt"
	"os"
)

func reorder(first_name string, last_name string, language string) (first string, last string) {
	switch language {

	case "vn":
		{
			first = last_name
			last = first_name
		}
	case "us":
		{
			first = first_name
			last = last_name
		}
	default:
		{
			panic("Invalid language")
		}
	}

	return
}

func main() {
	args := os.Args

	length_of_args := int32(len(args))

	if length_of_args < 4 || length_of_args > 5 {
		fmt.Printf("Invalid number of arguments\n")

		return
	}

	first_name, last_name := args[1], args[2]

	var mid_name string
	var language string

	if length_of_args == 5 {
		mid_name = args[3]
		language = args[4]
	} else {
		language = args[3]
	}

	first, last := reorder(first_name, last_name, language)

	if mid_name != "" {
		fmt.Printf("Output: %s %s %s\n", first, mid_name, last)
	} else {
		fmt.Printf("Output: %s %s\n", first, last)
	}
}
