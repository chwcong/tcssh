package handler

import "testing"

func Test_splitAddress(t *testing.T) {
	testStr1 := "10.5.1.1"
	splitAddress(testStr1)
	testStr2 := "root@10.5.1.1"
	splitAddress(testStr2)
}
