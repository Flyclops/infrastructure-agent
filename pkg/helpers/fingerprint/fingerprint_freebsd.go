// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
package fingerprint

// TODO: porting over darwin but need to investigate how to get a boot id from
// freebsd the way the linux version of this file does
func GetBootId() string {
	return ""
}
