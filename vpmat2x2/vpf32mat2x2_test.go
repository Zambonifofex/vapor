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

package vpmat2x2

import (
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"math"
	"math/rand"
	"testing"
)

func TestF32Mat2x2Math(t *testing.T) {
	const f11 = 3.0
	const f12 = 3.0
	const f21 = -8.0
	const f22 = -4.0

	const f51 = -4.5
	const f52 = -4.1
	const f61 = 6.0
	const f62 = 6.6

	const fmul = 10.0

	var m1, m2, m3, m4 *F32Mat2x2

	m1 = F32Mat2x2New(f11, f12, f21, f22)
	if !m1.IsSimilar(m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToI32().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("I32 conversion error")
	}

	m2 = m1.ToI64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("I64 conversion error")
	}

	m2 = m1.ToX32().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("F64 conversion error")
	}

	m2 = F32Mat2x2New(f51, f52, f61, f62)
	m3 = F32Mat2x2Add(m1, m2)
	m4 = F32Mat2x2New(f11+f51, f12+f52, f21+f61, f22+f62)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = F32Mat2x2Sub(m1, m2)
	m4 = F32Mat2x2New(f11-f51, f12-f52, f21-f61, f22-f62)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = F32Mat2x2MulScale(m1, fmul)
	m4 = F32Mat2x2New(f11*fmul, f12*fmul, f21*fmul, f22*fmul)
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = F32Mat2x2DivScale(m3, fmul)
	if !m3.IsSimilar(m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func invertableF32Mat2x2() *F32Mat2x2 {
	var ret F32Mat2x2

	for math.Abs(float64(ret.Det())) < 0.5 {
		for i := range ret {
			ret[i] = rand.Float32()
		}
	}

	return &ret
}

func TestF32Mat2x2Comp(t *testing.T) {
	m1 := invertableF32Mat2x2()
	m2 := F32Mat2x2Inv(m1)
	id := F32Mat2x2Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestF32Mat2x2Aff(t *testing.T) {
	const p1 = 3.0
	const t1 = 6.0

	v1 := vpvec2.F32Vec2New(p1, vpnumber.F32Const1)
	mt := F32Mat2x2Trans(t1)
	t.Logf("translation mat2 for %f is %s", p1, mt.String())
	v2 := mt.MulVec(v1)
	t.Logf("mat2 MulVec %s * %s = %s", mt.String(), v1.String(), v2.String())
	v3 := vpvec2.F32Vec2New(p1+t1, vpnumber.F32Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat2 MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos := mt.MulVecPos(p1)
	v3pos := float32(p1 + t1)
	if !vpnumber.F32IsSimilar(v2pos, v3pos) {
		t.Errorf("mat2 MulVecPos error v2pos=%f v3pos=%f", v2pos, v3pos)
	}
	v2dir := mt.MulVecDir(p1)
	v3dir := float32(p1)
	if !vpnumber.F32IsSimilar(v2dir, v3dir) {
		t.Errorf("mat2 MulVecDir error v2dir=%f v3dir=%f", v2dir, v3dir)
	}
}

func TestF32Mat2x2JSON(t *testing.T) {
	m1 := invertableF32Mat2x2()
	m2 := F32Mat2x2Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F32Mat2x2 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F32Mat2x2")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F32Mat2x2, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F32Mat2x2")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkF32Mat2x2Add(b *testing.B) {
	mat := F32Mat2x2New(vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkF32Mat2x2Inv(b *testing.B) {
	mat := invertableF32Mat2x2()

	for i := 0; i < b.N; i++ {
		_ = F32Mat2x2Inv(mat)
	}
}