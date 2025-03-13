package custom_types

import (
	"fmt"
	"strings"
	"time"
)

type DateOnly struct {
	time.Time
}

const dateOnlyLayout = "2006-01-02"

func (c *DateOnly) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`) // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(dateOnlyLayout, s)
	return
}

func (c DateOnly) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(dateOnlyLayout))), nil
}
