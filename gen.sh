#!/bin/bash

elegen folder -d v1/specs -o codegen || exit 1

rm -rf ./v1/golang/*.go
mv codegen/elemental/*.go ./v1/golang
rm -rf codegen
