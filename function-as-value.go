package main

import "fmt"

func saygoodbye(name string) string {
	return "good bye " + name
}

func main() {
	goodbye := saygoodbye("melda")
	fmt.Println(goodbye)
}
