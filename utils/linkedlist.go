package utils

type LLNode[T any] struct {
	Next, Prev *LLNode[T]
	Value      T
}

func (e *LLNode[T]) InsertNext(value T) *LLNode[T] {
	newEntry := &LLNode[T]{
		Prev:  e,
		Value: value,
	}
	if e != nil {
		newEntry.Next = e.Next
		if e.Next != nil {
			e.Next.Prev = newEntry
		}
		e.Next = newEntry
	}
	return newEntry
}

func (e *LLNode[T]) InsertPrev(value T) *LLNode[T] {
	newEntry := &LLNode[T]{
		Next:  e,
		Value: value,
	}
	if e != nil {
		newEntry.Prev = e.Prev
		if e.Prev != nil {
			e.Prev.Next = newEntry
		}
		e.Prev = newEntry
	}
	return newEntry
}

func (e *LLNode[T]) Pop() {
	if e != nil {
		if e.Next != nil {
			e.Next.Prev = e.Prev
		}
		if e.Prev != nil {
			e.Prev.Next = e.Next
		}
	}
}
