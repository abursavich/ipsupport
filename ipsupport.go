// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ipsupport exposes features of the platform's networking
// functionality using code extracted from the standard library's
// net package.
package ipsupport

import "net"

// OK returns a version of the IP that the platform supports.
// If it is not supported it returns nil.
func OK(ip net.IP) net.IP {
	v4 := ip.To4()
	if supportsIPv4 && v4 != nil {
		return v4
	}
	if supportsIPv6 && len(ip) == net.IPv6len && v4 == nil {
		return ip
	}
	return nil
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
