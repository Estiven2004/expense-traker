package main

import (
	"encoding/json"
	"errors"
	"os"
)

const fileName = "expenses.json"

func LoadExpenses() ([]Expense, error) {
	var expenses []Expense
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return expenses, nil // Return empty slice if file doesn't exist
		}
		return nil, err // Return error if any other issue occurs
	}
	err = json.Unmarshal(file, &expenses)
	return expenses, err
}

func SaveExpenses(expenses []Expense) error {
	file, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, file, 0644)
}

func FindExpenseIndex(expenses []Expense, id int) (int, error) {
	for i, e := range expenses {
		if e.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("expense not found")
}