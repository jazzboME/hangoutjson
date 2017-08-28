package hangoutjson

import (
	"encoding/json"
	"time"
	"strconv"
)

func Do(bytestream []byte) (Hangouts, int, error) {
	var jsontree Hangouts
	err := json.Unmarshal(bytestream, &jsontree)
	return jsontree, len(jsontree.ConversationState), err
}

func ConvTimestamp(timestamp string) (time.Time) {
	secs, err := strconv.ParseInt(timestamp[0:10], 10, 64)
	if err != nil {
		secs = 0
	}
	msecs, err := strconv.ParseInt(timestamp[10:16], 10, 64)
	if err != nil {
		msecs = 0
	}
	return time.Unix(secs,msecs)
}
