package main

import (
	"fmt"
)

var oneDigitNames = map[int]string{
	1: "bir",
	2: "iki",
	3: "üc",
	4: "dörd",
	5: "beş",
	6: "altı",
	7: "yeddi",
	8: "səkkiz",
	9: "doqquz",
}
var twoDigitsNames = map[int]string{
	1: "on",
	2: "iyirmi",
	3: "otuz",
	4: "qırx",
	5: "əlli",
	6: "altımış",
	7: "yetmiş",
	8: "səksən",
	9: "doxsan",
}

const threeDigitName = "yüz"
const fourDigitName = "min"
const sevenDigitName = "milion"

func convertNumToString(num int) []string {
	stack := []string{}
	if num == 0 {
		stack = append(stack, "sıfır")
		return stack
	}
	counter := 1

	for i := 1; num != 0; i++ {

		last := num % 10

		if i == 4 {
			stack = append(stack, fourDigitName)
		} else if i == 7 {
			stack = append(stack, sevenDigitName)
		}
		num /= 10

		if counter == 1 {
			if num != 0 {
				stack = append(stack, oneDigitNames[last])
			}
			counter++
		} else if counter == 2 {
			stack = append(stack, twoDigitsNames[last])
			counter++
		} else if counter == 3  {
			if last != 0 {
				stack = append(stack, threeDigitName)
				if last != 1 { //2 yuz
					stack = append(stack, oneDigitNames[last])
				}
			}
			counter = 1
		}
	}
	return stack
}

func main() {

	num1 := 1011
	num2 := 251
	num3 := 45
	num4 := 10
	num5 := 88921
	num6 := 123456
	num7 := 0
	num8 := 10
	num9 := 100
	num10 := 1000
	num11 := 100001
	num12 := 10000001




	res := convertNumToString(num1)
	print(res)
	res = convertNumToString(num2)
	print(res)
	res = convertNumToString(num3)
	print(res)
	res = convertNumToString(num4)
	print(res)
	res = convertNumToString(num5)
	print(res)


	res = convertNumToString(num6)
	print(res)
	res = convertNumToString(num7)
	print(res)
	res = convertNumToString(num8)
	print(res)
	res = convertNumToString(num9)
	print(res)
	res = convertNumToString(num10)
	print(res)

	res = convertNumToString(num11)
	print(res)

	res = convertNumToString(num12)
	print(res)
}

func print(arr []string) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
