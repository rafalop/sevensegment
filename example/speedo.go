// Note your ELM device has to be already paired outside this app (eg. using bluetoothctl) 
// arg[1] is device address, arg[2] is the device file you want to bind it to
// in order to do serial comms.

package main

import (
	"time"
  "github.com/rafalop/sevensegment"
  "os"
  "os/exec"
  "strconv"
  "github.com/rzetterberg/elmobd"
  "flag"
  "log"
)



func main() {
  // Pass device mac addr as first arg
  dev_mac := os.Args[1]
  // Pass device to bind to as second arg
  bind_dev := os.Args[2]
  disp := sevensegment.NewSevenSegment(0x70)
  disp.SetBrightness(0)
  disp.SetNum(0)
  disp.WriteData()

  //var out []byte
  var err error
  sleeptime := 1
  cmd := exec.Command(`sudo`, `l2ping`, dev_mac, `-c1`, `-t2`)
  for {
    err = cmd.Run()
    if err != nil {
      log.Println("Could not ping ELM bluetooth dev, trying again in ", sleeptime, " seconds...")
      time.Sleep(time.Duration(sleeptime)*time.Second)
      sleeptime = sleeptime+1
    } else {
      break
    }
  }
  cmd = exec.Command(`sudo`, `bind`, `0`, dev_mac, bind_dev)
  for {
    err = cmd.Run()
    if err != nil {
      log.Println("Could not bind bluetooth device trying again in ", sleeptime, " seconds...")
      time.Sleep(time.Duration(sleeptime)*time.Second)
      sleeptime = sleeptime+1
    } else {
      break
    }
  }


//  current_second := now.Second()
	serialPath := flag.String(
		"serial",
		bind_dev,
		"Path to the serial device to use",
	)

	flag.Parse()

	dev, err := elmobd.NewDevice(*serialPath, false)
	//dev, err := elmobd.NewTestDevice(*serialPath, false)

	if err != nil {
		log.Println("Failed to create new device", err)
		return
	}

  var speed_s string
  var speed int
  brightness_hours := []int{0,0,0,0,0,0,
                            7,8,9,10,11,12,
                            15,15,15,15,15,
                            15,10,7,0,0,0}
  var old_h = time.Now().Hour()
  var new_h int
  disp.SetBrightness(brightness_hours[old_h])
  for {
    new_h = time.Now().Hour()
    if new_h != old_h {
      disp.SetBrightness(brightness_hours[new_h])
      old_h = new_h
    }
	  speed_t, err := dev.RunOBDCommand(elmobd.NewVehicleSpeed())
	  if err != nil {
		  log.Println("Failed to get speed", err)
		  return
	  }
    speed_s = speed_t.ValueAsLit()
    speed,_ = strconv.Atoi(speed_s)
    disp.SetNum(speed)
    disp.WriteData()
    time.Sleep(50*time.Millisecond)
  }
}
