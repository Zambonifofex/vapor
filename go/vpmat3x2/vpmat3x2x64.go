// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
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

package vpmat3x2

import (
	"encoding/json"
	"github.com/ufoot/vapor/go/vperror"
	"github.com/ufoot/vapor/go/vpmath"
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec2"
	"github.com/ufoot/vapor/go/vpvec3"
)

// X64 is a matrix containing 3x2 fixed point 64 bit values.
// Can hold the values of a point in space.
type X64 [Size]vpnumber.X64

// X64New creates a new matrix containing 3x2 fixed point 64 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X64New(x1, x2, x3, x4, x5, x6 vpnumber.X64) *X64 {
	return &X64{x1, x2, x3, x4, x5, x6}
}

// X64Identity creates a new identity matrix.
func X64Identity() *X64 {
	return &X64{vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0}
}

// X64Translation creates a new translation matrix.
func X64Translation(vec *vpvec2.X64) *X64 {
	return &X64{vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vec[0], vec[1]}
}

// X64Scale creates a new scale matrix.
func X64Scale(vec *vpvec2.X64) *X64 {
	return &X64{vec[0], vpnumber.X64Const0, vpnumber.X64Const0, vec[1], vpnumber.X64Const0, vpnumber.X64Const0}
}

// X64Rot creates a new rotation matrix.
// The rotation is done in 2D over a virtual z axis, such as z = cross(x,y).
// Angle is given in radians.
func X64Rot(r vpnumber.X64) *X64 {
	cos := vpmath.X64Cos(r)
	sin := vpmath.X64Sin(r)

	return &X64{cos, sin, -sin, cos, vpnumber.X64Const0, vpnumber.X64Const0}
}

// X64RebaseOXY creates a matrix that translates from the default
// O=(0,0), X=(1,0), Y=(0,1) basis to the given
// basis. It assumes f(a+b) equals f(a)+f(b).
func X64RebaseOXY(Origin, PosX, PosY *vpvec2.X64) *X64 {
	return &X64{PosX[0] - Origin[0], PosX[1] - Origin[1], PosY[0] - Origin[0], PosY[1] - Origin[1], Origin[0], Origin[1]}
}

// ToX32 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *X64) ToX32() *X32 {
	var ret X32

	for i, v := range mat {
		ret[i] = vpnumber.X64ToX32(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *X64) ToF32() *F32 {
	var ret F32

	for i, v := range mat {
		ret[i] = vpnumber.X64ToF32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *X64) ToF64() *F64 {
	var ret F64

	for i, v := range mat {
		ret[i] = vpnumber.X64ToF64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *X64) Set(col, row int, val vpnumber.X64) {
	mat[col*Height+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *X64) Get(col, row int) vpnumber.X64 {
	return mat[col*Height+row]
}

// SetCol sets a column to the values contained in a vector.
func (mat *X64) SetCol(col int, vec *vpvec2.X64) {
	for row, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetCol gets a column and returns it in a vector.
func (mat *X64) GetCol(col int) *vpvec2.X64 {
	var ret vpvec2.X64

	for row := range ret {
		ret[row] = mat[col*Height+row]
	}

	return &ret
}

// SetRow sets a row to the values contained in a vector.
func (mat *X64) SetRow(row int, vec *vpvec3.X64) {
	for col, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetRow gets a row and returns it in a vector.
func (mat *X64) GetRow(row int) *vpvec3.X64 {
	var ret vpvec3.X64

	for col := range ret {
		ret[col] = mat[col*Height+row]
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *X64) MarshalJSON() ([]byte, error) {
	var tmpArray [Width][Height]int64

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = int64(mat[col*Height+row])
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vperror.Chain(err, "unable to marshal X64")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *X64) UnmarshalJSON(data []byte) error {
	var tmpArray [Width][Height]int64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vperror.Chain(err, "unable to unmarshal X64")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*Height+row] = vpnumber.X64(tmpArray[col][row])
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *X64) String() string {
	buf, err := mat.ToF64().MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) Add(op *X64) *X64 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) Sub(op *X64) *X64 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) MulScale(factor vpnumber.X64) *X64 {
	for i, v := range mat {
		mat[i] = vpnumber.X64Mul(v, factor)
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) DivScale(factor vpnumber.X64) *X64 {
	for i, v := range mat {
		mat[i] = vpnumber.X64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *X64) IsSimilar(op *X64) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) MulComp(op *X64) *X64 {
	*mat = *X64MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *X64) Det() vpnumber.X64 {
	return vpnumber.X64Mul(mat[Col0Row0], mat[Col1Row1]) - vpnumber.X64Mul(mat[Col0Row1], mat[Col1Row0])
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) Inv() *X64 {
	*mat = *X64Inv(mat)

	return mat
}

// MulVecPos performs a multiplication of a vector by a 3x2 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 2 (a point in a plane) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *X64) MulVecPos(vec *vpvec2.X64) *vpvec2.X64 {
	var ret vpvec2.X64

	for i := range vec {
		ret[i] = vpnumber.X64Mul(mat.Get(0, i), vec[0]) + vpnumber.X64Mul(mat.Get(1, i), vec[1]) + mat.Get(2, i)
	}

	return &ret
}

// MulVecDir performs a multiplication of a vector by a 3x2 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 2 (a point in a plane) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *X64) MulVecDir(vec *vpvec2.X64) *vpvec2.X64 {
	var ret vpvec2.X64

	for i := range vec {
		ret[i] = vpnumber.X64Mul(mat.Get(0, i), vec[0]) + vpnumber.X64Mul(mat.Get(1, i), vec[1])
	}

	return &ret
}

// X64Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func X64Add(mata, matb *X64) *X64 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// X64Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func X64Sub(mata, matb *X64) *X64 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// X64MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64MulScale(mat *X64, factor vpnumber.X64) *X64 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// X64DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64DivScale(mat *X64, factor vpnumber.X64) *X64 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// X64MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func X64MulComp(a, b *X64) *X64 {
	var ret X64

	for c := 0; c < Width-1; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, vpnumber.X64Mul(a.Get(0, r), b.Get(c, 0))+vpnumber.X64Mul(a.Get(1, r), b.Get(c, 1)))
		}
	}
	for r := 0; r < Height; r++ {
		ret.Set(2, r, vpnumber.X64Mul(a.Get(0, r), b[Col2Row0])+vpnumber.X64Mul(a.Get(1, r), b[Col2Row1])+a.Get(2, r))
	}

	return &ret
}

// X64Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func X64Inv(mat *X64) *X64 {
	ret := X64{
		mat[Col1Row1],
		-mat[Col0Row1],
		-mat[Col1Row0],
		mat[Col0Row0],
		vpnumber.X64Mul(mat[Col1Row0], mat[Col2Row1]) - vpnumber.X64Mul(mat[Col1Row1], mat[Col2Row0]),
		vpnumber.X64Mul(mat[Col0Row1], mat[Col2Row0]) - vpnumber.X64Mul(mat[Col0Row0], mat[Col2Row1]),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
