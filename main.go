package hexconvgo

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var args []string = os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go This is a test")
		fmt.Println("Output: 54 68 69 73 20 69 73 20 61 20 74 65 73 74")
	}

	args_joined := strings.Join(args[1:], " ")
	fmt.Println(StrToHex(args_joined))
}
