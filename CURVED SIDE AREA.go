package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	intro()
	main_menu()
	farewell()
}

func hapusLayar() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main_menu() {
	var menu int
	for {
		fmt.Println("=====================")
		fmt.Println("  CURVED SIDE AREA   ")
		fmt.Println("=====================")
		fmt.Println("1. Calculate the area of the cone")
		fmt.Println("2. Calculate the area of the tube")
		fmt.Println("3. Calculate the area of the ball")
		fmt.Println("4. Exit")
		fmt.Println("=====================")
		fmt.Println("Choose the number (1/2/3/4)")
		fmt.Print("Enter your choice : ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			hapusLayar()
			hitung_luas_kerucut()
		case 2:
			hapusLayar()
			hitung_luas_tabung()
		case 3:
			hapusLayar()
			hitung_luas_bola()
		}
		if menu == 4 {
			break
		}
	}
}

func intro() {
	fmt.Println("Welcome to Calculating the Area of ​​Constructed Curved Sides")
}

func hitung_luas_kerucut() {
	var pi, r, s, luas_kerucut float64
	var sudah string
	var selesai bool
	pi = 3.14
	selesai = false
	for !selesai {
		fmt.Println("= YOU WILL CALCULATE THE AREA OF THE CONES =")
		fmt.Println("Enter the value in meters ^-^")
		fmt.Print("Enter the radius value : ")
		fmt.Scan(&r)
		fmt.Print("Enter painter line value : ")
		fmt.Scan(&s)
		luas_kerucut = (pi * r) * (pi + s)
		fmt.Printf("The area of ​​your cone is %v square meter", luas_kerucut)
		fmt.Println()
		fmt.Println("Want to count again? (y/t) : ")
		fmt.Scan(&sudah)
		hapusLayar()
		if sudah == "t" {
			selesai = true
		}
	}

}

func hitung_luas_bola() {
	var pi, r, luas_bola float64
	var sudah string
	var selesai bool
	pi = 3.14
	selesai = false
	for !selesai {
		fmt.Println("= YOU WILL CALCULATE THE AREA OF THE BALL =")
		fmt.Println("Enter Value r : ")
		fmt.Scanln(&r)
		luas_bola = (4 * pi) * (r * r)
		fmt.Println("The surface area of ​​your sphere is ", luas_bola)
		fmt.Println()
		fmt.Println("Want to count again? (y/t) : ")
		fmt.Scan(&sudah)
		hapusLayar()
		if sudah == "t" {
			selesai = true
		}
	}
}

func hitung_luas_tabung() {
	var pi, r, t, luas_tabung float64
	var sudah string
	var selesai bool
	pi = 3.14
	selesai = false
	for !selesai {
		fmt.Println("= KAMU AKAN MENGHITUNG LUAS TABUNG =")
		fmt.Println("Enter Value r : ")
		fmt.Scanln(&r)
		fmt.Println("Enter Value t : ")
		fmt.Scanln(&t)
		luas_tabung = (2 * pi * r) * (r + t)
		fmt.Println("The surface area of ​​your cylinder is ", luas_tabung)
		fmt.Println()
		fmt.Println("Want to count again?? (y/t) : ")
		fmt.Scan(&sudah)
		hapusLayar()
		if sudah == "t" {
			selesai = true
		}
	}
}

func farewell() {
	fmt.Println("Thankss. See you")
}
