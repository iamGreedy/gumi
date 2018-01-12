package gumi

import (
	"fmt"
	"math"
)

var AUTOSIZE = Size{AUTOLENGTH, AUTOLENGTH}

//
type Size struct {
	Vertical   Length
	Horizontal Length
}

func (s *Size) VModify(vertical Length) Size {
	return Size{
		Vertical:   vertical,
		Horizontal: s.Horizontal,
	}
}
func (s *Size) HModify(horizontal Length) Size {
	return Size{
		Vertical:   s.Vertical,
		Horizontal: horizontal,
	}
}
func (s Size) String() string {
	ret := "Size("
	if s.Horizontal.Min == 0 {
		ret += "Horizontal.Min : Auto"
	} else {
		ret += fmt.Sprintf("Horizontal.Min : %d", s.Horizontal.Min)
	}
	ret += ", "
	if s.Horizontal.Max == math.MaxUint16 {
		ret += "Horizontal.Max : Auto"
	} else {
		ret += fmt.Sprintf("Horizontal.Max : %d", s.Horizontal.Max)
	}
	ret += ", "
	if s.Vertical.Min == 0 {
		ret += "Vertical.Max : Auto"
	} else {
		ret += fmt.Sprintf("Vertical.Min : %d", s.Vertical.Min)
	}
	ret += ", "
	if s.Vertical.Max == math.MaxUint16 {
		ret += "Vertical.Max : Auto"
	} else {
		ret += fmt.Sprintf("Vertical.Max : %d", s.Vertical.Max)
	}
	ret += ")"

	return ret
}
func FixedSize(h, v uint16) Size {
	return Size{
		Vertical:   FixedLength(v),
		Horizontal: FixedLength(h),
	}
}
func MinSize(h, v uint16) Size {
	return Size{
		Vertical:   Length{v, AUTOLENGTH.Max},
		Horizontal: Length{h, AUTOLENGTH.Max},
	}
}
func MaxSize(h, v uint16) Size {
	return Size{
		Vertical:   Length{AUTOLENGTH.Min, v},
		Horizontal: Length{AUTOLENGTH.Min, h},
	}
}
