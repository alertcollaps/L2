package main

import "fmt"

type Master struct {
	balance int
	name    string
	Account
}

func NewMaster(mon int) *Master {
	ms := new(Master)
	ms.name = "Master card"
	ms.balance = mon
	return ms
}

type Qiwi struct {
	balance int
	name    string
	Account
}

func NewQiwi(mon int) *Master {
	ms := new(Master)
	ms.name = "QIWI card"
	ms.balance = mon
	return ms
}

func (ms Master) checkBalance() int {
	return ms.balance
}
func (qw Qiwi) checkBalance() int {
	return qw.balance
}

func (ms *Master) decrease(d int) {
	ms.balance = ms.balance - d
}
func (qw *Qiwi) decrease(d int) {
	qw.balance = qw.balance - d
}
func (ms *Master) getName() string {
	return ms.name
}
func (qw *Qiwi) getName() string {
	return qw.name
}

type Account struct {
	incomer Cards
}
type Cards interface {
	checkBalance() int
	decrease(int)
	getName() string
	pay(int, Cards)
}

func (acc Account) pay(cash int, crd Cards) {
	if cash <= crd.checkBalance() {
		fmt.Printf("Paid %v using %v\n", cash, crd.getName())
		crd.decrease(cash)
	} else if acc.incomer != nil {
		fmt.Printf("Can't pay using %v\n", crd.getName())
		acc.incomer.pay(cash, acc.incomer)
	} else {
		fmt.Printf("No method to pay\n")
	}
}

func (acc *Account) setNext(crd Cards) {
	acc.incomer = crd
}

func main() {
	master := NewMaster(300)
	qiwi := NewQiwi(500)
	master.setNext(qiwi)
	fmt.Println("Current balance:", master.checkBalance())
	master.pay(400, master)
	fmt.Println("Balance now:", master.checkBalance())
}
