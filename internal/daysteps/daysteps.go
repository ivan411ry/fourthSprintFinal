package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
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

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("error: expect two values")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error: convert steps fail %w", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("error: expect steps more than zero")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error: convert duration fail %w", err)
	}
	if duration <= 0 {
		return 0, 0, errors.New("error: expect duration more than zero ")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	metre := float64(steps) * stepLength
	kilometre := metre / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		kilometre,
		calories,
	)
}
