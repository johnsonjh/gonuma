// Copyright 2020 Jeffrey H. Johnson.
// Copyright 2020 Gridfinity, LLC.
// Copyright 2012 The Go Authors.
// All rights reserved.

package gonuma // import "go.gridfinity.dev/gonuma"

import (
	"fmt"
	"testing"

	u "go.gridfinity.dev/leaktestfe"
	licn "go4.org/legal"
)

func TestLicense(
	t *testing.T,
) {
	defer u.Leakplug(
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
