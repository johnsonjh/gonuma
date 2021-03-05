# gonuma

----

`gonuma` is a Go utility library for writing NUMA-aware applications

----

## Availability

- [GitHub](https://github.com/johnsonjh/gonuma)
- [GitLab](https://gitlab.com/johnsonjh/gonuma)
- [SourceHut](https://sr.ht/~trn/gonuma)
- [NotABug](https://notabug.org/trn/gonuma)

----

## Original Author

- [lrita](https://github.com/lrita/numa)
  \<[lrita@163.com](mailto:lrita@163.com)\>

----

## License

- [MIT License](https://tldrlegal.com/license/mit-license)

----

## Usage

```go
package main

import (
        gonuma "github.com/johnsonjh/gonuma"
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

----
