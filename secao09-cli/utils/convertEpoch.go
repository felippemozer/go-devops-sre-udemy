package utils

import (
	"strconv"
	"time"
)

func ConvertEpoch(epoch string) time.Time {
	epochInteger, _ := strconv.ParseInt(epoch, 10, 64)
	timeResponse := time.Unix(epochInteger, 0)

	return timeResponse
}
