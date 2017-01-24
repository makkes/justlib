package logging

import (
	"testing"

	"github.com/justsocialapps/assert"
)

func TestLevelFromStringShouldBeCaseInsensitive(t *testing.T) {
	a := assert.NewAssert(t)
	l, _ := LevelFromString("debug")
	a.Equal(*l, DEBUG, "Wrong level returned")

	l, _ = LevelFromString("DEBUG")
	a.Equal(*l, DEBUG, "Wrong level returned")
}

func TestLevelFromStringShouldReturnNilWithUnknownString(t *testing.T) {
	a := assert.NewAssert(t)
	l, _ := LevelFromString("oink")
	a.Nil(l, "Level is not nil")
}

func TestLevelFromStringShouldReturnNilWithEmptyString(t *testing.T) {
	a := assert.NewAssert(t)
	l, _ := LevelFromString("")
	a.Nil(l, "Level is not nil")
}
