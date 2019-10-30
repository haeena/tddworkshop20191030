package closedrange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClosedrange(t *testing.T) {
	// 初期化
	lower := 0
	upper := 0
	obj := NewClosedrange(lower, upper)
	// 動作
	// 検証
	assert.NotNil(t, obj, "作成されたObjectがNil")

	// 掃除

}
