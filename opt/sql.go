package opt

import (
	"database/sql"
	"time"
)

func FromSQLNull[T any](n sql.Null[T]) Opt[T] {
	if n.Valid {
		return Of(n.V)
	}
	return Empty[T]()
}

func FromSQLNullString(n sql.NullString) Opt[string] {
	if n.Valid {
		return Of(n.String)
	}
	return Empty[string]()
}

func FromSQLNullInt64(n sql.NullInt64) Opt[int64] {
	if n.Valid {
		return Of(n.Int64)
	}
	return Empty[int64]()
}

func FromSQLNullInt32(n sql.NullInt32) Opt[int32] {
	if n.Valid {
		return Of(n.Int32)
	}
	return Empty[int32]()
}

func FromSQLNullInt16(n sql.NullInt16) Opt[int16] {
	if n.Valid {
		return Of(n.Int16)
	}
	return Empty[int16]()
}

func FromSQLNullByte(n sql.NullByte) Opt[byte] {
	if n.Valid {
		return Of(n.Byte)
	}
	return Empty[byte]()
}

func FromSQLNullFloat64(n sql.NullFloat64) Opt[float64] {
	if n.Valid {
		return Of(n.Float64)
	}
	return Empty[float64]()
}

func FromSQLNullBool(n sql.NullBool) Opt[bool] {
	if n.Valid {
		return Of(n.Bool)
	}
	return Empty[bool]()
}

func FromSQLNullTime(n sql.NullTime) Opt[time.Time] {
	if n.Valid {
		return Of(n.Time)
	}
	return Empty[time.Time]()
}

func ToSQLNull[T any](o Opt[T]) sql.Null[T] {
	if o.IsPresent() {
		return sql.Null[T]{V: o.value, Valid: true}
	}
	return sql.Null[T]{}
}

func ToSQLNullString(o Opt[string]) sql.NullString {
	if o.IsPresent() {
		return sql.NullString{String: o.value, Valid: true}
	}
	return sql.NullString{}
}

func ToSQLNullInt64(o Opt[int64]) sql.NullInt64 {
	if o.IsPresent() {
		return sql.NullInt64{Int64: o.value, Valid: true}
	}
	return sql.NullInt64{}
}

func ToSQLNullInt32(o Opt[int32]) sql.NullInt32 {
	if o.IsPresent() {
		return sql.NullInt32{Int32: o.value, Valid: true}
	}
	return sql.NullInt32{}
}

func ToSQLNullInt16(o Opt[int16]) sql.NullInt16 {
	if o.IsPresent() {
		return sql.NullInt16{Int16: o.value, Valid: true}
	}
	return sql.NullInt16{}
}

func ToSQLNullByte(o Opt[byte]) sql.NullByte {
	if o.IsPresent() {
		return sql.NullByte{Byte: o.value, Valid: true}
	}
	return sql.NullByte{}
}

func ToSQLNullFloat64(o Opt[float64]) sql.NullFloat64 {
	if o.IsPresent() {
		return sql.NullFloat64{Float64: o.value, Valid: true}
	}
	return sql.NullFloat64{}
}

func ToSQLNullBool(o Opt[bool]) sql.NullBool {
	if o.IsPresent() {
		return sql.NullBool{Bool: o.value, Valid: true}
	}
	return sql.NullBool{}
}

func ToSQLNullTime(o Opt[time.Time]) sql.NullTime {
	if o.IsPresent() {
		return sql.NullTime{Time: o.value, Valid: true}
	}
	return sql.NullTime{}
}
