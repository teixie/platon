package types

import (
	"encoding/json"
)

type String string

func (m String) JSONToArrayString() []string {
	var res []string
	_ = json.Unmarshal([]byte(m), &res)
	return res
}

func (m String) JSONToMapStringString() map[string]string {
	res := make(map[string]string)
	_ = json.Unmarshal([]byte(m), &res)
	return res
}
