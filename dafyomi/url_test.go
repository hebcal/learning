package dafyomi_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/event"
	"github.com/hebcal/learning/dafyomi"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.New(5784, hdate.Nisan, 1)
	cases := []struct {
		daf  dafyomi.Daf
		want string
	}{
		{dafyomi.Daf{Name: "Shabbat", Blatt: 104}, "https://www.sefaria.org/Shabbat.104a?lang=bi"},
		{dafyomi.Daf{Name: "Berachot", Blatt: 2}, "https://www.sefaria.org/Berakhot.2a?lang=bi"},
		{dafyomi.Daf{Name: "Rosh Hashana", Blatt: 2}, "https://www.sefaria.org/Rosh_Hashanah.2a?lang=bi"},
		{dafyomi.Daf{Name: "Baba Kamma", Blatt: 2}, "https://www.sefaria.org/Bava_Kamma.2a?lang=bi"},
		{dafyomi.Daf{Name: "Shekalim", Blatt: 2}, "https://www.sefaria.org/Jerusalem_Talmud_Shekalim.1.1.1-10?lang=bi"},
		{dafyomi.Daf{Name: "Kinnim", Blatt: 3}, "https://www.dafyomi.org/index.php?masechta=meilah&daf=3a"},
		{dafyomi.Daf{Name: "Midot", Blatt: 2}, "https://www.dafyomi.org/index.php?masechta=meilah&daf=2a"},
	}
	for _, tc := range cases {
		ev := dafyomi.NewDafYomiEvent(hd, tc.daf)
		assert.Equal(tc.want, event.URL(ev), tc.daf.String())
	}
}
