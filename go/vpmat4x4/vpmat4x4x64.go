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

package vpmat4x4

import (
	"encoding/json"
	"github.com/ufoot/vapor/go/vperror"
	"github.com/ufoot/vapor/go/vpmat3x3"
	"github.com/ufoot/vapor/go/vpmath"
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec3"
	"github.com/ufoot/vapor/go/vpvec4"
)

// X64 is a matrix containing 4x4 fixed point 64 bit values.
// Can be used in 3D matrix transformations.
type X64 [Size]vpnumber.X64

// X64New creates a new matrix containing 4x4 fixed point 64 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X64New(x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16 vpnumber.X64) *X64 {
	return &X64{x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16}
}

// X64Identity creates a new identity matrix.
func X64Identity() *X64 {
	return &X64{vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1}
}

// X64Translation creates a new translation matrix.
func X64Translation(vec *vpvec3.X64) *X64 {
	return &X64{vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0, vec[0], vec[1], vec[2], vpnumber.X64Const1}
}

// X64Scale creates a new scale matrix.
func X64Scale(vec *vpvec3.X64) *X64 {
	return &X64{vec[0], vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vec[1], vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vec[2], vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1}
}

// X64RotX creates a new rotation matrix.
// The rotation is done in 3D over the x (1st) axis.
// Angle is given in radians.
func X64RotX(r vpnumber.X64) *X64 {
	cos := vpmath.X64Cos(r)
	sin := vpmath.X64Sin(r)

	return &X64{vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, cos, sin, vpnumber.X64Const0, vpnumber.X64Const0, -sin, cos, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1}
}

// X64RotY creates a new rotation matrix.
// The rotation is done in 3D over the y (2nd) axis.
// Angle is given in radians.
func X64RotY(r vpnumber.X64) *X64 {
	cos := vpmath.X64Cos(r)
	sin := vpmath.X64Sin(r)

	return &X64{cos, vpnumber.X64Const0, -sin, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, sin, vpnumber.X64Const0, cos, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1}
}

// X64RotZ creates a new rotation matrix.
// The rotation is done in 3D over the z (3rd) axis.
// Angle is given in radians.
func X64RotZ(r vpnumber.X64) *X64 {
	cos := vpmath.X64Cos(r)
	sin := vpmath.X64Sin(r)

	return &X64{cos, sin, vpnumber.X64Const0, vpnumber.X64Const0, -sin, cos, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1}
}

// X64RebaseOXYZ creates a matrix that translates from the default
// O=(0,0,0), X=(1,0,0), Y=(0,1,0), Z=(0,0,1) basis to the given
// basis. It assumes f(a+b) equals f(a)+f(b).
func X64RebaseOXYZ(Origin, PosX, PosY, PosZ *vpvec3.X64) *X64 {
	return &X64{PosX[0] - Origin[0], PosX[1] - Origin[1], PosX[2] - Origin[2], vpnumber.X64Const0, PosY[0] - Origin[0], PosY[1] - Origin[1], PosY[2] - Origin[2], vpnumber.X64Const0, PosZ[0] - Origin[0], PosZ[1] - Origin[1], PosZ[2] - Origin[2], vpnumber.X64Const0, Origin[0], Origin[1], Origin[2], vpnumber.X64Const1}
}

// X64RebaseOXYZP creates a matrix that translates from the default
// O=(0,0,0), X=(1,0,0), Y=(0,1,0), Z=(0,0,1), P=(1,1,1) basis to the given
// basis. Note that there can be a projection, so  f(a+b) is not f(a)+f(b).
func X64RebaseOXYZP(Origin, PosX, PosY, PosZ, PosP *vpvec3.X64) *X64 {
	var tmpMat vpmat3x3.X64
	projMat := X64Identity()

	dX := vpvec3.X64Sub(PosX, Origin)
	dY := vpvec3.X64Sub(PosY, Origin)
	dZ := vpvec3.X64Sub(PosZ, Origin)
	dP := vpvec3.X64Sub(PosP, Origin)
	tmpMat.SetCol(0, vpvec3.X64Sub(dX, dP))
	tmpMat.SetCol(1, vpvec3.X64Sub(dY, dP))
	tmpMat.SetCol(2, vpvec3.X64Sub(dZ, dP))
	tmpMat.Inv()
	tmpVec := vpvec3.X64Sub(dP, vpvec3.X64Add(dX, vpvec3.X64Add(dY, dZ)))
	lastRow := tmpMat.MulVec(tmpVec)
	colX := vpvec3.X64MulScale(dX, vpnumber.X64Const1+lastRow[0])
	colY := vpvec3.X64MulScale(dY, vpnumber.X64Const1+lastRow[1])
	colZ := vpvec3.X64MulScale(dZ, vpnumber.X64Const1+lastRow[2])
	projMat.SetCol(0, vpvec4.X64FromVec3(colX, lastRow[0]))
	projMat.SetCol(1, vpvec4.X64FromVec3(colY, lastRow[1]))
	projMat.SetCol(2, vpvec4.X64FromVec3(colZ, lastRow[2]))
	transMat := X64Translation(Origin)

	ret := X64MulComp(transMat, projMat)

	return ret
}

// X64Ortho creates a projection matrix the way the standard OpenGL glOrtho
// would (see https://www.opengl.org/sdk/docs/man2/xhtml/glOrtho.xml).
// Note: use -nearVal and -farVal to initialize.
// It's a little akward, if you expect to pass vectors with positions
// ranging from nearVal to farVal then you need to pass -nearVal and
// -farVal to this function. This is probably due to the fact that
// with a right-handed basis and X,Y set up "as usual", then Z is negative
// when going farther and farther. This tweak allows farVal to yield
// +1 and nearVal -1. We keep this function as is here, as this is the
// way OpenGL functions seem to work.
func X64Ortho(left, right, bottom, top, nearVal, farVal vpnumber.X64) *X64 {
	var ret X64

	ret[Col0Row0] = vpnumber.X64Div(vpnumber.I64ToX64(2), right-left)
	ret[Col1Row1] = vpnumber.X64Div(vpnumber.I64ToX64(2), top-bottom)
	ret[Col2Row2] = vpnumber.X64Div(-vpnumber.I64ToX64(2), farVal-nearVal)
	ret[Col3Row0] = -vpnumber.X64Div(right+left, right-left)
	ret[Col3Row1] = -vpnumber.X64Div(top+bottom, top-bottom)
	ret[Col3Row2] = -vpnumber.X64Div(farVal+nearVal, farVal-nearVal)
	ret[Col3Row3] = vpnumber.X64Const1

	return &ret
}

// X64Perspective creates a projection matrix the way the standard GLU
// gluPerspective function would (see
// https://www.opengl.org/sdk/docs/man2/xhtml/gluPerspective.xml).
// Beware, fovy is in degrees, not radians.
func X64Perspective(fovy, aspect, zNear, zFar vpnumber.X64) *X64 {
	var ret X64

	radFovy2 := vpmath.X64DegToRad(vpmath.X64DegMod(fovy) >> 1)
	f := vpnumber.X64Div(vpmath.X64Cos(radFovy2), vpmath.X64Sin(radFovy2))

	ret[Col0Row0] = vpnumber.X64Div(f, aspect)
	ret[Col1Row1] = f
	ret[Col2Row2] = vpnumber.X64Div(zFar+zNear, zNear-zFar)
	ret[Col2Row3] = -vpnumber.X64Const1
	ret[Col3Row2] = vpnumber.X64Div(vpnumber.X64Mul(zFar, zNear)<<1, zNear-zFar)

	return &ret
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
func (mat *X64) SetCol(col int, vec *vpvec4.X64) {
	for row, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetCol gets a column and returns it in a vector.
func (mat *X64) GetCol(col int) *vpvec4.X64 {
	var ret vpvec4.X64

	for row := range ret {
		ret[row] = mat[col*Height+row]
	}

	return &ret
}

// SetRow sets a row to the values contained in a vector.
func (mat *X64) SetRow(row int, vec *vpvec4.X64) {
	for col, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetRow gets a row and returns it in a vector.
func (mat *X64) GetRow(row int) *vpvec4.X64 {
	var ret vpvec4.X64

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

// Transpose inverts rows and columns (matrix transposition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) Transpose(op *X64) *X64 {
	*mat = *X64Transpose(op)

	return mat
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) MulComp(op *X64) *X64 {
	*mat = *X64MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *X64) Det() vpnumber.X64 {
	return vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row2], mat[Col2Row1], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row3], mat[Col2Row1], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row1], mat[Col2Row2], mat[Col3Row0]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row3], mat[Col2Row2], mat[Col3Row0]) + vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row1], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row2], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row2], mat[Col2Row0], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row3], mat[Col2Row0], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row0], mat[Col2Row2], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row3], mat[Col2Row2], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row0], mat[Col2Row3], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row2], mat[Col2Row3], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row1], mat[Col2Row0], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row3], mat[Col2Row0], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row0], mat[Col2Row1], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row3], mat[Col2Row1], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row0], mat[Col2Row3], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row1], mat[Col2Row3], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row1], mat[Col2Row0], mat[Col3Row3]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row2], mat[Col2Row0], mat[Col3Row3]) + vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row0], mat[Col2Row1], mat[Col3Row3]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row2], mat[Col2Row1], mat[Col3Row3]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row0], mat[Col2Row2], mat[Col3Row3]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row1], mat[Col2Row2], mat[Col3Row3])
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64) Inv() *X64 {
	*mat = *X64Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *X64) MulVec(vec *vpvec4.X64) *vpvec4.X64 {
	var ret vpvec4.X64

	for i := range vec {
		ret[i] = vpnumber.X64Mul(mat.Get(0, i), vec[0]) + vpnumber.X64Mul(mat.Get(1, i), vec[1]) + vpnumber.X64Mul(mat.Get(2, i), vec[2]) + vpnumber.X64Mul(mat.Get(3, i), vec[3])
	}

	return &ret
}

// MulVecPos performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *X64) MulVecPos(vec *vpvec3.X64) *vpvec3.X64 {
	var ret vpvec3.X64

	for i := range vec {
		ret[i] = vpnumber.X64Mul(mat.Get(0, i), vec[0]) + vpnumber.X64Mul(mat.Get(1, i), vec[1]) + vpnumber.X64Mul(mat.Get(2, i), vec[2]) + mat.Get(3, i)
	}

	return ret.DivScale(vpnumber.X64Mul(mat[Col0Row3], vec[0]) + vpnumber.X64Mul(mat[Col1Row3], vec[1]) + vpnumber.X64Mul(mat[Col2Row3], vec[2]) + mat[Col3Row3])
}

// MulVecDir performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *X64) MulVecDir(vec *vpvec3.X64) *vpvec3.X64 {
	var ret vpvec3.X64

	for i := range vec {
		ret[i] = vpnumber.X64Mul(mat.Get(0, i), vec[0]) + vpnumber.X64Mul(mat.Get(1, i), vec[1]) + vpnumber.X64Mul(mat.Get(2, i), vec[2])
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

// X64Transpose inverts rows and columns (matrix transposition).
// Args is left untouched, a pointer on a new object is returned.
func X64Transpose(mat *X64) *X64 {
	var ret X64

	for c := 0; c < Width; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, mat.Get(r, c))
		}
	}

	return &ret
}

// X64MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func X64MulComp(a, b *X64) *X64 {
	var ret X64

	for c := 0; c < Width; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, vpnumber.X64Mul(a.Get(0, r), b.Get(c, 0))+vpnumber.X64Mul(a.Get(1, r), b.Get(c, 1))+vpnumber.X64Mul(a.Get(2, r), b.Get(c, 2))+vpnumber.X64Mul(a.Get(3, r), b.Get(c, 3)))
		}
	}

	return &ret
}

// X64Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func X64Inv(mat *X64) *X64 {
	ret := X64{
		vpnumber.X64Muln(mat[Col1Row2], mat[Col2Row3], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col1Row3], mat[Col2Row2], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col1Row3], mat[Col2Row1], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col1Row1], mat[Col2Row3], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col1Row2], mat[Col2Row1], mat[Col3Row3]) + vpnumber.X64Muln(mat[Col1Row1], mat[Col2Row2], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row3], mat[Col2Row2], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col2Row3], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col2Row1], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col2Row3], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col0Row2], mat[Col2Row1], mat[Col3Row3]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col2Row2], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row3], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row2], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row1], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row3], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row1], mat[Col3Row3]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row2], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row2], mat[Col2Row1]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row3], mat[Col2Row1]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row1], mat[Col2Row2]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row3], mat[Col2Row2]) + vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row1], mat[Col2Row3]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row2], mat[Col2Row3]),
		vpnumber.X64Muln(mat[Col1Row3], mat[Col2Row2], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col1Row2], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col1Row3], mat[Col2Row0], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col1Row0], mat[Col2Row3], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col1Row2], mat[Col2Row0], mat[Col3Row3]) - vpnumber.X64Muln(mat[Col1Row0], mat[Col2Row2], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row2], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col2Row2], mat[Col3Row0]) + vpnumber.X64Muln(mat[Col0Row3], mat[Col2Row0], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col2Row3], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col2Row0], mat[Col3Row3]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col2Row2], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row2], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row3], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row0], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row3], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row0], mat[Col3Row3]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row2], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row3], mat[Col2Row0]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row2], mat[Col2Row0]) + vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row0], mat[Col2Row2]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row3], mat[Col2Row2]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row0], mat[Col2Row3]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row2], mat[Col2Row3]),
		vpnumber.X64Muln(mat[Col1Row1], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col1Row3], mat[Col2Row1], mat[Col3Row0]) + vpnumber.X64Muln(mat[Col1Row3], mat[Col2Row0], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col1Row0], mat[Col2Row3], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col1Row1], mat[Col2Row0], mat[Col3Row3]) + vpnumber.X64Muln(mat[Col1Row0], mat[Col2Row1], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row3], mat[Col2Row1], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col2Row0], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col2Row3], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col2Row0], mat[Col3Row3]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col2Row1], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row3], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row1], mat[Col3Row0]) + vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row0], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row3], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row0], mat[Col3Row3]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row1], mat[Col3Row3]),
		vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row1], mat[Col2Row0]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row3], mat[Col2Row0]) - vpnumber.X64Muln(mat[Col0Row3], mat[Col1Row0], mat[Col2Row1]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row3], mat[Col2Row1]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row0], mat[Col2Row3]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row1], mat[Col2Row3]),
		vpnumber.X64Muln(mat[Col1Row2], mat[Col2Row1], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col1Row1], mat[Col2Row2], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col1Row2], mat[Col2Row0], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col1Row0], mat[Col2Row2], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col1Row1], mat[Col2Row0], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col1Row0], mat[Col2Row1], mat[Col3Row2]),
		vpnumber.X64Muln(mat[Col0Row1], mat[Col2Row2], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col2Row1], mat[Col3Row0]) + vpnumber.X64Muln(mat[Col0Row2], mat[Col2Row0], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col2Row2], mat[Col3Row1]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col2Row0], mat[Col3Row2]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col2Row1], mat[Col3Row2]),
		vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row1], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row2], mat[Col3Row0]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row0], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row2], mat[Col3Row1]) + vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row0], mat[Col3Row2]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row1], mat[Col3Row2]),
		vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row2], mat[Col2Row0]) - vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row1], mat[Col2Row0]) + vpnumber.X64Muln(mat[Col0Row2], mat[Col1Row0], mat[Col2Row1]) - vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row2], mat[Col2Row1]) - vpnumber.X64Muln(mat[Col0Row1], mat[Col1Row0], mat[Col2Row2]) + vpnumber.X64Muln(mat[Col0Row0], mat[Col1Row1], mat[Col2Row2]),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
