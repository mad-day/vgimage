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

/*
A Component for working with Normal-Maps.
*/
package normalmap

import "image"
import "github.com/mad-day/vgimage"
import "math"

func clampf(f float64) float64 {
	if math.IsNaN(f) || math.IsInf(f,0) { return 0.0 }
	if f< -1.0 { return -1.0 }
	if f>  1.0 { return 1.0 }
	return f
}

func clamp16(i int32) uint32 {
	if i<0 { return 0 }
	if i>0xffff { return 0xffff }
	return uint32(i)
}

type rectangular struct{
	rect image.Rectangle
}
func (r *rectangular) clamp(x, y int) (int,int) {
	if x<r.rect.Min.X { x=r.rect.Min.X }
	if x>r.rect.Max.X { x=r.rect.Max.X }
	if y<r.rect.Min.Y { y=r.rect.Min.Y }
	if y>r.rect.Max.Y { y=r.rect.Max.Y }
	return x,y
}

type NormalSink interface{
	SetNormal64(x,y int, c vgimage.Normal64)
	SetNormal(x,y int, c vgimage.Normal)
}

