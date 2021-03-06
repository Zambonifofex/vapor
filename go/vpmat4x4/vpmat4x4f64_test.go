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
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec3"
	"github.com/ufoot/vapor/go/vpvec4"
	"math"
	"math/rand"
	"testing"
)

func TestF64Math(t *testing.T) {
	const f11 = 13.0
	const f12 = 23.0
	const f13 = 33.0
	const f14 = 43.0
	const f21 = -14.0
	const f22 = -24.0
	const f23 = -34.0
	const f24 = -44.0
	const f31 = 11.0
	const f32 = 21.0
	const f33 = 31.0
	const f34 = 41.0
	const f41 = 110.0
	const f42 = 210.0
	const f43 = 310.0
	const f44 = 410.0

	const f51 = -64.15
	const f52 = -74.25
	const f53 = -84.35
	const f54 = -94.45
	const f61 = 66.4
	const f62 = 76.3
	const f63 = 86.2
	const f64 = 96.1
	const f71 = 62.4
	const f72 = 72.3
	const f73 = 82.2
	const f74 = 92.1
	const f81 = -630.01
	const f82 = -730.02
	const f83 = -830.03
	const f84 = -930.04

	const fmul = 10.0

	var m1, m2, m3, m4 *F64

	m1 = F64New(f11, f12, f13, f14, f21, f22, f23, f24, f31, f32, f33, f34, f41, f42, f43, f44)
	if !m1.IsSimilar(m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToX32().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToF32().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("F32 conversion error")
	}

	m2 = F64New(f51, f52, f53, f54, f61, f62, f63, f64, f71, f72, f73, f74, f81, f82, f83, f84)
	m3 = F64Add(m1, m2)
	m4 = F64New(f11+f51, f12+f52, f13+f53, f14+f54, f21+f61, f22+f62, f23+f63, f24+f64, f31+f71, f32+f72, f33+f73, f34+f74, f41+f81, f42+f82, f43+f83, f44+f84)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = F64Sub(m1, m2)
	m4 = F64New(f11-f51, f12-f52, f13-f53, f14-f54, f21-f61, f22-f62, f23-f63, f24-f64, f31-f71, f32-f72, f33-f73, f34-f74, f41-f81, f42-f82, f43-f83, f44-f84)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = F64MulScale(m1, fmul)
	m4 = F64New(f11*fmul, f12*fmul, f13*fmul, f14*fmul, f21*fmul, f22*fmul, f23*fmul, f24*fmul, f31*fmul, f32*fmul, f33*fmul, f34*fmul, f41*fmul, f42*fmul, f43*fmul, f44*fmul)
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = F64DivScale(m3, fmul)
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

func invertableF64() *F64 {
	var ret F64

	for math.Abs(float64(ret.Det())) < 0.25 {
		for i := range ret {
			ret[i] = rand.Float64()
		}
	}

	return &ret
}

func randomVecF64() *vpvec3.F64 {
	var ret vpvec3.F64

	for i := range ret {
		ret[i] = rand.Float64()
	}

	return &ret
}

func TestF64Comp(t *testing.T) {
	m1 := invertableF64()
	m2 := F64Inv(m1)
	id := F64Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestF64Aff(t *testing.T) {
	const p1 = 3.0
	const p2 = 4.0
	const p3 = 5.0
	const t1 = 6.5
	const t2 = 8.5
	const t3 = 10.5

	v1 := vpvec4.F64New(p1, p2, p3, vpnumber.F64Const1)
	vt := vpvec3.F64New(t1, t2, t3)
	mt := F64Translation(vt)
	t.Logf("translation mat4x4 for %s is %s", vt.String(), mt.String())
	v2 := mt.MulVec(v1)
	t.Logf("mat4x4 MulVec %s * %s = %s", mt.String(), v1.String(), v2.String())
	v3 := vpvec4.F64New(p1+t1, p2+t2, p3+t3, vpnumber.F64Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4x4 MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos := mt.MulVecPos(v1.ToVec3())
	v3pos := v1.ToVec3().Add(vt)
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4x4 MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir := mt.MulVecDir(v1.ToVec3())
	v3dir := v1.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4x4 MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}

	mr := F64RotX(math.Pi / 2)
	t.Logf("rotation mat4x4 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat4x4 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec4.F64New(v1[0], -v1[2], v1[1], vpnumber.F64Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4x4 Z rotation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos = mr.MulVecPos(v1.ToVec3())
	v3pos = v3.ToVec3()
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4x4 Z rotation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir = mr.MulVecDir(v1.ToVec3())
	v3dir = v3.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4x4 Z rotation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}

	mr = F64RotY(math.Pi / 2)
	t.Logf("rotation mat4x4 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat4x4 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec4.F64New(v1[2], v1[1], -v1[0], vpnumber.F64Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4x4 Z rotation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos = mr.MulVecPos(v1.ToVec3())
	v3pos = v3.ToVec3()
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4x4 Z rotation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir = mr.MulVecDir(v1.ToVec3())
	v3dir = v3.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4x4 Z rotation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}

	mr = F64RotZ(math.Pi / 2)
	t.Logf("rotation mat4x4 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat4x4 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec4.F64New(-v1[1], v1[0], v1[2], vpnumber.F64Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4x4 Z rotation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos = mr.MulVecPos(v1.ToVec3())
	v3pos = v3.ToVec3()
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4x4 Z rotation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir = mr.MulVecDir(v1.ToVec3())
	v3dir = v3.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4x4 Z rotation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}
}

func TestF64Rebase(t *testing.T) {
	const cx = 1.5
	const cy = 2.0
	const cz = 2.5
	m1 := invertableF64()
	var vo1 vpvec3.F64
	vx1 := vpvec3.F64AxisX()
	vy1 := vpvec3.F64AxisY()
	vz1 := vpvec3.F64AxisZ()
	vp1 := vpvec3.F64Add(vx1, vpvec3.F64Add(vy1, vz1))
	vo2 := randomVecF64()
	vx2 := vpvec3.F64Add(m1.GetCol(0).ToVec3(), vo2)
	vy2 := vpvec3.F64Add(m1.GetCol(1).ToVec3(), vo2)
	vz2 := vpvec3.F64Add(m1.GetCol(2).ToVec3(), vo2)
	vp2 := vpvec3.F64Add(vo2, vpvec3.F64Add(vpvec3.F64MulScale(vpvec3.F64Sub(vx2, vo2), cx), vpvec3.F64Add(vpvec3.F64MulScale(vpvec3.F64Sub(vy2, vo2), cy), vpvec3.F64MulScale(vpvec3.F64Sub(vz2, vo2), cz))))

	m2 := F64RebaseOXYZ(vo2, vx2, vy2, vz2)
	t.Logf("transformation matrix for O=%s X=%s Y=%s Z=%s is M=%s", vo2.String(), vx2.String(), vy2.String(), vz2.String(), m2.String())

	vo3 := m2.MulVecPos(&vo1)
	if !vo3.IsSimilar(vo2) {
		t.Errorf("vo1 -> vo2 error vo1=%s vo2=%s vo3=%s", vo1.String(), vo2.String(), vo3.String())
	}
	vx3 := m2.MulVecPos(vx1)
	if !vx3.IsSimilar(vx2) {
		t.Errorf("vx1 -> vx2 error vx1=%s vx2=%s vx3=%s", vx1.String(), vx2.String(), vx3.String())
	}
	vy3 := m2.MulVecPos(vy1)
	if !vy3.IsSimilar(vy2) {
		t.Errorf("vy1 -> vy2 error vy1=%s vy2=%s vy3=%s", vy1.String(), vy2.String(), vy3.String())
	}
	vz3 := m2.MulVecPos(vz1)
	if !vz3.IsSimilar(vz2) {
		t.Errorf("vz1 -> vz2 error vz1=%s vz2=%s vz3=%s", vz1.String(), vz2.String(), vz3.String())
	}

	m2 = F64RebaseOXYZP(vo2, vx2, vy2, vz2, vp2)
	t.Logf("transformation matrix for O=%s X=%s Y=%s Z=%s P=%s is M=%s", vo2.String(), vx2.String(), vy2.String(), vz2.String(), vp2.String(), m2.String())
	vo3 = m2.MulVecPos(&vo1)
	if !vo3.IsSimilar(vo2) {
		t.Errorf("vo1 -> vo2 error vo1=%s vo2=%s vo3=%s", vo1.String(), vo2.String(), vo3.String())
	}
	vx3 = m2.MulVecPos(vx1)
	if !vx3.IsSimilar(vx2) {
		t.Errorf("vx1 -> vx2 error vx1=%s vx2=%s vx3=%s", vx1.String(), vx2.String(), vx3.String())
	}
	vy3 = m2.MulVecPos(vy1)
	if !vy3.IsSimilar(vy2) {
		t.Errorf("vy1 -> vy2 error vy1=%s vy2=%s vy3=%s", vy1.String(), vy2.String(), vy3.String())
	}
	vz3 = m2.MulVecPos(vz1)
	if !vz3.IsSimilar(vz2) {
		t.Errorf("vz1 -> vz2 error vz1=%s vz2=%s vz3=%s", vz1.String(), vz2.String(), vz3.String())
	}
	vp3 := m2.MulVecPos(vp1)
	if !vp3.IsSimilar(vp2) {
		t.Errorf("vp1 -> vp2 error vp1=%s vp2=%s vp3=%s", vp1.String(), vp2.String(), vp3.String())
	}
}

func TestF64Ortho(t *testing.T) {
	const left float64 = -100
	const right float64 = 900
	const bottom float64 = -100
	const top float64 = 900
	const nearVal float64 = -3
	const farVal float64 = -10
	v0 := vpvec3.F64New(left, bottom, nearVal)
	v1 := vpvec3.F64New(right, top, farVal)
	v0OrthoCheck := vpvec3.F64New(-vpnumber.F64Const1, -vpnumber.F64Const1, -vpnumber.F64Const1)
	v1OrthoCheck := vpvec3.F64New(vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1)

	mOrtho := F64Ortho(left, right, bottom, top, -nearVal, -farVal)
	t.Logf("mOrtho=%s", mOrtho.String())
	v0Ortho := mOrtho.MulVecPos(v0)
	t.Logf("%s * %s = %s", mOrtho.String(), v0.String(), v0Ortho.String())
	if !v0Ortho.IsSimilar(v0OrthoCheck) {
		t.Errorf("problem with ortho proj, got %s should be %s", v0Ortho.String(), v0OrthoCheck.String())
	}
	v1Ortho := mOrtho.MulVecPos(v1)
	t.Logf("%s * %s = %s", mOrtho.String(), v1.String(), v1Ortho.String())
	if !v1Ortho.IsSimilar(v1OrthoCheck) {
		t.Errorf("problem with ortho proj, got %s should be %s", v1Ortho.String(), v1OrthoCheck.String())
	}
}

func TestF64Perspective(t *testing.T) {
	const fovy float64 = 45
	const aspect float64 = 1.5
	const zNear float64 = 3
	const zFar float64 = 10
	const xNearCheck float64 = 0.5364919
	const xFarCheck float64 = 0.16094756
	const yNearCheck float64 = 0.8047378
	const yFarCheck float64 = 0.24142134

	v000 := vpvec3.F64New(vpnumber.F64Const0, vpnumber.F64Const0, -zNear)
	v001 := vpvec3.F64New(vpnumber.F64Const0, vpnumber.F64Const0, -zFar)
	v100 := vpvec3.F64New(vpnumber.F64Const1, vpnumber.F64Const0, -zNear)
	v101 := vpvec3.F64New(vpnumber.F64Const1, vpnumber.F64Const0, -zFar)
	v010 := vpvec3.F64New(vpnumber.F64Const0, vpnumber.F64Const1, -zNear)
	v011 := vpvec3.F64New(vpnumber.F64Const0, vpnumber.F64Const1, -zFar)

	v000PerspectiveCheck := vpvec3.F64New(vpnumber.F64Const0, vpnumber.F64Const0, -vpnumber.F64Const1)
	v001PerspectiveCheck := vpvec3.F64New(vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1)
	v100PerspectiveCheck := vpvec3.F64New(xNearCheck, vpnumber.F64Const0, -vpnumber.F64Const1)
	v101PerspectiveCheck := vpvec3.F64New(xFarCheck, vpnumber.F64Const0, vpnumber.F64Const1)
	v010PerspectiveCheck := vpvec3.F64New(vpnumber.F64Const0, yNearCheck, -vpnumber.F64Const1)
	v011PerspectiveCheck := vpvec3.F64New(vpnumber.F64Const0, yFarCheck, vpnumber.F64Const1)

	mPerspective := F64Perspective(fovy, aspect, zNear, zFar)
	t.Logf("mPerspective=%s", mPerspective.String())
	v000Perspective := mPerspective.MulVecPos(v000)
	t.Logf("%s * %s = %s", mPerspective.String(), v000.String(), v000Perspective.String())
	if !v000Perspective.IsSimilar(v000PerspectiveCheck) {
		t.Errorf("problem with perspective proj, got %s should be %s", v000Perspective.String(), v000PerspectiveCheck.String())
	}
	v001Perspective := mPerspective.MulVecPos(v001)
	t.Logf("%s * %s = %s", mPerspective.String(), v001.String(), v001Perspective.String())
	if !v001Perspective.IsSimilar(v001PerspectiveCheck) {
		t.Errorf("problem with perspective proj, got %s should be %s", v001Perspective.String(), v001PerspectiveCheck.String())
	}
	v100Perspective := mPerspective.MulVecPos(v100)
	t.Logf("%s * %s = %s", mPerspective.String(), v100.String(), v100Perspective.String())
	if !v100Perspective.IsSimilar(v100PerspectiveCheck) {
		t.Errorf("problem with perspective proj, got %s should be %s", v100Perspective.String(), v100PerspectiveCheck.String())
	}
	v101Perspective := mPerspective.MulVecPos(v101)
	t.Logf("%s * %s = %s", mPerspective.String(), v101.String(), v101Perspective.String())
	if !v101Perspective.IsSimilar(v101PerspectiveCheck) {
		t.Errorf("problem with perspective proj, got %s should be %s", v101Perspective.String(), v101PerspectiveCheck.String())
	}
	v010Perspective := mPerspective.MulVecPos(v010)
	t.Logf("%s * %s = %s", mPerspective.String(), v010.String(), v010Perspective.String())
	if !v010Perspective.IsSimilar(v010PerspectiveCheck) {
		t.Errorf("problem with perspective proj, got %s should be %s", v010Perspective.String(), v010PerspectiveCheck.String())
	}
	v011Perspective := mPerspective.MulVecPos(v011)
	t.Logf("%s * %s = %s", mPerspective.String(), v011.String(), v011Perspective.String())
	if !v011Perspective.IsSimilar(v011PerspectiveCheck) {
		t.Errorf("problem with perspective proj, got %s should be %s", v011Perspective.String(), v011PerspectiveCheck.String())
	}
}

func TestF64JSON(t *testing.T) {
	m1 := invertableF64()
	m2 := F64Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F64 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F64")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F64, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F64")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkF64Add(b *testing.B) {
	mat := F64New(vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkF64InvPlot(b *testing.B) {
	mat := invertableF64()

	for i := 0; i < b.N; i++ {
		_ = F64Inv(mat)
	}
}
