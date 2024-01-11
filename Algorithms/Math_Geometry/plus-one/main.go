package main

import "fmt"

// Constraints:

// 	1 <= digits.length <= 100
// 	0 <= digits[i] <= 9
// 	digits does not contain any leading 0's.

// int n = digits.length;
//     for(int i=n-1; i>=0; i--) {
//         if(digits[i] < 9) {
//             digits[i]++;
//             return digits;
//         }
        
//         digits[i] = 0;
//     }
    
//     int[] newNumber = new int [n+1];
//     newNumber[0] = 1;
    
//     return newNumber;

func plusOne(digits []int) []int {
	n := len(digits)	
	for i:= n - 1; i >=0; i--{
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	newDigits:= make([]int, n + 1)
	newDigits[0] = 1
	return newDigits
}

func main() {
	digits:= []int{1,9}
	fmt.Println(plusOne(digits))
}
