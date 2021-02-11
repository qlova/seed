package dating

import "time"

type HolidayJSON struct {
	Date       string    `json:"date"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Rule       string    `json:"rule"`
	Weekday    string    `json:"_weekday"`
	Substitute bool      `json:"substitute,omitempty"`
}
