// Code generated by "stringer -type=Type"; DO NOT EDIT.

package etensor

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _Type_name = "NULLBOOlUINT8INT8UINT16INT16UINT32INT32UINT64INT64FLOAT16FLOAT32FLOAT64STRINGCOMPLEX64COMPLEX128"

var _Type_index = [...]uint8{0, 4, 8, 13, 17, 23, 28, 34, 39, 45, 50, 57, 64, 71, 77, 86, 96}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}

func (i *Type) FromString(s string) error {
	for j := 0; j < len(_Type_index)-1; j++ {
		if s == _Type_name[_Type_index[j]:_Type_index[j+1]] {
			*i = Type(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Type")
}
