package gopq

import (
	"sync"
)

type MaxQueue struct {
	Name string
	mu sync.RWMutex
	n    int
	data []Element
}

func NewPriorityQueue(name string) *MaxQueue {
	pq := new(MaxQueue)
	pq.Name = name
	pq.data = append(pq.data, nil)
	return pq
}

func (pq *MaxQueue) Insert(v Element) {
	pq.mu.Lock()
	pq.data = append(pq.data, v)
	pq.n++
	pq.swim(pq.n)
	pq.mu.Unlock()
}

func (pq *MaxQueue) Max() Element {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.data[1]
}

func (pq *MaxQueue) DelMax() Element {
	if pq.n == 0 {
		return nil
	}
	pq.mu.Lock()
	max := pq.data[1]
	pq.exchange(1, pq.n)
	pq.data = pq.data[:pq.n]
	pq.n--
	pq.sink(1, pq.n)
	pq.mu.Unlock()
	return max
}

func (pq *MaxQueue) Sort() {
	for pq.n > 1 {
		// exchange 1 pq.n
		pq.exchange(1, pq.n)
		// pq.n--
		pq.n--
		// adjust 1 pq.n
		pq.sink(1, pq.n)
	}
}

func (pq *MaxQueue) sink(parent, length int) {
	for 2*parent <= length {
		child := 2 * parent
		// search max value with child
		if child+1 <= pq.n && pq.less(child, child+1) {
			child++
		}
		if pq.less(child, parent) {
			break
		}
		pq.exchange(parent, child)
		parent = child
	}
}

func (pq *MaxQueue) swim(n int) {
	for ; n/2 >= 1; n /= 2 {
		if pq.less(n/2, n) {
			pq.exchange(n/2, n)
		}
	}
}

func (pq *MaxQueue) less(i, j int) bool {
	if pq.data[i].CompareTo(pq.data[j]) == -1 {
		return true
	}
	return false
}

func (pq *MaxQueue) exchange(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}
