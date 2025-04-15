package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// parsePackage принимает строку вида "678,0h50m",
// где "678" - количество шагов, а "0h50m" - продолжительность прогулки
func parsePackage(data string) (int, time.Duration, error) {
	comma := strings.Split(data, ",")
	if len(comma) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных, ожидается 'steps, duration', получили: %s", data)
	}

	steps, err := strconv.Atoi(comma[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования количества шагов (%s): %w", comma[0], err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть > 0, получено: %d", steps)
	}

	duration, err := time.ParseDuration(comma[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преоборзаования продолжительности (%s): %w", comma[1], err)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
