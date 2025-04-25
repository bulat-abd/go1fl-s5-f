package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const trainingInfoFormat = "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"

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
	var err error
	calories := 0.0
	switch t.TrainingType {
		case "Бег":
			calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		case "Ходьба":
			calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		default:
			err = fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", fmt.Errorf("bad data: error while calculating calories - %w", err)
	}
	result := fmt.Sprintf(trainingInfoFormat, t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height), spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), calories)
	return result, nil
}
