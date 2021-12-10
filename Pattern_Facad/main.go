package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BaseData struct {
	name string
	data []int
}

func NewBaseData(name string) *BaseData {
	bd := new(BaseData)
	bd.name = name
	return bd
}

func (bd *BaseData)Generate()  {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) + 10
	for i := 0; i < n; i++{
		bd.data = append(bd.data, rand.Intn(10))
	}
}

func (bd BaseData)Print()  {
	fmt.Printf("name = %v, data = %v\n", bd.name, bd.data)
}


type test struct {
	sync.Mutex
	count int
}

func (t test)PrintGenerateSlice()  {
	t.Lock()
	bd := NewBaseData(fmt.Sprintf("Test%v", t.count))
	t.Unlock()
	bd.Generate()
	bd.Print()

}

func main() {
	t := new(test)
	t.PrintGenerateSlice()
}
