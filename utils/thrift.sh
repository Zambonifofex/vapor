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

THRIFT_GO_OPTIONS="package_prefix=github.com/ufoot/vapor/go/"
THRIFT_CPP_OPTIONS=""
THRIFT_JS_OPTIONS="jquery"
THRIFT_HTML_OPTIONS="standalone"

for i in common bus p2p ; do
    cd thrift && \
	thrift --gen "go:$THRIFT_GO_OPTIONS" --gen "cpp:$THIFT_CPP_OPTIONS" --gen "js:$THRIFT_JS_OPTIONS" --gen "html:$THRIFT_HTML_OPTIONS" vp${i}api.thrift && \
	cp ./gen-go/vp${i}api/*.go ../go/vp${i}api/ && \
	cp ./gen-go/vp${i}api/vp_${i}_api-remote/vp_${i}_api-remote.go ../go/vp${i}client/vp${i}client.go && \
	sed -i "s/\"vp${i}api\"/\"github.com\/ufoot\/vapor\/go\/vp${i}api\"/" ../go/vp${i}client/vp${i}client.go && \
	cp ./gen-html/vp${i}api.html ../doc/thrift/vp${i}api-thrift.html || exit 1
    cd ..
    rm -f go/vp*/vp_*.go
done

if [ -f ./src/git.apache.org/thrift.git/lib/js/src/thrift.js ] && [ -d ./thrift/gen-js/ ] ; then cp ./src/git.apache.org/thrift.git/lib/js/src/thrift.js ./thrift/gen-js/ ; fi
if which wget > /dev/null 2>&1 ; then wget https://code.jquery.com/jquery.js -O ./thrift/gen-js/jquery.js ; fi

