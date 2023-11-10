package frontrpc

import (
	"encoding/json"
	"fmt"
)

type RequestId json.Number

func (id RequestId) String() string {
	return string(id)
}

func (id RequestId) MarshalJSON() ([]byte, error) {
	if id == "" {
		return []byte("0"), nil
	}
	out := fmt.Sprintf("%s", id)
	return []byte(out), nil
}

func (id *RequestId) UnmarshalJSON(data []byte) error {
	var dc []byte
	fc := true
	for _, b := range data {
		if b < 58 && b > 47 {
			if fc && b != 48 {
				fc = false
			}
			if !fc {
				dc = append(dc, b)
			}
		}
	}
	if len(dc) == 0 {
		*id = RequestId("0")
		return nil
	}
	*id = RequestId(dc)
	return nil
}
