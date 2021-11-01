#!/bin/sh

go run github.com/vektra/mockery/v2@v2.9.0 --case=underscore --disable-version-string $*
