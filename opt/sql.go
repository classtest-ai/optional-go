package opt

import "database/sql"

func (o Opt[T]) ToSQLNull() sql.Null[T] {
	if o.IsPresent() {
		return sql.Null[T]{V: o.value, Valid: true}
	}
	return sql.Null[T]{}
}

func FromSQLNull[T any](n sql.Null[T]) Opt[T] {
	if n.Valid {
		return Of(n.V)
	}
	return Empty[T]()
}
