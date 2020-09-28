package main

import "fmt"

func main() {
	var day int
	var month int
	fmt.Scan(&day)
	fmt.Scan(&month)

	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		fmt.Print("acuario")
	} else if (month == 2 && day >= 19) || (month == 3 && day <= 10) {
		fmt.Print("piscis")
	} else if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		fmt.Print("aries")
	} else if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		fmt.Print("tauro")
	} else if (month == 5 && day >= 21) || (month == 6 && day <= 20) {
		fmt.Print("geminis")
	} else if (month == 6 && day >= 21) || (month == 7 && day <= 22) {
		fmt.Print("cancer")
	} else if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		fmt.Print("leo")
	} else if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		fmt.Print("virgo")
	} else if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		fmt.Print("libra")
	} else if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		fmt.Print("escorpio")
	} else if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		fmt.Print("sagitario")
	} else if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		fmt.Print("capricornio")
	}
}
