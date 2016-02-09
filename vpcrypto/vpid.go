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
	"github.com/ufoot/vapor/vperror"
	"math/big"
	"time"
)

// FilterChecker is used to filter and check wether a number,
// typically an id, verifies a property or not.
type FilterChecker interface {
	// Filter processes the value and returns the
	// filtered value.
	Filter(*big.Int) *big.Int
	// Check should return true if number matches property,
	// false if not.
	Check(*big.Int) bool
}

// GenerateID512 generates a 512 bits id, and signs it.
// If key is nil, no signature is generated.
// If filterChecker is not nil, it is garanteed that the property
// is verified by the id.
// If seconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
func GenerateID512(key *Key, filterChecker FilterChecker, seconds int) (*big.Int, []byte, int, error) {
	r := NewRand()
	var ret, tmpInt *big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().Unix(); ret == nil || time.Now().Unix() < start+int64(seconds); {
		tmpInt = Rand512(r, nil)
		if filterChecker != nil {
			tmpInt = filterChecker.Filter(tmpInt)
		}
		if filterChecker == nil || filterChecker.Check(tmpInt) {
			if key != nil {
				tmpData = IntToBuf512(tmpInt)
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return nil, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(Checksum512(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}

// GenerateID256 generates a 256 bits id, and signs it.
// If key is nil, no signature is generated.
// If filterChecker is not nil, it is garanteed that the property
// is verified by the id.
// If seconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
func GenerateID256(key *Key, filterChecker FilterChecker, seconds int) (*big.Int, []byte, int, error) {
	r := NewRand()
	var ret, tmpInt *big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().Unix(); ret == nil || time.Now().Unix() < start+int64(seconds); {
		tmpInt = Rand256(r, nil)
		if filterChecker != nil {
			tmpInt = filterChecker.Filter(tmpInt)
		}
		if filterChecker == nil || filterChecker.Check(tmpInt) {
			if key != nil {
				tmpData = IntToBuf256(tmpInt)
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return nil, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(Checksum256(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}

// GenerateID128 generates a 128 bits id, and signs it.
// If key is nil, no signature is generated.
// If filterChecker is not nil, it is garanteed that the property
// is verified by the id.
// If seconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
func GenerateID128(key *Key, filterChecker FilterChecker, seconds int) (*big.Int, []byte, int, error) {
	r := NewRand()
	var ret, tmpInt *big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().Unix(); ret == nil || time.Now().Unix() < start+int64(seconds); {
		tmpInt = Rand128(r, nil)
		if filterChecker != nil {
			tmpInt = filterChecker.Filter(tmpInt)
		}
		if filterChecker == nil || filterChecker.Check(tmpInt) {
			if key != nil {
				tmpData = IntToBuf128(tmpInt)
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return nil, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(Checksum128(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}

// GenerateID64 generates a 64 bits id, and signs it.
// If key is nil, no signature is generated.
// If filterChecker is not nil, it is garanteed that the property
// is verified by the id.
// If seconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
func GenerateID64(key *Key, filterChecker FilterChecker, seconds int) (uint64, []byte, int, error) {
	r := NewRand()
	var ret, tmpInt uint64
	var tmpBig big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().Unix(); ret == 0 || time.Now().Unix() < start+int64(seconds); {
		tmpInt = Rand64(r, 0)
		tmpBig.SetUint64(tmpInt)
		if filterChecker != nil {
			tmpBig = *filterChecker.Filter(&tmpBig)
		}
		tmpInt = tmpBig.Uint64()
		if filterChecker == nil || filterChecker.Check(&tmpBig) {
			if key != nil {
				tmpData = IntToBuf64(tmpInt)
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return 0, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(Checksum64(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}

// GenerateID32 generates a 32 bits id, and signs it.
// If key is nil, no signature is generated.
// If filterChecker is not nil, it is garanteed that the property
// is verified by the id.
// If seconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
func GenerateID32(key *Key, filterChecker FilterChecker, seconds int) (uint32, []byte, int, error) {
	r := NewRand()
	var ret, tmpInt uint32
	var tmpBig big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().Unix(); ret == 0 || time.Now().Unix() < start+int64(seconds); {
		tmpInt = Rand32(r, 0)
		tmpBig.SetUint64(uint64(tmpInt))
		if filterChecker != nil {
			tmpBig = *filterChecker.Filter(&tmpBig)
		}
		tmpInt = uint32(tmpBig.Uint64())
		if filterChecker == nil || filterChecker.Check(&tmpBig) {
			if key != nil {
				tmpData = IntToBuf32(tmpInt)
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return 0, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(Checksum32(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}
