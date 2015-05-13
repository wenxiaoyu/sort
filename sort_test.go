package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

type Integer struct {
	v int
}

func (i *Integer) Compared(v interface{}) int {
	if _, ok := v.(*Integer); !ok {
		return 0
	}
	vv := v.(*Integer)
	if i.v > vv.v {
		return -1
	}
	return 1
}

func TestBubbleSort(t *testing.T) {
	c := NewCollection(BubbleSort)
	fmt.Println("TestBubbleSort:before sorted :")
	for i, v := range c.array {
		x := (v).(*Integer)
		fmt.Println("i->", i, "v->", x.v)
	}
	err := c.Sort()
	fmt.Println("TestBubbleSort:after sorted :")
	for i, v := range c.array {
		x := (v).(*Integer)
		fmt.Println("i->", i, "v->", x.v)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestSelectionSort(t *testing.T) {
	c := NewCollection(SelectionSort)
	fmt.Println("SelectionSort:before sorted :")
	for i, v := range c.array {
		x := (v).(*Integer)
		fmt.Println("i->", i, "v->", x.v)
	}
	err := c.Sort()
	fmt.Println("SelectionSort:after sorted :")
	for i, v := range c.array {
		x := (v).(*Integer)
		fmt.Println("i->", i, "v->", x.v)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertionSort(t *testing.T) {
	c := NewCollection(InsertionSort)
	fmt.Println("InsertionSort:before sorted :")
	for i, v := range c.array {
		x := (v).(*Integer)
		fmt.Println("i->", i, "v->", x.v)
	}
	err := c.Sort()
	fmt.Println("InsertionSort:after sorted :")
	for i, v := range c.array {
		x := (v).(*Integer)
		fmt.Println("i->", i, "v->", x.v)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestShellSort(t *testing.T) {
	c := NewCollection(ShellSort)
	fmt.Println("ShellSort:before sorted :")
	for i, v := range c.array {
		x := (v).(*Integer)
		fmt.Println("i->", i, "v->", x.v)
	}
	err := c.Sort()
	fmt.Println("ShellSort:after sorted :")
	for i, v := range c.array {
		x := (v).(*Integer)
		fmt.Println("i->", i, "v->", x.v)
	}
	if err != nil {
		t.Fatal(err)
	}
}


func NewCollection(way SortWay) *collection {
	array := make([]Element, 0, 100)
	for i := 0; i < 20; i++ {
		array = append(array, &Integer{int(rand.Int31n(1000))})
	}
	c := &collection{array, way}
	return c
}

func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer()
	c := NewCollection(BubbleSort)
	b.StartTimer()
	for i:=0;i<b.N;i++{
		c.Sort()
	}
}
func BenchmarkSelectionSort(b *testing.B) {
	b.StopTimer()
	c := NewCollection(SelectionSort)
	b.StartTimer()
	for i:=0;i<b.N;i++{
		c.Sort()
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.StopTimer()
	c := NewCollection(InsertionSort)
	b.StartTimer()
	for i:=0;i<b.N;i++{
		c.Sort()
	}
}

func BenchmarkShellSort(b *testing.B) {
	b.StopTimer()
	c := NewCollection(ShellSort)
	b.StartTimer()
	for i:=0;i<b.N;i++{
		c.Sort()
	}
}
