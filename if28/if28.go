package if28

import "fmt"

func Task28(year int64) int64 {
	if (year % 4) == 0 {
		if year%100 == 0 && year%400 != 0 {
			return 365
		} else {
			return 366
		}
	} else {
		return 365
	}
}

func Task29(number int64) {

	if number > 0 {
		if number%2 == 0 {
			fmt.Println("\"Положительное четное число\"")
			return
		} else {
			fmt.Println("\"Положительное нечетное число\"")
			return
		}
		return
	} else if number < 0 {
		if number%2 == 0 {
			fmt.Println("\"отрицательное четное число\"")
			return
		} else {
			fmt.Println("\"отрицательное нечетное число\"")
			return
		}
	} else {
		fmt.Println("\"нулевое число\"")
		return
	}
}

func Task30(number int64) {

	if number > 0 && number < 1000 {
		if number < 10 {
			if number%2 == 0 {
				fmt.Println("четное однозначное число")
				return
			} else {
				fmt.Println("нечетное однозначное число")
				return
			}
		} else if number > 9 && number < 100 {
			if number%2 == 0 {
				fmt.Println("четное двузначное число")
				return
			} else {
				fmt.Println("нечетное двузначное число")
				return
			}
		} else {
			if number%2 == 0 {
				fmt.Println("четное трехзначное число")
				return
			} else {
				fmt.Println("нечетное трехзначное число")
				return
			}
		}

	} else {
		fmt.Println("Неверное число. Вводите число в диапазоне 1-999")
		return
	}
}
