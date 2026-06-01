package tanakhyomi_test

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/tanakhyomi"
	"github.com/stretchr/testify/assert"
)

func TestTanakhYomi(t *testing.T) {
	assert := assert.New(t)

	r, ok := tanakhyomi.New(hdate.FromGregorian(2021, time.February, 28))
	assert.True(ok)
	assert.Equal(tanakhyomi.Reading{Name: "Isaiah", Blatt: "23", Verses: "55:13-58:13"}, r)

	// Skipped on Shabbat
	_, ok = tanakhyomi.New(hdate.FromGregorian(2020, time.October, 17))
	assert.False(ok)

	// Composite book: Minor Prophets seder embeds the sub-book
	r, _ = tanakhyomi.New(hdate.FromGregorian(2021, time.June, 2))
	assert.Equal(tanakhyomi.Reading{Name: "Minor Prophets", Blatt: "15", Verses: "Zephaniah 3:20"}, r)

	// "Before 23 Tishrei" Chronicles tail branch
	r, _ = tanakhyomi.New(hdate.FromGregorian(2021, time.September, 22))
	assert.Equal(tanakhyomi.Reading{Name: "Chronicles", Blatt: "21", Verses: "II Chronicles 26:2-29:10"}, r)
}

func TestTanakhYomiBeforeStart(t *testing.T) {
	_, ok := tanakhyomi.New(hdate.FromGregorian(1948, time.October, 25))
	assert.False(t, ok)
}

func TestTanakhYomiBookTransition(t *testing.T) {
	assert := assert.New(t)
	first := func(y int, m time.Month, d int) string {
		r, ok := tanakhyomi.New(hdate.FromGregorian(y, m, d))
		assert.True(ok)
		return r.Name + " " + r.Blatt
	}
	assert.Equal("Joshua 1", first(2020, time.October, 11)) // cycle start, 23 Tishrei 5781
	assert.Equal("Joshua 14", first(2020, time.October, 26))
	assert.Equal("Judges 1", first(2020, time.October, 27))
}

func TestTanakhYomiLongJoshua(t *testing.T) {
	// 5787 is a year where Joshua seder 4 is split across two days.
	assert := assert.New(t)
	cases := []struct {
		d    int
		want tanakhyomi.Reading
	}{
		{6, tanakhyomi.Reading{Name: "Joshua", Blatt: "3", Verses: "4:24-6:26"}},
		{7, tanakhyomi.Reading{Name: "Joshua", Blatt: "4.1", Verses: "6.27-7.26"}},
		{8, tanakhyomi.Reading{Name: "Joshua", Blatt: "4.2", Verses: "8.1-32"}},
		{9, tanakhyomi.Reading{Name: "Joshua", Blatt: "5", Verses: "8:33-10:7"}},
	}
	for _, c := range cases {
		r, ok := tanakhyomi.New(hdate.FromGregorian(2026, time.October, c.d))
		assert.True(ok)
		assert.Equal(c.want, r, "2026-10-%02d", c.d)
	}
}

func TestTanakhYomiRender(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.FromGregorian(2021, time.February, 28)
	r, _ := tanakhyomi.New(hd)
	ev := tanakhyomi.NewTanakhYomiEvent(hd, r)
	assert.Equal("Isaiah Seder 23 (55:13-58:13)", ev.Render("en"))
	assert.Equal("Isaiah Seder 23 (55:13-58:13)", ev.Basename())
	// Hebrew uses gematriya for the seder and keeps the verse range.
	assert.Contains(ev.Render("he"), "ס׳ כג (55:13-58:13)")

	hd2 := hdate.FromGregorian(2021, time.June, 2)
	r2, _ := tanakhyomi.New(hd2)
	ev2 := tanakhyomi.NewTanakhYomiEvent(hd2, r2)
	assert.Equal("Minor Prophets Seder 15 (Zephaniah 3:20)", ev2.Render("en"))
}
