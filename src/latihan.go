package main

import "fmt" //package fmt untuk input/output

func main() {
	var umur = 21                  //variabel dan tipe data
	var nama = "Rony Samuel Tamba" //variabel dan tipe data

	fmt.Println("Hello World, My name is", nama, ", im", umur) //Cetak data
	fmt.Printf("100 is binary from : %b", 100)                 //Print console fmt binary
	fmt.Println()
	fmt.Printf("This is : %t", false) //Print console fmt boolean
	fmt.Println()
	fmt.Printf("Number float : %f", 2.10901) //Print console fmt float
	fmt.Println()
	fmt.Printf("This is : %s", "Samuel") //Print console fmt string (deafult)
	fmt.Println()
	fmt.Printf("This is : %q", "Samuel") //Print console fmt string (double quote)
	fmt.Println()
	fmt.Printf("This is 9 width & percision : %9s", "Sam") //Print console fmt width & precision
	fmt.Println()
	fmt.Printf("This : %-9s is cool", "Sam") //Print console fmt width & precision (with spacing)

}
