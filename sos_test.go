// Copyright 2014 Tom Grennan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sos_test

import (
	"github.com/tgrennan/sos"
	"testing"
)

func Test(t *testing.T) {
	sos := sos.New("sos", "-a", "A", "--b=B", "-c", "-d", "--e",
		"-t", "NAME", "VALUE", "X", "Y", "Z")
	var prog, aS, bS, cS, tName, tValue string
	var b, c, d, e, f bool
	if sos, prog = sos.Pop(); prog != "sos" {
		t.Fatal("prog:", prog)
	}
	if sos, aS = sos.Arg("a"); aS != "A" {
		t.Fatal("aS:", aS)
	}
	if sos, b = sos.Flag("b"); !b {
		if sos, bS = sos.Arg("b"); bS != "B" {
			t.Fatal("bS:", bS)
		}
	} else {
		t.Fatal("b:", b)
	}
	if sos, c = sos.Flag("c"); !c {
		t.Fatal("c:", c)
	}
	if sos, cS = sos.Arg("c"); cS != "" {
		t.Fatal("cS:", cS)
	}
	if sos, d = sos.Flag("d"); !d {
		t.Fatal("d:", d)
	}
	if sos, e = sos.Flag("e"); !e {
		t.Fatal("e:", e)
	}
	if sos, f = sos.Flag("f"); f {
		t.Fatal("f:", f)
	}
	if sos, tName, tValue = sos.Ternary("t"); tName != "NAME" {
		t.Fatal("tName:", tName)
	} else if tValue != "VALUE" {
		t.Fatal("tValue:", tValue)
	}
	if i := sos.Mismatch("X", "Y", "Z"); i >= 0 {
		t.Fatal("mismatch at", i, "of:", sos)
	}
	if i := sos.Index("Y"); i != 1 {
		t.Fatal("index of Y:", i)
	}
	if s := sos.Join(" "); s != "X Y Z" {
		t.Fatal(`Join(" "):`, s)
	}
	sos = sos.Push("push")
	if i := sos.Mismatch("push", "X", "Y", "Z"); i >= 0 {
		t.Fatal("mismatch at", i, "of:", sos)
	}
	var popped string
	sos, popped = sos.Pop()
	if i := sos.Mismatch("X", "Y", "Z"); i >= 0 {
		t.Fatal("mismatch at", i, "of:", sos)
	}
	if popped != "push" {
		t.Fatal("popped:", popped)
	}
	sos = sos.Remove(1, 1)
	if n := sos.Len(); n != 2 {
		t.Fatal("sos.Len():", n)
	}
	if i := sos.Mismatch("X", "Z"); i >= 0 {
		t.Fatal("mismatch at", i, "of:", sos)
	}
	sos = sos.Insert(1, "Y")
	if i := sos.Mismatch("X", "Y", "Z"); i >= 0 {
		t.Fatal("mismatch at", i, "of:", sos)
	}
}
