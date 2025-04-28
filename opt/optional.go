package opt

type Opt[T any] struct {
	value     T
	isPresent bool
}

func Of[T any](value T) Opt[T] {
	return Opt[T]{value, true}
}

func OfPtr[T any](value *T) Opt[T] {
	if value != nil {
		return Opt[T]{*value, true}
	}
	return Opt[T]{*new(T), false}
}

func From[T any](from func() *T) Opt[T] {
	if value := from(); value != nil {
		return Opt[T]{*value, true}
	}
	return Opt[T]{*new(T), false}
}

func Empty[T any]() Opt[T] {
	return Opt[T]{*new(T), false}
}

func (o Opt[T]) Predicate(predicate func(value T) bool) bool {
	if o.isPresent {
		return predicate(o.value)
	}
	return false
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

func (o *Opt[T]) EqualsTo(other Opt[T], equals func(first T, second T) bool) bool {
	if o.IsPresent() && other.IsPresent() {
		return equals(o.value, other.value)
	}
	return !o.isPresent && !other.isPresent
}
