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

package vpkoorde

import (
	"math/big"
)

// NodeInfo contains the basics for a node, id and url.
type NodeInfo struct {
	// The node Id, is technically the same as a key Id.
	NodeID *big.Int
	// Url can be used and should be enough to connect to node.
	ConnectURL string
	// Title is a human-readable text identifying the node, need not be unique.
	Title string
	// Description is a human-readable description of the node.
	Description string
}

// Node implements basic funcs over a node
type Node struct {
	// The node details.
	Info NodeInfo
}