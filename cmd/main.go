package main

import (
	"fmt"

	"github.com/DouglasSouzaGomes/expense-tracker/pkg/utils"
)

func main() {

	choice := ""

	fmt.Println("💰 Expense Tracker")
	for choice != "5" {
		fmt.Println("----------------------")
		fmt.Println("1️⃣  Adicionar gasto")
		fmt.Println("2️⃣  Listar gastos")
		fmt.Println("3️⃣  Calcular total")
		fmt.Println("4️⃣  Excluir gasto")
		fmt.Println("5️⃣  Sair")

		choice = utils.ReadInput("Digite a opção desejada: ")

		switch choice {
		case "1":
			utils.ClearScreen()
			fmt.Println("Adicionar gasto")
		case "2":
			utils.ClearScreen()
			fmt.Println("Listar gastos")
		case "3":
			utils.ClearScreen()
			fmt.Println("Calcular total")
		case "4":
			utils.ClearScreen()
			fmt.Println("Excluir gasto")
		case "5":
			utils.ClearScreen()
			fmt.Println("Até mais!")
			return
		default:
			utils.ClearScreen()
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
