package logging

import (
	"bytes"
	"log"
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

func TestDefaultDebugShouldLogWithDebugLevel(t *testing.T) {
	a := assert.NewAssert(t)
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	Debug("did that")

	a.Match("^\\d\\d\\d\\d\\/\\d\\d\\/\\d\\d \\d\\d:\\d\\d:\\d\\d logging_test\\.go:\\d+: \\[DEBUG\\] did that\n$", buf.String(), "Unexpected log output")
}

func TestDefaultInfoShouldLogWithInfoLevel(t *testing.T) {
	a := assert.NewAssert(t)
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	Info("did it")

	a.Match("^\\d\\d\\d\\d\\/\\d\\d\\/\\d\\d \\d\\d:\\d\\d:\\d\\d logging_test\\.go:\\d+: \\[INFO\\] did it\n$", buf.String(), "Unexpected log output")
}

func TestDefaultWarnShouldLogWithWarnLevel(t *testing.T) {
	a := assert.NewAssert(t)
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	Warn("warning")

	a.Match("^\\d\\d\\d\\d\\/\\d\\d\\/\\d\\d \\d\\d:\\d\\d:\\d\\d logging_test\\.go:\\d+: \\[WARN\\] warning\n$", buf.String(), "Unexpected log output")
}

func TestDefaultErrorShouldLogWithErrorLevel(t *testing.T) {
	a := assert.NewAssert(t)
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	Error("erroaarrr!!")

	a.Match("^\\d\\d\\d\\d\\/\\d\\d\\/\\d\\d \\d\\d:\\d\\d:\\d\\d logging_test\\.go:\\d+: \\[ERROR\\] erroaarrr!!\n$", buf.String(), "Unexpected log output")
}
