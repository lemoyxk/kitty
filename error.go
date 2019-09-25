/**
* @program: lemo
*
* @description:
*
* @author: Mr.Wang
*
* @create: 2019-09-25 20:37
**/

package lemo

import (
	"runtime"
)

type Error struct {
	File  string
	Line  int
	Error error
}

func New(err error) func() *Error {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return nil
	}

	return func() *Error {
		return &Error{file, line, err}
	}
}
