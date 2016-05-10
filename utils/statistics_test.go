package utils

import (
	"testing"
	"time"
)

const (
	statPeriod    = 500 * time.Millisecond
	expectedCount = 1
	emptyCount    = 0
)

func TestStatistics(t *testing.T) {

	statistics := NewStatistics(2 * statPeriod)
	statistics.PlusOne()

	time.Sleep(statPeriod)

	if statistics.counter != expectedCount {
		t.Errorf("Wrong statictics count %d is different from expected %d", statistics.counter, expectedCount)
	}

	time.Sleep(3 * statPeriod)

	if statistics.counter != emptyCount {
		t.Errorf("Wrong statictics count %d is different from expected %d", statistics.counter, emptyCount)
	}

}
