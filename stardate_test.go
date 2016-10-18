package stardate_test

import (
	"github.com/mckayward/stardate"
	"testing"
	"time"
)

func TestStringDoesNotRound(t *testing.T) {
	var s = stardate.Stardate(12345.05)

	actual := s.String()
	expected := "12345.0"

	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestTimeToStardateOfStardateZero(t *testing.T) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	date := time.Date(2318, 7, 5, 12, 0, 0, 0, loc)

	actual := stardate.TimeToStardate(date)
	expected := stardate.Stardate(0.0)

	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestStardateToTimeOfStardateZero(t *testing.T) {
	loc, _ := time.LoadLocation("America/Los_Angeles")

	actual := stardate.StardateToTime(stardate.Stardate(0.0))
	expected := time.Date(2318, 7, 5, 12, 0, 0, 0, loc)

	if !actual.Equal(expected) {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
