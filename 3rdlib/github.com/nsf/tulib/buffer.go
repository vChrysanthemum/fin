package tulib

import "github.com/nsf/termbox-go"
import "unicode/utf8"

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)

type Buffer struct {
	Cells []termbox.Cell
	Rect
}

func NewBuffer(w, h int) Buffer {
	return Buffer{
		Cells: make([]termbox.Cell, w*h),
		Rect:  Rect{0, 0, w, h},
	}
}

func TermboxBuffer() Buffer {
	w, h := termbox.Size()
	return Buffer{
		Cells: termbox.CellBuffer(),
		Rect:  Rect{0, 0, w, h},
	}
}

// Fills an area which is an intersection between buffer and 'dest' with 'proto'.
func (this *Buffer) Fill(dst Rect, proto termbox.Cell) {
	this.unsafe_fill(this.Rect.Intersection(dst), proto)
}

// Sets a cell at specified position
func (this *Buffer) Set(x, y int, proto termbox.Cell) {
	if x < 0 || x >= this.Width {
		return
	}
	if y < 0 || y >= this.Height {
		return
	}
	off := this.Width*y + x
	this.Cells[off] = proto
}

// Gets a pointer to the cell at specified position or nil if it's out
//  of range.
func (this *Buffer) Get(x, y int) *termbox.Cell {
	if x < 0 || x >= this.Width {
		return nil
	}
	if y < 0 || y >= this.Height {
		return nil
	}
	off := this.Width*y + x
	return &this.Cells[off]
}

// Resizes the Buffer, buffer contents are invalid after the resize.
func (this *Buffer) Resize(nw, nh int) {
	this.Width = nw
	this.Height = nh

	nsize := nw * nh
	if nsize <= cap(this.Cells) {
		this.Cells = this.Cells[:nsize]
	} else {
		this.Cells = make([]termbox.Cell, nsize)
	}
}

func (this *Buffer) Blit(dstr Rect, srcx, srcy int, src *Buffer) {
	srcr := Rect{srcx, srcy, 0, 0}

	// first adjust 'srcr' if 'dstr' has negatives
	if dstr.X < 0 {
		srcr.X -= dstr.X
	}
	if dstr.Y < 0 {
		srcr.Y -= dstr.Y
	}

	// adjust 'dstr' against 'this.Rect', copy 'dstr' size to 'srcr'
	dstr = this.Rect.Intersection(dstr)
	srcr.Width = dstr.Width
	srcr.Height = dstr.Height

	// adjust 'srcr' against 'src.Rect', copy 'srcr' size to 'dstr'
	srcr = src.Rect.Intersection(srcr)
	dstr.Width = srcr.Width
	dstr.Height = srcr.Height

	if dstr.IsEmpty() {
		return
	}

	// blit!
	srcstride := src.Width
	dststride := this.Width
	linew := dstr.Width
	srcoff := src.Width*srcr.Y + srcr.X
	dstoff := this.Width*dstr.Y + dstr.X
	for i := 0; i < dstr.Height; i++ {
		linesrc := src.Cells[srcoff : srcoff+linew]
		linedst := this.Cells[dstoff : dstoff+linew]
		copy(linedst, linesrc)
		srcoff += srcstride
		dstoff += dststride
	}
}

// Unsafe part of the fill operation, doesn't check for bounds.
func (this *Buffer) unsafe_fill(dest Rect, proto termbox.Cell) {
	stride := this.Width
	off := this.Width*dest.Y + dest.X
	for y := 0; y < dest.Height; y++ {
		for x := 0; x < dest.Width; x++ {
			this.Cells[off+x] = proto
		}
		off += stride
	}
}

// draws from left to right, 'off' is the beginning position
// (DrawLabel uses that method)
func (this *Buffer) draw_n_first_runes(off, n int, params *LabelParams, text []byte) {
	for n > 0 {
		r, size := utf8.DecodeRune(text)
		this.Cells[off] = termbox.Cell{
			Ch: r,
			Fg: params.Fg,
			Bg: params.Bg,
		}
		text = text[size:]
		off++
		n--
	}
}

// draws from right to left, 'off' is the end position
// (DrawLabel uses that method)
func (this *Buffer) draw_n_last_runes(off, n int, params *LabelParams, text []byte) {
	for n > 0 {
		r, size := utf8.DecodeLastRune(text)
		this.Cells[off] = termbox.Cell{
			Ch: r,
			Fg: params.Fg,
			Bg: params.Bg,
		}
		text = text[:len(text)-size]
		off--
		n--
	}
}

type LabelParams struct {
	Fg             termbox.Attribute
	Bg             termbox.Attribute
	Align          Alignment
	Ellipsis       rune
	CenterEllipsis bool
}

var DefaultLabelParams = LabelParams{
	termbox.ColorDefault,
	termbox.ColorDefault,
	AlignLeft,
	'…',
	false,
}

func skip_n_runes(x []byte, n int) []byte {
	if n <= 0 {
		return x
	}

	for n > 0 {
		_, size := utf8.DecodeRune(x)
		x = x[size:]
		n--
	}
	return x
}

func (this *Buffer) DrawLabel(dest Rect, params *LabelParams, text []byte) {
	if dest.Height != 1 {
		dest.Height = 1
	}

	dest = this.Rect.Intersection(dest)
	if dest.Height == 0 || dest.Width == 0 {
		return
	}

	ellipsis := termbox.Cell{Ch: params.Ellipsis, Fg: params.Fg, Bg: params.Bg}
	off := dest.Y*this.Width + dest.X
	textlen := utf8.RuneCount(text)
	n := textlen
	if n > dest.Width {
		// string doesn't fit in the dest rectangle, draw ellipsis
		n = dest.Width - 1

		// if user asks for ellipsis in the center, alignment doesn't matter
		if params.CenterEllipsis {
			this.Cells[off+dest.Width/2] = ellipsis
		} else {
			switch params.Align {
			case AlignLeft:
				this.Cells[off+dest.Width-1] = ellipsis
			case AlignCenter:
				this.Cells[off] = ellipsis
				this.Cells[off+dest.Width-1] = ellipsis
				n--
			case AlignRight:
				this.Cells[off] = ellipsis
			}
		}
	}

	if n <= 0 {
		return
	}

	if params.CenterEllipsis && textlen != n {
		firsthalf := dest.Width / 2
		secondhalf := dest.Width - 1 - firsthalf
		this.draw_n_first_runes(off, firsthalf, params, text)
		off += dest.Width - 1
		this.draw_n_last_runes(off, secondhalf, params, text)
		return
	}

	switch params.Align {
	case AlignLeft:
		this.draw_n_first_runes(off, n, params, text)
	case AlignCenter:
		if textlen == n {
			off += (dest.Width - n) / 2
			this.draw_n_first_runes(off, n, params, text)
		} else {
			off++
			mid := (textlen - n) / 2
			text = skip_n_runes(text, mid)
			this.draw_n_first_runes(off, n, params, text)
		}
	case AlignRight:
		off += dest.Width - 1
		this.draw_n_last_runes(off, n, params, text)
	}
}
