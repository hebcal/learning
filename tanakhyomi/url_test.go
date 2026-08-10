package tanakhyomi_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/tanakhyomi"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	cases := []struct {
		reading tanakhyomi.Reading
		want    string
	}{
		{tanakhyomi.Reading{Name: "Isaiah", Blatt: "23", Verses: "55:13-58:13"},
			"https://www.sefaria.org/Isaiah.55.13-58.13?lang=bi"},
		{tanakhyomi.Reading{Name: "Chronicles", Blatt: "21", Verses: "II Chronicles 26:2-29:10"},
			"https://www.sefaria.org/II_Chronicles.26.2-29.10?lang=bi"},
		{tanakhyomi.Reading{Name: "Minor Prophets", Blatt: "15", Verses: "Zephaniah 3:20"},
			"https://www.sefaria.org/Zephaniah.3.20?lang=bi"},
	}
	for _, tc := range cases {
		ev := tanakhyomi.NewTanakhYomiEvent(hd, tc.reading)
		assert.Equal(tc.want, event.URL(ev), tc.reading.String())
	}
}
