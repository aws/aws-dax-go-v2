/*
  Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.

  Licensed under the Apache License, Version 2.0 (the "License").
  You may not use this file except in compliance with the License.
  A copy of the License is located at

      http://www.apache.org/licenses/LICENSE-2.0

  or in the "license" file accompanying this file. This file is distributed
  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
  express or implied. See the License for the specific language governing
  permissions and limitations under the License.
*/

package types

import "strings"

// Define constant string values for possible inputs of the user provided IpDiscovery
type IpDiscovery string

const (
	IpDiscoveryIPv4 IpDiscovery = "ipv4"
	IpDiscoveryIPv6 IpDiscovery = "ipv6"
)

// String implements fmt.Stringer interface

func (ipD IpDiscovery) String() string {
	return string(ipD)
}

// IsValid represents a validation function on the user-inserted value for ipDiscovery
// Returns bool true if the value matches "ipv4", "ipv6" or empty string regardless the capitalization. False, otherwise.
func (ipD IpDiscovery) IsValid() bool {
	s := ipD.String()
	return strings.EqualFold(IpDiscoveryIPv4.String(), s) ||
		strings.EqualFold(IpDiscoveryIPv6.String(), s) ||
		s == ""
}
