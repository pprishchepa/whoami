package random_test

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/pprishchepa/whoami/internal/random"
	"github.com/stretchr/testify/assert"
)

func TestUniformWrite(t *testing.T) {
	//nolint:staticcheck
	//goland:noinspection GoDeprecation
	rand.Seed(1)
	random.Randomize(8) // randomize default charset with length of 8 bytes

	var buf bytes.Buffer

	buf.Reset()
	random.Write(&buf, 10) // write rand 10 bytes
	assert.Len(t, buf.String(), 10)
	//goland:noinspection SpellCheckingInspection
	assert.Equal(t, "BpLnfgDsBp", buf.String())

	buf.Reset()
	random.Write(&buf, 5) // write rand 5 bytes
	assert.Len(t, buf.String(), 5)
	//goland:noinspection SpellCheckingInspection
	assert.Equal(t, "fgDsB", buf.String())
}
