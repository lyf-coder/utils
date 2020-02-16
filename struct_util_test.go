/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package utils

import (
	"log"
	"testing"
)

type People struct {
	Name  string
	Age   int
}

type Student struct {
	People
	Grade int
}

func TestCopy(t *testing.T) {
	s := Student{}
	s.Name = "zhangsan"
	s.Age = 20
	s.Grade = 100

	p := s.People
	log.Println(p.Name)
}

