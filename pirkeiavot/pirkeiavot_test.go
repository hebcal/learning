package pirkeiavot_test

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/pirkeiavot"
	"github.com/stretchr/testify/assert"
)

// collect returns the Pirkei Avot readings for every Shabbat from the
// first Saturday on or after 21 Nisan through Erev Rosh Hashana, mirroring
// the makeSaturdays helper in the @hebcal/learning test suite. A nil entry
// represents a skipped Shabbat.
func collect(year int, il bool) [][]int {
	pesach7 := hdate.New(year, hdate.Nisan, 21)
	fs := pesach7.OnOrAfter(time.Saturday)
	firstSat := fs.Abs()
	rh := hdate.New(year+1, hdate.Tishrei, 1)
	erevRH := rh.Abs() - 1
	var out [][]int
	for abs := firstSat; abs <= erevRH; abs += 7 {
		r, ok := pirkeiavot.New(hdate.FromRD(abs), il)
		if !ok {
			out = append(out, nil)
		} else {
			out = append(out, r)
		}
	}
	return out
}

func TestPirkeiAvot5784Israel(t *testing.T) {
	assert.Equal(t, [][]int{
		{1}, {2}, {3}, {4}, {5}, {6},
		{1}, {2}, {3}, {4}, {5}, {6},
		{1}, {2}, {3}, {4}, {5}, {6},
		{1}, {2}, {3, 4}, {5, 6},
	}, collect(5784, true))
}

func TestPirkeiAvot5783Diaspora(t *testing.T) {
	assert.Equal(t, [][]int{
		{1}, {2}, {3}, {4}, {5}, {6},
		nil, // 7 Sivan: 2nd day Shavuot
		{1}, {2}, {3}, {4}, {5}, {6},
		{1}, {2}, {3}, {4}, {5}, {6},
		{1, 2}, {3, 4}, {5, 6},
	}, collect(5783, false))
}

func TestPirkeiAvot5783Israel(t *testing.T) {
	assert.Equal(t, [][]int{
		{1}, {2}, {3}, {4}, {5}, {6},
		{1}, {2}, {3}, {4}, {5}, {6},
		{1}, {2}, {3}, {4}, {5}, {6},
		{1}, {2}, {3, 4}, {5, 6},
	}, collect(5783, true))
}

func TestPirkeiAvotNotShabbat(t *testing.T) {
	// 24 Nisan 3759 is not a Pirkei Avot Shabbat.
	_, ok := pirkeiavot.New(hdate.New(3759, hdate.Nisan, 24), true)
	assert.False(t, ok)
}

func TestPirkeiAvotRender(t *testing.T) {
	ev2 := pirkeiavot.NewPirkeiAvotSummerEvent(hdate.New(5783, hdate.Elul, 9), []int{2})
	ev34 := pirkeiavot.NewPirkeiAvotSummerEvent(hdate.New(5783, hdate.Elul, 16), []int{3, 4})
	assert.Equal(t, "Pirkei Avot 2", ev2.Render("en"))
	assert.Equal(t, "Pirkei Avot 3-4", ev34.Render("en"))
	assert.Equal(t, "Pirkei Avot 3-4", ev34.Basename())
	assert.Equal(t, "פִּרְקֵי אָבוֹת ב׳", ev2.Render("he"))
	assert.Equal(t, "פִּרְקֵי אָבוֹת ג׳-ד׳", ev34.Render("he"))
}
