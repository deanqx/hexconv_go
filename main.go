package main

import (
	"fmt"
	"os"
	"strings"
)

// Range: 0 to 15
func numberToHex(x uint8) byte {
	switch x {
	case 10:
		return 'A'
	case 11:
		return 'B'
	case 12:
		return 'C'
	case 13:
		return 'D'
	case 14:
		return 'E'
	case 15:
		return 'F'
	default:
		return x + 48 // ascii: 0 is at 48th place
	}
}

func strToHex(str string) string {
	// Every byte of str gets two hex characters and one delimiter
	hex_str := make([]byte, 3*len(str))

	for i, k := 0, 0; i < len(str); i, k = i+1, k+3 {
		var char_a uint8 = str[i] >> 4
		var char_b uint8 = str[i] & 0x0F

		hex_str[k] = numberToHex(char_a)
		hex_str[k+1] = numberToHex(char_b)
		hex_str[k+2] = ' '
	}

	return string(hex_str)
}

func main() {
	var args []string = os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go This is a test")
		fmt.Println("Output: 54 68 69 73 20 69 73 20 61 20 74 65 73 74")
	}

	args_joined := strings.Join(args[1:], " ")
	fmt.Println(strToHex(args_joined))
}
