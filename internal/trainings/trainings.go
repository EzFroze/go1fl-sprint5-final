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

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parsedStr := strings.Split(datastring, ",")

	if len(parsedStr) != 3 {
		return errors.New("invalid training data")
	}

	stepsStr := parsedStr[0]
	trainingType := parsedStr[1]
	durationStr := parsedStr[2]

	t.TrainingType = trainingType

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("invalid steps")
	}

	t.Steps = steps

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("invalid duration")
	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
			"Тип тренировки: %s\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км.\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %.2f\n",
			t.TrainingType,
			t.Duration.Hours(),
			distance,
			meanSpeed,
			calories,
		),
		nil
}
