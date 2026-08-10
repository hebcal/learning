package perekyomi_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/perekyomi"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	ev := perekyomi.NewPerekYomiEvent(hd, perekyomi.Reading{Tractate: "Oholot", Chap: 9})
	assert.Equal("https://www.sefaria.org/Mishnah_Oholot.9?lang=bi", event.URL(ev))
	ev = perekyomi.NewPerekYomiEvent(hd, perekyomi.Reading{Tractate: "Avot", Chap: 1})
	assert.Equal("https://www.sefaria.org/Pirkei_Avot.1?lang=bi", event.URL(ev))
}
