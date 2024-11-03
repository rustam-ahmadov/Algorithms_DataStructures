package main

import (
	"log"
	"strings"
)

//Constraints:
//
//1 <= num <= 3999

var romans = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	return helper(num, 1)
}

func helper(num int, place int) string {
	if num == 0 {
		return ""
	}

	last := num % 10
	var res strings.Builder

	if last >= 5 && last <= 8 {
		res.WriteString(romans[5*place])
		last -= 5
	}

	if last <= 3 {
		for last != 0 {
			res.WriteString(romans[1*place])
			last--
		}
	} else if last == 4 || last == 9 {
		res.WriteString(romans[last*place])
	}

	return helper(num/10, place*10) + res.String()
}

func main() {
	if intToRoman(3749) != "MMMDCCXLIX" {
		log.Fatal(3749)
	}

	if intToRoman(58) != "LVIII" {
		log.Fatal(58)
	}

	if intToRoman(1994) != "MCMXCIV" {
		log.Fatal(1994)
	}

	if intToRoman(382) != "CCCLXXXII" {
		log.Fatal(382)
	}
}
