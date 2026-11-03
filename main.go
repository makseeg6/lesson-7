package main

import "fmt"

func main() {
	/// #1
	temperature := 5

	if temperature < 0 {
		fmt.Println("сегодня холодно")
	} else if temperature <= 20 {
		fmt.Println("сегодня тепло")
	} else {
		fmt.Println("сегодня жарко")
	}

	/// #2
	score := 75

	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 70 {
		fmt.Println("Хорошо")
	} else if score >= 50 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Не сдал")
	}

	/// #3
	hour := 12

	switch hour {
	case 0, 1, 2, 3, 4, 5:
		fmt.Println("night")

	case 6, 7, 8, 9, 10, 11:
		fmt.Println("morning")

	case 12, 13, 14, 15, 16, 17:
		fmt.Println("day")

	case 18, 19, 20, 21, 22, 23:
		fmt.Println("evening")
	}

	/// #4
	var number int

	fmt.Println("Введите число:")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Четное число")
	} else {
		fmt.Println("Нечетное число")
	}

	/// #5
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Будний день")
	case "Saturday", "Sunday":
		fmt.Println("Выходной")
	default:
		fmt.Println("Некорректный день")
	}

	/// #6
	balance := 1000

	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}

	/// #7
	var age int

	fmt.Println("Введите возраст:")
	fmt.Scan(&age)

	if age < 13 {
		fmt.Println("Ребенок")
	} else if age < 18 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}

	/// #8
	var command string

	fmt.Println("Введите команду:")
	fmt.Scan(&command)

	switch command {
	case "start":
		fmt.Println("Запуск программы")
	case "stop":
		fmt.Println("Остановка программы")
	case "restart":
		fmt.Println("Перезапуск программы")
	default:
		fmt.Println("Неизвестная коммнда")
	}

	/// #9
	grade := 5

	switch grade {
	case 5:
		fmt.Println("A")
	case 4:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 2:
		fmt.Println("D")
	case 1:
		fmt.Println("F")
	default:
		fmt.Println("Некорректная оценка")
	}
}
