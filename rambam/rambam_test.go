package rambam_test

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/rambam"
	"github.com/stretchr/testify/assert"
)

func TestRambam1Start(t *testing.T) {
	hd := hdate.FromRD(rambam.Rambam1Start)
	assert.Equal(t, time.Sunday, hd.Weekday())
	y, m, d := hd.Greg()
	assert.Equal(t, 1984, y)
	assert.Equal(t, time.April, m)
	assert.Equal(t, 29, d)
}

func TestRambam1(t *testing.T) {
	assert := assert.New(t)

	// 1 Feb 1987 => Kings and Wars 4
	r, ok := rambam.New1(hdate.FromGregorian(1987, time.February, 1))
	assert.True(ok)
	assert.Equal(rambam.Reading{Name: "Kings and Wars", Perek: "4"}, r)

	// 10 Feb 1987 => Transmission of the Oral Law 1-21 (intro section)
	r, _ = rambam.New1(hdate.FromGregorian(1987, time.February, 10))
	assert.Equal(rambam.Reading{Name: "Transmission of the Oral Law", Perek: "1-21"}, r)

	ev := rambam.NewDailyRambam1Event(hdate.FromGregorian(1987, time.February, 1), rambam.Reading{Name: "Kings and Wars", Perek: "4"})
	assert.Equal("Kings and Wars 4", ev.Render("en"))
	assert.Equal("Kings and Wars 4", ev.Render("ashkenazi"))
	assert.Equal("Kings and Wars 4", ev.Basename())
	// Hebrew chapter numbers carry no geresh / gershayim, matching @hebcal/learning
	assert.Equal("הלכות מלכים ומלחמות פרק ד", ev.Render("he"))
}

func TestRambam1BeforeStart(t *testing.T) {
	_, ok := rambam.New1(hdate.FromGregorian(1984, time.April, 28))
	assert.False(t, ok)
}

func TestRambam3(t *testing.T) {
	assert := assert.New(t)

	r, ok := rambam.New3(hdate.FromGregorian(2020, time.July, 9))
	assert.True(ok)
	assert.Equal([]rambam.Reading{
		{Name: "Kings and Wars", Perek: "10"},
		{Name: "Kings and Wars", Perek: "11"},
		{Name: "Kings and Wars", Perek: "12"},
	}, r)

	r, _ = rambam.New3(hdate.FromGregorian(1984, time.May, 6))
	assert.Equal([]rambam.Reading{
		{Name: "Foundations of the Torah", Perek: "10"},
		{Name: "Human Dispositions", Perek: "1"},
		{Name: "Human Dispositions", Perek: "2"},
	}, r)

	r, _ = rambam.New3(hdate.FromGregorian(2024, time.March, 27))
	assert.Equal([]rambam.Reading{
		{Name: "Transmission of the Oral Law", Perek: "1-21"},
		{Name: "Transmission of the Oral Law", Perek: "22-33"},
		{Name: "Transmission of the Oral Law", Perek: "34-45"},
	}, r)
}

func TestRambam3Spotcheck(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		y    int
		m    time.Month
		d    int
		want []rambam.Reading
	}{
		{2020, time.August, 15, []rambam.Reading{{"The Order of Prayer", "5"}, {"Sabbath", "1"}, {"Sabbath", "2"}}},
		{2020, time.August, 16, []rambam.Reading{{"Sabbath", "3"}, {"Sabbath", "4"}, {"Sabbath", "5"}}},
		{2020, time.September, 1, []rambam.Reading{{"Leavened and Unleavened Bread", "2"}, {"Leavened and Unleavened Bread", "3"}, {"Leavened and Unleavened Bread", "4"}}},
		{2020, time.December, 5, []rambam.Reading{{"The Chosen Temple", "5"}, {"The Chosen Temple", "6"}, {"The Chosen Temple", "7"}}},
		{2021, time.January, 2, []rambam.Reading{{"Trespass", "2"}, {"Trespass", "3"}, {"Trespass", "4"}}},
	}
	for _, c := range cases {
		r, _ := rambam.New3(hdate.FromGregorian(c.y, c.m, c.d))
		assert.Equal(c.want, r, "%d-%02d-%02d", c.y, c.m, c.d)
	}
}

func TestRambam3Render(t *testing.T) {
	assert := assert.New(t)

	// Three chapters in the same section collapse to a single range.
	r, _ := rambam.New3(hdate.FromGregorian(2020, time.July, 9))
	ev := rambam.NewDailyRambam3Event(hdate.FromGregorian(2020, time.July, 9), r)
	assert.Equal("Kings and Wars 10-12", ev.Render("en"))
	assert.Equal("Kings and Wars 10-12", ev.Basename())

	// Two sections: first collapses, third stands alone.
	r, _ = rambam.New3(hdate.FromGregorian(1984, time.May, 6))
	ev = rambam.NewDailyRambam3Event(hdate.FromGregorian(1984, time.May, 6), r)
	assert.Equal("Foundations of the Torah 10, Human Dispositions 1-2", ev.Render("en"))

	// Intro section verse ranges collapse from first start to last end.
	r, _ = rambam.New3(hdate.FromGregorian(2024, time.March, 27))
	ev = rambam.NewDailyRambam3Event(hdate.FromGregorian(2024, time.March, 27), r)
	assert.Equal("Transmission of the Oral Law 1-45", ev.Render("en"))
}
