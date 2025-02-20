package spentcalories

import (
	"fmt"
	"time"
	"strings"
	"errors"
	"strconv"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	val := strings.Split(data, ",")

	if len(val) != 3 {return 0, "", 0, errors.New("в строке не 3 элемента \n")}

	steps, err := strconv.Atoi(val[0])
    if err != nil {
        return 0, "", 0, err
    }

	activity := (val[1])

	duration, err := time.ParseDuration(val[2])
	if err != nil {
        return 0, "", 0, err
    }

	return steps, activity, duration, nil

}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	return (float64(steps) * lenStep) / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {return 0}
	distance:=distance(steps)
	hours := float64(duration.Hours())
	return distance / hours
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
        return err.Error()
    }


	switch {
	case activity == "Ходьба":
		distance:=distance(steps)
		meanSpeed:=meanSpeed(steps,duration)
		spentCalories:=WalkingSpentCalories(steps, weight, height, duration)
		str := `Тип тренировки: ` + activity + "\n" +
    	`Длительность: ` + fmt.Sprintf("%.2f", duration.Hours())  + " ч.\n" +
    	`Дистанция: ` + fmt.Sprintf("%.2f", distance) + " км.\n" +
    	`Скорость: ` + fmt.Sprintf("%.2f", meanSpeed) + " км/ч\n" +
    	`Сожгли калорий: ` + fmt.Sprintf("%.2f", spentCalories)	+ "\n"	
		return str
	case activity == "Бег":
		distance:=distance(steps)
		meanSpeed:=meanSpeed(steps,duration)
		spentCalories:=RunningSpentCalories(steps, weight, duration)
		str := `Тип тренировки: ` + activity + "\n" +
    	`Длительность: ` + fmt.Sprintf("%.2f", duration.Hours())  + " ч.\n" +
    	`Дистанция: ` + fmt.Sprintf("%.2f", distance) + " км.\n" +
    	`Скорость: ` + fmt.Sprintf("%.2f", meanSpeed) + " км/ч\n" +
    	`Сожгли калорий: ` + fmt.Sprintf("%.2f", spentCalories)	 + "\n"	
		return str
	default:
		return "неизвестный тип тренировки"
	} 


}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	meanSpeed:=meanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier*meanSpeed)-runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	meanSpeed:=meanSpeed(steps, duration)
	hours := float64(duration.Hours())
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * hours * minInH
}
