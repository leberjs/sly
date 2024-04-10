#!/bin/sh

printf "%s" "$(git describe --tags --abbrev=0)" > .version
