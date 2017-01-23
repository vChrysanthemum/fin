package terminal

func grow_byte_slice(s []byte, desired_cap int) []byte {
	if cap(s) < desired_cap {
		ns := make([]byte, len(s), desired_cap)
		copy(ns, s)
		return ns
	}
	return s
}

func insert_bytes(s []byte, offset int, data []byte) []byte {
	n := len(s) + len(data)
	s = grow_byte_slice(s, n)
	s = s[:n]
	copy(s[offset+len(data):], s[offset:])
	copy(s[offset:], data)
	return s
}

func copy_byte_slice(dst, src []byte) []byte {
	if cap(dst) < len(src) {
		dst = clone_byte_slice(src)
	}
	dst = dst[:len(src)]
	copy(dst, src)
	return dst
}

func clone_byte_slice(s []byte) []byte {
	c := make([]byte, len(s))
	copy(c, s)
	return c
}
