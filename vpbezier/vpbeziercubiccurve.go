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
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
)

// F32CubicCurve1d returns the cubic Bezier curve from p0 to p3,
// going through p1 and p2. Points p1 and p2 are typically not reached,
// but they influence the curve, as it heads to p1 from p0, and conversely
// from p3 to p2. Works on float32 scalars.
// First returned value is the position, second is the derivative.
func F32CubicCurve1d(p0, p1, p2, p3 float32, beta float32) (float32, float32) {
	switch {
	case vpnumber.F32IsSimilar(p0, p1):
		return F32QuadraticCurve1d(p0, p2, p3, beta)
	case vpnumber.F32IsSimilar(p2, p3):
		return F32QuadraticCurve1d(p0, p1, p3, beta)
	case beta < vpnumber.F32Const0:
		return p0, vpnumber.F32Const0
	case beta > vpnumber.F32Const1:
		return p1, vpnumber.F32Const0
	}

	oneMinusBeta := vpnumber.F32Const1 - beta

	return p0*oneMinusBeta*oneMinusBeta*oneMinusBeta + p1*3*oneMinusBeta*oneMinusBeta*beta + p2*3*oneMinusBeta*beta*beta + p3*beta*beta*beta, (p1-p0)*3*oneMinusBeta*oneMinusBeta + (p2-p1)*6*oneMinusBeta*beta + (p3-p2)*3*beta*beta
}

// F32CubicCurve2d returns the cubic Bezier curve from p0 to p3,
// going through p1 and p2. Points p1 and p2 are typically not reached,
// but they influence the curve, as it heads to p1 from p0, and conversely
// from p3 to p2. Works on float32 2d vectors.
// First returned value is the position, second is the derivative.
func F32CubicCurve2d(p0, p1, p2, p3 *vpvec2.F32, beta float32) (*vpvec2.F32, *vpvec2.F32) {
	switch {
	case p0.IsSimilar(p1):
		return F32QuadraticCurve2d(p0, p2, p3, beta)
	case p2.IsSimilar(p3):
		return F32QuadraticCurve2d(p0, p1, p3, beta)
	case beta < vpnumber.F32Const0:
		return p0, vpvec2.F32New(vpnumber.F32Const0, vpnumber.F32Const0)
	case beta > vpnumber.F32Const1:
		return p1, vpvec2.F32New(vpnumber.F32Const0, vpnumber.F32Const0)
	}

	oneMinusBeta := vpnumber.F32Const1 - beta

	return vpvec2.F32Add(vpvec2.F32MulScale(p0, oneMinusBeta*oneMinusBeta*oneMinusBeta), vpvec2.F32MulScale(p1, 3*oneMinusBeta*oneMinusBeta*beta)).Add(vpvec2.F32MulScale(p2, 3*oneMinusBeta*beta*beta)).Add(vpvec2.F32MulScale(p3, beta*beta*beta)), vpvec2.F32Sub(p1, p0).MulScale(3 * oneMinusBeta * oneMinusBeta).Add(vpvec2.F32Sub(p2, p1).MulScale(6 * oneMinusBeta * beta)).Add(vpvec2.F32Sub(p3, p2).MulScale(3 * beta * beta))
}

// F32CubicCurve3d returns the cubic Bezier curve from p0 to p3,
// going through p1 and p2. Points p1 and p2 are typically not reached,
// but they influence the curve, as it heads to p1 from p0, and conversely
// from p3 to p2. Works on float32 3d vectors.
// First returned value is the position, second is the derivative.
func F32CubicCurve3d(p0, p1, p2, p3 *vpvec3.F32, beta float32) (*vpvec3.F32, *vpvec3.F32) {
	switch {
	case p0.IsSimilar(p1):
		return F32QuadraticCurve3d(p0, p2, p3, beta)
	case p2.IsSimilar(p3):
		return F32QuadraticCurve3d(p0, p1, p3, beta)
	case beta < vpnumber.F32Const0:
		return p0, vpvec3.F32New(vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0)
	case beta > vpnumber.F32Const1:
		return p1, vpvec3.F32New(vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0)
	}

	oneMinusBeta := vpnumber.F32Const1 - beta

	return vpvec3.F32Add(vpvec3.F32MulScale(p0, oneMinusBeta*oneMinusBeta*oneMinusBeta), vpvec3.F32MulScale(p1, 3*oneMinusBeta*oneMinusBeta*beta)).Add(vpvec3.F32MulScale(p2, 3*oneMinusBeta*beta*beta)).Add(vpvec3.F32MulScale(p3, beta*beta*beta)), vpvec3.F32Sub(p1, p0).MulScale(3 * oneMinusBeta * oneMinusBeta).Add(vpvec3.F32Sub(p2, p1).MulScale(6 * oneMinusBeta * beta)).Add(vpvec3.F32Sub(p3, p2).MulScale(3 * beta * beta))
}