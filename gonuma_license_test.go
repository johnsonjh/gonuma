// Copyright 2020 Jeffrey H. Johnson.
// Copyright 2020 Gridfinity, LLC.
// Copyright 2012 The Go Authors.
// All rights reserved.
// Use of this source code is governed by the BSD-style
// license that can be found in the testutil/LICENSE file.

package gonuma

import (
	"fmt"
	"testing"

	u "github.com/johnsonjh/gonuma/testutil"
	licn "go4.org/legal"
)

func TestLicense(
	t *testing.T,
) {
	defer u.LeakVerifyNone(
		t,
	)
	licenses := licn.Licenses()
	if len(
		licenses,
	) == 0 {
		t.Fatal(
			"\ngonuma_license_test.TestLicense.Licenses FAILURE",
		)
	} else {
		t.Log(
			fmt.Sprintf(
				"\n%v\n",
				licenses,
			),
		)
	}
}
