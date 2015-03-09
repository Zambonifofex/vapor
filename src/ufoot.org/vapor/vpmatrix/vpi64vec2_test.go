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

func TestI64Vec2Math(t *testing.T) {
	const i1 = 0
	const i2 = -4

	const i5 = -10
	const i6 = 1000

	var v1, v2, v3, v4 *I64Vec2

	v1 = I64Vec2New(i1, i2)

	v2=v1.ToI32().ToI64()
	if *v1!=*v2 {
		t.Error("I32 conversion error")
	}	

	v2=v1.ToX32().ToI64()
	if *v1!=*v2 {
		t.Error("X32 conversion error")
	}	

	v2=v1.ToX64().ToI64()
	if *v1!=*v2 {
		t.Error("X64 conversion error")
	}	

	v2=v1.ToF32().ToI64()
	if *v1!=*v2 {
		t.Error("F32 conversion error")
	}	
	
	v2=v1.ToF64().ToI64()
	if *v1!=*v2 {
		t.Error("F64 conversion error")
	}	
	
	v2 = I64Vec2New(i5, i6)
	v3 = I64Vec2Add(v1, v2)
	v4 = I64Vec2New(i1+i5, i2+i6)
	if *v3 != *v4 {
		t.Error("Add error")
	}

	v3 = I64Vec2Sub(v1, v2)
	v4 = I64Vec2New(i1-i5, i2-i6)
	if *v3 != *v4 {
		t.Error("Sub error")
	}
}
