// sort project sort.go
package sort

import (
	"errors"
	//"math"
)

type SortWay int

const (
	BubbleSort SortWay = iota
	SelectionSort
	InsertionSort
	ShellSort
	MergeSort
	QuickSort
	HeapSort
)

//compare interface
//sort metho object must implented this interface
type Element interface {
	Compared(v interface{}) int
}

type collection struct {
	array []Element
	way   SortWay
}

var (
	notImplementCompareError = errors.New("the object of collection not implements the compare interface.")
	notDefineTheSortWay      = errors.New("no defined the sort way.")
	notImplementTheSoryWay   = errors.New("sorry,this sort way not implement by auther")
)

//provieded the sort method
func (c *collection) Sort() error {
	if len(c.array) == 0 {
		return nil
	}
	for _, v := range c.array {
		if _, ok := v.(Element); !ok {
			return notImplementCompareError
		}
	}
	switch c.way {
	case BubbleSort:
		return c.bubbleSort()
	case SelectionSort:
		return c.selectionSort()
	case InsertionSort:
		return c.insertionSort()
	case ShellSort:
		return c.shellSort()
	case MergeSort:
		return c.mergeSort()
	case QuickSort:
		return c.quickSort()
	case HeapSort:
		return c.heapSort()
	default:
		return notDefineTheSortWay
	}
}

func (c *collection) bubbleSort() error {
	n := len(c.array)
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if c.array[j-1].Compared(c.array[j]) == 1 {
				c.array[j-1], c.array[j] = c.array[j], c.array[j-1]
			}
		}
	}
	return nil
}
func (c *collection) selectionSort() error {
	n := len(c.array)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if c.array[j].Compared(c.array[min]) == -1 {
				min = j
			}
		}
		c.array[min], c.array[i] = c.array[i], c.array[min]
	}
	return nil
}
func (c *collection) insertionSort() error {
	n := len(c.array)
	for i := 1; i < n; i++ {
		tmp := c.array[i]
		tmpIndex := i
		for j := i - 1; j >= 0; j-- {
			if tmp.Compared(c.array[j]) == -1 {
				c.array[j+1] = c.array[j]
				tmpIndex = j
			}
		}
		c.array[tmpIndex] = tmp
	}
	return nil
}
func (c *collection) shellSort() error {
	n := len(c.array)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			tmp := c.array[i]
			j := i
			for j >= gap && c.array[j-gap].Compared(tmp) == 1 {
				c.array[j] = c.array[j-gap]
				j = j - gap
			}
			c.array[j] = tmp
		}
		gap = gap / 2
	}
	return nil
}
func (c *collection) mergeSort() error {
	return notImplementTheSoryWay
}

func (c *collection) quickSort() error {
	return nil
}
func (c *collection) heapSort() error {
	return notImplementTheSoryWay
}
