package psalms_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/psalms"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	ev := psalms.NewPsalmsEvent(hd, psalms.Reading{Begin: "10", End: "17"})
	assert.Equal("https://www.sefaria.org/Psalms.10-17?lang=bi", event.URL(ev))
	// Psalm 119 is split across two days using verse-prefixed references.
	ev = psalms.NewPsalmsEvent(hd, psalms.Reading{Begin: "119:1", End: "119:96"})
	assert.Equal("https://www.sefaria.org/Psalms.119.1-119.96?lang=bi", event.URL(ev))
}
