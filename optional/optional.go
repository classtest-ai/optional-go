package optional

type Optional[T any] struct {
	value     T
	isPresent bool
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{value, true}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{*new(T), false}
}

func (o Optional[T]) IsPresent() bool {
	return o.isPresent
}

func (o Optional[T]) Get() (T, bool) {
	return o.value, o.isPresent
}

func (o Optional[T]) OrElse(other T) T {
	if o.isPresent {
		return o.value
	}
	return other
}

func (o Optional[T]) OrElseGet(other func() T) T {
	if o.isPresent {
		return o.value
	}
	return other()
}

func (o Optional[T]) OrElsePanic() T {
	if o.isPresent {
		return o.value
	}
	panic("Optional is empty")
}
