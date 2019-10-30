package closedrange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClosedrange(t *testing.T) {
	// 初期化
	lower := 0
	upper := 1
	obj := NewClosedrange(lower, upper)
	// 動作

	// 検証
	assert.NotNil(t, obj, "作成されたObjectがNil")
	assert.Equal(t, obj.lower, lower, "作成されたObjectのLowerが異なる")
	assert.Equal(t, obj.upper, upper, "作成されたObjectのUpperが異なる")
	// 掃除
}

func TestCheckLowerUpper(t *testing.T) {
	// 初期化
	lower := 2
	upper := 1
	obj := NewClosedrange(lower, upper)
	// 動作

	// 検証
	assert.Nil(t, obj, "作成されたObjectがNilじゃない")
}

func TestLowerUpperEqual(t *testing.T) {
	// 初期化
	lower := 1
	upper := 1
	obj := NewClosedrange(lower, upper)
	// 動作

	// 検証
	assert.NotNil(t, obj, "作成されたObjectがNilじゃない")
}

func TestInclude(t *testing.T) {
	// 初期化
	lower := 1
	upper := 3
	number := 2
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Include(number)

	// 検証
	assert.Equal(t, result, true, "与えた数字が閉区間に含まれる")
}

func TestNotInclude(t *testing.T) {
	// 初期化
	lower := 1
	upper := 3
	number := 0
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Include(number)

	// 検証
	assert.Equal(t, result, false, "与えた数字が閉区間に含まれない")
}
