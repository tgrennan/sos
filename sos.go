// Copyright 2014 Tom Grennan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sos manipulates a slice of strings.
// Use it to process command line line arguments like this:
//
//	$ sos -a A --b=B -c -d --e -t NAME VALUE X Y Z
//
// ...
//	var prog, aS, bS, cS, tName, tValue string
//	var b, c, d, e, f bool
//	sos := sos.New(os.Args[:]...)
//	sos, prog = sos.Pop()
//	sos, aS = sos.Arg("a")
//	if sos, b = sos.Flag("b"); !b {
//		sos, bS = sos.Arg("b")
//	}
//	if sos, c = sos.Flag("c"); !c {
//		sos, cS = sos.Arg("c")
//	}
//	sos, d = sos.Flag("d")
//	sos, e = sos.Flag("e")
//	sos, f = sos.Flag("f")
//	sos, tName, tValue = sos.Ternary("t")
//
// Results:
//	prog == "sos"
//	aS == "A"
//	b == false
//	bS == "B"
//	c == true
//	cS == ""
//	d == true
//	e == true
//	f == false
//	tName == "NAME"
//	tValue == "VALUE"
//	sos.Mismatch("X", "Y", "Z") == -1
package sos

import "strings"

type SoS []string

// Create a Slice of Strings from variadic strings.
func New(a ...string) SoS {
	sos := make([]string, 0)
	sos = append(sos, a...)
	return (SoS)(sos)
}

// Returns and strips the argument of the matching string.
func (sos SoS) Arg(flag string) (SoS, string) {
	var arg string
	for i, s := range sos {
		if strings.HasPrefix(s, "-") {
			if t := strings.TrimLeft(s, "-"); t == flag {
				arg = sos.String(i + 1)
				sos = sos.Remove(i, 2)
				break
			} else if iequal := strings.Index(t, "="); iequal > 0 {
				if t[:iequal] == flag {
					arg = t[iequal+1:]
					sos = sos.Remove(i, 1)
					break
				}
			}
		}
	}
	return sos, arg
}

// Returns and strips the boolean flag of the matching string.
func (sos SoS) Flag(flag string) (SoS, bool) {
	var found bool
	for i, s := range sos {
		if strings.HasPrefix(s, "-") {
			if t := strings.TrimLeft(s, "-"); t == flag {
				found = true
				sos = sos.Remove(i, 1)
				break
			}
		}
	}
	return sos, found
}

// Returns the index of the first instance of s,
// or -1 if not found.
func (sos SoS) Index(s string) int {
	for i, x := range sos {
		if x == s {
			return i
		}
	}
	return -1
}

// Insert a slice at the given index.
func (sos SoS) Insert(i int, slice ...string) SoS {
	if i < 0 || i >= sos.Len() {
		sos = append(sos, slice[:]...)
	} else {
		sos = append(sos[:i], append(slice, (sos[i:])...)...)
	}
	return sos
}

// Concatenate the slice with the given separator.
func (sos SoS) Join(sep string) string {
	return strings.Join(sos, sep)
}

// Length of the appointed slice.
func (sos SoS) Len() int {
	return len(sos)
}

// Returns index of first mismatched string,
// or -1 if all match.
func (sos SoS) Mismatch(slice ...string) int {
	for i, s := range slice {
		if s != sos.String(i) {
			return i
		}
	}
	return -1
}

// Return and remove first string in slice.
// The returned string is empty if slice is empty.
func (sos SoS) Pop() (SoS, string) {
	if len(sos) > 0 {
		return sos[1:], sos[0]
	}
	return sos, ""
}

// Insert string(s) at SoS beginning.
func (sos SoS) Push(a ...string) SoS {
	na := len(a)
	nsos := len(sos)
	if nsos == 0 {
		sos = append(sos, a...)
	} else {
		sos = append(sos, a...)
		copy(sos[na:], sos[:nsos])
		copy(sos[0:na], a)
	}
	return sos
}

// Remove the slice of strings at the given index, length.
func (sos SoS) Remove(i, n int) SoS {
	if l := sos.Len(); i >= 0 && i < l && n >= 1 {
		if i+n < l {
			sos = append(sos[:i], (sos[i+n:])...)[:l-n]
		} else {
			sos = sos[:i]
		}
	}
	return sos
}

// Returns the slice of strings at the given index.
// sos.Slice(0, -1) returns a copy of the entire slice.
func (sos SoS) Slice(i, n int) []string {
	if l := sos.Len(); i >= 0 && i < l {
		if n < 0 || n > l {
			n = l
		}
		dst := make([]string, n)
		copy(dst, sos[i:n])
		return dst
	}
	return []string{}
}

// Returns the string at the given index,
// or an empty string if index is out of range.
func (sos SoS) String(i int) string {
	if i >= 0 && i < len(sos) {
		return sos[i]
	}
	return ""
}

// Returns and strips the paired arguments of the matching string.
func (sos SoS) Ternary(flag string) (SoS, string, string) {
	var name, value string
	for i, s := range sos {
		if strings.HasPrefix(s, "-") {
			if t := strings.TrimLeft(s, "-"); t == flag {
				name = sos.String(i + 1)
				value = sos.String(i + 2)
				sos = sos.Remove(i, 3)
				break
			}
		}
	}
	return sos, name, value
}
