# lantern
Command line tool for pattern search in files

## Installation
Install with [Gobinaries](https://gobinaries.com/):
```
$ curl -sf https://gobinaries.com/alveflo/lantern | sh
```

## Usage
```
$ lantern [searchpattern]
```
Note that regular expressions is supported.

## Example
Example showing searching for `package \w+$` in this
repository
```
$ lantern "package \w+$"
/lantern/main.go (1:0):
package main
/lantern/pkg/crawler/crawler.go (1:0):
package crawler
/lantern/pkg/crawler/filereader.go (1:0):
package crawler
/lantern/pkg/crawler/filereader_test.go (1:0):
package crawler
```
