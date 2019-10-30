package closedrange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var lower, upper int = 3, 8

func createBaseClosedrange() *Closedrange {
	return NewClosedrange(lower, upper)
}

func TestCreateClosedrange(t *testing.T) {
	// 初期化
	obj := createBaseClosedrange()

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

func TestString(t *testing.T) {
	// 初期化
	/*
		lower := 3
		upper := 8
		obj := NewClosedrange(lower, upper)
	*/
	obj := createBaseClosedrange()

	// 動作
	result := obj.String()

	// 検証
	assert.Equal(t, result, "[3,8]", "閉区間の文字列表現が正しい")
}

func TestInclude(t *testing.T) {
	// 初期化
	lower := 3
	upper := 8
	number := 5
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Include(number)

	// 検証
	assert.Equal(t, result, true, "与えた数字が閉区間に含まれる")
}

func TestNotInclude(t *testing.T) {
	// 初期化
	lower := 3
	upper := 8
	number := 0
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Include(number)

	// 検証
	assert.Equal(t, result, false, "与えた数字が閉区間に含まれない")
}

func TestIncludeUppperEqual(t *testing.T) {
	// 初期化
	lower := 3
	upper := 8
	number := 8
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Include(number)

	// 検証
	assert.Equal(t, result, true, "与えた数字が閉区間の上端と等しい")
}

func TestIncludeLowerEqual(t *testing.T) {
	// 初期化
	lower := 3
	upper := 8
	number := 3
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Include(number)

	// 検証
	assert.Equal(t, result, true, "与えた数字が閉区間の下端と等しい")
}

func TestEqual(t *testing.T) {
	// 初期化
	lower := 3
	upper := 8
	obj1 := NewClosedrange(lower, upper)

	obj2 := NewClosedrange(lower, upper)

	// 動作
	result := obj1.Equal(obj2)

	// 検証
	assert.Equal(t, result, true, "与えた閉区間が元の閉区間と等しい")
}

func TestNotEqual(t *testing.T) {
	// 初期化
	lower1 := 3
	upper1 := 8
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
	lower := 3
	upper := 8
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.Equal(nil)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間がNilである")
}
func TestIncludeClosedrangeNil(t *testing.T) {
	// 初期化
	lower := 3
	upper := 8
	obj := NewClosedrange(lower, upper)

	// 動作
	result := obj.IncludeClosedrange(nil)

	// 検証
	assert.Equal(t, result, false, "与えた閉区間がnilであり、含まれない")
}

type testcaseInclude struct {
	obj1     *Closedrange
	obj2     *Closedrange
	expected bool
}

func TestIncludeClosedrangePattern(t *testing.T) {
	// 初期化
	testcases := []*testcaseInclude{
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(1, 2), false},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(1, 3), false},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(1, 4), false},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(1, 8), false},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(1, 9), false},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(3, 3), true},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(3, 4), true},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(3, 8), true},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(3, 9), false},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(4, 5), true},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(4, 8), true},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(4, 9), false},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(8, 8), true},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(8, 9), false},
		&testcaseInclude{NewClosedrange(3, 8), NewClosedrange(9, 10), false},
	}

	// 動作
	for _, testcase := range testcases {
		obj1 := testcase.obj1
		obj2 := testcase.obj2
		expected := testcase.expected
		expected_str := ""
		if expected {
			expected_str = "included"
		} else {
			expected_str = "not included"
		}
		test_desc := fmt.Sprintf("%v should %s in %v", obj1, expected_str, obj2)

		t.Run(test_desc, func(t *testing.T) {
			//t.Parallel()
			result := obj1.IncludeClosedrange(obj2)

			assert.Equal(t, result, testcase.expected)
		})
	}

	// 検証
}
