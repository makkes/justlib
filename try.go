package justlib

import (
	"log"
	"time"
)

// Try calls the given function f at most »times« times as long as it returns
// an error, sleeping the given backoffTime in between subsequent calls.  If f
// returns an error on the last try, that error is returned by Try.
func Try(times uint16, backoffTime time.Duration, f func() error) error {
	if times == 0 {
		return nil
	}
	var err error
	var i uint16
	for i = 1; ; i++ {
		err = f()
		if err == nil {
			return nil
		}
		if i >= times {
			return err
		}
		log.Printf("Error at try %d: %v. Trying again in %v. %d tries left", i, err, backoffTime, times-i)
		time.Sleep(backoffTime)
	}
}
