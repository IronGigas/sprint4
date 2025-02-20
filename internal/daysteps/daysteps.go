package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	val := strings.Split(data, ",")
	if len(val) != 2 {return 0, 0, errors.New("в строке не 2 элемента \n")}

    steps, err := strconv.Atoi(val[0])
    if err != nil {
        return 0, 0, err
    }

	duration, err := time.ParseDuration(val[1])
	if err != nil {
        return 0, 0, err
    } 
    
    return steps, duration, nil
	
}



// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
        fmt.Println(err)
		return ""
    }

	if steps < 0 {
		return ""
	}

	metres := float64(steps) * StepLength
	kilometres := metres / 1000
	calories :=  spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	str := 	fmt.Sprintf("Количество шагов: %d.\n", steps) +
	fmt.Sprintf("Дистанция составила %.2f км.\n", kilometres) +
	fmt.Sprintf("Вы сожгли %.2f ккал.\n",  calories)	
	return str

}
