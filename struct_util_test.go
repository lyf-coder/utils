/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package utils

import (
	"testing"
)

type People struct {
	Name string
	Age  int
}

type Student struct {
	Name  string
	Grade int
}

func TestCopyStruct(t *testing.T) {
	s := Student{}
	s.Name = "jack"
	s.Grade = 100
	p := People{}
	CopyStruct(s, &p)
	if p.Name != s.Name {
		t.Error("CopyStruct fail!")
	}
}
