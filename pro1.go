package main

import "math"

func convertNumber(num int) string {
	var roman string

	if num == 1 {
		roman = "I"
	} else if num == 2 {
		roman = "II"
	} else if num == 3 {
		roman = "III"
	} else if num == 4 {
		roman = "IV"
	} else if num == 5 {
		roman = "V"
	} else if num == 6 {
		roman = "VI"
	} else if num == 7 {
		roman = "VII"
	} else if num == 8 {
		roman = "VIII"
	} else if num == 9 {
		roman = "IX"
	} else if num == 10 {
		roman = "X"
	} else if num == 20 {
		roman = "XX"
	} else if num == 30 {
		roman = "XX"
	} else if num == 40 {
		roman = "XL"
	} else if num == 50 {
		roman = "L"
	} else if num == 60 {
		roman = "LX"
	} else if num == 70 {
		roman = "LXX"
	} else if num == 80 {
		roman = "LXXX"
	} else if num == 90 {
		roman = "XC"
	} else if num == 100 {
		roman = "C"
	} else if num == 200 {
		roman = "CC"
	} else if num == 300 {
		roman = "CCC"
	} else if num == 400 {
		roman = "CD"
	} else if num == 500 {
		roman = "D"
	} else if num == 600 {
		roman = "DC"
	} else if num == 700 {
		roman = "DCC"
	} else if num == 800 {
		roman = "DCCC"
	} else if num == 900 {
		roman = "CM"
	} else if num == 1000 {
		roman = "M"
	} else if num == 2000 {
		roman = "MM"
	} else if num == 3000 {
		roman = "MMM"
	} else if num == 4000 {
		roman = "MMMM"
	}

	return roman
}

func myfunc(num int, str string, i int) string {
	if num == 0 {
		return str
	} else {
		return myfunc(num/10, convertNumber(num%10*int(math.Pow(10, float64(i))))+str, i+1)
	}
}

func main() {
	// implementation is not yet
	//var i = 0
	//fmt.Println(convertNumber(2))
}
