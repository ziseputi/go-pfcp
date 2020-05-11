// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// SteeringMode definitions.
const (
	SteeringModeActiveStandby uint8 = 0
	SteeringModeSmallestDelay uint8 = 1
	SteeringModeLoadBalancing uint8 = 2
	SteeringModePriorityBased uint8 = 3
)

// NewSteeringMode creates a new SteeringMode IE.
func NewSteeringMode(mode uint8) *IE {
	return newUint8ValIE(SteeringMode, mode&0x0f)
}

// SteeringMode returns SteeringMode in uint8 if the type of IE matches.
func (i *IE) SteeringMode() (uint8, error) {
	if i.Type != SteeringMode {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
