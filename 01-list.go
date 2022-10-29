package main

import (
	"errors"
	"fmt"
)

func main() {
	l := new()
	count := l.count()
	fmt.Printf("count: %v\n", count)

	fmt.Println("append 10 at tail")
	l.insert(-1, 10)
	fmt.Println(l)
	fmt.Println("append 11 at tail")
	l.insert(-1, 11)
	fmt.Println(l)
	fmt.Println("append 12 at tail")
	l.insert(-1, 12)
	fmt.Println(l)
	fmt.Println("insert 9 at 0")
	l.insert(0, 9)
	fmt.Println(l)
	fmt.Println("insert 9 at 1")
	l.insert(1, 9)
	fmt.Println(l)
	fmt.Println("remove at 1")
	err := l.remove(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(l)
	fmt.Println("append 13 at tail")
	l.insert(-1, 13)
	fmt.Println(l)
	fmt.Println("set 8 at 1")
	l.set(1, 8)
	fmt.Println(l)

	fmt.Println("get at 0")
	first, err := l.get(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("first: %v\n", first)
	}
	fmt.Println("get at 10000")
	_, err = l.get(10000)
	if err != nil {
		fmt.Println(err)
	}
}

const (
	length int = 1
)

type list struct {
	a   []int
	end int
}

// New creates a list
func new() *list {
	list := list{
		a:   make([]int, length),
		end: -1,
	}
	return &list
}

// count
func (l *list) count() int {
	return l.end + 1
}

// get
func (l *list) get(poi int) (int, error) {
	if l.end < 0 || l.end < poi {
		return -1, errors.New("out of range")
	}

	return l.a[poi], nil
}

// set
func (l *list) set(poi int, v int) error {
	if l.end < 0 || l.end < poi {
		return errors.New("out of range")
	}

	l.a[poi] = v

	return nil
}

// insert a value, if poi is -1 will append
func (l *list) insert(poi int, v int) {
	if l.count() == len(l.a) {
		temp := make([]int, len(l.a)*2)
		for i := 0; i <= l.end; i++ {
			temp[i] = l.a[i]
		}
		l.a = temp
	}

	if poi == -1 {
		poi = l.end+1
	} else {
		for i := l.end; i >= poi; i-- {
			l.a[i+1] = l.a[i]
		}
	}
	l.end++
	l.a[poi] = v
}

// remove
func (l *list) remove(poi int) error {
	if l.end < 0 || l.end < poi {
		return errors.New("out of range")
	}

	for i := poi; i < l.end; i++ {
		l.a[i] = l.a[i+1]
	}
	l.a[l.end] = 0
	l.end--

	return nil
}
