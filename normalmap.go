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
	"image"
	"image/color"
)

type NormalMap struct{
	BasePicture
	Normals []float32
}
func NewNormalMap(r image.Rectangle) *NormalMap {
	nm := new(NormalMap)
	nm.BasePicture = NewBasePicture(r)
	nm.Normals = make([]float32,nm.Length)
	return nm
}
func (n NormalMap) SetNormal(x,y int, c Normal) {
	i := n.PixOffset(x,y)*3
	if i<0 { return }
	n.Normals[i  ] = c.X
	n.Normals[i+1] = c.Y
	n.Normals[i+2] = c.Z
}
func (n NormalMap) NormalAt(x,y int) Normal {
	i := n.PixOffset(x,y)*3
	if i<0 { return Normal{0,0,1} }
	return Normal{n.Normals[i],n.Normals[i+1],n.Normals[i+2]}
}
func (n NormalMap) SetNormal64(x,y int, c Normal64) {
	n.SetNormal(x,y,c.Normal())
}
func (n NormalMap) Normal64At(x,y int) Normal64 {
	return n.NormalAt(x,y).Normal64()
}

func (n NormalMap) Set(x,y int, c color.Color) {
	n.SetNormal(x,y,ColorToNormal(c))
}
func (n NormalMap) At(x, y int) color.Color {
	return n.NormalAt(x,y)
}
func (n NormalMap) ColorModel() color.Model { return NormalModel }
func (n NormalMap) Opaque() bool { return true }


type Normal64Map struct{
	BasePicture
	Normals []float64
}
func NewNormal64Map(r image.Rectangle) *Normal64Map {
	nm := new(Normal64Map)
	nm.BasePicture = NewBasePicture(r)
	nm.Normals = make([]float64,nm.Length)
	return nm
}
func (n Normal64Map) SetNormal64(x,y int, c Normal64) {
	i := n.PixOffset(x,y)*3
	if i<0 { return }
	n.Normals[i  ] = c.X
	n.Normals[i+1] = c.Y
	n.Normals[i+2] = c.Z
}
func (n Normal64Map) Normal64At(x,y int) Normal64 {
	i := n.PixOffset(x,y)*3
	if i<0 { return Normal64{0,0,1} }
	return Normal64{n.Normals[i],n.Normals[i+1],n.Normals[i+2]}
}
func (n Normal64Map) SetNormal(x,y int, c Normal) {
	n.SetNormal64(x,y,c.Normal64())
}
func (n Normal64Map) NormalAt(x,y int) Normal {
	return n.Normal64At(x,y).Normal()
}

func (n Normal64Map) Set(x,y int, c color.Color) {
	n.SetNormal64(x,y,ColorToNormal64(c))
}
func (n Normal64Map) At(x, y int) color.Color {
	return n.Normal64At(x,y)
}
func (n Normal64Map) ColorModel() color.Model { return Normal64Model }
func (n Normal64Map) Opaque() bool { return true }

