package main

import "fmt"

type nxtState interface {
	getName() string
	next() *nxtState
}

type orderStatus struct {
	name       string
	nextStatus *nxtState
}

func (ord *orderStatus) SetorderStatus(name string, next *nxtState) *orderStatus {
	ord.nextStatus = next
	ord.name = name
	return ord
}

func (ord orderStatus) next() *nxtState {
	return ord.nextStatus
}

func (ord orderStatus) getName() string {
	return ord.name
}

type WaitPayment struct {
	orderStatus
}

func NewWaitPayment() *WaitPayment {
	wt := new(WaitPayment)
	var i nxtState
	var t = NewDelivered()
	i = t
	wt.SetorderStatus("WaitPayment", &i)
	return wt
}

type Delivered struct {
	orderStatus
}

func NewDelivered() *Delivered {
	wt := new(Delivered)
	var i nxtState
	var t = wt
	i = t
	wt.SetorderStatus("Delivered", &i)
	return wt
}

type Order struct {
	state nxtState
}

func NewOrder() *Order {
	ord := new(Order)
	ord.state = NewWaitPayment()
	return ord
}

func (ord *Order) nextState() {
	ord.state = *ord.state.next()
}

func main() {
	myOrder := NewOrder()
	fmt.Println(myOrder.state.getName())
	myOrder.nextState()
	fmt.Println(myOrder.state.getName())

}
