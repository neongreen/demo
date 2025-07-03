package demo2

import (
    "encoding/json"
    "time"
)

// Event represents a named event with a timestamp.
type Event struct {
    Name string    `json:"name"`
    Time time.Time `json:"time"`
}

// MarshalJSON ensures Time is formatted RFC3339 in UTC.
func (e Event) MarshalJSON() ([]byte, error) {
    type alias struct {
        Name string `json:"name"`
        Time string `json:"time"`
    }
    a := alias{
        Name: e.Name,
        Time: e.Time.UTC().Format(time.RFC3339),
    }
    return json.Marshal(a)
}

// UnmarshalJSON parses RFC3339 time into Time.
func (e *Event) UnmarshalJSON(data []byte) error {
    type alias struct {
        Name string `json:"name"`
        Time string `json:"time"`
    }
    var a alias
    if err := json.Unmarshal(data, &a); err != nil {
        return err
    }
    t, err := time.Parse(time.RFC3339, a.Time)
    if err != nil {
        return err
    }
    e.Name = a.Name
    e.Time = t.UTC()
    return nil
}

