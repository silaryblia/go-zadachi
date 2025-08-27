package main

import (
	"errors"
	"fmt"
)

var errNegativeBalance = errors.New("error: баланс не может быть отрицательным")
var errNegativeQuantity = errors.New("error: количество не может быть отрицательным")

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner}
}

//

func (a *Account) SetBalance(balance float64) error {
	if balance < 0 {
		return errNegativeBalance
	}

	a.balance = balance
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

//

// метод зачисления средств на счет
func (a *Account) Deposit(quantity float64) error {
	if quantity < 0 {
		return errNegativeQuantity
	}
	a.balance += quantity // зачисление средств
	return nil
}

// метод снятия средств со счета
func (a *Account) Withdraw(quantity float64) error {
	if quantity < 0 {
		return errNegativeQuantity
	}

	a.balance -= quantity // снятие средств
	return nil
}

func main() {
	a := &Account{owner: "ownername"}
	a.SetBalance(100)

	fmt.Println(a.GetBalance()) // 100

	err := a.Deposit(10)
	if err != nil {
		panic(err) // вывод сообщения об ошибке с её описанием
	}

	err = a.Withdraw(100)
	if err != nil {
		panic(err) // вывод сообщения об ошибке с её описанием
	}

	fmt.Println(a.GetBalance()) // 10
	fmt.Println(a.GetBalance())
}
