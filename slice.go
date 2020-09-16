package main 

import "fmt"
//Slice refresh dari elem array -> dibuat dan dihasilkan dari manipulkasi array atau slice lain nya
func main(){
	//ini adalah array- dan bersifat dinamis dan array itu statis
	var buah = []string{"apel", "mangga", "jeruk", "anggur"}
	//membuat menjadi dinami example
	buah = append(buah, "durian")
	fmt.Println(buah)
	fmt.Println("Jumlah Elmen :" , len(buah))
}
