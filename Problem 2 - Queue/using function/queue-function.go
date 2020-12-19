package main

import "fmt"

var testValues = []string{
	"lorem",
	"ipsum",
	"1",
	"2",
	"3",
	"jack",
	"jill",
	"felix",
	"donking",
}

func main() {
	//create queue
	size := 5
	q := New()

	fmt.Println("========== Test Push Queue ==========")
	for i := 0; i < len(testValues); i++ {
		//push a new item to queue
		q = Push(q, testValues[i], size)

		//show the current list items in queue
		fmt.Print("current: ")
		Keys(q)

		// validate item existence
		if !Contains(q, testValues[i]) {
			fmt.Printf("policy: newly inserted %v must be exists\n", testValues[i])
		}

		//validate the queue's size
		if i < 5 && len(q) != (i+1) {
			fmt.Printf("expected length %d but actual: %d\n", i+1, len(q))
		} else if i >= 5 && len(q) != 5 {
			fmt.Printf("expexted length: %d but actual: %d\n", size, len(q))
		}
	}

	fmt.Println("\n\n========== Test Pop Queue ==========")
	for len(q) > 0 {
		//show the current list items in queue
		fmt.Print("current: ")
		Keys(q)

		//pop the first queue's data
		var v string
		v, q = Pop(q)

		//validate the pop data
		expect := testValues[size-(len(q)+1)]
		if v != expect {
			fmt.Printf("expected %v but recevied %v", expect, v)
		}
	}
}

func New() []string {
	queue := make([]string, 0)
	return queue
}

func Push(queue []string, item string, size int) []string {
	if len(queue) < size { //check whether the queue's length is > size
		queue = append(queue, item)
		return queue
	}

	fmt.Printf("Cannot insert %v because the queue is just %v length\n", item, size)
	return queue
}

func Keys(queue []string) {
	for i := 0; i < len(queue); i++ {
		fmt.Print(queue[i], " ")
	}

	fmt.Println()
}

func Contains(queue []string, item string) bool {
	for i := 0; i < len(queue); i++ {
		if queue[i] == item { //if item founded
			return false
		}
	}
	return true
}

func Pop(queue []string) (string, []string) {
	x := queue[0]     //get the first data from queue
	queue = queue[1:] //renew the queue without the first data

	return x, queue
}
