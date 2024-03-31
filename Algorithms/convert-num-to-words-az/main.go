package main

import "fmt"

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

var twoDigitNames = map[int]string{
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

const fourDigitName = " min"
const sevenDigitName = " milyon"
const tenDigitsName = " milyard"

func addAdditionalWordInFrontIfNeeded(last, next int, current, additional string) string {
	if last == 1 && next == 0 {
		return additional + current
	}
	return current
}

func convert(num int) string {
	if num == 0 {
		return oneDigitNames[num]
	}
	return helper(num, true, 1, 1)
}

func helper(num int, isCategoryWordAdded bool, digitPosition int, digitCount int) string {
	cur := num % 10
	leftNum := num / 10

	if digitPosition > 3 {
		digitPosition = 1
		isCategoryWordAdded = false
	}

	if cur == 0 { //we just pass it to the next lvl
		return helper(leftNum, isCategoryWordAdded, digitPosition+1, digitCount+1)
	}

	var res string
	switch digitPosition {
	case 1:
		res = oneDigitNames[cur] //1
	case 2:
		res = twoDigitNames[cur] //11
	case 3:
		res = threeDigitName[cur] //111
		res = addAdditionalWordInFrontIfNeeded(cur, leftNum, res, "bir ")
	}

	if !isCategoryWordAdded && digitCount > 3 {
		switch digitCount {
		case 4, 5, 6: //1_000 ... 100_000
			res += fourDigitName
		case 7, 8, 9: //1_000_000 ... 100_000_000
			res += sevenDigitName
		case 10, 11, 12: //1_000_000_000 ... 100_000_000_000
			res += tenDigitsName
		}
		isCategoryWordAdded = true
	}

	if leftNum == 0 { //if cur digit is the last one
		return res
	}
	return helper(leftNum, isCategoryWordAdded, digitPosition+1, digitCount+1) + " " + res
}

func main() {
	fmt.Printf("%s%c\n", convert(1), '*')               //ok
	fmt.Printf("%s%c\n", convert(10), '*')              //1 2 3
	fmt.Printf("%s%c\n", convert(100), '*')             //1 2 3
	fmt.Printf("%s%c\n", convert(1_000), '*')           //1 2 3
	fmt.Printf("%s%c\n", convert(10_000), '*')          //1 2 3
	fmt.Printf("%s%c\n", convert(100_000), '*')         //1 2 3
	fmt.Printf("%s%c\n", convert(1_000_000), '*')       //1 2 3
	fmt.Printf("%s%c\n", convert(10_000_000), '*')      //1 2 3
	fmt.Printf("%s%c\n", convert(100_000_000), '*')     //1 2 3
	fmt.Printf("%s%c\n", convert(1_000_000_000), '*')   //1 2 3
	fmt.Printf("%s%c\n", convert(10_000_000_000), '*')  //1 2 3
	fmt.Printf("%s%c\n", convert(113_168_135_431), '*') //1 2 3
	fmt.Printf("%s%c\n", convert(1123), '*')            //1 2 3
	fmt.Printf("%s%c\n", convert(353), '*')             //1 2 3
	fmt.Printf("%s%c\n", convert(349_111_123), '*')     //1 2 3
	fmt.Printf("%s%c\n", convert(88_999_123), '*')      //1 2 3
	fmt.Printf("%s%c\n", convert(0), '*')               //1 2 3
	//
	//num, err := strconv.Atoi("40")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(convert(num))

}
