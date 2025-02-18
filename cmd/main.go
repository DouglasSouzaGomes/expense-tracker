package main

import (
	"fmt"

	expense "github.com/DouglasSouzaGomes/expense-tracker/internal"
	"github.com/DouglasSouzaGomes/expense-tracker/pkg/utils"
)

func main() {

	choice := ""

	fmt.Println("💹 Expense Tracker")
	for choice != "5" {
		fmt.Println("----------------------")
		fmt.Println("1️⃣  Adicionar transação")
		fmt.Println("2️⃣  Listar transações")
		fmt.Println("3️⃣  Calcular total")
		fmt.Println("4️⃣  Excluir transação")
		fmt.Println("5️⃣  Sair")
		utils.Space()

		choice = utils.ReadInput("Digite a opção desejada: ")

		switch choice {
		case "1":
			utils.ClearScreen()
			fmt.Println("Adicionar transação")
			utils.Space()

			description := utils.ReadInput("Descrição: ")
			amountStr := utils.ReadInput("Valor: ")
			date := utils.ReadInput("Data: ")

			expense.AddExpense(description, amountStr, date)
		case "2":
			expense.ListExpenses()
		case "3":
			expense.CalculateTotal()
		case "4":
			expense.ListExpenses()
			utils.Space()
			idStr := utils.ReadInput("Qual transação deseja deletar? ")
			expense.DeleteExpense(idStr)
		case "5":
			utils.ClearScreen()
			fmt.Println("Até mais!")
			return
		case "6":
			expense.AddExpense("Lanche", "30", "17/02/2025")
			expense.AddExpense("Torneira", "250", "17/02/2025")
			expense.AddExpense("Abastecimento", "300", "17/02/2025")
		default:
			utils.ClearScreen()
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
