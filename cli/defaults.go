package cli

import "go.nandlabs.io/golly/l3"

var (
	logger = l3.Get()
)

type ActionFunc func(conTxt *Context) error
