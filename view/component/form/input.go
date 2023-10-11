package form

import (
	"strings"
)

type WrapOption int // TextArea option

const (
	SoftWrap WrapOption = iota
	HardWrap
)

func MarshalWrapOption(wrapOption string) WrapOption {
	switch strings.ToLower(wrapOption) {
	case HardWrap.String(): // HardWrap must have ColsAttribute defined
		return HardWrap
	default: // Soft is the default
		return SoftWrap
	}
}

func (self WrapOption) String() string {
	switch self {
	case HardWrap:
		return "hard"
	default:
		return "soft"
	}
}
