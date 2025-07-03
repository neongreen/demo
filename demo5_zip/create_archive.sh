#!/bin/sh
set -e
mkdir -p tmp/b
printf 'a' > tmp/a.txt
printf 'c' > tmp/b/c.txt
(cd tmp && zip -r ../archive.zip a.txt b/c.txt > /dev/null)
rm -r tmp
