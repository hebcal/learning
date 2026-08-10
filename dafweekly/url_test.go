package dafweekly_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/dafweekly"
	"github.com/hebcal/learning/dafyomi"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	ev := dafweekly.NewDafWeeklyEvent(hd, dafyomi.Daf{Name: "Shabbat", Blatt: 104})
	assert.Equal("https://www.sefaria.org/Shabbat.104a?lang=bi", event.URL(ev))
}
