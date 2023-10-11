# ccwc

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/elliotwutingfeng/ccwc)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliotwutingfeng/ccwc?style=for-the-badge)](https://goreportcard.com/report/github.com/elliotwutingfeng/ccwc)
[![Coveralls](https://img.shields.io/coverallsCoverage/github/elliotwutingfeng/ccwc?logo=coveralls&style=for-the-badge)](https://coveralls.io/github/elliotwutingfeng/ccwc?branch=main)<img src='https://coveralls.io/repos/github/elliotwutingfeng/ccwc/badge.svg?branch=main' alt='' width="0" height="0" />

[![GitHub license](https://img.shields.io/badge/LICENSE-BSD--3--CLAUSE-GREEN?style=for-the-badge)](LICENSE)

## Summary

**ccwc** is a simplified clone of [wc](https://en.wikipedia.org/wiki/Wc_(Unix)) written in Go.

## Usage

First, build the CLI application.

```sh
# `git clone` and `cd` to the ccwc repository folder first
make build_cli
```

Then use the `c`, `l`, `w` and/or `m` flags to count number of bytes, lines, words, and characters in **test.txt**.

```sh
./dist/ccwc -c test.txt # Count bytes
```

```sh
./dist/ccwc -l test.txt # Count lines
```

```sh
./dist/ccwc -w test.txt # Count words
```

```sh
./dist/ccwc -m test.txt # Count characters
```

Piping in from standard input is also supported.

```sh
cat test.txt | ./dist/ccwc -l
```
