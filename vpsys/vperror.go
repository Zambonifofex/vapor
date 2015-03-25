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

package vpsys

import (
	"fmt"
)

// ErrorChain generates an error which combines the error and
// the text passed as an argument. This enables a caller to annotate
// a previous error, generated by a callee.
func ErrorChain(err error, comment string) error {
	if err == nil || comment == "" {
		return err
	}
	return fmt.Errorf("%s (\"%s\")", comment, err.Error())
}

// ErrorChainf generates an error which combines the error and
// the text passed as an argument. This enables a caller to annotate
// a previous error, generated by a callee.
// Allows usage of "à la" printf formatting.
func ErrorChainf(err error, comment string, v ...interface{}) error {
	if err == nil || comment == "" {
		return err
	}
	return fmt.Errorf("%s (\"%s\")", fmt.Sprintf(comment, v...), err.Error())
}