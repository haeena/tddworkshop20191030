package closedrange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClosedrange(t *testing.T) {
	// 初期化
	lower := 1
	upper := 1
	obj := NewClosedrange(lower, upper)
	// 動作

	// 検証
	assert.NotNil(t, obj, "作成されたObjectがNil")
	assert.Equal(t, obj.lower, lower, "作成されたObjectのLowerが異なる")
	assert.Equal(t, obj.upper, upper, "作成されたObjectのUpperが異なる")
	// 掃除

}
