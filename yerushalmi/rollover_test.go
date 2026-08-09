package yerushalmi

// These tests are in-package because they exercise the unexported cycle
// arithmetic directly. Yerushalmi Yomi used to recompute its cycle by stepping
// forward from the epoch one cycle at a time, which both grew without bound for
// far-future dates and, for Vilna, drifted into an unreachable state. The tests
// below pin the cycle boundaries and exercise the far-future range.

import (
	"testing"

	"github.com/hebcal/greg"
	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/dafyomi"
	"github.com/stretchr/testify/assert"
)

// TestVilnaRollover2172 covers 18 Av 5932 (2172-08-08), which used to panic:
// the cycle-start walk under-advanced and produced total == numDapim, one past
// the end of the shas table.
func TestVilnaRollover2172(t *testing.T) {
	assert := assert.New(t)
	const abs = 793162
	hd := hdate.FromRD(abs)
	assert.Equal("18 Av 5932", hd.String())
	assert.Equal(int64(abs), greg.ToRD(2172, 8, 8))

	// A new cycle begins on this day rather than blowing up.
	assert.Equal(dafyomi.Daf{Name: "Berakhot", Blatt: 1}, New(hd, Vilna))
	assert.Equal(int64(abs), cycleStart(Vilna, abs))

	// The preceding day is the last daf of the outgoing cycle.
	assert.Equal(dafyomi.Daf{Name: "Niddah", Blatt: 13}, New(hdate.FromRD(abs-1), Vilna))

	// ...and the days after continue in order.
	assert.Equal(dafyomi.Daf{Name: "Berakhot", Blatt: 2}, New(hdate.FromRD(abs+1), Vilna))
}

// TestVilnaRollover2420 covers the second date where the old cycle walk broke
// down.
func TestVilnaRollover2420(t *testing.T) {
	assert := assert.New(t)
	const abs = 883792
	assert.Equal("20 Tishrei 6181", hdate.FromRD(abs).String())
	assert.NotEqual(0, New(hdate.FromRD(abs), Vilna).Blatt)
}

// TestVilnaEveryCycleBoundaryIsBerakhot1 walks 200 consecutive cycles and
// confirms each starts at Berakhot 1 and the day before ends at the last daf
// of Niddah.
func TestVilnaEveryCycleBoundaryIsBerakhot1(t *testing.T) {
	assert := assert.New(t)
	numDapim := countDapim(Vilna)
	assert.Equal(1554, numDapim)
	abs := VilnaStartRD
	for cycle := 0; cycle < 200; cycle++ {
		assert.Equal(dafyomi.Daf{Name: "Berakhot", Blatt: 1}, New(hdate.FromRD(abs), Vilna))
		assert.Equal(abs, cycleStart(Vilna, abs))
		if cycle > 0 {
			// find the previous reading day (skipping YK / Tisha B'Av)
			prev := abs - 1
			for New(hdate.FromRD(prev), Vilna).Blatt == 0 {
				prev--
			}
			assert.Equal(dafyomi.Daf{Name: "Niddah", Blatt: 13}, New(hdate.FromRD(prev), Vilna))
		}
		// advance past the end of this cycle
		next := abs + int64(numDapim)
		for cycleStart(Vilna, next) == abs {
			next++
		}
		abs = cycleStart(Vilna, next)
	}
	// 200 cycles is roughly 850 years past 1980
	year, _, _ := greg.FromRD(abs)
	assert.Greater(year, 2800)
}

// TestNumSpecialDaysBoundaries confirms there is exactly one Yom Kippur and
// one (observed) Tisha B'Av per Hebrew year.
func TestNumSpecialDaysBoundaries(t *testing.T) {
	assert := assert.New(t)
	for year := 5750; year < 6250; year++ {
		rh := hdate.ToRD(year, hdate.Tishrei, 1)
		nextRh := hdate.ToRD(year+1, hdate.Tishrei, 1)
		assert.Equal(2, numSpecialDays(Vilna, rh, nextRh-1))
	}
	// Schottenstein never skips
	assert.Equal(0, numSpecialDays(Schottenstein, 738473, 803533))
}
