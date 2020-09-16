package main

import "fmt"

func main(){
	var buah=[4]string{"apel", "jeruk",  "anggur", "melon"}
	// Change by array
	buah[1] = "kurma"
	fmt.Println("Jumlah buah adalah :" , len(buah))
	fmt.Println("ISI DARI ARRAY ADALAH: " , buah)
}