// Code generated by "stringer -type=lobTypecode"; DO NOT EDIT.

package protocol

import "strconv"

const _lobTypecode_name = "ltcUndefinedltcBlobltcClobltcNclob"

var _lobTypecode_index = [...]uint8{0, 12, 19, 26, 34}

func (i lobTypecode) String() string {
	if i < 0 || i >= lobTypecode(len(_lobTypecode_index)-1) {
		return "lobTypecode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _lobTypecode_name[_lobTypecode_index[i]:_lobTypecode_index[i+1]]
}
