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
	"ufoot.org/vapor/vpnumber"
)

func TestX32Mat3Math(t *testing.T) {
	var x1 = vpnumber.F32ToX32(3.0)
	var x2 = vpnumber.F32ToX32(-4.0)
	var x3 = vpnumber.F32ToX32(1.0)

	var x5 = vpnumber.F32ToX32(-4.5)
	var x6 = vpnumber.F32ToX32(6.0)
	var x7 = vpnumber.F32ToX32(2.0)

	var xmul = vpnumber.F32ToX32(10.0)

	var m1, m2, m3, m4 *X32Mat3

	m1 = X32Mat3New(x1, x2, x3)
	if !X32Mat3IsSimilar(m1, m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToI32().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("I32 conversion error")
	}

	m2 = m1.ToI64().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("I64 conversion error")
	}

	m2 = m1.ToX64().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF32().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("F32 conversion error")
	}

	m2 = m1.ToF64().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("F64 conversion error")
	}

	m2 = X32Mat3New(x5, x6, x7)
	m3 = X32Mat3Add(m1, m2)
	m4 = X32Mat3New(x1+x5, x2+x6, x3+x7)
	if !X32Mat3IsSimilar(m3, m4) {
		t.Error("Add error")
	}

	m3 = X32Mat3Sub(m1, m2)
	m4 = X32Mat3New(x1-x5, x2-x6, x3-x7)
	if !X32Mat3IsSimilar(m3, m4) {
		t.Error("Sub error")
	}

	m3 = X32Mat3MulScale(m1, xmul)
	m4 = X32Mat3New(vpnumber.X32Mul(x1, xmul), vpnumber.X32Mul(x2, xmul), vpnumber.X32Mul(x3, xmul))
	if !X32Mat3IsSimilar(m3, m4) {
		t.Error("MulScale error")
	}

	m3 = X32Mat3DivScale(m3, xmul)
	if !X32Mat3IsSimilar(m3, m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func BenchmarkX32Mat3Add(b *testing.B) {
	mat := X32Mat3New(vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}