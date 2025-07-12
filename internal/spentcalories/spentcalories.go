package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("error: expect three values")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("error: step convert fail: %w", err)
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("error: expect steps more than zero")
	}
	activity := parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("error: convert duration fail: %w", err)
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("error: expect duration more than zero")
	}
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	metre := float64(steps) * stepLength
	kilometre := metre / mInKm
	return kilometre
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	hours := duration.Hours()
	speed := dist / hours
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var calories float64
	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		log.Println(err)
		return "", err
	}
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity,
		duration.Hours(),
		dist,
		speed,
		calories,
	), nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("error: expect steps more than zero")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("error: expect weight more than zero")
	}
	if height <= 0 {
		return 0, fmt.Errorf("error: expect height more than zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("error: expect duration more than zero")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("error: expect steps more than zero")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("error: expect weight more than zero")
	}
	if height <= 0 {
		return 0, fmt.Errorf("error: expect height more than zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("error: expect duration more than zero")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	walkingCalories := calories * walkingCaloriesCoefficient
	return walkingCalories, nil
}
