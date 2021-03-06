#!/bin/bash

# Vapor is a toolkit designed to support Liquid War 7.
# Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it wil/l be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
#
# Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
# Contact author: ufoot@ufoot.org

if [ -d ../utils ] && [ ! -d utils ] ; then
    cd ..
fi
if [ ! -d utils ] ; then
    echo "$0 should be run in srcdir"
    exit 1
fi

export GOPATH="$(pwd)/go"
export PATH="$(pwd)/go/bin:$PATH"

get() {
    echo "go get $1"
    go get -u $1
}

get github.com/alecthomas/gometalinter
get github.com/tools/godep
get github.com/axw/gocov/gocov
get github.com/AlekSi/gocov-xml
get github.com/jstemmer/go-junit-report
get github.com/ryancox/gobench2plot
get github.com/dineshappavoo/basex
get golang.org/x/crypto/ripemd160
get golang.org/x/crypto/openpgp
get golang.org/x/crypto/openpgp/packet
get git.apache.org/thrift.git/lib/go/thrift
get github.com/llgcode/draw2d

gometalinter --install --update

rm -rf go/src/github.com/ufoot/vapor
install -d go/src/github.com/ufoot/vapor/go
for i in go/vp*; do ln -s $(pwd)/$i go/src/github.com/ufoot/vapor/$i ; done

