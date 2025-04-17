package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	return steps, duration, nil
}

// DayActionInfo функция возвращает инофрмацию о кол-ве шагов, пройденную дистанцию и сожженные калории
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Ошибка парсинга данных: %v", err)
		return ""
	}
	if steps <= 0 {
		log.Printf("Некорректное количество шагов: %d", steps)
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("Ошибка расчёта калорий: %v", err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
	return result
}
