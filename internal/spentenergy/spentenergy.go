package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func calculateCalories(steps int, weight, height float64, duration time.Duration) float64 {
	return weight * MeanSpeed(steps, height, duration) * duration.Hours()
}

func validateParams(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("bad data: step count must be positive")
	}
	if weight <= 0 {
		return fmt.Errorf("bad data: weight must be positive")
	}
	if height <= 0 {
		return fmt.Errorf("bad data: height must be positive")
	}
	if duration <= 0 {
		return fmt.Errorf("bad data: duration must be positive")
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validateParams(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	calories := calculateCalories(steps, weight, height, duration) * walkingCaloriesCoefficient
	return  calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validateParams(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calculateCalories(steps, weight, height, duration), nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps < 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
