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

package vpmat2x1

import (
	"encoding/json"
	"github.com/ufoot/vapor/go/vperror"
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec2"
)

// F32 is a matrix containing 2x1 float32 values.
// Can hold the values of a point in a plane.
type F32 [Size]float32

// F32New creates a new matrix containing 2x1 float32 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F32New(f1, f2 float32) *F32 {
	return &F32{f1, f2}
}

// F32Identity creates a new identity matrix.
func F32Identity() *F32 {
	return &F32{vpnumber.F32Const1, vpnumber.F32Const0}
}

// F32Translation creates a new translation matrix.
func F32Translation(f float32) *F32 {
	return &F32{vpnumber.F32Const1, f}
}

// F32Scale creates a new scale matrix.
func F32Scale(f float32) *F32 {
	return &F32{f, vpnumber.F32Const0}
}

// F32RebaseOX creates a matrix that translates from the default
// O=(0), X=(1) basis to the given
// basis. It assumes f(a+b) equals f(a)+f(b).
func F32RebaseOX(Origin, PosX float32) *F32 {
	return &F32{PosX - Origin, Origin}
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F32) ToX32() *X32 {
	var ret X32

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F32) ToX64() *X64 {
	var ret X64

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX64(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *F32) ToF64() *F64 {
	var ret F64

	for i, v := range mat {
		ret[i] = float64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F32) Set(col, row int, val float32) {
	mat[col+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F32) Get(col, row int) float32 {
	return mat[col+row]
}

// SetCol sets a column to the values contained in a vector.
func (mat *F32) SetCol(col int, vec float32) {
	mat[col] = vec
}

// GetCol gets a column and returns it in a vector.
func (mat *F32) GetCol(col int) float32 {
	return mat[col]
}

// SetRow sets a row to the values contained in a vector.
func (mat *F32) SetRow(vec *vpvec2.F32) {
	*mat = F32(*vec)
}

// GetRow gets a row and returns it in a vector.
func (mat *F32) GetRow() *vpvec2.F32 {
	ret := vpvec2.F32(*mat)

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *F32) MarshalJSON() ([]byte, error) {
	var tmpArray [Width][Height]float32

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vperror.Chain(err, "unable to marshal F32")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *F32) UnmarshalJSON(data []byte) error {
	var tmpArray [Width][Height]float32

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vperror.Chain(err, "unable to unmarshal F32")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *F32) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) Add(op *F32) *F32 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) Sub(op *F32) *F32 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) MulScale(factor float32) *F32 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) DivScale(factor float32) *F32 {
	for i, v := range mat {
		mat[i] = vpnumber.F32Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F32) IsSimilar(op *F32) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F32IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) MulComp(op *F32) *F32 {
	*mat = *F32MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *F32) Det() float32 {
	return mat[Col0Row0]
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) Inv() *F32 {
	*mat = *F32Inv(mat)

	return mat
}

// MulVecPos performs a multiplication of a vector by a 2x1 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *F32) MulVecPos(vec float32) float32 {
	return mat[Col0Row0]*vec + mat[Col1Row0]
}

// MulVecDir performs a multiplication of a vector by a 2x1 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *F32) MulVecDir(vec float32) float32 {
	return mat[Col0Row0] * vec
}

// F32Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func F32Add(mata, matb *F32) *F32 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F32Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F32Sub(mata, matb *F32) *F32 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F32MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32MulScale(mat *F32, factor float32) *F32 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F32DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32DivScale(mat *F32, factor float32) *F32 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F32MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func F32MulComp(a, b *F32) *F32 {
	ret := F32{a[Col0Row0] * b[Col0Row0],
		a[Col0Row0]*b[Col1Row0] + a[Col1Row0]}

	return &ret
}

// F32Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func F32Inv(mat *F32) *F32 {
	ret := F32{
		vpnumber.F32Const1,
		-mat[Col1Row0],
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
