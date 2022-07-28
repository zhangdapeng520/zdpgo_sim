package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhangdapeng520/zdpgo_sim/chroma"
)

func TestClosestColour(t *testing.T) {
	actual := findClosest(ttyTables[256], chroma.MustParseColour("#e06c75"))
	assert.Equal(t, chroma.MustParseColour("#d75f87"), actual)
}
