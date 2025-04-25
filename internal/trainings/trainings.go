package trainings

import (
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"time"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	return "", nil
}
