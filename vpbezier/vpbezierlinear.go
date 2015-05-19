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

package vpbezier

import (
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
)

// F32Linear1d returns the linear Bezier curve from p0 to p1,
// which is basically a Lerp. Works on float32 scalars.
// First returned value is the position, second is the derivative.
func F32Linear1d(p0, p1 float32, beta float32) (float32, float32) {
	switch {
	case vpnumber.F32IsSimilar(p0, p1):
		return vpmath.F32Lerp(p0, p1, beta), vpnumber.F32Const0
	case beta < vpnumber.F32Const0:
		return p0, vpnumber.F32Const0
	case beta > vpnumber.F32Const1:
		return p1, vpnumber.F32Const0
	}

	return vpmath.F32Lerp(p0, p1, beta), (p1 - p0) / beta
}

// F32Linear2d returns the linear Bezier curve from p0 to p1,
// which is basically a Lerp. Works on float32 2d vectors.
// First returned value is the position, second is the derivative.
func F32Linear2d(p0, p1 *vpvec2.F32, beta float32) (*vpvec2.F32, *vpvec2.F32) {
	switch {
	case p0.IsSimilar(p1):
		return vpvec2.F32Lerp(p0, p1, beta), vpvec2.F32New(vpnumber.F32Const0, vpnumber.F32Const0)
	case beta < vpnumber.F32Const0:
		return p0, vpvec2.F32New(vpnumber.F32Const0, vpnumber.F32Const0)
	case beta > vpnumber.F32Const1:
		return p1, vpvec2.F32New(vpnumber.F32Const0, vpnumber.F32Const0)
	}

	return vpvec2.F32Lerp(p0, p1, beta), vpvec2.F32Sub(p1, p0).DivScale(beta)
}

// F32Linear3d returns the linear Bezier curve from p0 to p1,
// which is basically a Lerp. Works on float32 3d vectors.
// First returned value is the position, second is the derivative.
func F32Linear3d(p0, p1 *vpvec3.F32, beta float32) (*vpvec3.F32, *vpvec3.F32) {
	switch {
	case p0.IsSimilar(p1):
		return vpvec3.F32Lerp(p0, p1, beta), vpvec3.F32New(vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0)
	case beta < vpnumber.F32Const0:
		return p0, vpvec3.F32New(vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0)
	case beta > vpnumber.F32Const1:
		return p1, vpvec3.F32New(vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0)
	}

	return vpvec3.F32Lerp(p0, p1, beta), vpvec3.F32Sub(p1, p0).DivScale(beta)
}
