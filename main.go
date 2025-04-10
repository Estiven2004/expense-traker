package main

import(
	"flag"
	"fmt"
	"os"
	"time"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Comando requerido: add, list, delete")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		desc := addCmd.String("description", "", "Descripción del gasto")
		amt := addCmd.Float64("amount", 0, "Monto del gasto")
		addCmd.Parse(os.Args[2:])

		expenses,_:= LoadExpenses()
		id := len(expenses) + 1

		expense := Expense{
			ID:          id,
			Description: *desc,
			Amount:      *amt,
			Date:        time.Now(),
	}
	expenses = append(expenses, expense)
	SaveExpenses(expenses)

	fmt.Printf("Gasto agregado exitosamente (ID: %d)\n", id)

	case "list":
		expenses,_ := LoadExpenses()
		fmt.Println("ID	Fecha 		Descripción	 Monto")
		for _, e := range expenses {
			fmt.Printf("%-3d	%-12s	%-12s	$%.2f\n", e.ID, e.Date.Format("2006-01-02"), e.Description, e.Amount)
		}

	case "delete":
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		idFlag := delCmd.Int("id", 0, "ID del gasto a eliminar")
		delCmd.Parse(os.Args[2:])

		expenses, _ := LoadExpenses()
		index, err := FindExpenseIndex(expenses, *idFlag)
		if err != nil {
			fmt.Println("ID no encontrado:")
			return
		}
		expenses = append(expenses[:index], expenses[index+1:]...)
		SaveExpenses(expenses)
		fmt.Printf("Gasto eliminado exitosamente")
		
		case "summary":
			month := 0
			if len(os.Args) > 2 {
				monthArg := os.Args[2]
				if monthArg == "--month" && len(os.Args) > 3 {
					month, _ = strconv.Atoi(os.Args[3])
				}
			}

			expenses, _ := LoadExpenses()
			total := 0.0
			for _, e := range expenses {
				if month == 0 || int(e.Date.Month()) == month {
					total += e.Amount
				}
			}

			if month == 0 {
				fmt.Printf("Total de gastos: $%.2f\n", total)
			} else {
				fmt.Printf("Total de gastos en el mes %d: $%.2f\n", month, total)
			}

		default:
			fmt.Println("Comando no reconocido.")
	}
}