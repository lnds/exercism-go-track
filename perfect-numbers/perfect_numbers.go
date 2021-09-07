package perfect

type Classification int

const (
	None = iota
	ClassificationPerfect
	ClassificationAbundant
	ClassificationDeficient
)

type ClassificationError string

const ErrOnlyPositive ClassificationError = "only classify positive numbers"

func (e ClassificationError) Error() string {
	return string(e)
}

func Classify(number int64) (Classification, error) {
	if number <= 0 {
		return None, ErrOnlyPositive
	}
	sum := int64(0)
	for i := int64(1); i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	switch {
	case sum == number:
		return ClassificationPerfect, nil
	case sum < number:
		return ClassificationDeficient, nil
	default:
		return ClassificationAbundant, nil

	}
}
