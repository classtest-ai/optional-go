package optional

import (
	"encoding/json"
)

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return []byte("null"), nil
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = *new(T)
		return nil
	}
	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	o.value = value
	o.isPresent = true
	return nil
}
