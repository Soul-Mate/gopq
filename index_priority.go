package gopq

type IndexPriorityQueue struct {
	n    int
	p    []int
	rp   []int
	data []Element
}

func NewIndexPriorityQueue(n int) *IndexPriorityQueue {
	ipq := new(IndexPriorityQueue)
	ipq.data = make([]Element, n+1)
	ipq.p = make([]int, n+1)
	ipq.rp = make([]int, n+1)
	return ipq
}

func (ipq *IndexPriorityQueue) Insert(key int, value Element) {
	// index out of range
	if key <= 0 || key > len(ipq.rp) {
		return
	}
	ipq.data[key] = value
	ipq.n++
	ipq.p[ipq.n] = key
	ipq.rp[key] = ipq.n
	ipq.swim(ipq.n)
}

func (ipq *IndexPriorityQueue) DelMax() Element {
	if ipq.n < 1 {
		return nil
	}
	key := ipq.p[1]
	max := ipq.data[key]
	ipq.exchange(1, ipq.n)
	// del element
	ipq.p = ipq.p[:ipq.n]
	ipq.data[key] = nil
	ipq.rp[key] = -1
	ipq.n--
	ipq.sink(1, ipq.n)
	return max
}

func (ipq *IndexPriorityQueue) sink(parent, length int) {
	for ; 2*parent <= length; {
		child := 2 * parent
		if child+1 <= length && ipq.less(child, child+1) {
			child++
		}
		if ipq.less(child, parent) {
			break
		}
		ipq.exchange(child, parent)
		parent = child
	}
}

func (ipq *IndexPriorityQueue) swim(length int) {
	for ; length/2 >= 1; length /= 2 {
		// parent < child
		if ipq.less(length/2, length) {
			ipq.exchange(length/2, length)
		}
	}
}

func (ipq *IndexPriorityQueue) less(i, j int) bool {
	v1 := ipq.data[ipq.p[i]]
	v2 := ipq.data[ipq.p[j]]
	if v1.CompareTo(v2) == -1 {
		return true
	}
	return false
}

func (ipq *IndexPriorityQueue) exchange(i, j int) {
	ipq.p[i], ipq.p[j] = ipq.p[j], ipq.p[i]
}
