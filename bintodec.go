package main

import (
	"fmt"
)

func dectobin(num int) { //fungsi konversi bilangan desimal ke binary
	var binary []int //bikin array untuk menyimpan data binary

	for num != 0 { // bikin loop untuk secara konstan membagi 2 dan mengambil sisa bagi untuk disimpan di array binary, nilai yang hanya bisa disimpan yaitu 0 dan 1
		binary = append(binary, num%2)
		num = num / 2 //menambahkan hasil modulo 2 dan hasil bagi 2 ke array binary
	}
	if len(binary) == 0 { //ketika di array binary itu kosong, maka print 0
		fmt.Print("0")
	} else {
		for i := len(binary) - 1; i >= 0; i-- { //ketika di array ada nilai selain 0, maka dijalankan loop untuk membaca array dari urutan belakang
			fmt.Printf("%d", binary[i]) // i yang berkurang karena loop for diatas tadi, diassign ke index array binary
		}
	}
}

func main() { //fungsi input bilangan desimal dan output binary
	var decimal int
	fmt.Scan(&decimal)
	dectobin(decimal)
}
