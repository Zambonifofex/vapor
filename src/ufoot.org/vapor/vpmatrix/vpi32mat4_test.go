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
	"testing"
)

func TestI32Mat4Math(t *testing.T) {
	const i1 = 0
	const i2 = -4
	const i3 = 42
	const i4 = 10000

	const i5 = -10
	const i6 = 1000
	const i7 = 222
	const i8 = -1

	var m1, m2, m3, m4 *I32Mat4

	m1 = I32Mat4New(i1, i2, i3, i4)

	m2 = m1.ToI64().ToI32()
	if *m1 != *m2 {
		t.Error("I64 conversion error")
	}

	m2 = m1.ToX32().ToI32()
	if *m1 != *m2 {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToI32()
	if *m1 != *m2 {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF32().ToI32()
	if *m1 != *m2 {
		t.Error("F32 conversion error")
	}

	m2 = m1.ToF64().ToI32()
	if *m1 != *m2 {
		t.Error("F64 conversion error")
	}

	m2 = I32Mat4New(i5, i6, i7, i8)
	m3 = I32Mat4Add(m1, m2)
	m4 = I32Mat4New(i1+i5, i2+i6, i3+i7, i4+i8)
	if *m3 != *m4 {
		t.Error("Add error")
	}

	m3 = I32Mat4Sub(m1, m2)
	m4 = I32Mat4New(i1-i5, i2-i6, i3-i7, i4-i8)
	if *m3 != *m4 {
		t.Error("Sub error")
	}
}