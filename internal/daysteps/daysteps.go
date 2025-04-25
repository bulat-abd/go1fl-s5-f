package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 2 {
		return fmt.Errorf("bad format: exactly one comma must be present, separating exactly two fields")
	}
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		fmt.Errorf("bad format: could not parse int in the first field - %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("bad data: step count must be positive")
	}
	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return fmt.Errorf("bad format: could not parse duration - %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("bad data: duration must be positive")
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	return "", nil
}
