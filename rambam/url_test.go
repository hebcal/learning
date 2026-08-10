package rambam_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/rambam"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)

	ev := rambam.NewDailyRambam1Event(hd, rambam.Reading{Name: "Kings and Wars", Perek: "4"})
	assert.Equal("https://www.sefaria.org/Mishneh_Torah%2C_Kings_and_Wars.4?lang=bi", event.URL(ev))

	ev = rambam.NewDailyRambam1Event(hd, rambam.Reading{Name: "Transmission of the Oral Law", Perek: "1-21"})
	assert.Equal("https://www.sefaria.org/Mishneh_Torah%2C_Transmission_of_the_Oral_Law.1-21?lang=bi", event.URL(ev))
}

func TestURL3(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)

	// Three chapters in the same section collapse to one reading, which
	// carries a Sefaria URL.
	ev := rambam.NewDailyRambam3Event(hd, []rambam.Reading{
		{Name: "Overview of Mishneh Torah Contents", Perek: "1:1-4:8"},
		{Name: "Overview of Mishneh Torah Contents", Perek: "5:1-9:9"},
		{Name: "Overview of Mishneh Torah Contents", Perek: "10:1-14:10"},
	})
	assert.Equal("https://www.sefaria.org/Mishneh_Torah%2C_Overview_of_Mishneh_Torah_Contents.1.1-14.10?lang=bi", event.URL(ev))

	// A day spanning multiple sections has no single URL.
	ev = rambam.NewDailyRambam3Event(hd, []rambam.Reading{
		{Name: "Foundations of the Torah", Perek: "10"},
		{Name: "Human Dispositions", Perek: "1"},
		{Name: "Human Dispositions", Perek: "2"},
	})
	assert.Equal("", event.URL(ev))
}
