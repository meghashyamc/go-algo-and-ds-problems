package onstrings

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

/*
Input:Our company wants to create a data entry verification system.

Given an amount in words and an amount indicated by data entry person in numbers, you have to detect whether the amounts are same or not.

Note: There are a lot of corner cases to be considered. Take care of them.

Input:

String num: Amount written in digits as a string. This string will be an integer number without having any commas in between the digits.
String words: Amount written in words according to Indian Numbering System.
Output:

An integer
1: Values match
0: Otherwise
Please note :Every word needs to be separated using "-" rather than a space character
 https://en.wikipedia.org/wiki/Indian_numbering_system

Constraints

The number will fall in integer range(10^9)

Example :

Input :
String  num = "1234"
String words  = "one-thousand-two-hundred-and-thirty-four"

Output:
1
*/

const (
	valZero  = "0"
	nameZero = "zero"
)
const and = "and"

var numberNamesByZeroes = map[int]string{
	2: "hundred",
	3: "thousand",
	5: "lakh",
	7: "crore",
	9: "arab",
}
var numberNames = map[string]string{
	"1":  "one",
	"2":  "two",
	"3":  "three",
	"4":  "four",
	"5":  "five",
	"6":  "six",
	"7":  "seven",
	"8":  "eight",
	"9":  "nine",
	"10": "ten",
	"11": "eleven",
	"12": "twelve",
	"13": "thirteen",
	"14": "fourteen",
	"15": "fifteen",
	"16": "sixteen",
	"17": "seventeen",
	"18": "eighteen",
	"19": "nineteen",
	"20": "twenty",
	"30": "thirty",
	"40": "forty",
	"50": "fifty",
	"60": "sixty",
	"70": "seventy",
	"80": "eighty",
	"90": "ninety",
}

var errUnknownDigits = errors.New("Unknown digit(s) encountered: %s")

func validateNumbersAreTheSame(num1, num2 string) int {
	num1 = strings.TrimSpace(num1)
	num1 = removeLeadingZeroes(num1)
	if num2 == getNumberInWords(num1) {
		return 1
	}

	return 0
}

func removeLeadingZeroes(num string) string {
	if num == valZero {
		return num
	}

	leadingZeroesCount := 0
	for _, ch := range num {
		if ch == '0' {
			leadingZeroesCount++
		}
	}

	return num[leadingZeroesCount:]
}

func getNumberInWords(numStr string) string {

	if len(numStr) == 0 {
		return ""
	}
	if len(numStr) == 1 && numStr == valZero {
		return nameZero
	}
	if len(numStr) == 1 {
		return numberNames[numStr]
	}

	var result string
	for digitsConsidered := 1; digitsConsidered < len(numStr); {
		if digitsConsidered == 1 {
			result = getTwoDigitNumberInWords(numStr[len(numStr)-2:])
			digitsConsidered++
			continue
		}
		if digitsConsidered == 2 {
			thirdLastDigit := numStr[len(numStr)-3 : len(numStr)-2]
			val, ok := numberNames[thirdLastDigit]
			if !ok && thirdLastDigit != valZero {
				log.Fatalf(errUnknownDigits.Error(), thirdLastDigit)
			}

			if thirdLastDigit == valZero && result != "" {
				result = fmt.Sprintf("%s-%s", and, result)
			} else if result != "" {
				result = fmt.Sprintf("%s-%s-%s-%s", val, numberNamesByZeroes[digitsConsidered], and, result)
			} else if thirdLastDigit != valZero && result == "" {
				result = fmt.Sprintf("%s-%s", val, numberNamesByZeroes[digitsConsidered])
			}

			digitsConsidered++
			continue

		}

		if len(numStr) == digitsConsidered+1 {
			numStr = valZero + numStr
		}

		nextDigitsInWords := getTwoDigitNumberInWords(numStr[len(numStr)-digitsConsidered-2 : len(numStr)-digitsConsidered])
		if nextDigitsInWords != "" {
			if result != "" {
				result = fmt.Sprintf("%s-%s-%s", nextDigitsInWords, numberNamesByZeroes[digitsConsidered], result)
			} else {
				result = fmt.Sprintf("%s-%s", nextDigitsInWords, numberNamesByZeroes[digitsConsidered])

			}
		}
		digitsConsidered += 2

	}

	return result
}

func getTwoDigitNumberInWords(numStr string) string {
	var result string
	val, ok := numberNames[numStr[len(numStr)-2:]]
	if ok {
		result = val + result
	} else {
		lastDigit := numStr[len(numStr)-1:]
		val, ok = numberNames[lastDigit]
		if !ok && lastDigit != valZero {
			log.Fatalf(errUnknownDigits.Error(), lastDigit)
		}
		result = val + result
		secondLastDigit := numStr[len(numStr)-2 : len(numStr)-1]
		if secondLastDigit != valZero {
			val, ok = numberNames[secondLastDigit+valZero]
			if !ok {
				log.Fatalf(errUnknownDigits.Error(), secondLastDigit)
			}
			result = fmt.Sprintf("%s-%s", val, result)

		}
	}
	return result
}
