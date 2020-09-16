package main
import "fmt"
// map adalah tipe data assosiatif nilai nya berupa key / value (harus unik)
// klarena digunakan sebagai identifier
// Membuat sebuah array harus memanggil index, tapi di map tinggal panggil nama identifier nya
func main(){
	var harga_makanan = map[string]int{"ayam":20000, "sate":15000}
	fmt.Println("harga ayam", harga_makanan["ayam"])
	//Key nya harus sama percis yaa karena sebagai indetifier
}