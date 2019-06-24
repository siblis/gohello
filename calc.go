package main

import (
	"fmt"
	"strconv"
)

func calc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	println("Простой калькулятор. Для выхода введите exit")

	for {
		// Ввод первого числа
		firstNum := enterNumber("Введите первое число: ")

		action := inputData("Укажите действие: ")

		// Ввод второго числа
		secondNum := enterNumber("Введите второе число: ")

		fmt.Printf("Результат: %v \n", calculate(firstNum, action, secondNum))
	}
}

func calculate(firstNum float32, action string, secondNum float32) float32 {
	switch action {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	default:
		fmt.Println("Не удалось распознать действие")
		return 0
	}
}

func inputData(msg string) (data string) {
	fmt.Print(msg)
	fmt.Scanln(&data)
	if data == "exit" {
		panic(data)
	}
	return data
}

func enterNumber(msg string) float32 {
	for {
		num, err := strconv.ParseFloat(inputData(msg), 32)
		if err == nil {
			// Ввод числа прошел без ошибок
			return float32(num)
		}
		fmt.Println("Не удалось распознать число")
	}
}
