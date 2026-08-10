package mishnayomi_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/mishnayomi"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	ev := mishnayomi.NewMishnaYomiEvent(hd, mishnayomi.MishnaPair{
		{Tractate: "Berakhot", Chap: 3, Verse: 6},
		{Tractate: "Berakhot", Chap: 4, Verse: 1},
	})
	assert.Equal("https://www.sefaria.org/Mishnah_Berakhot.3.6-4.1?lang=bi", event.URL(ev))
	// When the pair spans two tractates only the first is linked.
	ev = mishnayomi.NewMishnaYomiEvent(hd, mishnayomi.MishnaPair{
		{Tractate: "Berakhot", Chap: 9, Verse: 5},
		{Tractate: "Peah", Chap: 1, Verse: 1},
	})
	assert.Equal("https://www.sefaria.org/Mishnah_Berakhot.9.5?lang=bi", event.URL(ev))
	// Avot uses the Pirkei_ prefix.
	ev = mishnayomi.NewMishnaYomiEvent(hd, mishnayomi.MishnaPair{
		{Tractate: "Avot", Chap: 1, Verse: 1},
		{Tractate: "Avot", Chap: 1, Verse: 2},
	})
	assert.Equal("https://www.sefaria.org/Pirkei_Avot.1.1-2?lang=bi", event.URL(ev))
}
