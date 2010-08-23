# Copyright 2010 Nigel Gourlay. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.


include $(GOROOT)/src/Make.$(GOARCH)

TARG=yes
GOFILES=yes.go

include $(GOROOT)/src/Make.cmd
