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
	//"image/color"
)

func translate(x,y int , r image.Rectangle) (nx int,ny int) {
	nx = x + r.Min.X
	ny = y + r.Min.Y
	return
}

type BasePicture struct{
	Rect image.Rectangle
	Stride int
	Length int
}
func NewBasePicture(r image.Rectangle) BasePicture {
	return BasePicture{r,r.Dx(),r.Dx()*r.Dy()}
}
func (b *BasePicture) PixOffset(x,y int) int {
	if !(image.Point{x, y}.In(b.Rect)) { return -1 }
	nx,ny := translate(x,y,b.Rect)
	return (nx*b.Stride)+ny
}
func (b *BasePicture) Bounds() image.Rectangle { return b.Rect }
