package main

import (
	"bytes"
	"image"
	"image/png"

	"fyne.io/fyne/v2"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

// SVG خورشید لوکس Suppersun
const sunSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256" width="256" height="256">
  <defs>
    <radialGradient id="bg" cx="50%" cy="50%" r="70%">
      <stop offset="0%" stop-color="#0e3b3a"/>
      <stop offset="100%" stop-color="#072423"/>
    </radialGradient>
    <radialGradient id="core" cx="50%" cy="45%" r="55%">
      <stop offset="0%" stop-color="#fff6c9"/>
      <stop offset="45%" stop-color="#ffd447"/>
      <stop offset="100%" stop-color="#e88b1a"/>
    </radialGradient>
    <linearGradient id="ray" x1="0" y1="0" x2="0" y2="1">
      <stop offset="0%" stop-color="#ffe27a"/>
      <stop offset="100%" stop-color="#f2a516"/>
    </linearGradient>
    <filter id="glow" x="-50%" y="-50%" width="200%" height="200%">
      <feGaussianBlur stdDeviation="4" result="b"/>
      <feMerge><feMergeNode in="b"/><feMergeNode in="SourceGraphic"/></feMerge>
    </filter>
  </defs>

  <rect width="256" height="256" rx="56" fill="url(#bg)"/>
  <circle cx="128" cy="128" r="104" fill="none" stroke="#f3e6c4" stroke-width="2" opacity="0.55"/>
  <circle cx="128" cy="128" r="96"  fill="none" stroke="#3fd0c9" stroke-width="1" opacity="0.35"/>

  <g filter="url(#glow)" fill="url(#ray)">
    <g transform="translate(128 128)">
      <g id="r8">
        <rect x="-5" y="-108" width="10" height="30" rx="5"/>
        <rect x="-3" y="-108" width="6"  height="18" rx="3" fill="#fff2b0" opacity="0.8"/>
      </g>
      <use href="#r8" transform="rotate(45)"/>
      <use href="#r8" transform="rotate(90)"/>
      <use href="#r8" transform="rotate(135)"/>
      <use href="#r8" transform="rotate(180)"/>
      <use href="#r8" transform="rotate(225)"/>
      <use href="#r8" transform="rotate(270)"/>
      <use href="#r8" transform="rotate(315)"/>
    </g>
  </g>

  <circle cx="128" cy="128" r="52" fill="url(#core)" filter="url(#glow)"/>
  <circle cx="128" cy="128" r="52" fill="none" stroke="#fff6c9" stroke-width="1.5" opacity="0.7"/>
  <ellipse cx="114" cy="112" rx="18" ry="10" fill="#ffffff" opacity="0.35"/>
</svg>`

// loadSunIcon آیکون SVG را به fyne.Resource تبدیل می‌کند
func loadSunIcon() fyne.Resource {
	icon, err := oksvg.ReadIconStream(bytes.NewReader([]byte(sunSVG)))
	if err != nil {
		return fyne.NewStaticResource("sun.svg", []byte(sunSVG))
	}

	const size = 256
	icon.SetTarget(0, 0, size, size)
	rgba := image.NewRGBA(image.Rect(0, 0, size, size))
	scanner := rasterx.NewScannerGV(size, size, rgba, rgba.Bounds())
	dasher := rasterx.NewDasher(size, size, scanner)
	icon.Draw(dasher, 1.0)

	var buf bytes.Buffer
	if err := png.Encode(&buf, rgba); err != nil {
		return fyne.NewStaticResource("sun.svg", []byte(sunSVG))
	}
	return fyne.NewStaticResource("sun.png", buf.Bytes())
}
