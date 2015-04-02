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
// Vapor homepage: https://github.com/ufoot/vapor
// Contact author: ufoot@ufoot.org

package vpmatrix3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpmatrix2"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
)

// X32Mat3x3 is a matrix containing 3x3 fixed point 32 bit values.
// Can hold the values of a point in space.
type X32Mat3x3 [9]vpnumber.X32

// X32Mat3x3New creates a new matrix containing 3x3 fixed point 32 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X32Mat3x3New(x1, x2, x3, x4, x5, x6, x7, x8, x9 vpnumber.X32) *X32Mat3x3 {
	return &X32Mat3x3{x1, x2, x3, x4, x5, x6, x7, x8, x9}
}

// X32Mat3x3Identity creates a new identity matrix.
func X32Mat3x3Identity() *X32Mat3x3 {
	return &X32Mat3x3{vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1}
}

// X32Mat3x3Trans creates a new translation matrix.
func X32Mat3x3Trans(vec *vpmatrix2.X32Vec2) *X32Mat3x3 {
	return &X32Mat3x3{vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0, vec[0], vec[1], vpnumber.X32Const1}
}

// X32Mat3x3Rot creates a new rotation matrix.
// The rotation is done in 2D over a virtual z axis, such as z = cross(x,y).
// Angle is given in radians.
func X32Mat3x3Rot(r vpnumber.X32) *X32Mat3x3 {
	return &X32Mat3x3{vpmath.X32Cos(r), vpmath.X32Sin(r), vpnumber.X32Const0, -vpmath.X32Sin(r), vpmath.X32Cos(r), vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *X32Mat3x3) ToI32() *I32Mat3x3 {
	var ret I32Mat3x3

	for i, v := range mat {
		ret[i] = vpnumber.X32ToI32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int32 matrix.
func (mat *X32Mat3x3) ToI64() *I64Mat3x3 {
	var ret I64Mat3x3

	for i, v := range mat {
		ret[i] = vpnumber.X32ToI64(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *X32Mat3x3) ToX64() *X64Mat3x3 {
	var ret X64Mat3x3

	for i, v := range mat {
		ret[i] = vpnumber.X32ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *X32Mat3x3) ToF32() *F32Mat3x3 {
	var ret F32Mat3x3

	for i, v := range mat {
		ret[i] = vpnumber.X32ToF32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *X32Mat3x3) ToF64() *F64Mat3x3 {
	var ret F64Mat3x3

	for i, v := range mat {
		ret[i] = vpnumber.X32ToF64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *X32Mat3x3) Set(col, row int, val vpnumber.X32) {
	mat[col*3+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *X32Mat3x3) Get(col, row int) vpnumber.X32 {
	return mat[col*3+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *X32Mat3x3) MarshalJSON() ([]byte, error) {
	var tmpArray [3][3]int32

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = int32(mat[col*3+row])
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal X32Mat3x3")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *X32Mat3x3) UnmarshalJSON(data []byte) error {
	var tmpArray [3][3]int32

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal X32Mat3x3")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*3+row] = vpnumber.X32(tmpArray[col][row])
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *X32Mat3x3) String() string {
	buf, err := mat.ToF32().MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat3x3) Add(op *X32Mat3x3) *X32Mat3x3 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat3x3) Sub(op *X32Mat3x3) *X32Mat3x3 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat3x3) MulScale(factor vpnumber.X32) *X32Mat3x3 {
	for i, v := range mat {
		mat[i] = vpnumber.X32Mul(v, factor)
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat3x3) DivScale(factor vpnumber.X32) *X32Mat3x3 {
	for i, v := range mat {
		mat[i] = vpnumber.X32Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *X32Mat3x3) IsSimilar(op *X32Mat3x3) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.X32IsSimilar(v, op[i])
	}

	return ret
}

// Transpose inverts rows and columns (matrix transposition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat3x3) Transpose(op *X32Mat3x3) *X32Mat3x3 {
	*mat = *X32Mat3x3Transpose(op)

	return mat
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat3x3) MulComp(op *X32Mat3x3) *X32Mat3x3 {
	*mat = *X32Mat3x3MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *X32Mat3x3) Det() vpnumber.X32 {
	return vpnumber.X32Muln(mat.Get(0, 0), mat.Get(1, 1), mat.Get(2, 2)) + vpnumber.X32Muln(mat.Get(0, 1), mat.Get(1, 2), mat.Get(2, 0)) + vpnumber.X32Muln(mat.Get(0, 2), mat.Get(1, 0), mat.Get(2, 1)) - vpnumber.X32Muln(mat.Get(0, 0), mat.Get(1, 2), mat.Get(2, 1)) - vpnumber.X32Muln(mat.Get(0, 1), mat.Get(1, 0), mat.Get(2, 2)) - vpnumber.X32Muln(mat.Get(0, 2), mat.Get(1, 1), mat.Get(2, 0))
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat3x3) Inv() *X32Mat3x3 {
	*mat = *X32Mat3x3Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 3x3 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *X32Mat3x3) MulVec(vec *X32Vec3) *X32Vec3 {
	var ret X32Vec3

	for i := range vec {
		ret[i] = vpnumber.X32Mul(mat.Get(0, i), vec[0]) + vpnumber.X32Mul(mat.Get(1, i), vec[1]) + vpnumber.X32Mul(mat.Get(2, i), vec[2])
	}

	return &ret
}

// MulVecPos performs a multiplication of a vector by a 3x3 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 2 (a point in a plane) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *X32Mat3x3) MulVecPos(vec *vpmatrix2.X32Vec2) *vpmatrix2.X32Vec2 {
	var ret vpmatrix2.X32Vec2

	for i := range vec {
		ret[i] = vpnumber.X32Mul(mat.Get(0, i), vec[0]) + vpnumber.X32Mul(mat.Get(1, i), vec[1]) + mat.Get(2, i)
	}

	return &ret
}

// MulVecDir performs a multiplication of a vector by a 3x3 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 2 (a point in a plane) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *X32Mat3x3) MulVecDir(vec *vpmatrix2.X32Vec2) *vpmatrix2.X32Vec2 {
	var ret vpmatrix2.X32Vec2

	for i := range vec {
		ret[i] = vpnumber.X32Mul(mat.Get(0, i), vec[0]) + vpnumber.X32Mul(mat.Get(1, i), vec[1])
	}

	return &ret
}

// X32Mat3x3Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat3x3Add(mata, matb *X32Mat3x3) *X32Mat3x3 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// X32Mat3x3Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat3x3Sub(mata, matb *X32Mat3x3) *X32Mat3x3 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// X32Mat3x3MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat3x3MulScale(mat *X32Mat3x3, factor vpnumber.X32) *X32Mat3x3 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// X32Mat3x3DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat3x3DivScale(mat *X32Mat3x3, factor vpnumber.X32) *X32Mat3x3 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// X32Mat3x3Transpose inverts rows and columns (matrix transposition).
// Args is left untouched, a pointer on a new object is returned.
func X32Mat3x3Transpose(mat *X32Mat3x3) *X32Mat3x3 {
	var ret X32Mat3x3

	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			ret.Set(c, r, mat.Get(r, c))
		}
	}

	return &ret
}

// X32Mat3x3MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func X32Mat3x3MulComp(a, b *X32Mat3x3) *X32Mat3x3 {
	var ret X32Mat3x3

	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			ret.Set(c, r, vpnumber.X32Mul(a.Get(0, r), b.Get(c, 0))+vpnumber.X32Mul(a.Get(1, r), b.Get(c, 1))+vpnumber.X32Mul(a.Get(2, r), b.Get(c, 2)))
		}
	}

	return &ret
}

// X32Mat3x3Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func X32Mat3x3Inv(mat *X32Mat3x3) *X32Mat3x3 {
	ret := X32Mat3x3{
		vpnumber.X32Mul(mat.Get(1, 1), mat.Get(2, 2)) - vpnumber.X32Mul(mat.Get(1, 2), mat.Get(2, 1)),
		vpnumber.X32Mul(mat.Get(0, 2), mat.Get(2, 1)) - vpnumber.X32Mul(mat.Get(0, 1), mat.Get(2, 2)),
		vpnumber.X32Mul(mat.Get(0, 1), mat.Get(1, 2)) - vpnumber.X32Mul(mat.Get(0, 2), mat.Get(1, 1)),
		vpnumber.X32Mul(mat.Get(1, 2), mat.Get(2, 0)) - vpnumber.X32Mul(mat.Get(1, 0), mat.Get(2, 2)),
		vpnumber.X32Mul(mat.Get(0, 0), mat.Get(2, 2)) - vpnumber.X32Mul(mat.Get(0, 2), mat.Get(2, 0)),
		vpnumber.X32Mul(mat.Get(0, 2), mat.Get(1, 0)) - vpnumber.X32Mul(mat.Get(0, 0), mat.Get(1, 2)),
		vpnumber.X32Mul(mat.Get(1, 0), mat.Get(2, 1)) - vpnumber.X32Mul(mat.Get(1, 1), mat.Get(2, 0)),
		vpnumber.X32Mul(mat.Get(0, 1), mat.Get(2, 0)) - vpnumber.X32Mul(mat.Get(0, 0), mat.Get(2, 1)),
		vpnumber.X32Mul(mat.Get(0, 0), mat.Get(1, 1)) - vpnumber.X32Mul(mat.Get(0, 1), mat.Get(1, 0)),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
