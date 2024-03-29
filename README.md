# goodreads crawler

[![Go Report Card](https://goreportcard.com/badge/github.com/eval-exec/goodreads?style=flat-square)](https://goreportcard.com/report/github.com/eval-exec/goodreads)
[![Coverage](https://codecov.io/gh/eval-exec/goodreads/branch/main/graph/badge.svg)](https://codecov.io/gh/eval-exec/goodreads)
[![codeql-analysis](https://github.com/eval-exec/goodreads/actions/workflows/codeql.yml/badge.svg)](https://github.com/eval-exec/goodreads/actions/workflows/codeql.yml)
[![LICENSE](https://img.shields.io/github/license/eval-exec/goodreads.svg?style=flat-square)](https://github.com/eval-exec/goodreads/blob/main/LICENSE)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/eval-exec/goodreads)


## About The Project


## usage:

- proxy-server: optional
- tag: optional, if empty, all tags will be visited
- debug: optional, enable debug mode?
```bash

> go build .
> ./goodreads -h
a simple goodread quote crawler

Usage:
  goodreads [flags]

Examples:
./goodreads --proxy-server http://127.0.0.1:1081 --tag hope --debug

Flags:
      --debug                 debug mode?
  -h, --help                  help for goodreads
      --proxy-server string   proxy server to use
      --tag string            optional: specify a tag to scrape?

```
```bash
> ./goodreads --proxy-server http://127.0.0.1:1081 --tag hope --debug
```
then check the `goodreads.csv` file
