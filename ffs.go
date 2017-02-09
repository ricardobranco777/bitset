package bitset

func ffs(w word) (bit uint) {
	if w&0xffffffff == 0 {
		bit += 32
		w >>= 32
	}
	if w&0xffff == 0 {
		bit += 16
		w >>= 16
	}
	if w&0xff == 0 {
		bit += 8
		w >>= 8
	}
	if w&0xf == 0 {
		bit += 4
		w >>= 4
	}
	if w&0x3 == 0 {
		bit += 2
		w >>= 2
	}
	if w&0x1 == 0 {
		bit++
	}
	return
}
