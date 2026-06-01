package dafweekly_test

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/dafweekly"
	"github.com/hebcal/learning/dafyomi"
	"github.com/stretchr/testify/assert"
)

func TestDafWeekly(t *testing.T) {
	assert := assert.New(t)

	daf, ok := dafweekly.New(hdate.FromGregorian(2015, time.September, 26))
	assert.True(ok)
	assert.Equal(dafyomi.Daf{Name: "Yoma", Blatt: 88}, daf)

	daf, _ = dafweekly.New(hdate.FromGregorian(2015, time.September, 27))
	assert.Equal(dafyomi.Daf{Name: "Sukkah", Blatt: 2}, daf)

	daf, _ = dafweekly.New(hdate.FromGregorian(2057, time.February, 17))
	assert.Equal(dafyomi.Daf{Name: "Niddah", Blatt: 73}, daf)
}

func TestDafWeeklySameDafAllWeek(t *testing.T) {
	// 6 March 2005 is a Sunday; the whole week studies the first daf.
	want := dafyomi.Daf{Name: "Berachot", Blatt: 2}
	for d := 6; d <= 12; d++ {
		daf, ok := dafweekly.New(hdate.FromGregorian(2005, time.March, d))
		assert.True(t, ok)
		assert.Equal(t, want, daf, "2005-03-%02d", d)
	}
	// The next Sunday advances to the second daf.
	daf, _ := dafweekly.New(hdate.FromGregorian(2005, time.March, 13))
	assert.Equal(t, dafyomi.Daf{Name: "Berachot", Blatt: 3}, daf)
}

func TestDafWeeklyBeforeStart(t *testing.T) {
	_, ok := dafweekly.New(hdate.FromGregorian(2005, time.March, 5))
	assert.False(t, ok)
}

func TestDafWeeklyRender(t *testing.T) {
	hd := hdate.FromGregorian(2008, time.May, 9)
	daf, _ := dafweekly.New(hd)
	ev := dafweekly.NewDafWeeklyEvent(hd, daf)
	assert.Equal(t, "Shabbat 104", ev.Render("en"))
	assert.Equal(t, "Shabbat 104", ev.Basename())
}
