// Package actioninfo предоставляет общий интерфейс
// для обработки данных о физической активности.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser определяет методы разбора данных
// и формирования информации об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает записи об активности, выводит результаты
// и записывает возникающие ошибки в журнал.
func Info(dataset []string, dp DataParser) {
	for _, record := range dataset {
		if err := dp.Parse(record); err != nil {
			log.Printf(
				"ошибка разбора записи %q: %v",
				record,
				err,
			)
			continue
		}

		description, err := dp.ActionInfo()
		if err != nil {
			log.Printf(
				"ошибка обработки активности %q: %v",
				record,
				err,
			)
			continue
		}

		fmt.Print(description)
	}
}
