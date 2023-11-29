package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a int

func main() {
	// Принимаем выражение
	fmt.Println("Введите выражение формата a + b:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Делим на части
	parts := strings.Split(input, " ")

	// Ошибка если неправильный формат
	if len(parts) != 3 {
		fmt.Println("Некорректный ввод. Пожалуйста, введите выражение в формате: a + b.")
		return
	}

	RimChisla := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	a := RimChisla[parts[0]]
	b := RimChisla[parts[2]]

	oneFormat(a, b)

	c := trans(a, parts[0])
	d := trans(b, parts[2])
	operator := parts[1]

	//fmt.Println(c)
	//fmt.Println(d)

	oshibka10(c)
	oshibka10(d)

	switch operator {
	case "+":
		result := c + d
		R := resultRim(result, a)
		fmt.Println(R)
	case "-":
		result := c - d
		R := resultRim(result, a)
		fmt.Println(R)
	case "*":
		result := c * d
		R := resultRim(result, a)
		fmt.Println(R)
	case "/":
		result := c / d
		R := resultRim(result, a)
		fmt.Println(R)
	default:
		fmt.Println("Неверная арифметическая операция")

	}

}

// Преобразованмя чисел
func trans(a int, val string) int {
	if a == 0 {
		num, err := strconv.Atoi(val)
		if err == nil {
			return num
		} else {
			fmt.Println("Введите целые числа")
			os.Exit(1)
			return 0
		}
	} else {
		return a
	}
}

// Ошибка если числа меньше 1 или больше 10
func oshibka10(x int) {
	if x < 1 || x > 10 {
		fmt.Println("Введите числа от 1 до 10")
		os.Exit(1)
	}
}

// Проверка что числа введены одного формата
func oneFormat(x, y int) {
	if (x == 0 && y != 0) || (x != 0 && y == 0) {
		fmt.Println("Введите чмсла от 1 до 10 в одгой системе счисления")
		os.Exit(1)
	}
}

func resultRim(result, a int) string {
	if a != 0 {
		//Карта для ответа римскими числами
		RimChisla2 := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX", 21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX", 31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL", 41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L", 51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX", 61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX", 71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX", 81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC", 91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C"}
		if result < 1 {
			fmt.Println("Результат меньше 1, в римской системе счисления нет чисел меньше 1")
			os.Exit(1)
		} else {
			return RimChisla2[result]
		}
	}
	return fmt.Sprint(result)
}
