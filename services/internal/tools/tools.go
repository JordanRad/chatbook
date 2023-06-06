//go:build tools
// +build tools

package tools

import (
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
	_ "goa.design/goa/v3/cmd/goa"
	_ "goa.design/goa/v3/codegen"
	_ "goa.design/goa/v3/codegen/generator"
)

// This file imports packages that are used when running go generate, or used
// during the development process but not otherwise depended on by built code.
