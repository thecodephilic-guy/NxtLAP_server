package structs

import "time"

type Session struct {
	MeetingKey       int       `json:"meeting_key"`
	SessionKey       int       `json:"session_key"`
	Location         string    `json:"location"`
	DateStart        time.Time `json:"date_start"`
	DateEnd          time.Time `json:"date_end"`
	SessionType      string    `json:"session_type"`
	SessionName      string    `json:"session_name"`
	CountryKey       int       `json:"country_key"`
	CountryCode      string    `json:"country_code"`
	CountryName      string    `json:"country_name"`
	CircuitKey       int       `json:"circuit_key"`
	CircuitShortName string    `json:"circuit_short_name"`
	GmtOffset        string    `json:"gmt_offset"` // Stored as string to keep "03:00:00" format
	Year             int       `json:"year"`
}
