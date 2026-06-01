package nine29_test

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/nine29"
	"github.com/stretchr/testify/assert"
)

func read(t *testing.T, y int, m time.Month, d int) (nine29.Reading, bool) {
	t.Helper()
	return nine29.New(hdate.FromGregorian(y, m, d))
}

func TestNine29Start(t *testing.T) {
	hd := hdate.FromRD(nine29.Nine29Start)
	assert.Equal(t, time.Sunday, hd.Weekday())
	y, m, d := hd.Greg()
	assert.Equal(t, 2014, y)
	assert.Equal(t, time.December, m)
	assert.Equal(t, 21, d)
}

func TestNine29Basic(t *testing.T) {
	assert := assert.New(t)

	r, ok := read(t, 2014, time.December, 21) // Sun, first day
	assert.True(ok)
	assert.Equal(nine29.Reading{CycleChap: 1, CycleNum: 1, Book: "Genesis", BookChap: 1}, r)

	r, _ = read(t, 2014, time.December, 22) // Mon
	assert.Equal(nine29.Reading{CycleChap: 2, CycleNum: 1, Book: "Genesis", BookChap: 2}, r)

	r, _ = read(t, 2014, time.December, 25) // Thu
	assert.Equal(nine29.Reading{CycleChap: 5, CycleNum: 1, Book: "Genesis", BookChap: 5}, r)

	_, ok = read(t, 2014, time.December, 26) // Fri - skip
	assert.False(ok)
	_, ok = read(t, 2014, time.December, 27) // Sat - skip
	assert.False(ok)

	r, _ = read(t, 2014, time.December, 28) // Sun - resumes
	assert.Equal(nine29.Reading{CycleChap: 6, CycleNum: 1, Book: "Genesis", BookChap: 6}, r)

	r, _ = read(t, 2014, time.December, 29) // Mon
	assert.Equal(nine29.Reading{CycleChap: 7, CycleNum: 1, Book: "Genesis", BookChap: 7}, r)
}

func TestNine29BeforeStart(t *testing.T) {
	_, ok := read(t, 2014, time.December, 20)
	assert.False(t, ok)
}

func TestNine29CycleTransition(t *testing.T) {
	assert := assert.New(t)

	// Day after the historical end of cycle 1 is in the gap: no reading.
	_, ok := nine29.New(hdate.FromRD(nine29.Nine29Start)) // sanity
	assert.True(ok)
	_, ok = read(t, 2018, time.April, 19) // gap to cycle 2
	assert.False(ok)
	_, ok = read(t, 2018, time.July, 14) // Sat before cycle 2
	assert.False(ok)

	// Cycle 2 starts Sun 15 Jul 2018 at chapter 1.
	r, ok := read(t, 2018, time.July, 15)
	assert.True(ok)
	assert.Equal(1, r.CycleChap)
	assert.Equal(2, r.CycleNum)
	assert.Equal("Genesis", r.Book)
	assert.Equal(1, r.BookChap)

	// Thu 20 Sep 2018 is chapter 50 of cycle 2.
	r, _ = read(t, 2018, time.September, 20)
	assert.Equal(50, r.CycleChap)
	assert.Equal("Genesis", r.Book)
	assert.Equal(50, r.BookChap)

	// Cycle 2 ends Wed 2 Feb 2022 with chapter 929.
	r, _ = read(t, 2022, time.February, 2)
	assert.Equal(929, r.CycleChap)
	assert.Equal("II Chronicles", r.Book)
	assert.Equal(36, r.BookChap)
}

func TestNine29Render(t *testing.T) {
	hd := hdate.FromGregorian(2014, time.December, 21)
	r, _ := nine29.New(hd)
	ev := nine29.NewNine29Event(hd, r)
	assert.Equal(t, "Genesis 1 (1)", ev.Render("en"))
	assert.Equal(t, "Genesis 1 (1)", ev.Basename())
}
