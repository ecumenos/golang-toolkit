package list

// Element ...
type Element[T any] struct {
	Prev  *Element[T]
	Value T
	Next  *Element[T]
}

// FIFOUniqueList ...
type FIFOUniqueList[T any] struct {
	FirstElem   *Element[T]
	LastElem    *Element[T]
	Size        int
	MaxSize     int
	CompareFunc func(T, T) bool
}

// NewFIFOUniqueList ...
func NewFIFOUniqueList[T any](maxSize int, compareFnc func(T, T) bool, elements ...T) *FIFOUniqueList[T] {
	out := &FIFOUniqueList[T]{
		MaxSize:     maxSize,
		CompareFunc: compareFnc,
	}
	for i := 0; i < len(elements); i++ {
		out.Push(elements[i])
	}

	return out
}

// Push ...
func (l *FIFOUniqueList[T]) Push(in T) {
	if l.FirstElem == nil {
		el := &Element[T]{
			Value: in,
		}
		l.FirstElem = el
		l.LastElem = el
		l.Size++

		return
	}
	l.Remove(in)

	el := &Element[T]{
		Value: in,
		Next:  l.FirstElem,
	}
	l.FirstElem.Prev = el
	l.FirstElem = el
	l.Size++

	if l.MaxSize < l.Size {
		l.Pop()
	}
}

// Remove ...
func (l *FIFOUniqueList[T]) Remove(in T) bool {
	for el := l.FirstElem; el != nil; el = el.Next {
		if l.CompareFunc(el.Value, in) {
			if el.Prev != nil && el.Next != nil {
				el.Next.Prev = el.Prev
				el.Prev.Next = el.Next
				l.Size--

				return true
			}
			if el.Prev == nil && el.Next != nil {
				el.Next.Prev = nil
				l.FirstElem = el.Next
				l.Size--

				return true
			}
			if el.Prev != nil && el.Next == nil {
				el.Prev.Next = nil
				l.LastElem = el.Prev
				l.Size--

				return true
			}
		}
	}

	return false
}

// RemoveOne ...
func (l *FIFOUniqueList[T]) RemoveOne(cb func(prev, curr, next T) bool) bool {
	for el := l.FirstElem; el != nil; el = el.Next {
		var prev, next T
		if el.Prev != nil {
			prev = el.Prev.Value
		}
		if el.Next != nil {
			next = el.Next.Value
		}
		if cb(prev, el.Value, next) {
			if el.Prev != nil && el.Next != nil {
				el.Next.Prev = el.Prev
				el.Prev.Next = el.Next
				l.Size--

				return true
			}
			if el.Prev == nil && el.Next != nil {
				el.Next.Prev = nil
				l.FirstElem = el.Next
				l.Size--

				return true
			}
			if el.Prev != nil && el.Next == nil {
				el.Prev.Next = nil
				l.LastElem = el.Prev
				l.Size--

				return true
			}
		}
	}

	return false
}

// RemoveMany ...
func (l *FIFOUniqueList[T]) RemoveMany(cb func(prev, curr, next T) bool) bool {
	var out bool
	for el := l.FirstElem; el != nil; el = el.Next {
		var prev, next T
		if el.Prev != nil {
			prev = el.Prev.Value
		}
		if el.Next != nil {
			next = el.Next.Value
		}
		if cb(prev, el.Value, next) {
			if el.Prev != nil && el.Next != nil {
				el.Next.Prev = el.Prev
				el.Prev.Next = el.Next
				l.Size--
				out = true
				continue
			}
			if el.Prev == nil && el.Next != nil {
				el.Next.Prev = nil
				l.FirstElem = el.Next
				l.Size--
				out = true
				continue
			}
			if el.Prev != nil && el.Next == nil {
				el.Prev.Next = nil
				l.LastElem = el.Prev
				l.Size--
				out = true
				continue
			}
		}
	}

	return out
}

// Pop ...
func (l *FIFOUniqueList[T]) Pop() {
	if l.Size == 0 {
		return
	}
	if l.Size == 1 {
		l.FirstElem = nil
		l.LastElem = nil
		l.Size = 0

		return
	}
	l.LastElem = l.LastElem.Prev
	l.LastElem.Next = nil
	l.Size--
}

// FindOne ...
func (l *FIFOUniqueList[T]) FindOne(cb func(prev, curr, next T) bool) (T, bool) {
	for el := l.FirstElem; el != nil; el = el.Next {
		var prev, next T
		if el.Prev != nil {
			prev = el.Prev.Value
		}
		if el.Next != nil {
			next = el.Next.Value
		}
		if cb(prev, el.Value, next) {
			return el.Value, true
		}
	}
	var zero T

	return zero, false
}

// FindAll ...
func (l *FIFOUniqueList[T]) FindAll() []T {
	out := make([]T, 0, l.Size)
	for el := l.FirstElem; el != nil; el = el.Next {
		out = append(out, el.Value)
	}

	return out
}

// FindMany ...
func (l *FIFOUniqueList[T]) FindMany(cb func(prev, curr, next T) bool) []T {
	out := make([]T, 0, l.Size)
	for el := l.FirstElem; el != nil; el = el.Next {
		var prev, next T
		if el.Prev != nil {
			prev = el.Prev.Value
		}
		if el.Next != nil {
			next = el.Next.Value
		}
		if cb(prev, el.Value, next) {
			out = append(out, el.Value)
		}
	}

	return out
}
