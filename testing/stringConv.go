package main

import "fmt"
import "regexp"

//import "strconv"

func main() {
	fmt.Println("Hello, playground")

	var validID = regexp.MustCompile(`\d{1,}[.]\d{1,}|\d{1,}`)
	str_numbers := validID.FindAllString("Hello, the car is $79000. But I pay only $100.50.", -1)
	fmt.Println(str_numbers)
	//	total := 0
	fmt.Println(str_numbers)
	for _, elem := range str_numbers {
		fmt.Println(elem)
		// total += elem
	}
	//	fmt.Println(total)

	//	var numbers []float64

	//	for _, elem := range str_numbers {
	//		i, err := strconv.ParseFloat(elem, 64)
	//		if err == nil {
	//			numbers = append(numbers, i)
	//		}
	//	}
	//	fmt.Println(numbers)
}
