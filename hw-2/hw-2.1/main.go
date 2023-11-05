package main

import (
	"fmt"
	"strings"
)
//Создаю структру roman
type roman struct {
	number int
	sign   string
}
//Создаю массив из объектов структры roman
var m = [13]roman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
    {1, "I"},
}
/*
В данной функции временная сложность составляет O(n) где n - входное число, и для каждой итерации цикла будет 
производиться поиск подходящей римской цифры в массиве m.
*/
func intToRoman(input int) string {
    if input > 3999 || input <= 0 { //больше 3999 мы не можем вывести с существующими данными 
		return ""
	}

    var str strings.Builder // Использую чтобы легче конкатенировать строки

	for i := 0; input != 0; {
		now := m[i]

		if input >= now.number {
			str.WriteString(now.sign)
			input -= now.number
		} else {
			i++
		}
	}

	return str.String()
    
}


func main() {
	fmt.Println(intToRoman(33))
}