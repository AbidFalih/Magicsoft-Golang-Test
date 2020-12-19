package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type QQ struct {
	data []interface{}
	size int
}

func New(size int) QQ {
	q := QQ{make([]interface{}, 0), size}
	return q
}

func (q *QQ) Push(key interface{}) {
	// if len(q.data) < q.size {
	// 	q.data = append(q.data, key)
	// }

	q.data = append(q.data, key)
}

func (q *QQ) Pop() interface{} {
	x := q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *QQ) Keys() []interface{} {
	return q.data
}

func (q *QQ) Contains(key interface{}) bool {
	for _, v := range q.data {
		if v == key {
			return true
		}
	}
	return false
}

func (q *QQ) Len() int {
	return len(q.data)
}
