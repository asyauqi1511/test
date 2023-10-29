package pkg

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	t time.Time
}

func NewDateTime(t time.Time) DateTime {
	return DateTime{t}
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	tt, err := time.Parse(`"2006-01-02"`, string(data))
	*d = DateTime{tt}
	return err
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.t.Format("2006-01-02"))
}

func (d DateTime) IsZero() bool {
	return d.t.IsZero()
}

func (d DateTime) Value() time.Time {
	return d.t
}
