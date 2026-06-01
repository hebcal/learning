package perekyomi_test

import (
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/perekyomi"
	"github.com/stretchr/testify/assert"
)

func TestPerekYomi(t *testing.T) {
	assert := assert.New(t)

	r, ok := perekyomi.New(hdate.FromGregorian(2025, time.February, 8))
	assert.True(ok)
	assert.Equal(perekyomi.Reading{Tractate: "Berakhot", Chap: 1}, r)

	r, _ = perekyomi.New(hdate.FromGregorian(2025, time.February, 19))
	assert.Equal(perekyomi.Reading{Tractate: "Peah", Chap: 3}, r)

	r, _ = perekyomi.New(hdate.FromGregorian(2024, time.April, 8))
	assert.Equal(perekyomi.Reading{Tractate: "Sotah", Chap: 8}, r)
}

func TestPerekYomiStart(t *testing.T) {
	hd := hdate.FromRD(perekyomi.PerekYomiStart)
	y, m, d := hd.Greg()
	assert.Equal(t, 2002, y)
	assert.Equal(t, time.February, m)
	assert.Equal(t, 9, d)
	r, ok := perekyomi.New(hd)
	assert.True(t, ok)
	assert.Equal(t, perekyomi.Reading{Tractate: "Berakhot", Chap: 1}, r)
}

func TestPerekYomiBeforeStart(t *testing.T) {
	_, ok := perekyomi.New(hdate.FromGregorian(2002, time.February, 8))
	assert.False(t, ok)
}

func TestPerekYomiRender(t *testing.T) {
	hd := hdate.FromGregorian(2024, time.April, 8)
	r, _ := perekyomi.New(hd)
	ev := perekyomi.NewPerekYomiEvent(hd, r)
	assert.Equal(t, "Sotah 8", ev.Render("en"))
	assert.Equal(t, "Sotah 8", ev.Basename())
}
