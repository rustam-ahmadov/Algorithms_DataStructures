package main

var oneDigitNames = map[int]string{
	0: "sıfır",
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

var threeDigitName = map[int]string{
	1: "yüz",
	2: "iki yüz",
	3: "üc yüz",
	4: "dörd yüz",
	5: "beş yüz",
	6: "altı yüz",
	7: "yeddi yüz",
	8: "səkkiz yüz",
	9: "doqquz yüz",
}

const fourDigitName = "min"
const sevenDigitName = "milion"

func convert(num int) string {
	if num == 0 {
		return oneDigitNames[num]
	}
	counter, digitCount := 1, 1
	isCategoryWordAdded := false
	return helper(num, isCategoryWordAdded, counter, digitCount)
}

func increase(counter int) int {
	if counter == 3 {
		return 1
	}
	counter++
	return counter
}
func helper(num int, isCategoryWordAdded bool, counter int, digitCount int) string {
	if num == 0 {
		return "" //end
	}

	last := num % 10
	if last == 0 { //we just pass it to the next lvl
		counter = increase(counter)
		digitCount++
		return helper(num/10, isCategoryWordAdded, counter, digitCount)
	}

	var cur string
	isCategoryHasToChange := false
	switch counter {
	case 1:
		cur = oneDigitNames[last]
	case 2:
		cur = twoDigitsNames[last]
	case 3:
		cur = threeDigitName[last]
		isCategoryHasToChange = true
	}

	if !isCategoryWordAdded {
		if digitCount > 3 && digitCount < 7 {
			if last == 1 && num/10 == 0 && counter == 1 {
				cur = fourDigitName
			} else {
				cur += " " + fourDigitName
			}
			isCategoryWordAdded = true
		} else if digitCount > 6 && digitCount < 10 { 
			cur += " " + sevenDigitName
			isCategoryWordAdded = true
		}
	}

	if isCategoryHasToChange {
		isCategoryWordAdded = false
	}

	counter = increase(counter)
	digitCount++
	nextRes := helper(num/10, isCategoryWordAdded, counter, digitCount)
	return nextRes + cur + " "
}

func main() {
	println(convert(10))
	println(convert(1112))
	println(convert(251))
	println(convert(45))
	println(convert(88_921))
	println(convert(123_456))
	println(convert(0))
	println(convert(10))
	println(convert(100))
	println(convert(1_000))
	println(convert(100_001))
	println(convert(1_000_001)) //todo
	println(convert(10_000_001))
	println(convert(11_301_291))
}
