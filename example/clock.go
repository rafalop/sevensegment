package main

import (
	"time"
  "github.com/rafalop/sevensegment"
)

var colon_segments_on = [2][7]bool{
  {true, true, true, true, true, true, true},
  {false, false, false, false, false, false, false},
}

func flashColon(disp *sevensegment.SevenSegment){
  disp.SetSegments(4, colon_segments_on[0])
  disp.WriteData()
  time.Sleep(500*time.Millisecond)
  disp.SetSegments(4, colon_segments_on[1])
  disp.WriteData()
}

func main() {
  disp := sevensegment.NewSevenSegment(0x70)
  disp.SetBrightness(0)
  disp.SetSegments(4, colon_segments_on[0])
  var second int
  var minute int
  var hour int
  var now time.Time

  current_second := now.Second()
  for {
    now = time.Now()
    second = now.Second()
    minute = now.Minute()
    hour = now.Hour()
    disp.SetDigit(0, minute%10)
    disp.SetDigit(1, minute/10)
    disp.SetDigit(2, hour%10)
    disp.SetDigit(3, hour/10)

 
    if current_second != second {
      flashColon(disp)
      current_second = second
    }
    time.Sleep(10*time.Millisecond)
  }
}
