package notificationscentre

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ParsedFrequency struct {
	Type     string         // "at" or "every"
	Time     *time.Time     // Specific time for "at"
	Interval *time.Duration // Duration for "every"
}

type FrequencyParser struct{}

func (fp *FrequencyParser) Parse(frequency string) (*ParsedFrequency, error) {
	frequency = strings.ToLower(strings.TrimSpace(frequency))

	// Match "at HH:MM"
	atRegex := regexp.MustCompile(`^at (\d{1,2}):(\d{2})$`)
	if matches := atRegex.FindStringSubmatch(frequency); matches != nil {
		hour, _ := strconv.Atoi(matches[1])
		minute, _ := strconv.Atoi(matches[2])
		now := time.Now()
		parsedTime := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
		return &ParsedFrequency{Type: "at", Time: &parsedTime}, nil
	}

	// Match "every X hour(s)"
	everyRegex := regexp.MustCompile(`^every (\d+) hour(s)?$`)
	if matches := everyRegex.FindStringSubmatch(frequency); matches != nil {
		hours, _ := strconv.Atoi(matches[1])
		interval := time.Duration(hours) * time.Hour
		return &ParsedFrequency{Type: "every", Interval: &interval}, nil
	}

	return nil, fmt.Errorf("invalid frequency format: %s", frequency)
}
