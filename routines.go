package hangoutjson

import (
	"encoding/json"
)

func Do(bytestream []byte) (Hangouts, int, error) {
	var jsontree Hangouts
	err := json.Unmarshal(bytestream, &jsontree)
	return jsontree, len(jsontree.ConversationState), err
}
