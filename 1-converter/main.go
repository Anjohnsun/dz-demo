package main

import "fmt"

const (
	usdTOeur = 0.8
	usdTOrub = 75
	eurTorub = usdTOrub / usdTOeur
)

func main() {
	fmt.Println("Значения констант:")
	fmt.Println("usd to eur:", usdTOeur)
	fmt.Println("usd to rub:", usdTOrub)
	fmt.Println("eur to rub:", eurTorub)
}
