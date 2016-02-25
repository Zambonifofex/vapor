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
	"github.com/ufoot/vapor/vpcrypto"
	"testing"
)

func init() {
	testKey, err := vpcrypto.NewKey()
	if err != nil {
		panic("vpcrypto.NewKey failed")
	}
	testPubKey, err = testKey.ExportPub()
	if err != nil {
		panic("vpcrypto.ExportPub failed")
	}
}

func TestNewHost(t *testing.T) {
	t.Logf("testPubKey len=%d", len(testPubKey))
	host, err := NewHost(testHostTitle, testHostURL, testPubKey)
	if err != nil {
		t.Error("unable to create host with a valid pubKey", err)
	}
	if host.CanSign() == false {
		t.Error("host can't sign but it should")
	}
	t.Logf("testFakeKey len=%d", len(testFakeKey))
	host, err = NewHost(testHostTitle, testHostURL, testFakeKey)
	if err != nil {
		t.Error("unable to create host with a dummy pubKey", err)
	}
	if host.CanSign() == true {
		t.Error("host can sign but it shouldn't")
	}
}
