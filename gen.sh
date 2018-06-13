#!/bin/bash

elegen folder -d specs -o codegen || exit 1

rm -rf ./*.go
mv codegen/elemental/*.go ./
rm -rf codegen
