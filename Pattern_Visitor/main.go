package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Smartphone struct {
	inch float64
	year int
}

type Table struct {
	length int
	width int
}

func NewSmartphone() *Smartphone {
	sm := new(Smartphone)
	sm.inch = 5 + rand.Float64()*(7 - 5)
	i, err := strconv.Atoi(time.Now().Format("2006"))
	if err == nil {
		sm.year = i
	} else {
		sm.year = 2000
	}
	return sm
}

func NewTable() *Table {
	tb := new(Table)
	tb.length = 2 + rand.Intn(5)
	tb.width = 1 + rand.Intn(4)
	return tb
}
func exportData(i interface{}){
	if tb, ok := i.(*Table); ok{
		fmt.Printf("Length of table = %v. Width of table = %v\n", tb.width, tb.width)
	}else if sm, ok := i.(*Smartphone); ok{
		fmt.Printf("Inch of smartphone = %v. Year of smartphone = %v\n", sm.inch, sm.year)
	}else{
		fmt.Printf("Can't find a need type")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	smart := NewSmartphone()
	tbl := NewTable()
	exportData(smart)
	exportData(tbl)
}
