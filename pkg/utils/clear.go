package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Println("Erro! ", err)
	}
}
