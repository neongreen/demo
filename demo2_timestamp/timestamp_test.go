package demo2

import (
    "encoding/json"
    "testing"
    "time"
)

func TestEventJSON(t *testing.T) {
    original := Event{"start", time.Date(2025, 7, 2, 15, 4, 5, 0, time.UTC)}
    data, err := json.Marshal(original)
    if err != nil {
        t.Fatal(err)
    }
    expected := "{\"name\":\"start\",\"time\":\"2025-07-02T15:04:05Z\"}"
    if string(data) != expected {
        t.Fatalf("got %s want %s", data, expected)
    }

    var decoded Event
    if err := json.Unmarshal(data, &decoded); err != nil {
        t.Fatal(err)
    }
    if decoded.Name != original.Name || !decoded.Time.Equal(original.Time) {
        t.Fatalf("decoded %+v not equal to original %+v", decoded, original)
    }
}

