package main

import "fmt"

func main() {
	// Exercise 1
	// for i := 0; i < 10001; i++ {
	// 	fmt.Println(i)
	// }

	// Exercise 2
	// for i := 65; i <= 90; i++ {
	// 	fmt.Println(i)
	// 	for n := 0; n < 3; n++ {
	// 		fmt.Printf("\t%U\n", i)
	// 	}
	// }

	// Exercise 3
	// for n := 1985; n < 2020; n++ {
	// 	fmt.Printf("%v\n", n)
	// }
	// x := 1985
	// for {

	// 	if x > 2019 {
	// 		break
	// 	} else {
	// 		fmt.Printf("%v\n", x)
	// 		x = x + 1
	// 	}

	// }

	// Exercise 4
	// for x := 10; x < 101; x++ {
	// 	y := x % 4
	// 	fmt.Println(y)
	// }

	// Exercise 5
	// for i := 1; i < 20; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%v is an even number\n", i)

	// 	}
	// }

	// Exercise 6, 7
	// for i := 1; i < 21; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%v is an even number\n", i)

	// 	} else if i < 5 {
	// 		fmt.Printf("%v is an odd number and less than 5\n", i)
	// 	} else {
	// 		fmt.Printf("%v is an odd number and greater than 5\n", i)
	// 	}
	// }

	// Exercise 8
	// switch {
	// case 1 > 10:
	// 	fmt.Printf("This will not print\n")
	// 	fallthrough

	// case 15 > 5:
	// 	fmt.Printf("This will print\n")

	// }

	// Exercise 9
	// favSport := "Football"
	// switch favSport {
	// case "Soccer":
	// 	fmt.Println("That's not my favorite")
	// case "Baseball", "Football", "Basketball":
	// 	fmt.Println("It's in there")
	// case "Bowling":
	// 	fmt.Println("That's not it!!!")
	// 	fallthrough
	// case "Hockey":
	// 	fmt.Println("Used to be")
	// }

	// Exercise 10

	fmt.Println(true && true)
	fmt.Println("THis should print true")
	fmt.Println(true && false)
	fmt.Println("This will print false")
	fmt.Println(true || true)
	fmt.Println("This will print true")
	fmt.Println(true || false)
	fmt.Println("This will print true")
	fmt.Println(!true)
	fmt.Println("This will print false")
}
