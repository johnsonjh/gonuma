# gonuma

`gonuma` is a Go utility library for writing NUMA-aware applications

---

## Original Author

- [lrita](https://github.com/lrita/numa)
  \<[lrita@163.com](mailto:lrita@163.com)\>

## License

- [MIT License](https://tldrlegal.com/license/mit-license)

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
