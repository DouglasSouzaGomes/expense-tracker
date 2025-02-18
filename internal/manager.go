package expense

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/DouglasSouzaGomes/expense-tracker/pkg/utils"
)

var expenses []Expense
var nextId = 1

func AddExpense(description, amountStr, date string) error {
	utils.ClearScreen()

	amountFloat, err := strconv.ParseFloat(amountStr, 64)

	if err != nil {
		fmt.Println("Erro ao adicionar o valor!")
		fmt.Println("Por favor, tente novamente.")
		return err
	}

	newExpense := Expense{Id: nextId, Description: description, Date: date, Amount: amountFloat}
	expenses = append(expenses, newExpense)
	nextId++

	utils.Space()
	fmt.Println("✅ Adicionado com sucesso!")
	return nil
}

func ListExpenses() {
	utils.ClearScreen()
	fmt.Println("💸 Gastos:")
	utils.Space()
	for _, exp := range expenses {
		fmt.Printf("[%v]	%s\nValor:	R$%.2f\nData:	%s\n", exp.Id, exp.Description, exp.Amount, exp.Date)
		utils.Space()
	}
}

func CalculateTotal() {
	utils.ClearScreen()
	var total float64

	for _, exp := range expenses {
		total += exp.Amount
	}

	fmt.Printf("💰 Total de gastos: R$%.2f\n", total)
}

func DeleteExpense(idStr string) error {
	utils.ClearScreen()

	idInt, err := strconv.ParseInt(idStr, 10, 8)

	if err != nil {
		fmt.Println("Erro ao selecionar transação! ")
		fmt.Println("Por favor, tente novamente.")
		return err
	}

	for i, exp := range expenses {
		if exp.Id == int(idInt) {
			expenses = slices.Delete(expenses, i, i+1)
		}
	}

	utils.Space()
	fmt.Println("✅ Transação deletada com sucesso")
	return nil
}
