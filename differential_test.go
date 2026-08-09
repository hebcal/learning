package learning_test

// Exhaustive day-by-day regression coverage for the three schedules whose
// cycle math was converted from "walk forward from the epoch" to closed-form
// arithmetic (929, Yerushalmi and Tanakh Yomi).
//
// Every day in the supported range is evaluated and folded into an FNV-1a
// checksum, one block at a time so a failure points at a date range. The
// expected checksums were generated from the original implementation, before
// the optimization, so any change in output for any single day in ~200 years
// will fail here.
//
// The one deliberate exception is documented on the Vilna table.

import (
	"strconv"
	"testing"
	"time"

	"github.com/hebcal/hdate"
	"github.com/hebcal/learning/dafyomi"
	"github.com/hebcal/learning/nine29"
	"github.com/hebcal/learning/tanakhyomi"
	"github.com/hebcal/learning/yerushalmi"
	"github.com/stretchr/testify/assert"
)

// fnv1a is the 32-bit FNV-1a hash of s.
func fnv1a(s string) uint32 {
	var h uint32 = 0x811c9dc5
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 0x01000193
	}
	return h
}

// block is an inclusive [startAbs, endAbs] range of R.D. days, the number of
// days it contains, and the expected checksum.
type block struct {
	startAbs, endAbs int64
	numDays          int
	checksum         uint32
}

func checkBlocks(t *testing.T, blocks []block, serialize func(abs int64) string) {
	t.Helper()
	assert := assert.New(t)
	for _, b := range blocks {
		var sb []byte
		n := 0
		for abs := b.startAbs; abs <= b.endAbs; abs++ {
			sb = strconv.AppendInt(sb, abs, 10)
			sb = append(sb, '\t')
			sb = append(sb, serialize(abs)...)
			sb = append(sb, '\n')
			n++
		}
		actual := fnv1a(string(sb))
		// Include the range in the failure message rather than asserting bare numbers
		assert.Equal(
			[]string{strconv.FormatInt(b.startAbs, 10), strconv.FormatInt(b.endAbs, 10),
				strconv.Itoa(b.numDays), strconv.FormatUint(uint64(b.checksum), 16)},
			[]string{strconv.FormatInt(b.startAbs, 10), strconv.FormatInt(b.endAbs, 10),
				strconv.Itoa(n), strconv.FormatUint(uint64(actual), 16)},
		)
	}
}

func serialize929(abs int64) string {
	r, ok := nine29.New(hdate.FromRD(abs))
	if !ok {
		return "null"
	}
	return strconv.Itoa(r.CycleNum) + "|" + strconv.Itoa(r.CycleChap) + "|" +
		r.Book + "|" + strconv.Itoa(r.BookChap)
}

func serializeYerushalmi(ed yerushalmi.Edition, name string) func(abs int64) string {
	return func(abs int64) string {
		d := yerushalmi.New(hdate.FromRD(abs), ed)
		if d.Blatt == 0 {
			return "null"
		}
		return name + "|" + d.Name + "|" + strconv.Itoa(d.Blatt)
	}
}

func serializeTanakhYomi(abs int64) string {
	r, ok := tanakhyomi.New(hdate.FromRD(abs))
	if !ok {
		return "null"
	}
	return r.Name + "|" + r.Blatt + "|" + r.Verses
}

func TestDifferential929(t *testing.T) {
	blocks := []block{
		{735588, 744364, 8777, 0xf892e9a3},
		{744365, 753495, 9131, 0xf1ca53cf},
		{753496, 762627, 9132, 0x9023d295},
		{762628, 771757, 9130, 0xb549556e},
		{771758, 780888, 9131, 0xb42e598f},
		{780889, 790019, 9131, 0xd46b8808},
		{790020, 799151, 9132, 0xcf2564b8},
		{799152, 803533, 4382, 0xafbc6327},
	}
	assert.Equal(t, nine29.Nine29Start, blocks[0].startAbs)
	checkBlocks(t, blocks, serialize929)
}

func TestDifferentialYerushalmiSchottenstein(t *testing.T) {
	blocks := []block{
		{738473, 747286, 8814, 0x3a69294b},
		{747287, 756417, 9131, 0x716f3c7e},
		{756418, 765549, 9132, 0xc580a991},
		{765550, 774679, 9130, 0x12a933a1},
		{774680, 783810, 9131, 0x4b82d25d},
		{783811, 792941, 9131, 0x5fd5e81f},
		{792942, 802073, 9132, 0xf8dc0ec1},
		{802074, 803533, 1460, 0x2c239f6b},
	}
	assert.Equal(t, yerushalmi.SchottensteinStartRD, blocks[0].startAbs)
	checkBlocks(t, blocks, serializeYerushalmi(yerushalmi.Schottenstein, "schottenstein"))
}

func TestDifferentialTanakhYomi(t *testing.T) {
	blocks := []block{
		{711426, 720258, 8833, 0xd58af23c},
		{720259, 729389, 9131, 0xd252a3ac},
		{729390, 738520, 9131, 0xf3bccb9e},
		{738521, 747651, 9131, 0xb8a47a3f},
		{747652, 756783, 9132, 0xe939e41f},
		{756784, 765914, 9131, 0x6faddd74},
		{765915, 775044, 9130, 0xd37f7312},
		{775045, 784175, 9131, 0x02574a54},
		{784176, 793307, 9132, 0x4f99fcd7},
		{793308, 802438, 9131, 0x9f0591c2},
		{802439, 803533, 1095, 0x2640b732},
	}
	assert.Equal(t, tanakhyomi.TanakhYomiStart, blocks[0].startAbs)
	checkBlocks(t, blocks, serializeTanakhYomi)
}

// TestDifferentialYerushalmiVilna is byte-for-byte identical to the
// pre-optimization implementation for the 70,315 days from 1980-02-02 up to
// the cycle rollover on 2172-08-08.
func TestDifferentialYerushalmiVilna(t *testing.T) {
	blocks := []block{
		{722847, 731977, 9131, 0xd972b71a},
		{731978, 741108, 9131, 0xd05f841f},
		{741109, 750239, 9131, 0x12e406fd},
		{750240, 759370, 9131, 0x80fa6c91},
		{759371, 768501, 9131, 0xbb2c2e13},
		{768502, 777632, 9131, 0x8f98a83d},
		{777633, 786763, 9131, 0xdc4f88c9},
		{786764, 793161, 6398, 0xf019973c},
	}
	assert.Equal(t, yerushalmi.VilnaStartRD, blocks[0].startAbs)
	checkBlocks(t, blocks, serializeYerushalmi(yerushalmi.Vilna, "vilna"))
}

// TestDifferentialYerushalmiVilnaAfterRollover covers the range past
// 2172-08-08, where the original implementation panicked with "this code
// should be unreachable" and was thereafter permanently one daf behind, so
// these checksums intentionally differ from the original.
func TestDifferentialYerushalmiVilnaAfterRollover(t *testing.T) {
	blocks := []block{
		{793162, 802292, 9131, 0x5e59804d},
		{802293, 803533, 1241, 0xf9e75eac},
	}
	checkBlocks(t, blocks, serializeYerushalmi(yerushalmi.Vilna, "vilna"))
}

// TestFarFutureYearIsComputable exercises a date roughly a millennium out,
// which the walk-forward implementations resolved only after grinding through
// every intervening cycle.
func TestFarFutureYearIsComputable(t *testing.T) {
	assert := assert.New(t)
	hd := hdate.FromGregorian(2999, time.June, 15)

	daf, err := dafyomi.New(hd)
	assert.NoError(err)
	assert.NotEmpty(daf.Name)
	if r, ok := nine29.New(hd); ok {
		assert.Greater(r.CycleChap, 0)
	}
	assert.NotEqual(0, yerushalmi.New(hd, yerushalmi.Schottenstein).Blatt)
	assert.NotEmpty(yerushalmi.New(hd, yerushalmi.Vilna).Name)
	if r, ok := tanakhyomi.New(hd); ok {
		assert.NotEmpty(r.Blatt)
	}
}
