package main

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
)

var (
	uniqueID int
)

type PooledObject struct {
	id int
}

func (o *PooledObject) String() string {
	return "PooledObject [id=" + strconv.Itoa(o.id) + "]"
}

type ObjectPool struct {
	mu     sync.Mutex
	idle   *list.List
	active *list.List
}

//
func NewObjectPool() *ObjectPool {
	idle := list.New()
	active := list.New()
	return &ObjectPool{idle: idle, active: active}
}

func (p *ObjectPool) BorrowObject() *PooledObject {
	p.mu.Lock()
	defer p.mu.Unlock()

	var object *PooledObject
	if p.idle.Len() <= 0 {
		object = &PooledObject{uniqueID}
		uniqueID++
	} else {
		object = p.removeAt(p.idle, 0)
	}
	fmt.Printf("Borrow: %s\n", object)
	p.active.PushBack(object)
	return object
}

func (p *ObjectPool) ReturnObject(object *PooledObject) {
	p.mu.Lock()
	defer p.mu.Unlock()

	fmt.Printf("Return: %s\n", object)
	p.idle.PushBack(object)
	p.remove(p.active, object)
}

func (p *ObjectPool) remove(list *list.List, object *PooledObject) {
	for e := list.Front(); e != nil; e = e.Next() {
		if object == e.Value.(*PooledObject) {
			list.Remove(e)
			return
		}
	}
}

func (p *ObjectPool) removeAt(list *list.List, index int) *PooledObject {
	for e, i := list.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		if index == i {
			return list.Remove(e).(*PooledObject)
		}
	}
	return nil
}

func main() {
	objectPool := NewObjectPool()
	object1 := objectPool.BorrowObject()
	objectPool.ReturnObject(object1)
	object2 := objectPool.BorrowObject()
	object3 := objectPool.BorrowObject()
	objectPool.ReturnObject(object2)
	objectPool.ReturnObject(object3)
	object4 := objectPool.BorrowObject()
	object5 := objectPool.BorrowObject()
	object6 := objectPool.BorrowObject()
	objectPool.ReturnObject(object4)
	objectPool.ReturnObject(object5)
	objectPool.ReturnObject(object6)
}

//pool object : a client access to an object poll, can avoid creating a new object by simplely asking pool object have already had the object instead.
// the pool will be a growing pool, we creat new object if pool empty. and we can restrict pool.
//when we borrow one object from pool , nobody can use the object untill it returns pool.
