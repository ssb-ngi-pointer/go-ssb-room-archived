// SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
//
// SPDX-License-Identifier: MIT

package aliases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	a := assert.New(t)

	cases := []struct {
		alias string
		valid bool
	}{
		{"basic", true},
		{"no spaces", false},
		{"no.dots", false},
		{"#*!(! nope", false},

		{"NoUpperCase", false},

		// too long
		{"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", false},
	}

	for i, tc := range cases {
		yes := IsValid(tc.alias)
		a.Equal(tc.valid, yes, "wrong for %d: %s", i, tc.alias)
	}
}
