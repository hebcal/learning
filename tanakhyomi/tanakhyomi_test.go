package tanakhyomi_test

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
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
	assert.Equal(tanakhyomi.Reading{Name: "Minor Prophets", Blatt: "15", Verses: "Zephaniah 3:20-Haggai 2:22"}, r)

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
		{7, tanakhyomi.Reading{Name: "Joshua", Blatt: "4.1", Verses: "6:27-7:26"}},
		{8, tanakhyomi.Reading{Name: "Joshua", Blatt: "4.2", Verses: "8:1-32"}},
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
	// Hebrew uses gematriya for the seder and the verse range.
	assert.Equal("יְשַׁעְיָהוּ ס׳ כג (נה:יג-נח:יג)", ev.Render("he"))
	assert.Equal("ישעיהו ס׳ כג (נה:יג-נח:יג)", ev.Render("he-x-NoNikud"))

	hd2 := hdate.FromGregorian(2021, time.June, 2)
	r2, _ := tanakhyomi.New(hd2)
	ev2 := tanakhyomi.NewTanakhYomiEvent(hd2, r2)
	assert.Equal("Minor Prophets Seder 15 (Zephaniah 3:20-Haggai 2:22)", ev2.Render("en"))
	assert.Equal("תְּרֵי עֲשַׂר ס׳ טו (צְפַנְיָה ג:כ-חַגַּי ב:כב)", ev2.Render("he"))

	hd3 := hdate.FromGregorian(2021, time.September, 22)
	r3, _ := tanakhyomi.New(hd3)
	ev3 := tanakhyomi.NewTanakhYomiEvent(hd3, r3)
	assert.Equal("Chronicles Seder 21 (II Chronicles 26:2-29:10)", ev3.Render("en"))
	assert.Equal("דִברֵי הַיָמִים ס׳ כא (דִברֵי הַיָמִים ב׳ כו:ב-כט:י)", ev3.Render("he"))

	hd4 := hdate.FromGregorian(2026, time.October, 7)
	r4, _ := tanakhyomi.New(hd4)
	ev4 := tanakhyomi.NewTanakhYomiEvent(hd4, r4)
	assert.Equal("Joshua Seder 4.1 (6:27-7:26)", ev4.Render("en"))
	assert.Equal("יהושע ס׳ ד1 (ו:כז-ז:כו)", ev4.Render("he-x-NoNikud"))
}

// Sedarim that run into the next book. Sefaria cannot express a range across
// two books, so these used to stop at the first book's last verse or at the
// seder's first verse, leaving e.g. Joel 1:1-2:26 in no seder at all.
func TestTanakhYomiCrossBook(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		y    int
		m    time.Month
		d    int
		en   string
		link string
	}{
		{2026, time.November, 29, "Samuel Seder 20 (I Samuel 30:25-II Samuel 2:6)", "I_Samuel.30.25-31.13"},
		{2027, time.May, 16, "Minor Prophets Seder 4 (Hosea 14:6-Joel 2:26)", "Hosea.14.6-14.10"},
		{2027, time.May, 21, "Minor Prophets Seder 9 (Obadiah 1:21-Jonah 4:11)", "Obadiah.1.21"},
		{2027, time.June, 4, "Minor Prophets Seder 21 (Zechariah 14:21-Malachi 3:24)", "Zechariah.14.21"},
	}
	for _, c := range cases {
		hd := hdate.FromGregorian(c.y, c.m, c.d)
		r, ok := tanakhyomi.New(hd)
		assert.True(ok)
		ev := tanakhyomi.NewTanakhYomiEvent(hd, r)
		assert.Equal(c.en, ev.Render("en"))
		assert.Equal(c.en, ev.Basename())
		assert.Equal("https://www.sefaria.org/"+c.link+"?lang=bi", event.URL(ev))
	}
	hd := hdate.FromGregorian(2027, time.May, 16)
	r, _ := tanakhyomi.New(hd)
	ev := tanakhyomi.NewTanakhYomiEvent(hd, r)
	assert.Equal("תרי עשר ס׳ ד (הושע יד:ו-יואל ב:כו)", ev.Render("he-x-NoNikud"))
	// Nahum is absent from the locales package
	hd = hdate.FromGregorian(2027, time.May, 25)
	r, _ = tanakhyomi.New(hd)
	ev = tanakhyomi.NewTanakhYomiEvent(hd, r)
	assert.Equal("תרי עשר ס׳ יב (מיכה ז:כ-נחום ג:יט)", ev.Render("he-x-NoNikud"))
}
