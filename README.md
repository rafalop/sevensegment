# sevensegment
Golang module for controlling the Adafruit seven segment LED display with backpack (HT16K33 i2c), using golang's periphio.

Note this was developed for use on a raspberry pi.

Examples of how to use are in example.go which does some simple tasks such as displaying numbers digits at specific positions, displaying a number up to 9999, letters, particular segments and setting the brightness.

## References and links
periphio (i2c)
https://periph.io/
https://godoc.org/periph.io/x/periph/conn/i2c

Adafruit LED backpack HT16k33 datasheet (info about commands/settings for the HT16K33)
https://cdn-shop.adafruit.com/datasheets/ht16K33v110.pdf

Adafruit LED backpack C/C++ library
https://github.com/adafruit/Adafruit_LED_Backpack
