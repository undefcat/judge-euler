package gen

import "fmt"

type heapPermutation struct {
	list []int
	callback func (list []int)
}

func NewHeapPermutation(list []int) *heapPermutation {
	return &heapPermutation{list, defaultCallback}
}

func defaultCallback(list []int) {
	fmt.Println(list)
}

func (it *heapPermutation) Callback(cb func ([]int)) {
	it.callback = cb
}

func (it *heapPermutation) Gen() {
	it.gen(len(it.list))
}

func (it *heapPermutation) gen(size int) {
	if size == 1 {
		it.callback(it.list)
		return
	}

	for i := 0; i < size; i++ {
		dec := size-1
		it.gen(dec)

		if size%2 == 1 {
			it.list[0], it.list[dec] = it.list[dec], it.list[0]
		} else {
			it.list[i], it.list[dec] = it.list[dec], it.list[i]
		}
	}
}
