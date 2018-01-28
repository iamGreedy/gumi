package main

import (
	"github.com/iamGreedy/gumi"
	"github.com/iamGreedy/gumi/gutl"
)

func main() {
	scr := gumi.NewScreen(gutl.DefinedResolutions.Get("VGA"))
	scr.Root(gumi.LinkingFrom(
		gumi.NDrawing1(gumi.BuildRuler(
			gumi.RULER_HINT_VERTICAL |gumi.RULER_HINT_HORIZONTAL,
			100,
		)...),
		gumi.NBackground(),
		gumi.NMargin(gumi.RegularBlank(gumi.MinLength(30))),
		gumi.NStyle(gumi.DefaultDarkTheme.Style(gumi.INTENSE3)),
		gumi.NBackground(),
		gumi.NVertical1(
			gumi.LinkingFrom(
				gumi.NMargin(gumi.RegularBlank(gumi.MinLength(30))),
				gumi.AText("Hello, World!", gumi.Align_CENTER),
			),
			gumi.AText("Hello, World!", gumi.Align_CENTER),
		),
	))

	gumi.Capture("out", scr.Frame())
}
