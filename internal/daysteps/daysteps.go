package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parsedStr := strings.Split(datastring, ",")

	if len(parsedStr) != 2 {
		return errors.New("invalid training data")
	}

	stepsStr := parsedStr[0]
	steps, err := strconv.Atoi(stepsStr)

	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("invalid steps")
	}

	ds.Steps = steps

	durationStr := parsedStr[1]
	duration, err := time.ParseDuration(durationStr)

	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("invalid duration")
	}

	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	var result string

	result += fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	result += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	result += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return result, nil
}
