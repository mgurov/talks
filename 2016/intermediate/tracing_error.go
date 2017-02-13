package main

import (
	"errors"

	"github.com/emicklei/tre"
)

// START OMIT
func main() {
	if err := doSomething("demo"); err != nil {
		println(tre.New(err, "doSomething failed").Error())
	}
}

func doSomething(with string) error {
	if err := somethingFaulty(with); err != nil {
		return tre.New(err, "somethingFaulty failed", "with", with) // pass error, message and context
	}
	return nil
}

func somethingFaulty(_ string) error {
	return errors.New("something bad has just happened")
}

// END OMIT
