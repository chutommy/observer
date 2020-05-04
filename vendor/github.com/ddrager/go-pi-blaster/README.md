pi-blaster.go
=============

This package allows go programs to interface with the excellent [pi-blaster](https://github.com/sarfata/pi-blaster) for Raspberry Pi. Since this provides a FIFO buffer typically modified by command line, I thought it would be useful to write a Go package to handle this interface.

Usage
=====

Include the repository and use like example.go:

```
package main

import (
  "fmt"
  "github.com/ddrager/go-pi-blaster"
)



func main() {
  fmt.Printf("Running\n")
  a := []int64{17, 22, 24}
  piblaster.Start(a)

  piblaster.Apply(17, 1);
}
```

**Start** will initialize the interface using the array of GPIO pins which are controlled by pi-blaster (or wish you want to interact with). 

**Apply** will set the GPIO pin.
