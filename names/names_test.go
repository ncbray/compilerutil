package names

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitSnakeCase(t *testing.T) {
	assert.Equal(t, []string{}, SplitSnakeCase(""))
	assert.Equal(t, []string{"a"}, SplitSnakeCase("a"))
	assert.Equal(t, []string{"ab"}, SplitSnakeCase("ab"))
	assert.Equal(t, []string{"a", "b"}, SplitSnakeCase("a_b"))
	assert.Equal(t, []string{"abc"}, SplitSnakeCase("abc"))
	assert.Equal(t, []string{"a", "bc"}, SplitSnakeCase("a_bc"))
	assert.Equal(t, []string{"ab", "c"}, SplitSnakeCase("ab_c"))
	assert.Equal(t, []string{"a", "b", "c"}, SplitSnakeCase("a_b_c"))
}

func TestJoinSnakeCase(t *testing.T) {
	assert.Equal(t, "", JoinSnakeCase([]string{}, false))
	assert.Equal(t, "", JoinSnakeCase([]string{}, true))
	assert.Equal(t, "foo_bar", JoinSnakeCase([]string{"Foo", "Bar"}, false))
	assert.Equal(t, "FOO_BAR", JoinSnakeCase([]string{"Foo", "Bar"}, true))
}

func TestSplitCamelCase(t *testing.T) {
	assert.Equal(t, []string{}, SplitCamelCase(""))
	assert.Equal(t, []string{"a"}, SplitCamelCase("a"))
	assert.Equal(t, []string{"A"}, SplitCamelCase("A"))
	assert.Equal(t, []string{"Ab"}, SplitCamelCase("Ab"))
	assert.Equal(t, []string{"AB"}, SplitCamelCase("AB"))
	assert.Equal(t, []string{"a", "B"}, SplitCamelCase("aB"))
	assert.Equal(t, []string{"abc"}, SplitCamelCase("abc"))
	assert.Equal(t, []string{"Abc"}, SplitCamelCase("Abc"))
	assert.Equal(t, []string{"a", "Bc"}, SplitCamelCase("aBc"))
	assert.Equal(t, []string{"ab", "C"}, SplitCamelCase("abC"))
	assert.Equal(t, []string{"A", "Bc"}, SplitCamelCase("ABc"))
	assert.Equal(t, []string{"ABC"}, SplitCamelCase("ABC"))
	assert.Equal(t, []string{"AB", "Cd"}, SplitCamelCase("ABCd"))
	assert.Equal(t, []string{"abc", "123", "F"}, SplitCamelCase("abc123F"))
	assert.Equal(t, []string{"FF", "123"}, SplitCamelCase("FF123"))
	assert.Equal(t, []string{"F", "Fabc"}, SplitCamelCase("FFabc"))
}

func TestJoinCamelCase(t *testing.T) {
	assert.Equal(t, "", JoinCamelCase([]string{}, false))
	assert.Equal(t, "", JoinCamelCase([]string{}, true))
	assert.Equal(t, "fooBar", JoinCamelCase([]string{"foo", "bar"}, false))
	assert.Equal(t, "FooBar", JoinCamelCase([]string{"foo", "bar"}, true))
	assert.Equal(t, "99Problems", JoinCamelCase([]string{"99", "problems"}, false))
	assert.Equal(t, "99Problems", JoinCamelCase([]string{"99", "Problems"}, true))
	assert.Equal(t, "xmlHttpRequest", JoinCamelCase([]string{"XML", "Http", "Request"}, false))
	assert.Equal(t, "XMLHttpRequest", JoinCamelCase([]string{"XML", "Http", "Request"}, true))
}
