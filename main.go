package main

import (
	"github.com/ArseniSkobelev/iactools/cmd/iactools"
	"github.com/Delta456/box-cli-maker/v2"
)

func main() {
	config := box.Config{Px: 2, Py: 1, Type: "Round", TitlePos: "Inside"}
	boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|",
		Config: config}
	boxNew.Println("IaCTools", "Streamline your manual IaC distribution")
	iactools.Execute()
}
