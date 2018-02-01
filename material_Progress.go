package gumi

import (
	"github.com/fogleman/gg"
	"image"
)

const (
	mtProgressMinWidth   = 40
	mtProgressMinHeight  = 8
	mtProgressAnimMillis = 600
)

type mtProgress struct {
	VoidStructure
	BoundStore
	StyleStore
	//
	mtColorFromTo
	//
	from float64
	cur  float64
	to   float64
	anim float64

	//
	cursorEnter, active bool
}

func (s *mtProgress) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	radius := h / 2
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	//
	mcl1 := s.style.Material.PalletteColor(s.mcl1)
	mcl2 := s.style.Material.PalletteColor(s.mcl2)
	//
	ctx.SetColor(phaseColor(mcl1[0], mcl2[0], s.anim))
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()

	//
	rectw := (w - radius*2) * s.anim
	//
	ctx.SetColor(phaseColor(mcl1[1], mcl2[1], s.anim))
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, rectw, h)
	ctx.DrawArc(radius+rectw, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
}

func (s *mtProgress) size() Size {
	return Size{
		MinLength(mtProgressMinHeight),
		MinLength(mtProgressMinWidth),
	}
}

func (s *mtProgress) rect(r image.Rectangle) {
	s.bound = r
}

func (s *mtProgress) update(info *Information, style *Style) {
	s.style = style
	if s.cur != s.to {
		if s.cur < s.to {
			s.cur += float64(info.Dt) / mtProgressAnimMillis
			if s.cur > s.to {
				s.cur = s.to
			}

			//
			s.anim = Material.Progress(s.cur)
		} else {
			s.cur -= float64(info.Dt) / mtProgressAnimMillis
			if s.cur < s.to {
				s.cur = s.to
			}
			s.anim = Material.Progress(s.cur)
		}

	}
}

func (s *mtProgress) Occur(event Event) {

}

//
func MTProgress(from, to MaterialColor, percent float64) *mtProgress {
	temp := &mtProgress{
		to: percent,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}
func MTProgress1(mcl MaterialColor, percent float64) *mtProgress {
	temp := &mtProgress{
		to: percent,
	}
	temp.SetFromMaterialColor(mcl)
	temp.SetToMaterialColor(mcl)
	return temp
}

//
func (s *mtProgress) Get() float64 {
	return s.to
}
func (s *mtProgress) Set(percent float64) {
	s.from = s.cur
	s.to = percent
}

func (s *mtProgress) GetFromMaterialColor() MaterialColor {
	return s.mcl1
}
func (s *mtProgress) SetFromMaterialColor(mcl MaterialColor) {
	s.mcl1 = mcl
}
func (s *mtProgress) GetToMaterialColor() MaterialColor {
	return s.mcl2
}
func (s *mtProgress) SetToMaterialColor(mcl MaterialColor) {
	s.mcl2 = mcl
}
