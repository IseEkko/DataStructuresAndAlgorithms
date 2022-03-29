package main

import "fmt"

/**
队列：
基本操作是入队(Enqueue)，在表的末端插入一个元素
出队(Dequeue)，删除(或返回)在表头的元素
*/

type Item interface{}

// 队列结构体
type Queue struct {
	Items []Item
}

type IQueue interface {
	New() Queue
	Enqueue(t Item) //入队
	Dequeue(t Item) //出队
	IsEmpty() bool  //是否为空
	Size() int      //大小
}

// 新建
func (q *Queue) New() *Queue {
	q.Items = []Item{}
	return q
}

// 入队
func (q *Queue) Enqueue(data Item) {
	q.Items = append(q.Items, data)
}

// 出队
func (q *Queue) Dequeue() *Item {
	// 由于是先进先出，0为队首
	item := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return &item
}

// 队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

// 队列长度
func (q *Queue) Size() int {
	return len(q.Items)
}

var q Queue

func initQueue() *Queue {
	if q.Items == nil {
		q = Queue{}
		q.New()
	}
	return &q
}

func main() {
	q := initQueue()
	fmt.Println("length:", q.Size())
	q.Enqueue("a")
	q.Enqueue(1)
	fmt.Println("length:", q.Size())
	fmt.Println(q)
	tmp := q.Dequeue()
	fmt.Println("出队：", *tmp)
}
