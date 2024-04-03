package main

import (
	"LeetCode-go/source/Go_LeetCode"
	"fmt"
)

func main() {
	box := Go_LeetCode.NewMessageBox()
	fmt.Println("这是2来到的时候")
	box.Receive(2, "B")
	fmt.Println("这是1来到的时候")
	box.Receive(1, "A")
	box.Receive(4, "D")
	box.Receive(5, "E")
	box.Receive(7, "G")
	box.Receive(8, "H")
	box.Receive(6, "F")
	box.Receive(3, "C")
	box.Receive(9, "I")
	box.Receive(10, "J")
	box.Receive(12, "L")
	box.Receive(13, "M")
	box.Receive(11, "K")
}
