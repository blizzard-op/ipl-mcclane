package brackets

import (
	"encoding/json"
)

func Format(bracket *Bracket) ([]byte, error) {
	result, err := json.Marshal(bracket)
	if err != nil {
		return make([]byte, 0), err
	}
	return result, nil
}
