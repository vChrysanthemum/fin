package ui

func growByteSlice(s []byte, desiredCap int) []byte {
	if cap(s) < desiredCap {
		ns := make([]byte, len(s), desiredCap)
		copy(ns, s)
		return ns
	}
	return s
}

func insertBytes(s []byte, offset int, data []byte) []byte {
	n := len(s) + len(data)
	s = growByteSlice(s, n)
	s = s[:n]
	copy(s[offset+len(data):], s[offset:])
	copy(s[offset:], data)
	return s
}

func copyByteSlice(dst, src []byte) []byte {
	if cap(dst) < len(src) {
		dst = cloneByteSlice(src)
	}
	dst = dst[:len(src)]
	copy(dst, src)
	return dst
}

func cloneByteSlice(s []byte) []byte {
	c := make([]byte, len(s))
	copy(c, s)
	return c
}
