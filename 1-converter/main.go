package main

import "fmt"

type curRate = map[string]float64

var curs = map[string]curRate{}

func main() {
	curs = make(map[string]curRate)
	initCurrences(&curs)

	/*	n, v1, v2 := SaveInput()
		count, currency := Calculate(n, v1, v2, &curs)
		fmt.Println(count, currency)*/

	changeOptions := map[string]func(float64, string) (float64, string){
		"usd": toUsd,
		"eur": toEur,
		"rub": toRub,
	}
	n, v1, v2 := SaveInput()
	op := changeOptions[v2]
	count, currency := op(n, v1)
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

func toUsd(n float64, from string) (float64, string) {
	k, _ := Calculate(n, from, "usd", &curs)
	return k, "usd"
}

func toEur(n float64, from string) (float64, string) {
	k, _ := toUsd(n, from)
	return Calculate(k, from, "eur", &curs)
}

func toRub(n float64, from string) (float64, string) {
	k, _ := toUsd(n, from)
	return Calculate(k, from, "rub", &curs)
}
