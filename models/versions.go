package models

import (
	"fmt"
	"sort"
	"strconv"
)

type Versions []VersionString

func (v Versions) Latest() string {
	sort.Sort(v)
	return string(v[0])
}

func (v Versions) Len() int      { return len(v) }
func (v Versions) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

func (v Versions) Less(i, j int) bool {
	iRunes := []rune(v[i])
	jRunes := []rune(v[j])

	max := len(iRunes)
	if max < len(jRunes) {
		max = len(jRunes)
	}

	for idx := 0; idx < max; idx++ {
		ir := iRunes[idx]
		jr := jRunes[idx]

		if ir != jr {
			return ir > jr
		}
	}

	return false
}

// See https://wiki.factorio.com/Version_string_format
type VersionString string

// Bit shift to the right + AND with a full uint16 to extract the binary info from the uint64.
func (v *VersionString) UnmarshalJSON(data []byte) error {
	ver, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}

	major := uint16(ver >> 48 & 65535)
	minor := uint16(ver >> 32 & 65535)
	hotfix := uint16(ver >> 16 & 65535)
	//dev := uint16(ver & 65535)

	// Set self to the string resulting from the unmarshalling.
	*v = VersionString(fmt.Sprintf("%d.%d.%d", major, minor, hotfix))

	return nil
}
