#!/bin/bash

#
# this function runs unit tests on all the packages that are defined
# here.
# 
function run_unit_tests() {
    echo "*** Package Unit Tests ***"
    go test ./...
}

#
# this function runs benchmark tests on all the packages that are
# defined here.
# 
function run_benchmarks() {
    echo "*** Package Benchmarks ***"
    pkg_list=`find . -maxdepth 2 -type d | egrep -v -e '(git|\.$|data|obj)'`
    for pkg in $pkg_list
    do
	echo "- `basename $pkg`"
	go test -run=NONE -bench=. $pkg | grep Benchmark | awk '{print "    -", $1, "\t", $3, "\t", $4}'
    done
}

function main() {
    run_unit_tests
    run_benchmarks
}

# start it all
main
