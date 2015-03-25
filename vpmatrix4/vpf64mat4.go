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

package vpmatrix4

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpmatrix3"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
)

// F64Mat4 is a matrix containing 4x4 float64 values.
// Can be used in 3D matrix transformations.
type F64Mat4 [16]float64

// F64Mat4New creates a new matrix containing 4x4 float64 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F64Mat4New(f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16 float64) *F64Mat4 {
	return &F64Mat4{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16}
}

// F64Mat4Identity creates a new identity matrix.
func F64Mat4Identity() *F64Mat4 {
	return &F64Mat4{vpnumber.F64Const1, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *F64Mat4) ToI32() *I32Mat4 {
	var ret I32Mat4

	for i, v := range mat {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int64 matrix.
func (mat *F64Mat4) ToI64() *I64Mat4 {
	var ret I64Mat4

	for i, v := range mat {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F64Mat4) ToX32() *X32Mat4 {
	var ret X32Mat4

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F64Mat4) ToX64() *X64Mat4 {
	var ret X64Mat4

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *F64Mat4) ToF32() *F32Mat4 {
	var ret F32Mat4

	for i, v := range mat {
		ret[i] = float32(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F64Mat4) Set(col, row int, val float64) {
	mat[col*4+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F64Mat4) Get(col, row int) float64 {
	return mat[col*4+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *F64Mat4) MarshalJSON() ([]byte, error) {
	var tmpArray [4][4]float64

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col*4+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal F64Mat4")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *F64Mat4) UnmarshalJSON(data []byte) error {
	var tmpArray [4][4]float64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal F64Mat4")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*4+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *F64Mat4) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat4) Add(op *F64Mat4) *F64Mat4 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat4) Sub(op *F64Mat4) *F64Mat4 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat4) MulScale(factor float64) *F64Mat4 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat4) DivScale(factor float64) *F64Mat4 {
	for i, v := range mat {
		mat[i] = vpnumber.F64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F64Mat4) IsSimilar(op *F64Mat4) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F64IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat4) MulComp(op *F64Mat4) *F64Mat4 {
	*mat = *F64Mat4MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *F64Mat4) Det() float64 {
	return mat.Get(0, 3)*mat.Get(1, 2)*mat.Get(2, 1)*mat.Get(3, 0) - mat.Get(0, 2)*mat.Get(1, 3)*mat.Get(2, 1)*mat.Get(3, 0) - mat.Get(0, 3)*mat.Get(1, 1)*mat.Get(2, 2)*mat.Get(3, 0) + mat.Get(0, 1)*mat.Get(1, 3)*mat.Get(2, 2)*mat.Get(3, 0) + mat.Get(0, 2)*mat.Get(1, 1)*mat.Get(2, 3)*mat.Get(3, 0) - mat.Get(0, 1)*mat.Get(1, 2)*mat.Get(2, 3)*mat.Get(3, 0) - mat.Get(0, 3)*mat.Get(1, 2)*mat.Get(2, 0)*mat.Get(3, 1) + mat.Get(0, 2)*mat.Get(1, 3)*mat.Get(2, 0)*mat.Get(3, 1) + mat.Get(0, 3)*mat.Get(1, 0)*mat.Get(2, 2)*mat.Get(3, 1) - mat.Get(0, 0)*mat.Get(1, 3)*mat.Get(2, 2)*mat.Get(3, 1) - mat.Get(0, 2)*mat.Get(1, 0)*mat.Get(2, 3)*mat.Get(3, 1) + mat.Get(0, 0)*mat.Get(1, 2)*mat.Get(2, 3)*mat.Get(3, 1) + mat.Get(0, 3)*mat.Get(1, 1)*mat.Get(2, 0)*mat.Get(3, 2) - mat.Get(0, 1)*mat.Get(1, 3)*mat.Get(2, 0)*mat.Get(3, 2) - mat.Get(0, 3)*mat.Get(1, 0)*mat.Get(2, 1)*mat.Get(3, 2) + mat.Get(0, 0)*mat.Get(1, 3)*mat.Get(2, 1)*mat.Get(3, 2) + mat.Get(0, 1)*mat.Get(1, 0)*mat.Get(2, 3)*mat.Get(3, 2) - mat.Get(0, 0)*mat.Get(1, 1)*mat.Get(2, 3)*mat.Get(3, 2) - mat.Get(0, 2)*mat.Get(1, 1)*mat.Get(2, 0)*mat.Get(3, 3) + mat.Get(0, 1)*mat.Get(1, 2)*mat.Get(2, 0)*mat.Get(3, 3) + mat.Get(0, 2)*mat.Get(1, 0)*mat.Get(2, 1)*mat.Get(3, 3) - mat.Get(0, 0)*mat.Get(1, 2)*mat.Get(2, 1)*mat.Get(3, 3) - mat.Get(0, 1)*mat.Get(1, 0)*mat.Get(2, 2)*mat.Get(3, 3) + mat.Get(0, 0)*mat.Get(1, 1)*mat.Get(2, 2)*mat.Get(3, 3)
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat4) Inv() *F64Mat4 {
	*mat = *F64Mat4Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *F64Mat4) MulVec(vec *F64Vec4) *F64Vec4 {
	var ret F64Vec4

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2] + mat.Get(3, i)*vec[3]
	}

	return &ret
}

// MulVecPos performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *F64Mat4) MulVecPos(vec *vpmatrix3.F64Vec3) *vpmatrix3.F64Vec3 {
	var ret vpmatrix3.F64Vec3

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2] + mat.Get(3, i)
	}

	return &ret
}

// MulVecDir performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *F64Mat4) MulVecDir(vec *vpmatrix3.F64Vec3) *vpmatrix3.F64Vec3 {
	var ret vpmatrix3.F64Vec3

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2]
	}

	return &ret
}

// F64Mat4Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat4Add(mata, matb *F64Mat4) *F64Mat4 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F64Mat4Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat4Sub(mata, matb *F64Mat4) *F64Mat4 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F64Mat4MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat4MulScale(mat *F64Mat4, factor float64) *F64Mat4 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F64Mat4DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat4DivScale(mat *F64Mat4, factor float64) *F64Mat4 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F64Mat4MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func F64Mat4MulComp(a, b *F64Mat4) *F64Mat4 {
	var ret F64Mat4

	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			ret.Set(c, r, a.Get(0, r)*b.Get(c, 0)+a.Get(1, r)*b.Get(c, 1)+a.Get(2, r)*b.Get(c, 2)+a.Get(3, r)*b.Get(c, 3))
		}
	}

	return &ret
}

// F64Mat4Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func F64Mat4Inv(mat *F64Mat4) *F64Mat4 {
	ret := F64Mat4{
		mat.Get(1, 2)*mat.Get(2, 3)*mat.Get(3, 1) - mat.Get(1, 3)*mat.Get(2, 2)*mat.Get(3, 1) + mat.Get(1, 3)*mat.Get(2, 1)*mat.Get(3, 2) - mat.Get(1, 1)*mat.Get(2, 3)*mat.Get(3, 2) - mat.Get(1, 2)*mat.Get(2, 1)*mat.Get(3, 3) + mat.Get(1, 1)*mat.Get(2, 2)*mat.Get(3, 3),
		mat.Get(0, 3)*mat.Get(2, 2)*mat.Get(3, 1) - mat.Get(0, 2)*mat.Get(2, 3)*mat.Get(3, 1) - mat.Get(0, 3)*mat.Get(2, 1)*mat.Get(3, 2) + mat.Get(0, 1)*mat.Get(2, 3)*mat.Get(3, 2) + mat.Get(0, 2)*mat.Get(2, 1)*mat.Get(3, 3) - mat.Get(0, 1)*mat.Get(2, 2)*mat.Get(3, 3),
		mat.Get(0, 2)*mat.Get(1, 3)*mat.Get(3, 1) - mat.Get(0, 3)*mat.Get(1, 2)*mat.Get(3, 1) + mat.Get(0, 3)*mat.Get(1, 1)*mat.Get(3, 2) - mat.Get(0, 1)*mat.Get(1, 3)*mat.Get(3, 2) - mat.Get(0, 2)*mat.Get(1, 1)*mat.Get(3, 3) + mat.Get(0, 1)*mat.Get(1, 2)*mat.Get(3, 3),
		mat.Get(0, 3)*mat.Get(1, 2)*mat.Get(2, 1) - mat.Get(0, 2)*mat.Get(1, 3)*mat.Get(2, 1) - mat.Get(0, 3)*mat.Get(1, 1)*mat.Get(2, 2) + mat.Get(0, 1)*mat.Get(1, 3)*mat.Get(2, 2) + mat.Get(0, 2)*mat.Get(1, 1)*mat.Get(2, 3) - mat.Get(0, 1)*mat.Get(1, 2)*mat.Get(2, 3),
		mat.Get(1, 3)*mat.Get(2, 2)*mat.Get(3, 0) - mat.Get(1, 2)*mat.Get(2, 3)*mat.Get(3, 0) - mat.Get(1, 3)*mat.Get(2, 0)*mat.Get(3, 2) + mat.Get(1, 0)*mat.Get(2, 3)*mat.Get(3, 2) + mat.Get(1, 2)*mat.Get(2, 0)*mat.Get(3, 3) - mat.Get(1, 0)*mat.Get(2, 2)*mat.Get(3, 3),
		mat.Get(0, 2)*mat.Get(2, 3)*mat.Get(3, 0) - mat.Get(0, 3)*mat.Get(2, 2)*mat.Get(3, 0) + mat.Get(0, 3)*mat.Get(2, 0)*mat.Get(3, 2) - mat.Get(0, 0)*mat.Get(2, 3)*mat.Get(3, 2) - mat.Get(0, 2)*mat.Get(2, 0)*mat.Get(3, 3) + mat.Get(0, 0)*mat.Get(2, 2)*mat.Get(3, 3),
		mat.Get(0, 3)*mat.Get(1, 2)*mat.Get(3, 0) - mat.Get(0, 2)*mat.Get(1, 3)*mat.Get(3, 0) - mat.Get(0, 3)*mat.Get(1, 0)*mat.Get(3, 2) + mat.Get(0, 0)*mat.Get(1, 3)*mat.Get(3, 2) + mat.Get(0, 2)*mat.Get(1, 0)*mat.Get(3, 3) - mat.Get(0, 0)*mat.Get(1, 2)*mat.Get(3, 3),
		mat.Get(0, 2)*mat.Get(1, 3)*mat.Get(2, 0) - mat.Get(0, 3)*mat.Get(1, 2)*mat.Get(2, 0) + mat.Get(0, 3)*mat.Get(1, 0)*mat.Get(2, 2) - mat.Get(0, 0)*mat.Get(1, 3)*mat.Get(2, 2) - mat.Get(0, 2)*mat.Get(1, 0)*mat.Get(2, 3) + mat.Get(0, 0)*mat.Get(1, 2)*mat.Get(2, 3),
		mat.Get(1, 1)*mat.Get(2, 3)*mat.Get(3, 0) - mat.Get(1, 3)*mat.Get(2, 1)*mat.Get(3, 0) + mat.Get(1, 3)*mat.Get(2, 0)*mat.Get(3, 1) - mat.Get(1, 0)*mat.Get(2, 3)*mat.Get(3, 1) - mat.Get(1, 1)*mat.Get(2, 0)*mat.Get(3, 3) + mat.Get(1, 0)*mat.Get(2, 1)*mat.Get(3, 3),
		mat.Get(0, 3)*mat.Get(2, 1)*mat.Get(3, 0) - mat.Get(0, 1)*mat.Get(2, 3)*mat.Get(3, 0) - mat.Get(0, 3)*mat.Get(2, 0)*mat.Get(3, 1) + mat.Get(0, 0)*mat.Get(2, 3)*mat.Get(3, 1) + mat.Get(0, 1)*mat.Get(2, 0)*mat.Get(3, 3) - mat.Get(0, 0)*mat.Get(2, 1)*mat.Get(3, 3),
		mat.Get(0, 1)*mat.Get(1, 3)*mat.Get(3, 0) - mat.Get(0, 3)*mat.Get(1, 1)*mat.Get(3, 0) + mat.Get(0, 3)*mat.Get(1, 0)*mat.Get(3, 1) - mat.Get(0, 0)*mat.Get(1, 3)*mat.Get(3, 1) - mat.Get(0, 1)*mat.Get(1, 0)*mat.Get(3, 3) + mat.Get(0, 0)*mat.Get(1, 1)*mat.Get(3, 3),
		mat.Get(0, 3)*mat.Get(1, 1)*mat.Get(2, 0) - mat.Get(0, 1)*mat.Get(1, 3)*mat.Get(2, 0) - mat.Get(0, 3)*mat.Get(1, 0)*mat.Get(2, 1) + mat.Get(0, 0)*mat.Get(1, 3)*mat.Get(2, 1) + mat.Get(0, 1)*mat.Get(1, 0)*mat.Get(2, 3) - mat.Get(0, 0)*mat.Get(1, 1)*mat.Get(2, 3),
		mat.Get(1, 2)*mat.Get(2, 1)*mat.Get(3, 0) - mat.Get(1, 1)*mat.Get(2, 2)*mat.Get(3, 0) - mat.Get(1, 2)*mat.Get(2, 0)*mat.Get(3, 1) + mat.Get(1, 0)*mat.Get(2, 2)*mat.Get(3, 1) + mat.Get(1, 1)*mat.Get(2, 0)*mat.Get(3, 2) - mat.Get(1, 0)*mat.Get(2, 1)*mat.Get(3, 2),
		mat.Get(0, 1)*mat.Get(2, 2)*mat.Get(3, 0) - mat.Get(0, 2)*mat.Get(2, 1)*mat.Get(3, 0) + mat.Get(0, 2)*mat.Get(2, 0)*mat.Get(3, 1) - mat.Get(0, 0)*mat.Get(2, 2)*mat.Get(3, 1) - mat.Get(0, 1)*mat.Get(2, 0)*mat.Get(3, 2) + mat.Get(0, 0)*mat.Get(2, 1)*mat.Get(3, 2),
		mat.Get(0, 2)*mat.Get(1, 1)*mat.Get(3, 0) - mat.Get(0, 1)*mat.Get(1, 2)*mat.Get(3, 0) - mat.Get(0, 2)*mat.Get(1, 0)*mat.Get(3, 1) + mat.Get(0, 0)*mat.Get(1, 2)*mat.Get(3, 1) + mat.Get(0, 1)*mat.Get(1, 0)*mat.Get(3, 2) - mat.Get(0, 0)*mat.Get(1, 1)*mat.Get(3, 2),
		mat.Get(0, 1)*mat.Get(1, 2)*mat.Get(2, 0) - mat.Get(0, 2)*mat.Get(1, 1)*mat.Get(2, 0) + mat.Get(0, 2)*mat.Get(1, 0)*mat.Get(2, 1) - mat.Get(0, 0)*mat.Get(1, 2)*mat.Get(2, 1) - mat.Get(0, 1)*mat.Get(1, 0)*mat.Get(2, 2) + mat.Get(0, 0)*mat.Get(1, 1)*mat.Get(2, 2),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
