package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erro ao ler comando: ", err)
		return ""
	}

	return strings.TrimSpace(input)
}
