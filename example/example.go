package main

import (
	"fmt"
	"github.com/rafalop/sevensegment"
	"time"
)

func main() {
	disp := sevensegment.NewSevenSegment(0x70)

	for i := 0; i <= 150; i++ {
		fmt.Println("Displaying", i)
		disp.SetNum(i)
		disp.WriteData()
		time.Sleep(50 * time.Millisecond)
	}
	disp.Clear()

	for _, char := range "AbcdefghIouS" {
		fmt.Printf("Displaying letter %c at pos 0\n", char)
		disp.SetAlpha(0, byte(char))
		disp.WriteData()
		time.Sleep(1000 * time.Millisecond)
	}
	segments_on := [7]bool{false, false, false, false, false, false, false}
	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			fmt.Println("Turning on segment", j, "at pos", i)
			segments_on[j] = true
			disp.SetSegments(i, segments_on)
			disp.WriteData()
			time.Sleep(200 * time.Millisecond)
		}
		for i, _ := range segments_on {
			segments_on[i] = false
		}
	}
	for i := 0; i <= 15; i++ {
		fmt.Println("Setting display brightness to", i)
		disp.SetBrightness(i)
		time.Sleep(200 * time.Millisecond)
	}
	for i := 15; i >= 0; i-- {
		fmt.Println("Setting display brightness to", i)
		disp.SetBrightness(i)
		time.Sleep(200 * time.Millisecond)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Clearing the board.")
	disp.Clear()

	fmt.Printf("Seven segment testing done.\n")
}
