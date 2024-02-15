package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
	//"strconv" strconv.A
)

var p = fmt.Println

var romNumConvert = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var strNumConv = map[string]int{
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
}

var romDigits = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IV", 10: "X",
}

var x, y int

// THE FUNCTION RECEIVES A STRING AS INPUT AND RETURNS A STRING CLEARED OF SPACES TO THE LEFT
func data_insert() string {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = s.TrimSpace(text)
	if s.Contains(text, " ") != true || s.Contains("+-/*", s.Split(text, " ")[1]) != true {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	p(text)
	return text
}

// The function checks the compliance of the Roman numerals entered by the user with the conditions of the task
func check_if_num_rom(n string, m string) bool {
	a, exist_x := romNumConvert[n]
	b, exist_y := romNumConvert[m]
	if exist_x && exist_y {
		x = a
		y = b
		return true
	} else {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
}

// The function checks the compliance of the Arabic numerals entered by the user with the conditions of the task
func check_if_num_arab(n, m string) bool {
	a, exist_x := strNumConv[n]
	b, exist_y := strNumConv[m]
	if exist_x && exist_y {
		x = a
		y = b
		return true
	} else {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
}

// ARITHMETIC OPERATION
// Operation function
type Operate func(int, int) int

var operators = map[string]Operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"*": func(x, y int) int { return x * y },
	"/": func(x, y int) int { return x / y },
}

// The function converts Arabic numbers in the range [1, 100] to Roman numerals
func convertArabToRom(num int) string {
	var res []string
	if num == 100 {
		res = append(res, "C")
	} else if num >= 90 {
		res = append(res, "XC")
		num -= 90
	} else if num >= 50 {
		res = append(res, "L")
		num -= 50
	} else if num >= 40 {
		res = append(res, "XL")
		num -= 40
	} else if num > 9 {
		i := num / 10
		num -= 10 * i
		for ; i > 0; i-- {
			res = append(res, "X")
		}

	} else if num > 0 {
		res = append(res, romDigits[num])
	}
	sentence := s.Join(res, "")
	return sentence
}

func main() {

	var i = 0
	for i < 1 {
		p("Введите данные для вычисления: ")
		// Get an array of strings separated by space
		l := s.Split(data_insert(), " ")
		if len(l) != 3 {
			panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}
		// Arithmetic Operation Sign
		aos := l[1]
		if check_if_num_rom(l[0], l[2]) {
			res := operators[aos](x, y)
			if res < 1 {
				panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
			} else {
				p(convertArabToRom(res))
			}

		} else if check_if_num_arab(l[0], l[2]) {
			p(operators[aos](x, y))
		}
	}
}
