// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Internet protocol family sockets

// Package ipsupport exposes features of the platform's networking
// functionality using code extraced from the standard library's
// net package.
package ipsupport

var (
	// supportsIPv4 reports whether the platform supports IPv4
	// networking functionality.
	supportsIPv4 bool

	// supportsIPv6 reports whether the platform supports IPv6
	// networking functionality.
	supportsIPv6 bool

	// supportsIPv4map reports whether the platform supports
	// mapping an IPv4 address inside an IPv6 address at transport
	// layer protocols.  See RFC 4291, RFC 4038 and RFC 3493.
	supportsIPv4map bool
)

func init() {
	supportsIPv4 = probeIPv4Stack()
	supportsIPv6, supportsIPv4map = probeIPv6Stack()
}

// V4 reports whether the platform supports IPv4
// networking functionality.
func V4() bool {
	return supportsIPv4
}

// V6 reports whether the platform supports IPv6
// networking functionality.
func V6() bool {
	return supportsIPv6
}

// V4Map reports whether the platform supports mapping an IPv4
// address inside an IPv6 address at transport layer protocols.
// See RFC 4291, RFC 4038 and RFC 3493.
func V4Map() bool {
	return supportsIPv4map
}
