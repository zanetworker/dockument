package kubesanityutils

import "github.com/fatih/color"

//ColorString colors a string with color of choice
func ColorString(colorToPrint, messageToPrint string) string {
	var c *color.Color

	switch colorToPrint {
	case "red":
		c = color.New(color.FgRed)
	case "green":
		c = color.New(color.FgGreen)
	}
	return c.Sprint(messageToPrint)
}
