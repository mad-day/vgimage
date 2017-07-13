/*
 * Copyright (c) 2017 Simon Schmidt
 * 
 * This software is provided 'as-is', without any express or implied
 * warranty. In no event will the authors be held liable for any damages
 * arising from the use of this software.
 * 
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 * 
 * 1. The origin of this software must not be misrepresented; you must not
 *    claim that you wrote the original software. If you use this software
 *    in a product, an acknowledgment in the product documentation would be
 *    appreciated but is not required.
 * 2. Altered source versions must be plainly marked as such, and must not be
 *    misrepresented as being the original software.
 * 3. This notice may not be removed or altered from any source distribution.
 */


package vgimage


import (
	//"image"
	"image/color"
	"math"
)
func normalToChannel64(f float64) uint32{
	// Get it from [-1.0,1.0] to [0.0,2.0]
	f += 1.0
	f *= float64(0xffff)/2.0
	s := int32(f)
	if s<0 { s = 0 } else if s>0xffff { s = 0xffff }
	return uint32(s)
}
func normalToChannel(f float32) uint32{
	// Get it from [-1.0,1.0] to [0.0,2.0]
	f += 1.0
	f *= float32(0xffff)/2.0
	s := int32(f)
	if s<0 { s = 0 } else if s>0xffff { s = 0xffff }
	return uint32(s)
}

type Normal64 struct{
	X,Y,Z float64
}

func ColorToNormal64(col color.Color) Normal64 {
	if n,ok := col.(Normal) ; ok { return n.Normal64() }
	if n,ok := col.(Normal64) ; ok { return n }
	r,g,b,a := col.RGBA()
	dev := 2.0/float64(a)
	
	// If invalid, Fallback to standard Normal
	if math.IsNaN(dev) || math.IsInf(dev,0) {
		return Normal64{0.0,0.0,1.0}
	}
	
	return Normal64{
		(float64(r)*dev)-1.,
		(float64(g)*dev)-1.,
		(float64(b)*dev)-1.,
	}
}
func normal64Model(col color.Color) color.Color {
	return ColorToNormal64(col)
}
var Normal64Model = color.ModelFunc(normal64Model)

func (n Normal64) Normalize() Normal64 {
	x,y,z := n.X,n.Y,n.Z
	sr := math.Sqrt((x*x)+(y*y)+(z*z))
	
	// If invalid, Fallback to standard Normal
	if math.IsNaN(sr) || math.IsInf(sr,0) {
		return Normal64{0.0,0.0,1.0}
	}
	return Normal64{x/sr,y/sr,z/sr}
}
func (n Normal64) RGBA() (r, g, b, a uint32) {
	r = normalToChannel64(n.X)
	g = normalToChannel64(n.Y)
	b = normalToChannel64(n.Z)
	a = 0xffff
	return
}
func (n Normal64) Normal64() Normal64 { return n }
func (n Normal64) Normal() (m Normal) {
	m.X = float32(n.X)
	m.Y = float32(n.Y)
	m.Z = float32(n.Z)
	return
}


type Normal struct{
	X,Y,Z float32
}

func ColorToNormal(col color.Color) Normal {
	if n,ok := col.(Normal) ; ok { return n }
	return ColorToNormal64(col).Normal()
}
func normalModel(col color.Color) color.Color {
	return ColorToNormal(col)
}
var NormalModel = color.ModelFunc(normalModel)

func (n Normal) Normal() Normal { return n }
func (n Normal) Normal64() (m Normal64) {
	m.X = float64(n.X)
	m.Y = float64(n.Y)
	m.Z = float64(n.Z)
	return
}
func (n Normal) Normalize() Normal {
	return n.Normal64().Normalize().Normal()
}
func (n Normal) RGBA() (r, g, b, a uint32) {
	r = normalToChannel(n.X)
	g = normalToChannel(n.Y)
	b = normalToChannel(n.Z)
	a = 0xffff
	return
}

