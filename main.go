package main

func main() {
	l1 := NewList()

	l1.PushFront(10) // [10]
	l1.PushFront(20) // [20 10]
	l1.PushFront(30) // [30 20 10]

	l1.Print() // Вывод должен быть 30 20 10

	l2 := NewList()

	l2.PushBack(10) // [10]
	l2.PushBack(20) // [10 20]
	l2.PushBack(30) // [10 20 30]

	l2.Print() // Вывод должен быть 10 20 30

}
