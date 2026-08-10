package nine29_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/nine29"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	cases := []struct {
		reading nine29.Reading
		want    string
	}{
		{nine29.Reading{Book: "Deuteronomy", BookChap: 34, CycleChap: 187}, "https://www.sefaria.org/Deuteronomy.34?lang=bi"},
		{nine29.Reading{Book: "II Chronicles", BookChap: 36, CycleChap: 929}, "https://www.sefaria.org/II_Chronicles.36?lang=bi"},
		{nine29.Reading{Book: "Song of Songs", BookChap: 8, CycleChap: 798}, "https://www.sefaria.org/Song_of_Songs.8?lang=bi"},
	}
	for _, tc := range cases {
		ev := nine29.NewNine29Event(hd, tc.reading)
		assert.Equal(tc.want, event.URL(ev), tc.reading.String())
	}
}
