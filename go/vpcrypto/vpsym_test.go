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

package vpcrypto

import (
	"encoding/hex"
	"fmt"
	"github.com/ufoot/vapor/go/vperror"
	"testing"
)

var benchSymPassword []byte
var benchSymContent []byte
var benchSymCrypted []byte

func init() {
	benchSymContent = make([]byte, 1500)
	benchSymPassword = []byte("abcdefghijklmnopqrstuvwxyz")
	fmt.Printf("Encrypting 0...0/%s", string(benchSymPassword))
	benchSymCrypted, _ = SymEncrypt(benchSymContent, benchSymPassword)
}

func TestSym(t *testing.T) {
	content := []byte("foo bar")
	password := []byte("0123456789abcdef")
	var encrypted []byte
	var decrypted []byte
	var err error

	t.Logf("Encrypting %s/%s", string(content), string(password))
	encrypted, err = SymEncrypt(content, password)
	if err == nil {
		t.Logf("encrypted content=\"%s\" encrypted=\"%s\"", string(content), hex.EncodeToString(encrypted))
		decrypted, err = SymDecrypt(encrypted, password)
		if err == nil {
			t.Logf("decrypted encrypted=\"%s\" decrypted=\"%s\"", hex.EncodeToString(encrypted), string(decrypted))
			if string(content) != string(decrypted) {
				t.Error("Content and decrypted differ")
			}
		} else {
			t.Error(err)
		}
		t.Log("test used to block here due to my not knowing how to use a prompt func, see https://go-review.googlesource.com/#/c/16865/")
		_, err = SymDecrypt(encrypted, []byte("this is a wrong password"))
		if err != nil {
			t.Log("OK, decrypt is impossible with a bad password")
		} else {
			t.Error("decrypt is possible with a bad password, this *should* be impossible")
		}
	} else {
		t.Error(err)
	}

}

func BenchmarkSymEncryptPlot(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = SymEncrypt(benchSymContent, benchSymPassword)
		if err != nil {
			b.Error(vperror.Chain(err, "unable to encrypt"))
		}
	}
}

func BenchmarkSymDecryptPlot(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = SymDecrypt(benchSymCrypted, benchSymPassword)
		if err != nil {
			b.Error(vperror.Chain(err, "unable to decrypt"))
		}
	}
}
