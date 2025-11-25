package main

import "fmt"

const (
	usdTOeur = 0.8
	usdTOrub = 75
	eurTorub = usdTOrub / usdTOeur
)

func main() {
	n, v1, v2 := SaveInput()
	Calculate(n, v1, v2)
}

func SaveInput() (n float64, v1 string, v2 string) {
	fmt.Print("Введите количество:")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	fmt.Print("Введите первую валюту(usd, eur, rub):")
	_, err = fmt.Scan(&v1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Введите вторую валюту(usd, eur, rub):")
	_, err = fmt.Scan(&v2)
	if err != nil {
		panic(err)
	}

	return
}

func Calculate(n float64, v1 string, v2 string) float64 {
	return 0
}
