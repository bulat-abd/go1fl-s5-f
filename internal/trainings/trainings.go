package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 3 {
		return fmt.Errorf("bad format: there must be exactly two commas separating three fields")
	}
	activity := fields[1]
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("bad format: could not parse int in first field - %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("bad data: step count must be positive")
	}
	duration, err := time.ParseDuration(fields[2])
	if err != nil {
		return fmt.Errorf("bad format: could not parse duration in third field - %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("bad data: duration must be positive")
	}
	t.Steps = steps
	t.TrainingType = activity
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	return "", nil
}
