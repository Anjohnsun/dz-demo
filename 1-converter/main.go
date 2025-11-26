package main

import "fmt"

type curRate = map[string]float64

func main() {
	curs := make(map[string]curRate)
	initCurrences(&curs)

	n, v1, v2 := SaveInput()
	count, currency := Calculate(n, v1, v2, &curs)
	fmt.Println(count, currency)
}

func initCurrences(curs *map[string]curRate) {
	usdMap := curRate{
		"usd": 1,
		"eur": 0.8,
		"rub": 70,
	}

	eurMap := curRate{
		"usd": 1.15,
		"eur": 1,
		"rub": 80,
	}

	rubMap := curRate{
		"usd": 0.014,
		"eur": 0.012,
		"rub": 1,
	}

	*curs = map[string]curRate{
		"usd": usdMap,
		"eur": eurMap,
		"rub": rubMap,
	}
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

func Calculate(n float64, v1 string, v2 string, curs *map[string]curRate) (float64, string) {
	return (*curs)[v1][v2] * n, v2
}
