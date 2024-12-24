/*
	   A template for adding descriptive error reports to a command line interface
																		 - kendfss
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	noError ErrorCode = iota
	internalError
	userInputError
)

type ErrorCode int

func (this ErrorCode) Error() string {
	return []string{"NO ERROR", "INTERNAL ERROR", "USER INPUT ERROR"}[this]
}

func (this ErrorCode) Abort(message error) {
	this.Warn(message)
	if this == userInputError {
		flag.Usage()
	}
	os.Exit(int(this))
}

func (this ErrorCode) Abortf(message string, args ...interface{}) {
	this.Warnf(message, args...)
	if this == userInputError {
		flag.Usage()
	}
	os.Exit(int(this))
}

func (this ErrorCode) Warn(message error) {
	if this >= internalError {
		stderr.Println(fmt.Errorf("[%s]: %s\n", this.Error(), message))
	}
}

func (this ErrorCode) Warnf(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	if this >= internalError {
		stderr.Println(fmt.Errorf("[%s]: %s\n", this.Error(), message))
	}
}
