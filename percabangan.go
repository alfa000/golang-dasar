package main

import "fmt"

func main() {
	var name = "Rahmat"

	if name == "Adnan" {
		fmt.Println("Hello Adnan")
	} else if name == "Alfa" {
		fmt.Println("Hello Alfa")
	} else if name == "Romlah" {
		fmt.Println("Hello Romlah")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
	name := "Rahmat"

	switch name {
	case "Adnan":
		fmt.Println("Hello Adnan")
		fmt.Println("Hello Adnan")
	case "Alfa":
		fmt.Println("Hello Alfa")
		fmt.Println("Hello Alfa")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}