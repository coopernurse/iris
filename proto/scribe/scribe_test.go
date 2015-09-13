// Iris - Decentralized cloud messaging
// Copyright (c) 2014 Project Iris. All rights reserved.
//
// Community license: for open source projects and services, Iris is free to use,
// redistribute and/or modify under the terms of the GNU Affero General Public
// License as published by the Free Software Foundation, either version 3, or (at
// your option) any later version.
//
// Evaluation license: you are free to privately evaluate Iris without adhering
// to either of the community or commercial licenses for as long as you like,
// however you are not permitted to publicly release any software or service
// built on top of it without a valid license.
//
// Commercial license: for commercial and/or closed source projects and services,
// the Iris cloud messaging system may be used in accordance with the terms and
// conditions contained in an individually negotiated signed written agreement
// between you and the author(s).

package scribe

import (
	"time"

	"github.com/coopernurse/iris/config"
)

// 512 bit RSA key in DER format
var privKeyDer = []byte{
	0x30, 0x82, 0x01, 0x39, 0x02, 0x01, 0x00, 0x02,
	0x41, 0x00, 0xbe, 0x89, 0x5d, 0x5c, 0xbe, 0x1d,
	0xef, 0xbc, 0x97, 0xab, 0xde, 0x90, 0xd2, 0x56,
	0xa1, 0xe2, 0x2f, 0x33, 0xb0, 0x4e, 0xdd, 0x54,
	0x97, 0x2b, 0xb8, 0xa8, 0xae, 0xfb, 0x11, 0x7c,
	0x7d, 0x8a, 0x9b, 0x22, 0x3e, 0xf3, 0xe4, 0xb5,
	0x1a, 0xe2, 0xed, 0xef, 0xc0, 0xaf, 0x8a, 0x6d,
	0xda, 0x6c, 0x81, 0x6e, 0x9a, 0xda, 0x36, 0x41,
	0x8b, 0xde, 0xdf, 0x6e, 0xef, 0x81, 0x91, 0x59,
	0x08, 0xb1, 0x02, 0x03, 0x01, 0x00, 0x01, 0x02,
	0x40, 0x0e, 0xf8, 0x41, 0xe2, 0x90, 0x79, 0x4f,
	0xa5, 0x94, 0x91, 0x07, 0x4a, 0x7f, 0x8c, 0x18,
	0xe9, 0xe9, 0x65, 0x79, 0x3b, 0xa8, 0xfe, 0x05,
	0x66, 0x84, 0xfa, 0x93, 0xcc, 0xdc, 0x01, 0xd8,
	0xe7, 0x11, 0x10, 0x4d, 0xee, 0x34, 0xf2, 0xbf,
	0x4d, 0xe9, 0xbb, 0x10, 0x26, 0x63, 0xbb, 0x33,
	0xe0, 0xdc, 0x16, 0x23, 0x58, 0x93, 0x44, 0x71,
	0xef, 0xd9, 0xb8, 0x4a, 0xe0, 0x56, 0x25, 0x60,
	0x55, 0x02, 0x21, 0x00, 0xf2, 0x6d, 0x07, 0x49,
	0x29, 0x10, 0xa2, 0xea, 0xb5, 0x12, 0x1e, 0xdf,
	0x14, 0x5b, 0x9d, 0xb4, 0x02, 0xe7, 0x9a, 0xc1,
	0x3d, 0xa9, 0xa7, 0x87, 0xc2, 0xe7, 0xee, 0x2b,
	0xc5, 0x3b, 0xca, 0x7f, 0x02, 0x21, 0x00, 0xc9,
	0x34, 0x8b, 0xea, 0x07, 0xd0, 0x35, 0x50, 0x6b,
	0xba, 0x96, 0x28, 0x5e, 0x86, 0x66, 0x15, 0x51,
	0xfa, 0xd2, 0x9e, 0x95, 0x67, 0x74, 0xc1, 0xec,
	0x71, 0x4c, 0x60, 0xee, 0xe1, 0xb4, 0xcf, 0x02,
	0x20, 0x13, 0x4d, 0x3f, 0x01, 0x42, 0x35, 0xc2,
	0xe2, 0xf1, 0x1b, 0xca, 0x3d, 0x74, 0xbf, 0x7e,
	0xa4, 0xf0, 0x7e, 0x44, 0x42, 0x12, 0x88, 0xc9,
	0x7f, 0xf3, 0xb2, 0xc7, 0xb1, 0xd0, 0x78, 0x5c,
	0x3d, 0x02, 0x20, 0x5b, 0xe2, 0x94, 0x56, 0xcf,
	0x34, 0xa5, 0x74, 0x51, 0x8e, 0x47, 0x4e, 0xae,
	0x44, 0x40, 0x50, 0x52, 0x3c, 0xf2, 0x7c, 0x9b,
	0x8c, 0x40, 0x84, 0xe3, 0x1e, 0xa6, 0x9b, 0xc9,
	0xdb, 0xe7, 0x7f, 0x02, 0x20, 0x75, 0x95, 0x8f,
	0xda, 0xf7, 0x42, 0x6d, 0x0a, 0x5f, 0xe5, 0x77,
	0x1e, 0x2a, 0xa9, 0xea, 0x21, 0x39, 0x4c, 0xcf,
	0x6b, 0xfe, 0x62, 0xd5, 0xd6, 0xa2, 0xd6, 0x35,
	0x19, 0x55, 0x63, 0x3a, 0xed,
}

// Id for connection filtering
var overId = "overlay.test"
var topicId = "topic.test"

// Configuration values for the pastry tests.
var bootTimeout = 500 * time.Millisecond
var convTimeout = 250 * time.Millisecond
var pastryLeaves = 4
var scribeBeat = 250 * time.Millisecond

func swapConfigs() {
	config.PastryBootTimeout, bootTimeout = bootTimeout, config.PastryBootTimeout
	config.PastryConvTimeout, convTimeout = convTimeout, config.PastryConvTimeout
	config.PastryLeaves, pastryLeaves = pastryLeaves, config.PastryLeaves
	config.ScribeBeatPeriod, scribeBeat = scribeBeat, config.ScribeBeatPeriod
}
