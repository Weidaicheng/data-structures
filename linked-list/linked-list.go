package linked_list

import (
	"errors"
	"fmt"
)

// linkedList is a type for linked list
type linkedList struct {
	data int
	next *linkedList
}

// new creates a linked list object
func new() *linkedList {
	linkedList := linkedList{
		data: 0,
		next: nil,
	}

	return &linkedList
}

// print prints data array of linked list
func (l *linkedList) print() {
	fmt.Println(l)

	fmt.Print("[")
	if l.next != nil {
		l = l.next
		for {
			fmt.Printf("%v", l.data)

			if l.next != nil {
				fmt.Print(", ")
				l = l.next
				continue
			} else {
				break
			}
		}
	}
	fmt.Println("]")
}

// get returns value at specific position
func (l *linkedList) get(poi int) (int, error) {
	i := -1
	for {
		if i == poi {
			return l.data, nil
		}

		l = l.next
		if l == nil {
			return 0, errors.New("out of range")
		}
		i++
	}
}

// insert inserts a new value at specific position
func (l *linkedList) insert(poi int, v int) error {
	temp := &linkedList{
		data: v,
		next: nil,
	}

	i := 0
	for {
		if i == poi {
			temp.next = l.next
			l.next = temp
			break
		}

		l = l.next
		if l == nil {
			return errors.New("out of range")
		}
		i++
	}

	return nil
}

// remove removes item at specific position
func (l *linkedList) remove(poi int) error {
	i := -1
	for {
		if i+1 == poi {
			l.next = l.next.next
			return nil
		}

		l = l.next
		if l == nil || l.next == nil {
			return errors.New("out of range")
		}
		i++
	}
}

func Test() {
	fmt.Println("------linked list------")
	l := new()
	l.print()
	fmt.Println("insert 10 at 0")
	l.insert(0, 10)
	l.print()
	fmt.Println("insert 9 at 0")
	l.insert(0, 9)
	l.print()
	fmt.Println("get at 1")
	v, err := l.get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("l[%v]: %v\n", 1, v)
	}
	fmt.Println("insert 8 at 0")
	l.insert(0, 8)
	l.print()
	fmt.Println("insert 1 at 1")
	l.insert(1, 1)
	l.print()
	fmt.Println("insert 2 at 4")
	l.insert(4, 2)
	l.print()
	fmt.Println("insert 18 at 6")
	err = l.insert(6, 18)
	if err != nil {
		fmt.Println(err)
	} else {
		l.print()
	}
	fmt.Println("remove at 0")
	l.remove(0)
	l.print()
	fmt.Println("remove at 3")
	l.remove(3)
	l.print()
	fmt.Println("remove at 1")
	l.remove(1)
	l.print()
	fmt.Println("remove at 2")
	err = l.remove(2)
	if err != nil {
		fmt.Println(err)
	} else {
		l.print()
	}
	fmt.Println("------linked list end------")
}
