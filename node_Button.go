package gumi

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
)

type nButton struct {
	SingleStructure
	BoundStore
	StyleStore
	//
	cursorEnter, active bool
	//
	onFocus NButtonFocus
	onClick NButtonClick
}

type NButtonClick func()
type NButtonFocus func(focus bool)

func (s *nButton) draw(frame *image.RGBA) {
	var ctx = GGContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	//
	radius := float64(s.bound.Dy() / 2)
	//
	var ok bool
	var clr color.Color
	if s.active {
		_, clr = IsColorImage(s.style.Line)
		ctx.SetColor(clr)
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h-1)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
	} else {
		ok, clr = IsColorImage(s.style.Face)
		if ok {
			ctx.SetColor(clr)
			ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
			ctx.DrawRectangle(radius, 0, w-radius*2, h-1)
			ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
			ctx.Fill()
		} else {
			draw.Draw(frame.SubImage(s.bound).(*image.RGBA), s.bound.Intersect(s.style.Face.Bounds()), s.style.Face, s.style.Face.Bounds().Min, draw.Over)
		}
		if s.cursorEnter {
			_, clr = IsColorImage(s.style.Line)
			ctx.SetColor(clr)
			ctx.DrawLine(radius, 0, w-radius, 0)
			ctx.Stroke()
			ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
			ctx.Stroke()
			ctx.DrawLine(radius, h-1, w-radius, h-1)
			ctx.Stroke()
			ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
			ctx.Stroke()
		}
	}
	//
	s.child.draw(frame)
}

const aBUTTONPADDING = 5

func (s *nButton) size() Size {
	sz := s.child.size()
	sz.Vertical.Min += aBUTTONPADDING * 2
	sz.Horizontal.Min += aBUTTONPADDING * 2
	if !(sz.Horizontal.Min < sz.Horizontal.Max) {
		sz.Horizontal.Max = sz.Horizontal.Min
	}
	if !(sz.Vertical.Min < sz.Vertical.Max) {
		sz.Vertical.Max = sz.Vertical.Min
	}
	return sz
}
func (s *nButton) rect(r image.Rectangle) {
	s.bound = r
	s.child.rect(image.Rect(
		r.Min.X+aBUTTONPADDING-1,
		r.Min.Y+aBUTTONPADDING-1,
		r.Max.X-aBUTTONPADDING+1,
		r.Max.Y-aBUTTONPADDING+1,
	))
}
func (s *nButton) update(info *Information, style *Style) {
	s.style = style
	s.child.update(info, style)
}
func (s *nButton) Occur(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		if ev.Key == KEY_MOUSE1 {
			if s.cursorEnter {
				s.active = true
			}
		}
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.active {
				if s.onClick != nil{
					s.onClick()
				}
				s.active = false
			}
		}
	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		if (s.bound.Min.X <= x && x < s.bound.Max.X) && (s.bound.Min.Y <= y && y < s.bound.Max.Y) {
			if s.onFocus != nil{
				s.onFocus(true)
			}
			s.cursorEnter = true
		} else {
			if s.onFocus != nil{
				s.onFocus(false)
			}
			s.cursorEnter = false
		}
	}
	s.child.Occur(event)
}

//
func NButton(click func()) *nButton {
	return &nButton{
		onClick: click,
	}
}
func NButtonEmpty() *nButton {
	return &nButton{}
}

func (s *nButton) OnClick(callback NButtonClick) {
	s.onClick = callback
}
func (s *nButton) ReferClick() NButtonClick {
	return s.onClick
}

func (s *nButton) OnEnter(callback NButtonFocus) {
	s.onFocus = callback
}
func (s *nButton) ReferEnter() NButtonFocus {
	return s.onFocus
}