package main

import "fmt"

func main() {

	people := []string{"Caio", "paochapado", "abidinhs", "Vasco14pt"} // [home](./home)

	for i, person := range people {
		fmt.Println("index:", i, "value:", person)
	}

}
