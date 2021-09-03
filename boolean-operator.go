package main

import "fmt"

func main() {
	var nilaiUjian = 90
	var absensi = 80

	var lulusNilaiUjian bool = nilaiUjian > 80
	var lulusAbsensi bool = absensi > 80
	fmt.Println(lulusNilaiUjian)
	fmt.Println(lulusAbsensi)

	var lulus bool = lulusNilaiUjian && lulusAbsensi
	fmt.Println(lulus)
}
