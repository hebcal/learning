package nine29

// These tests are in-package because they exercise the unexported cycle
// arithmetic directly. 929 used to recompute its cycle by stepping forward
// from the epoch one cycle at a time, counting days as it went, so cost grew
// without bound for far-future dates. The tests below pin the cycle
// boundaries.

import (
	"testing"
	"time"

	"github.com/hebcal/greg"
	"github.com/hebcal/hdate"
	"github.com/stretchr/testify/assert"
)

// TestCycleBoundaries confirms that from cycle 2 onward the cycles repeat on a
// fixed 1302-day period, with chapter 929 landing on the Wednesday at offset
// 1298 and the next cycle opening the Sunday after.
func TestCycleBoundaries(t *testing.T) {
	assert := assert.New(t)
	for cycle := int64(0); cycle < 300; cycle++ {
		start := nine29StartCycle2 + cycle*1302
		assert.Equal(int64(0), start%7) // Sunday

		first, ok := New(hdate.FromRD(start))
		assert.True(ok)
		assert.Equal(1, first.CycleChap)
		assert.Equal(2+int(cycle), first.CycleNum)
		assert.Equal("Genesis", first.Book)
		assert.Equal(1, first.BookChap)

		last, ok := New(hdate.FromRD(start + 1298))
		assert.True(ok)
		assert.Equal(Total929Chapters, last.CycleChap)
		assert.Equal(2+int(cycle), last.CycleNum)
		assert.Equal(int64(3), (start+1298)%7) // Wednesday

		// Thursday through Saturday after the last chapter have no reading
		for d := int64(1299); d <= 1301; d++ {
			_, ok := New(hdate.FromRD(start + d))
			assert.False(ok)
		}
	}
}

// TestCycle1TruncatedAtHistoricalEnd confirms that cycle 1, which followed a
// modified schedule, stopped early on Israel's 70th Independence Day without
// reaching chapter 929.
func TestCycle1TruncatedAtHistoricalEnd(t *testing.T) {
	assert := assert.New(t)
	first, ok := New(hdate.FromRD(Nine29Start))
	assert.True(ok)
	assert.Equal(1, first.CycleNum)
	assert.Equal(1, first.CycleChap)

	lastAbs := greg.ToRD(2018, time.April, 18)
	last, ok := New(hdate.FromRD(lastAbs))
	assert.True(ok)
	assert.Equal(1, last.CycleNum)
	assert.Equal(869, last.CycleChap)
	_, ok = New(hdate.FromRD(lastAbs + 1))
	assert.False(ok)

	// Gap until cycle 2 opens
	_, ok = New(hdate.FromRD(nine29StartCycle2 - 1))
	assert.False(ok)
	second, ok := New(hdate.FromRD(nine29StartCycle2))
	assert.True(ok)
	assert.Equal(2, second.CycleNum)
}

// TestNeverSkipsOrRepeatsAChapter sweeps 40 years and confirms chapters run
// 1..929 strictly in order.
func TestNeverSkipsOrRepeatsAChapter(t *testing.T) {
	assert := assert.New(t)
	expected := 0
	cycle := 2
	start := nine29StartCycle2
	for abs := start; abs < start+40*366; abs++ {
		r, ok := New(hdate.FromRD(abs))
		if !ok {
			continue
		}
		if r.CycleNum != cycle {
			assert.Equal(Total929Chapters, expected)
			cycle = r.CycleNum
			expected = 0
		}
		expected++
		assert.Equal(expected, r.CycleChap)
	}
}
