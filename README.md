# gonuma

`gonuma` is a Go utility library for writing NUMA-aware applications

---

[![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/gridfinity/gonuma/master/LICENSE)
[![GoVersion](https://img.shields.io/github/go-mod/go-version/gridfinity/gonuma.svg)](https://github.com/gridfinity/gonuma/blob/master/go.mod)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gridfinity/gonuma)](https://pkg.go.dev/github.com/gridfinity/gonuma)
[![GoReportCard](https://goreportcard.com/badge/github.com/gridfinity/gonuma)](https://goreportcard.com/report/github.com/gridfinity/gonuma)
[![LocCount](https://img.shields.io/tokei/lines/github/gridfinity/gonuma.svg)](https://github.com/XAMPPRocky/tokei)
[![GitHubCodeSize](https://img.shields.io/github/languages/code-size/johnsonjh/gonuma.svg)](https://github.com/gridfinity/gonuma)
[![CoverageStatus](https://coveralls.io/repos/github/gridfinity/gonuma/badge.svg)](https://coveralls.io/github/gridfinity/gonuma)
[![CodacyBadge](https://api.codacy.com/project/badge/Grade/6a688d07faaa4e848f59ec49fdb663bc)](https://app.codacy.com/gh/gridfinity/gonuma?utm_source=github.com&utm_medium=referral&utm_content=gridfinity/gonuma&utm_campaign=Badge_Grade)
[![CodeBeat](https://codebeat.co/badges/041414ca-af27-40f2-a5d6-13afc4ce9c6b)](https://codebeat.co/projects/github-com-gridfinity-gonuma-master)
[![CodeclimateMaintainability](https://api.codeclimate.com/v1/badges/61db603e26c07e0e9ee4/maintainability)](https://codeclimate.com/github/gridfinity/gonuma/maintainability)
[![TickgitTODOs](https://img.shields.io/endpoint?url=https://api.tickgit.com/badge?repo=github.com/gridfinity/gonuma)](https://www.tickgit.com/browse?repo=github.com/gridfinity/gonuma)
[![DeepSource](https://deepsource.io/gh/gridfinity/gonuma.svg/?label=active+issues)](https://deepsource.io/gh/gridfinity/gonuma/?ref=repository-badge)
[![FOSSAStatus](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fjohnsonjh%2Fgonuma.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%gridfinity%2Fgonuma?ref=badge_shield)

---

## Availability

### Go Modules

- [go.gridfinity.dev](https://go.gridfinity.dev/gonuma)
- [go.gridfinity.com](https://go.gridfinity.com/)

### Source Code

- [Gridfinity GitLab](https://gitlab.gridfinity.com/go/gonuma)
- [SourceHut](https://sr.ht/~trn/gonuma)
- [GitHub](https://github.com/gridfinity/gonuma)

## Original Author

- [lrita](https://github.com/lrita/numa)
  \<[lrita@163.com](mailto:lrita@163.com)\>

## Security

- [Security Policy and Vulnerability Reporting](https://github.com/gridfinity/gonuma/blob/master/SECURITY.md)

## Coverage Reports

- [GoCov](https://pktdist.gridfinity.com/coverage/gonuma/)
- [Coveralls](https://coveralls.io/github/johnsonjh/gonuma)

## License

- [MIT License](https://tldrlegal.com/license/mit-license)
- [![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fgridfinity%2Fgonuma.svg?type=small)](https://app.fossa.com/projects/git%2Bgithub.com%2Fgridfinity%2Fgonuma?ref=badge_small)

## Usage

```go
package main

import (
        gonuma "go.gridfinity.dev/gonuma"
)

type object struct {
        X int
        _ [...]byte // pad to page size
}

var objects = make([]object, gonuma.CPUCount())

func fnxxxx() {
        cpu, node := gonuma.GetCPUAndNode()
        objects[cpu].X = xx
}
```
