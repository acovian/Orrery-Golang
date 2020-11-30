package main

import "fmt"

func main() {

	var orbitNumber int
	fmt.Println("Enter an orbital number:")
	fmt.Scanln(&orbitNumber)
	if orbitNumber == 0 {
		fmt.Println("Sun")
	}
	if orbitNumber == 1 {
		fmt.Println("Mercury")
	}
	if orbitNumber == 2 {
		fmt.Println("Venus")
	}
	if orbitNumber == 3 {
		fmt.Println("Earth")
	}
	if orbitNumber == 4 {
		fmt.Println("Mars")
	}
	if orbitNumber == 5 {
		fmt.Println("Ceres")
	}
	if orbitNumber == 6 {
		fmt.Println("Jupiter")
	}
	if orbitNumber == 7 {
		fmt.Println("Saturn")
	}
	if orbitNumber == 8 {
		fmt.Println("Uranus")
	}
	if orbitNumber == 9 {
		fmt.Println("Neptune")
	}
	if orbitNumber == 10 {
		fmt.Println("Pluto")
	}
}
