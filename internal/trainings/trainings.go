// Package trainings отвечает за обработку данных о тренировках.
package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит сведения о тренировке и персональные данные пользователя.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку с количеством шагов,
// типом тренировки и её продолжительностью.
func (t *Training) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 3 {
		return fmt.Errorf(
			"тренировка должна содержать три поля: %q",
			datastring,
		)
	}

	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf(
			"не удалось прочитать число шагов %q: %w",
			fields[0],
			err,
		)
	}
	if steps <= 0 {
		return fmt.Errorf(
			"число шагов должно быть положительным: %d",
			steps,
		)
	}

	duration, err := time.ParseDuration(fields[2])
	if err != nil {
		return fmt.Errorf(
			"не удалось прочитать длительность %q: %w",
			fields[2],
			err,
		)
	}
	if duration <= 0 {
		return fmt.Errorf(
			"длительность тренировки должна быть положительной: %v",
			duration,
		)
	}

	// Записываем данные в структуру только после успешной проверки.
	t.Steps, t.TrainingType, t.Duration = steps, fields[1], duration
	return nil
}

// ActionInfo формирует строку с результатами тренировки.
func (t Training) ActionInfo() (string, error) {
	var (
		calories float64
		err      error
	)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", fmt.Errorf(
			"невозможно рассчитать расход энергии: %w",
			err,
		)
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(
		t.Steps, t.Height, t.Duration,
	)

	summary := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)
	return summary, nil
}
