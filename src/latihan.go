package main

import (
	"bufio"   //package scanner reader or write
	"fmt"     //package fmt untuk input/output
	"os"      //menyediakan antarmuka platform-independen untuk fungsionalitas sistem operasi
	"strconv" //conversi string ke tipe data lain
)

func main() {

	var umur = 21                  //variabel dan tipe data
	var nama = "Rony Samuel Tamba" //variabel dan tipe data
	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()

	//Conditions and Boolean Expressions

	x := 5
	y := 6.5

	val := float64(x)+1.5 == y
	fmt.Printf("%t", val)

	//Arithmetic and Math Go

	fmt.Println()
	//1. Integer
	var num1 int = 12
	var num2 int = 10
	var num3 int = 2

	answer := (num1 + num2) * num3
	fmt.Printf("%d", answer+6)
	fmt.Println()

	//1.1 Integer modulus
	var mod1 int = 10
	var mod2 int = 4

	answermod := mod1 / mod2
	fmt.Printf("%d", answermod)
	fmt.Println()

	//2. Float
	var numf1 float32 = 9
	var numf2 float32 = 4

	answerfloat := (numf1 / numf2)
	fmt.Printf("%g", answerfloat)
	fmt.Println()

	//End of Arithmetic and Math Go session

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
	fmt.Println()
	fmt.Println("================================================================================")
	fmt.Println("Scanner")
	fmt.Print("Type something : ")     //Mengetikan sesuatu
	Scanner.Scan()                     //Scanner tipe data String
	Sam := Scanner.Text()              //nama scanner
	fmt.Printf("Your typed : %q", Sam) //Cetak dari (type something), bisa juga langsung cetak tanpa cara yg pertama
	fmt.Println()
	fmt.Print("Type the year you were born : ")                            //Mengetikan sesuatu
	Scanner.Scan()                                                         //Scanner tipe data String
	input, _ := strconv.ParseInt(Scanner.Text(), 10, 64)                   //nama scanner
	fmt.Printf("Your will be %d years old at the end of 2020", 2020-input) //Cetak dari (type something), bisa juga langsung cetak tanpa cara yg

}
