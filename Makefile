# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include Make.$(GOARCH)

TARG=gohaml
GOFILES=\
	haml.go\
	states.go\
	stack.go\
	tree.go\

include Make.pkg
