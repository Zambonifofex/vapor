// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

package vpmatrix

import (
	"math"
	"ufoot.org/vapor/vpnumber"
)

// X64Vec3 is a vector containing 3 fixed point 64 bit values.
// Can hold the values of a point in space.
type X64Vec3 [3]vpnumber.X64

// X64Vec3New creates a new vector containing 3 fixed point 64 bit values.
func X64Vec3New(x1, x2, x3 vpnumber.X64) *X64Vec3 {
	return &X64Vec3{x1, x2, x3}
}

// Add adds operand to the vector.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec3) Add(op *X64Vec3) *X64Vec3 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec3) Sub(op *X64Vec3) *X64Vec3 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec3) MulScale(factor vpnumber.X64) *X64Vec3 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Mul(vec[i], v)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec3) DivScale(factor vpnumber.X64) *X64Vec3 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Div(vec[i], v)
	}

	return vec
}

// SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X64Vec3) SumSq() vpnumber.X64 {
	var sq vpnumber.X64

	for _, v := range vec {
		sq += vpnumber.X64Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X64Vec3) Length() vpnumber.X64 {
	return vpnumber.F64ToX64(math.Sqrt(vpnumber.X64ToF64(vec.SumSq())))
}

// Normalize scales the vector so that its length is 1.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec3) Normalize() *X64Vec3 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *X64Vec3) IsSimilar(op *X64Vec3) bool {
	ret:=true
	for i, v := range vec {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}
	
	return ret
}

// X64Vec3Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3Add(veca, vecb *X64Vec3) *X64Vec3 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X64Vec3Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3Sub(veca, vecb *X64Vec3) *X64Vec3 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X64Vec3MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3MulScale(vec *X64Vec3, factor vpnumber.X64) *X64Vec3 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X64Vec3DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3DivScale(vec *X64Vec3, factor vpnumber.X64) *X64Vec3 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X64Vec3SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func X64Vec3SumSq(vec *X64Vec3) vpnumber.X64 {
	return vec.SumSq()
}

// X64Vec3Length returns the length of a vector.
func X64Vec3Length(vec *X64Vec3) vpnumber.X64 {
	return vec.Length()
}

// X64Vec3Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X64Vec3Normalize(vec *X64Vec3) *X64Vec3 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// X64Vec3IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func X64Vec3IsSimilar(veca,vecb *X64Vec3) bool {
	return veca.IsSimilar(vecb)
}
