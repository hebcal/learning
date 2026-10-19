package learning_test

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/hebcal-go/dailylearning"
	"github.com/stretchr/testify/assert"

	_ "github.com/hebcal/learning"
)

type briefRenderer interface{ RenderBrief(string) string }
type memoer interface{ Memo(string) string }

// TestRenderBrief checks the brief renderings, the form @hebcal/learning's
// renderBrief() gives, against reference values from that package.
func TestRenderBrief(t *testing.T) {
	jan1 := hdate.FromGregorian(2025, time.January, 1)
	tests := []struct {
		name, locale, render, brief string
	}{
		{"929", "en", "Job 11 (759)", "Job 11"},
		{"929", "he", "אִיּוֹב י״א (759)", "אִיּוֹב י״א"},
		{"yerushalmi-vilna", "en", "Yerushalmi Shekalim 17", "Shekalim 17"},
		{"yerushalmi-vilna", "he", "יְרוּשַׁלְמִי שקלים דף יז", "שקלים דף יז"},
		{"tanakhYomi", "en", "Samuel Seder 31 (II Samuel 18:27-19:39)", "Samuel S31 (II Samuel 18:27-19:39)"},
		{"tanakhYomi", "he", "שְׁמוּאֵל ס׳ לא (שְׁמוּאֵל ב׳ יח:כז-יט:לט)", "שְׁמוּאֵל ס׳ לא (שְׁמוּאֵל ב׳ יח:כז-יט:לט)"},
	}
	for _, tc := range tests {
		ev := dailylearning.Lookup(tc.name, jan1, false)
		if !assert.NotNil(t, ev, tc.name) {
			continue
		}
		assert.Equal(t, tc.render, ev.Render(tc.locale), "%s %s", tc.name, tc.locale)
		b, ok := ev.(briefRenderer)
		if assert.True(t, ok, "%s has no RenderBrief", tc.name) {
			assert.Equal(t, tc.brief, b.RenderBrief(tc.locale), "%s %s", tc.name, tc.locale)
		}
	}
}

// TestMemo checks the memos against @hebcal/learning's.
func TestMemo(t *testing.T) {
	memo := func(name string, gd time.Time) string {
		ev := dailylearning.Lookup(name, hdate.FromTime(gd), false)
		m, ok := ev.(memoer)
		if !ok {
			t.Fatalf("%s has no Memo", name)
		}
		return m.Memo("en")
	}
	jan1 := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, "II Samuel 18:27-19:39", memo("tanakhYomi", jan1))
	// the seder's own book is named too
	assert.Equal(t, "Isaiah 1:1-4:2", memo("tanakhYomi", time.Date(2025, time.February, 16, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, "Neighbors 13-14\n"+
		"https://www.sefaria.org/Mishneh_Torah%2C_Neighbors.13-14?lang=bi\n\n"+
		"Agents and Partners 1\n"+
		"https://www.sefaria.org/Mishneh_Torah%2C_Agents_and_Partners.1?lang=bi",
		memo("rambam3", jan1))
	// a single reading has no memo; its link is the event's URL
	assert.Equal(t, "", memo("rambam3", jan1.AddDate(0, 0, 1)))
}
