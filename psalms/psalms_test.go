package psalms_test

import (
	"testing"

	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/psalms"
	"github.com/stretchr/testify/assert"
)

func TestPsalms(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(psalms.Reading{Begin: "1", End: "9"},
		psalms.New(hdate.New(5783, hdate.Av, 1)))
	assert.Equal(psalms.Reading{Begin: "88", End: "89"},
		psalms.New(hdate.New(5783, hdate.Sivan, 18)))
	// Sivan has 30 days
	assert.Equal(psalms.Reading{Begin: "140", End: "144"},
		psalms.New(hdate.New(5783, hdate.Sivan, 29)))
	assert.Equal(psalms.Reading{Begin: "145", End: "150"},
		psalms.New(hdate.New(5783, hdate.Sivan, 30)))
	// Tamuz has 29 days, so the 29th and 30th portions are combined
	assert.Equal(psalms.Reading{Begin: "140", End: "150"},
		psalms.New(hdate.New(5783, hdate.Tamuz, 29)))
}

func TestPsalms119(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(psalms.Reading{Begin: "119:1", End: "119:96"},
		psalms.New(hdate.New(5783, hdate.Av, 25)))
	assert.Equal(psalms.Reading{Begin: "119:97", End: "119:176"},
		psalms.New(hdate.New(5783, hdate.Av, 26)))
}

func TestPsalmsRender(t *testing.T) {
	hd := hdate.New(5783, hdate.Av, 3)
	ev := psalms.NewPsalmsEvent(hd, psalms.New(hd))
	assert.Equal(t, "Psalms 18-22", ev.Render("en"))
	assert.Equal(t, "Psalms 18-22", ev.Basename())
}
