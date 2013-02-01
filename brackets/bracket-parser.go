package brackets

import (
	"encoding/json"
)

func Parse(data []byte) (*Bracket, error) {
	var result *Bracket

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
