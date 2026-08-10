package pirkeiavot_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/pirkeiavot"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5783, hdate.Elul, 9)
	ev := pirkeiavot.NewPirkeiAvotSummerEvent(hd, []int{2})
	assert.Equal("https://www.sefaria.org/Pirkei_Avot.2?lang=bi", event.URL(ev))
	ev = pirkeiavot.NewPirkeiAvotSummerEvent(hd, []int{3, 4})
	assert.Equal("https://www.sefaria.org/Pirkei_Avot.3-4?lang=bi", event.URL(ev))
}
