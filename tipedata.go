package main

import "fmt"

func main(){
	// var number_positif wint64 = 50000000000000000
	var desimal=3.5
	// fmt.Println(desimal)

	// if desimal !=2{
	// 	fmt.Println("Tidasama Bro")
	// } else{
	// 	fmt.Println("Gak ada yang sama")
	// }
	

	switch desimal{
		case 10:
			fmt.Println("Ini adalah 10")
		case 3.55:
			fmt.Println("Ini adalah 3.55")
		// case 3.5:
		// 	fmt.Println("Yeayyyy ini 3.5")
		default:
			fmt.Println("Kok Gak Masuk?")
	}
}