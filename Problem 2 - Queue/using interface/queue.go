package queue

type Queue interface {
	Push(key interface{
		queue = append(queue, key)
	})

	Pop() interface{
		x = queue[0]
		queue = queue[1:]
	}

	Contains(key interface{
		for i:= 0; i<size; i++ {
			if queue[i] == key {
				return false
			}
		}
		return true
	}) bool

	Len() int {
		return size
	}

	Keys() []interface{
		for i := 0; i<size; i++ {
			fmt.Print(queue[i],)
		}
	}
}

func New(size int) Queue {
	queue := make([]int, size)
	return queue
}
