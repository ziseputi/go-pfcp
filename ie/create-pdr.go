// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewCreatePDR creates a new CreatePDR IE.
func NewCreatePDR(pdr, precedence, pdi, hdrRemoval, far, urr, qer, apr, atime, dtime, mar, prdInfo, mcInfo, pool *IE) *IE {
	return newGroupedIE(CreatePDR, 0, pdr, precedence, pdi, hdrRemoval, far, urr, qer, apr, atime, dtime, mar, prdInfo, mcInfo, pool)
}

// CreatePDR returns the IEs above CreatePDR if the type of IE matches.
func (i *IE) CreatePDR() ([]*IE, error) {
	if i.Type != CreatePDR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
