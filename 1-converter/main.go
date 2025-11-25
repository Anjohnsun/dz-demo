package main

import "fmt"

const (
	usdTOeur = 0.8
	usdTOrub = 75
	eurTorub = usdTOrub / usdTOeur
)

func main() {
	n, v1, v2 := SaveInput()
	count, currency := Calculate(n, v1, v2)
	fmt.Println(count, currency)
}

func SaveInput() (n float64, v1 string, v2 string) {
	fmt.Print("Введите количество:")

	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Некорректный ввод, попробуйте ещё раз.")
			continue
		}
		break
	}

	for {
		fmt.Print("Введите первую валюту(usd, eur, rub):")
		_, err := fmt.Scan(&v1)
		if err != nil || !correctCurrencyInput(v1) {
			fmt.Println("Некорректный ввод, попробуйте ещё раз.")
			continue
		}
		break
	}

	for {
		fmt.Print("Введите вторую валюту(usd, eur, rub):")
		_, err := fmt.Scan(&v2)
		if err != nil || !correctCurrencyInput(v2) {
			fmt.Println("Некорректный ввод, попробуйте ещё раз.")
			continue
		}
		break
	}
	return
}

func correctCurrencyInput(s string) bool {
	if s != "usd" && s != "eur" && s != "rub" {
		return false
	}
	return true
}

func Calculate(n float64, v1 string, v2 string) (float64, string) {
	if v1 == v2 {
		return n, v2
	}

	var usdValue float64 = 0
	switch v1 {
	case "usd":
		usdValue = n
	case "eur":
		usdValue = n / usdTOeur
	case "rub":
		usdValue = n / usdTOrub
	default:
		usdValue = 1
	}

	switch v2 {
	case "eur":
		return usdValue * usdTOeur, v2
	case "rub":
		return usdValue * usdTOrub, v2
	default:
		return usdValue, v2
	}
}
