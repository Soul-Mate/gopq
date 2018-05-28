package gopq

import "sync"

type MinPQ struct {
	mu       sync.RWMutex
	elements []Element
	n        int
	Name     string
}

func (pq *MinPQ) Insert(el Element) {
	pq.mu.Lock()
	pq.elements = append(pq.elements, el)
	pq.n++
	pq.swim(pq.n)
	pq.mu.Unlock()
}

func (pq *MinPQ) DelMin() Element {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	min := pq.elements[1]
	pq.exchange(1, pq.n)
	pq.elements = pq.elements[:pq.n]
	pq.n--
	pq.sink(1, pq.n)
	return min
}

func (pq *MinPQ) Min() Element {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.elements[1]
}

func (pq *MinPQ) swim(n int) {
	for n/2 >= 1 {
		if pq.less(n, n/2) {
			pq.exchange(n, n/2)
		}
		n /= 2
	}
}

func (pq *MinPQ) sink(p, n int) {
	for ; 2*p <= n; {
		child := 2 * p
		if child+1 <= n && pq.less(child+1, child) {
			child++
		}
		if pq.less(p, child) {
			break
		}
		pq.exchange(p, child)
		p = child
	}
}

func (pq *MinPQ) less(i, j int) bool {
	if pq.elements[i].CompareTo(pq.elements[j]) == -1 {
		return true
	}
	return false
}

func (pq *MinPQ) exchange(i, j int) {
	pq.elements[i], pq.elements[j] = pq.elements[j], pq.elements[i]
}
