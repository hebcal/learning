package yerushalmi_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/dafyomi"
	"github.com/hebcal/learning/yerushalmi"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	daf := dafyomi.Daf{Name: "Berakhot", Blatt: 10}

	ev := yerushalmi.NewYerushalmiYomiEvent(hd, daf, yerushalmi.Vilna)
	assert.Equal("https://www.sefaria.org/Jerusalem_Talmud_Berakhot.1.5.9-14?lang=bi", event.URL(ev))

	// The Schottenstein edition is not mapped to Sefaria references.
	ev = yerushalmi.NewYerushalmiYomiEvent(hd, daf, yerushalmi.Schottenstein)
	assert.Equal("", event.URL(ev))

	// A daf marked null in the Vilna map has no URL.
	ev = yerushalmi.NewYerushalmiYomiEvent(hd, dafyomi.Daf{Name: "Peah", Blatt: 10}, yerushalmi.Vilna)
	assert.Equal("", event.URL(ev))
}
