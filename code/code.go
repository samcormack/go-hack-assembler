package code

import (
	"strconv"
)

var compMap = map[string]string{
	"0": "0101010",
	"1": "0111111",
	"-1": "0111010",
	"D": "0001100",
	"A": "0110000",
	"M": "1110000",
	"!D": "0001101",
	"!A": "0110001",
	"!M": "1110001",
	"-D": "0001111",
	"-A": "0110011",
	"-M": "1110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"M+1": "1110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"M-1": "1110010",
	"D+A": "0000010",
	"D+M": "1000010",
	"D-A": "0010011",
	"D-M": "1010011",
	"A-D": "0000111",
	"M-D": "1000111",
	"D&A": "0000000",
	"D&M": "1000000",
	"D|A": "0010101",
	"D|M": "1010101",
}

var destMap = map[string]string{
	"": "000",
	"M": "001",
	"D": "010",
	"MD": "011",
	"A": "100",
	"AM": "101",
	"AD": "110",
	"AMD": "111",
}

var jumpMap = map[string]string{
	"": "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

func dest(d string) string {
	return destMap[d]
}

func comp(c string) string {
	return compMap[c]
}

func jump(j string) string {
	return jumpMap[j]
}

// Left pad string with zeros until it is "length" long
func zeroPad(num string, length int) string {
	for len(num) < length {
		num = "0" + num
	}
	return num
}

// Translate a C command to machine code
func TranslateC(d, c, j string) string {
	return "111" + comp(c) + dest(d) + jump(j)
}

// Translate an A command to machine code
func TranslateA(address int64) string {
	return zeroPad(strconv.FormatInt(address, 2), 16)
}

// func TranslateA(address string) string {
// 	address_dec, _ := strconv.ParseInt(address, 10, 64)
// 	return zeroPad(strconv.FormatInt(address_dec, 2), 16)
// }
