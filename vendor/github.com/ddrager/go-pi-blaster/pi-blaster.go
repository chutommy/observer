package piblaster

import (
  "os"
  "fmt"
  //"bytes"
  "strconv"
)

type Blaster struct {
  active []int64
  Pins []float64
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func (b *Blaster) Start(start []int64) {

  //TODO: Ensure we are running pi and have pi-blaster running

  const fifo = "/dev/pi-blaster"

  copy(start, b.active)
  b.active = start
  // initialize array used to keep track of pin value on updates
  b.Pins = make([]float64, 26)

  f, err := os.Create(fifo)
  check(err)
  defer f.Close()

  // set all active pin values to 0
  for i := range b.active {
    // fmt.Printf("Set pin %d to 0\n", b.active[i])
    f.WriteString(strconv.FormatInt(b.active[i], 10) + "=0\n") 
    b.Pins[b.active[i]-1] = 0.0
    f.Sync()
    //check(err)
  }

  f.Close()
}

func (b *Blaster) Apply(pin int64, value float64) {
  f, err := os.Create("/dev/pi-blaster")
  check(err)
  defer f.Close()

  // ensure set value > 0, < 1
  if value > 1.0 {
    fmt.Printf("Request value exceeds 1, setting to 1\n")
    value = 1.0
  } else if value < 0.0 {
    fmt.Printf("Requested value below 0, setting to 0\n")
    value = 0.0
  }

  var toVal string
  toVal = strconv.FormatFloat(value, 'f', 2, 64)
  f.WriteString(strconv.FormatInt(pin, 10) + "=" + toVal + "\n")
  // n1, err := f.WriteString(strconv.FormatInt(pin, 10) + "=" + toVal + "\n")
  b.Pins[pin-1] = value
  // fmt.Printf("wrote %d bytes (%d = %s)\n", n1, pin, toVal)
  f.Sync()
}

func (b *Blaster) DumpCurrent() {
  fmt.Printf("Dumping Current Values...\n")
  for i := range b.active {
    fmt.Printf("%d = %f\n", b.active[i], b.Pins[b.active[i]-1]) 
  }
}

func (b *Blaster) Reset() {
  fmt.Printf("Resetting Pins\n")
  //TODO: Ensure we are running pi and have pi-blaster running

  const fifo = "/dev/pi-blaster"

  b.active = make([]int64, 0)
  // initialize array used to keep track of pin value on updates
  b.Pins = make([]float64, 26)

  f, err := os.Create(fifo)
  check(err)
  defer f.Close()

  // set all active pin values to 0
  for i := range b.active {
    // fmt.Printf("Set pin %d to 0\n", b.active[i])
    f.WriteString(strconv.FormatInt(b.active[i], 10) + "=0\n") 
    b.Pins[b.active[i]-1] = 0.0
    f.Sync()
    //check(err)
  }

  f.Close()

}

