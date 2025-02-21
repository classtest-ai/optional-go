package opt

type Opt[T any] struct {
	value     T
	isPresent bool
}

func Of[T any](value T) Opt[T] {
	return Opt[T]{value, true}
}

func Empty[T any]() Opt[T] {
	return Opt[T]{*new(T), false}
}

func (o *Opt[T]) Clear() {
	o.value = *new(T)
	o.isPresent = false
}

func (o *Opt[T]) Set(value T) {
	o.value = value
	o.isPresent = true
}

func (o Opt[T]) IsPresent() bool {
	return o.isPresent
}

func (o Opt[T]) Get() (T, bool) {
	return o.value, o.isPresent
}

func (o Opt[T]) OrElse(other T) T {
	if o.isPresent {
		return o.value
	}
	return other
}

func (o Opt[T]) OrElseGet(other func() T) T {
	if o.isPresent {
		return o.value
	}
	return other()
}

func (o Opt[T]) OrElsePanic() T {
	if o.isPresent {
		return o.value
	}
	panic("Opt is empty")
}

func (o *Opt[T]) Default(other T) {
	if !o.isPresent {
		o.value = other
		o.isPresent = true
	}
}
