package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Filter(t *testing.T) {
	digits := List[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEvenNumber := func(item int) bool {
		return item%2 == 0
	}

	evenDigits := digits.Filter(isEvenNumber)

	assert.Equal(t, List[int]{0, 2, 4, 6, 8}, evenDigits)
}
