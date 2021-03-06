// (C) 2017 by Ricardo Branco
//
// MIT License

package bitset

import "fmt"
import "runtime"

type word uint64

const (
	bitsPerWord = 64
)

type Bitset struct {
	words map[int]word
	count int // Number of elements inserted
	max   int // Keep track of the maximum number ever inserted
}

// Returns a new Bitset. Set argument to the expected maximum number or -1.
func New(max int) *Bitset {
	bs := new(Bitset)
	bs.max = max
	bs.words = make(map[int]word)
	if max >= 0 {
		for i := max / bitsPerWord; i >= 0; i-- {
			bs.words[i] = 0
		}
	}
	runtime.SetFinalizer(bs, destroy)
	return bs
}

func (bs *Bitset) Add(i int) {
	if i < 0 {
		panic(fmt.Sprintf("Trying to Add a negative number: %d\n", i))
	}
	bs.words[i/bitsPerWord] |= 1 << uint(i%bitsPerWord)
	bs.count++
	if i > bs.max {
		bs.max = i
	}
}

func (bs *Bitset) Del(i int) {
	if i < 0 {
		panic(fmt.Sprintf("Trying to Del a negative number: %d\n", i))
	}
	bs.words[i/bitsPerWord] &= ^(1 << uint(i%bitsPerWord))
	bs.count--
}

func (bs *Bitset) Test(i int) bool {
	if i < 0 {
		panic(fmt.Sprintf("Trying to Test a negative number: %d\n", i))
	}
	return (bs.words[i/bitsPerWord] & (1 << uint(i%bitsPerWord))) != 0
}

func (bs *Bitset) SetAll() {
	for i := range bs.words {
		bs.words[i] = ^word(0)
	}
	bs.count = len(bs.words) * bitsPerWord
	if bs.max >= 0 {
		bs.count = bs.max + 1
	}
}

func (bs *Bitset) ClearAll() {
	for i := range bs.words {
		bs.words[i] = 0
	}
	bs.count = 0
}

func (bs *Bitset) GetCount() int {
	return bs.count
}

// Returns a slice of all numbers in the set
func (bs *Bitset) GetAll() (s []int) {
	if bs.count == 0 {
		return
	}
	s = make([]int, 0, bs.count)
	for i, w := range bs.words {
		for {
			if w == 0 {
				break
			}
			bit := ffs(w)
			num := i*bitsPerWord + int(bit)
			if num > bs.max {
				return
			}
			s = append(s, num)
			w &= ^(1 << bit)
		}
	}
	return
}

func destroy(bs *Bitset) {
	for i := range bs.words {
		delete(bs.words, i)
	}
	bs.count = 0
	bs.max = -1
}

func (bs *Bitset) Destroy() {
	destroy(bs)
}
