package sqs

import (
	"fmt"
	"go-sqs-worker/sqsworker"
)

func Recover() sqsworker.MiddlewareFunc {
	return func(next sqsworker.HandlerFunc) sqsworker.HandlerFunc {
		return func(c sqsworker.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
