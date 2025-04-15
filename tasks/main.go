package main

import (
	"fmt"

	//if you imports this with .  you do not have to repeat overflow everywhere
	. "github.com/bjartek/overflow/v2"
	"github.com/fatih/color"
)

// TESTING THE EGG WISDOM SMART CONTRACT

func main() {
	o := Overflow(
		WithGlobalPrintOptions(),
		// WithNetwork("testnet"),
	)

	fmt.Println("Testing Contract")

	color.Blue("Egg Wisdom Contract testing")

	color.Green("")

	// Create a new Artist struct
	o.Tx("admin/create_phrase",
		WithSigner("account"),
		WithArg("phrase", "ALL THE ZEN"),
		WithArg("base64Img", ""),
	)

	o.Script("get_all_phrases")

}
