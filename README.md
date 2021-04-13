# goodreads crawler

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