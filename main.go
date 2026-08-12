package main

import "fmt"
import "os" 

func main() {
	fmt.Println("sprout 🌱")
	entries, err := os.ReadDir(".")
	
	if err != nil {
		fmt.Println("Error",err)
		return
	}
	
	for _,entry := range entries{
		fmt.Println(entry.Name())
	}
}
