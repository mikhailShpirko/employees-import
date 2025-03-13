package common

import (
	"strconv"
	"strings"
	"time"
)

func IsStringEmptyOrWhiteSpace(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func IsTimeBefore(src time.Time, future time.Time) bool {
	return src.Compare(future) == -1
}

func IsStringNumeric(str string) bool {
	_, err := strconv.ParseInt(str, 10, 64)
	return err == nil
}
