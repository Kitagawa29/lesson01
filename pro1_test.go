package main

import "testing"

func TestNumConvert01(t *testing.T) {
	str := convertNumber(1)
	if str != "I" {
		t.Error("Test01 is failed")
	}
}

func TestNumConvert02(t *testing.T) {
	str := convertNumber(10)
	if str != "X" {
		t.Error("Test02 is failed")
	}
}

func TestNumConvert03(t *testing.T) {
	str := convertNumber(100)
	if str != "C" {
		t.Error("Test03 is failed")
	}
}

func TestNumConvert04(t *testing.T) {
	str := convertNumber(1000)
	if str != "M" {
		t.Error("Test04 is failed")
	}
}

/*func TestMyFunc01(t *testing.T) {
	str := myfunc(1990)
	if str != "MCMXC" {
		t.Error("Test05 is failed")
	}
}*/

/*func TestMyFunc02(t *testing.T) {
	str := myfunc(99)
	if str != "XCIX" {
		t.Error("Test06 is failed")
	}
}*/
