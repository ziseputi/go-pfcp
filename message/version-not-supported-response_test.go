// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestVersionNotSupportedResponse(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured:  message.NewVersionNotSupportedResponse(),
			Serialized: []byte{
				0x20, 0x0b, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseVersionNotSupportedResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
