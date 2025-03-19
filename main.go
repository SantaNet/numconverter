package main

import (
	"NumConvert/internal"
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a number with prefix (0x for hex, 0b for binary, 0o for octal) or 'stop' to exit.")
	fmt.Println("Examples: 0x1A, 0b1010, 0o777, or '0x1A to bin'")

	var input string

	for {
		fmt.Print("> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		input = sc.Text()

		// Преобразуем все буквы в нижний регистр
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "stop" {
			break
		}

		// Если в строке есть указанное преобразование 'to hex' возвращаем основание числа
		numStr, toBase := internal.GetWhatFormat(input)

		// Если в строке есть префикс (0x, 0b, 0o) возвращаем основание числа
		// Иначе используется автоформат, если число без префикса
		cleanNum, fromBase, err := internal.GetFormat(numStr)
		if err != nil {
			log.Println(err)
			continue
		}

		i := new(big.Int)
		if _, ok := i.SetString(cleanNum, fromBase); !ok {
			fmt.Println("Invalid value")
			continue
		}

		fmt.Println(i.Text(toBase))
	}
}
