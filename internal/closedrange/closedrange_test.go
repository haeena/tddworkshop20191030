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

func TestIncludeUppperEqual(t *testing.T) {
	// 初期化
	lower := 1
	upper := 3
	number := 3
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Include(number)

	// 検証
	assert.Equal(t, result, true, "与えた数字が閉区間の上端と等しい")
}

func TestIncludeLowerEqual(t *testing.T) {
	// 初期化
	lower := 1
	upper := 3
	number := 1
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Include(number)

	// 検証
	assert.Equal(t, result, true, "与えた数字が閉区間の下端と等しい")
}

func TestEqual(t *testing.T) {
	// 初期化
	lower := 1
	upper := 3
	obj1 := NewClosedrange(lower, upper)

	obj2 := NewClosedrange(lower, upper)

	// 動作
	result := obj1.Equal(obj2)

	// 検証
	assert.Equal(t, result, true, "与えた閉区間が元の閉区間と等しい")
}

func TestNotEqual(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 3
	obj1 := NewClosedrange(lower1, upper1)

	lower2 := 1
	upper2 := 2
	obj2 := NewClosedrange(lower2, upper2)

	// 動作
	result := obj1.Equal(obj2)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間が元の閉区間と等しくない")
}

func TestNotEqualToNil(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 3
	obj1 := NewClosedrange(lower1, upper1)

	// 動作
	result := obj1.Equal(nil)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間がNilである")
}

func TestIncludeClosedrange1(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 5
	obj1 := NewClosedrange(lower1, upper1)

	lower2 := 1
	upper2 := 5
	obj2 := NewClosedrange(lower2, upper2)

	// 動作
	result := obj1.IncludeClosedrange(obj2)

	// 検証
	assert.Equal(t, result, true, "与えた閉区間が閉区間に含まれる")
}

func TestIncludeClosedrange2(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 5
	obj1 := NewClosedrange(lower1, upper1)

	lower2 := -1
	upper2 := 0
	obj2 := NewClosedrange(lower2, upper2)

	// 動作
	result := obj1.IncludeClosedrange(obj2)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間が閉区間に含まれない")
}

func TestIncludeClosedrange3(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 5
	obj1 := NewClosedrange(lower1, upper1)

	lower2 := 0
	upper2 := 2
	obj2 := NewClosedrange(lower2, upper2)

	// 動作
	result := obj1.IncludeClosedrange(obj2)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間が閉区間に含まれない")
}

func TestIncludeClosedrange4(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 5
	obj1 := NewClosedrange(lower1, upper1)

	lower2 := 2
	upper2 := 3
	obj2 := NewClosedrange(lower2, upper2)

	// 動作
	result := obj1.IncludeClosedrange(obj2)

	// 検証
	assert.Equal(t, result, true, "与えた閉区間が閉区間に含まれる")
}

func TestIncludeClosedrange5(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 5
	obj1 := NewClosedrange(lower1, upper1)

	lower2 := 2
	upper2 := 6
	obj2 := NewClosedrange(lower2, upper2)

	// 動作
	result := obj1.IncludeClosedrange(obj2)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間が閉区間に含まれない")
}

func TestIncludeClosedrange6(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 5
	obj1 := NewClosedrange(lower1, upper1)

	lower2 := 6
	upper2 := 7
	obj2 := NewClosedrange(lower2, upper2)

	// 動作
	result := obj1.IncludeClosedrange(obj2)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間が閉区間に含まれない")
}

func TestIncludeClosedrangeNil(t *testing.T) {
	// 初期化
	lower1 := 1
	upper1 := 5
	obj1 := NewClosedrange(lower1, upper1)

	// 動作
	result := obj1.IncludeClosedrange(nil)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間がnilであり、含まれない")
}
