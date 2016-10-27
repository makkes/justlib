package justlib

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/justsocialapps/assert"
)

func TestMain(m *testing.M) {
	flag.Parse()
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestTryZeroTimesShouldNotCallFunctionAtAllAndReturnNil(t *testing.T) {
	a := assert.NewAssert(t)
	err := Try(0, 0, func() error {
		t.Fatal("f has been called")
		return nil
	})
	a.Equal(err, nil, "Try returned with an error")
}

func TestTryTwoTimesShouldCallPassingFunctionOnlyOnce(t *testing.T) {
	a := assert.NewAssert(t)
	called := 0
	err := Try(2, 0, func() error {
		called++
		return nil
	})
	a.Equal(called, 1, fmt.Sprintf("f has been called %d times", called))
	a.Equal(err, nil, "Try returned with an error")
}

func TestTryTwoTimesShouldCallFailingFunctionTwoTimes(t *testing.T) {
	a := assert.NewAssert(t)
	called := 0
	err := Try(2, 0, func() error {
		called++
		if called == 1 {
			return errors.New("Warning. Warp reactor core primary coolant failure.")
		}
		return nil
	})
	a.Equal(called, 2, fmt.Sprintf("f has been called %d times", called))
	a.Equal(err, nil, "Try returned with an error")
}
