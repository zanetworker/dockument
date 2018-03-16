package utils

import "github.com/fatih/color"

//ColorString colors a string with color of choice
func ColorString(colorToPrint, messageToPrint string) string {
	var c *color.Color

	switch colorToPrint {
	case "red":
		c = color.New(color.FgRed)
	case "green":
		c = color.New(color.FgGreen)
	case "blue":
		c = color.New(color.FgBlue)
	}

	return c.Sprint(messageToPrint)
}
