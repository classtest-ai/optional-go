package opt

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func FromPGText(n pgtype.Text) Opt[string] {
	if n.Valid {
		return Of(n.String)
	}
	return Empty[string]()
}

func FromPGInt64(n pgtype.Int8) Opt[int64] {
	if n.Valid {
		return Of(n.Int64)
	}
	return Empty[int64]()
}

func FromPGInt32(n pgtype.Int4) Opt[int32] {
	if n.Valid {
		return Of(n.Int32)
	}
	return Empty[int32]()
}

func FromPGInt(n pgtype.Int4) Opt[int] {
	if n.Valid {
		return Of(int(n.Int32))
	}
	return Empty[int]()
}

func FromPGInt16(n pgtype.Int2) Opt[int16] {
	if n.Valid {
		return Of(n.Int16)
	}
	return Empty[int16]()
}

func FromPGFloat64(n pgtype.Float8) Opt[float64] {
	if n.Valid {
		return Of(n.Float64)
	}
	return Empty[float64]()
}

func FromPGFloat32(n pgtype.Float4) Opt[float32] {
	if n.Valid {
		return Of(n.Float32)
	}
	return Empty[float32]()
}

func FromPGBool(n pgtype.Bool) Opt[bool] {
	if n.Valid {
		return Of(n.Bool)
	}
	return Empty[bool]()
}

func FromPGTime(n pgtype.Timestamp) Opt[time.Time] {
	if n.Valid {
		return Of(n.Time)
	}
	return Empty[time.Time]()
}

func ToPGString(o Opt[string]) pgtype.Text {
	if o.IsPresent() {
		return pgtype.Text{String: o.value, Valid: true}
	}
	return pgtype.Text{}
}

func ToPGInt64(o Opt[int64]) pgtype.Int8 {
	if o.IsPresent() {
		return pgtype.Int8{Int64: o.value, Valid: true}
	}
	return pgtype.Int8{}
}

func ToPGInt32(o Opt[int32]) pgtype.Int4 {
	if o.IsPresent() {
		return pgtype.Int4{Int32: o.value, Valid: true}
	}
	return pgtype.Int4{}
}

func ToPGInt(o Opt[int]) pgtype.Int4 {
	if o.IsPresent() {
		return pgtype.Int4{Int32: int32(o.value), Valid: true}
	}
	return pgtype.Int4{}
}

func ToPGInt16(o Opt[int16]) pgtype.Int2 {
	if o.IsPresent() {
		return pgtype.Int2{Int16: o.value, Valid: true}
	}
	return pgtype.Int2{}
}

func ToPGFloat64(o Opt[float64]) pgtype.Float8 {
	if o.IsPresent() {
		return pgtype.Float8{Float64: o.value, Valid: true}
	}
	return pgtype.Float8{}
}

func ToPGFloat32(o Opt[float32]) pgtype.Float4 {
	if o.IsPresent() {
		return pgtype.Float4{Float32: o.value, Valid: true}
	}
	return pgtype.Float4{}
}

func ToPGBool(o Opt[bool]) pgtype.Bool {
	if o.IsPresent() {
		return pgtype.Bool{Bool: o.value, Valid: true}
	}
	return pgtype.Bool{}
}

func ToPGTime(o Opt[time.Time]) pgtype.Timestamp {
	if o.IsPresent() {
		return pgtype.Timestamp{Time: o.value, Valid: true}
	}
	return pgtype.Timestamp{}
}
