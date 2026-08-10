package nachyomi_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/dafyomi"
	"github.com/hebcal/learning/nachyomi"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	ev := nachyomi.NewNachYomiEvent(hd, dafyomi.Daf{Name: "I Samuel", Blatt: 23})
	assert.Equal("https://www.sefaria.org/I_Samuel.23?lang=bi", event.URL(ev))
}
