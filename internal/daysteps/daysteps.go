// Package daysteps обрабатывает данные о дневной активности пользователя.
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит количество шагов, длительность прогулки
// и персональные данные пользователя.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку с количеством шагов и длительностью прогулки.
func (ds *DaySteps) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 2 {
		return fmt.Errorf(
			"прогулка должна содержать два поля: %q",
			datastring,
		)
	}

	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf(
			"ошибка чтения шагов %q: %w",
			fields[0],
			err,
		)
	}
	if steps <= 0 {
		return fmt.Errorf(
			"количество шагов должно быть положительным: %d",
			steps,
		)
	}

	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return fmt.Errorf(
			"ошибка чтения времени %q: %w",
			fields[1],
			err,
		)
	}
	if duration <= 0 {
		return fmt.Errorf(
			"время прогулки должно быть положительным: %v",
			duration,
		)
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo формирует строку с результатами дневной активности.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(
		ds.Steps, ds.Height,
	)
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps, ds.Weight, ds.Height, ds.Duration,
	)
	if err != nil {
		return "", fmt.Errorf(
			"невозможно рассчитать калории прогулки: %w",
			err,
		)
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
