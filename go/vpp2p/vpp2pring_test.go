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

package vpp2p

import (
	"github.com/ufoot/vapor/go/vpid"
	"github.com/ufoot/vapor/go/vpsum"
	"testing"
)

func TestNewRing(t *testing.T) {
	var host *Host
	var ring *Ring
	var err error
	var zeroes, zeroes2 int

	host, err = NewHost(testTitle, testURL, true, GlobalHostInfoCatalog())
	if err != nil {
		t.Error("unable to create host with a valid pubKey", err)
	}
	ring, err = NewRing(host, testTitle, testDescription, testID, nil, nil, nil)
	if err != nil {
		t.Error("unable to create ring with a valid pubKey", err)
	}
	if ring.IsSigned() == false {
		t.Error("ring is unsigned, when it should be")
	}
	zeroes = vpid.ZeroesInBuf(vpsum.Checksum512(ring.Info.RingSig))
	if zeroes < RingKeyZeroes {
		t.Errorf("Ring created, but not enough zeroes in sig (%d)", zeroes)
	}
	t.Logf("Ring created, number of zeroes in sig is %d", zeroes)
	zeroes2, err = ring.CheckSig()
	if err != nil {
		t.Error("wrong sig", err)
	}
	if zeroes != zeroes2 {
		t.Errorf("RingInfoCheckSig returned bad number of zeroes %d!=%d", zeroes, zeroes2)
	}
	ring.Info.RingSig = ring.Info.HostPubKey
	_, err = ring.CheckSig()
	if err == nil {
		t.Error("failed to report a broken sig", err)
	}

	host, err = NewHost(testTitle, testURL, false, GlobalHostInfoCatalog())
	if err != nil {
		t.Error("unable to create host with a valid pubKey", err)
	}
	ring, err = NewRing(host, testTitle, testDescription, testID, nil, nil, nil)
	if err != nil {
		t.Error("unable to create ring with a valid pubKey", err)
	}
	if ring.IsSigned() == true {
		t.Error("ring is signed, when it should not be")
	}
	zeroes = vpid.ZeroesInBuf(vpsum.Checksum512(ring.Info.RingSig))
	t.Logf("Ring created, number of zeroes in sig is %d", zeroes)
	zeroes2, err = ring.CheckSig()
	if err != nil {
		t.Error("sig reported as wrong when it's legal to have an empty sig")
	}
}

func TestBuiltinRing0(t *testing.T) {
	ring0, err := BuiltinRing0()
	if err != nil {
		t.Error("unable to create builtin ring0", err)
	}

	t.Logf("ring0, RingTitle=%s RingDescription=%s", ring0.Info.RingTitle, ring0.Info.RingDescription)

	if !ring0.IsSigned() {
		t.Error("ring0 is not signed")
	}

	_, err = ring0.CheckSig()
	if err != nil {
		t.Error("ring0 has wrong sig", err)
	}

	ring0.Info.HasPassword = true
	ring0.secret.PasswordHash = []byte("toto")

	_, err = ring0.CheckSig()
	if err == nil {
		t.Error("ring0 reported to have a right sig when it's not")
	}
}
