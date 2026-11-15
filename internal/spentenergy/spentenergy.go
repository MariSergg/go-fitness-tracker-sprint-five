// Package spentenergy содержит функции для расчёта дистанции,
// средней скорости и расхода калорий.
package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Distance вычисляет пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient
	distanceInMetres := float64(steps) * stepLength
	return distanceInMetres / mInKm
}

// MeanSpeed вычисляет среднюю скорость в километрах в час.
func MeanSpeed(
	steps int,
	height float64,
	duration time.Duration,
) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// RunningSpentCalories рассчитывает количество калорий,
// потраченных при беге.
func RunningSpentCalories(
	steps int,
	weight, height float64,
	duration time.Duration,
) (float64, error) {
	if err := validateParameters(
		steps, weight, height, duration,
	); err != nil {
		return 0, err
	}

	return calculateCalories(
		steps, weight, height, duration,
	), nil
}

// WalkingSpentCalories рассчитывает количество калорий,
// потраченных при ходьбе.
func WalkingSpentCalories(
	steps int,
	weight, height float64,
	duration time.Duration,
) (float64, error) {
	if err := validateParameters(
		steps, weight, height, duration,
	); err != nil {
		return 0, err
	}

	return calculateCalories(
		steps, weight, height, duration,
	) * walkingCaloriesCoefficient, nil
}

func validateParameters(
	steps int,
	weight, height float64,
	duration time.Duration,
) error {
	switch {
	case steps <= 0:
		return fmt.Errorf(
			"количество шагов должно быть больше нуля: %d",
			steps,
		)
	case weight <= 0:
		return fmt.Errorf(
			"вес должен быть больше нуля: %v",
			weight,
		)
	case height <= 0:
		return fmt.Errorf(
			"рост должен быть больше нуля: %v",
			height,
		)
	case duration <= 0:
		return fmt.Errorf(
			"продолжительность должна быть больше нуля: %v",
			duration,
		)
	}
	return nil
}

func calculateCalories(
	steps int,
	weight, height float64,
	duration time.Duration,
) float64 {
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	return (weight * speed * minutes) / minInH
}
