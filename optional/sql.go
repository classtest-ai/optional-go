package optional

import "database/sql"

func (o Optional[T]) ToSQLNull() sql.Null[T] {
	if o.IsPresent() {
		return sql.Null[T]{V: o.value, Valid: true}
	}
	return sql.Null[T]{}
}

func FromSQLNull[T any](n sql.Null[T]) Optional[T] {
	if n.Valid {
		return Of(n.V)
	}
	return Empty[T]()
}
