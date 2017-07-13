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

package normalmap


import (
	"image"
	"math"
	//"image/color"
	"github.com/mad-day/vgimage"
	"github.com/mad-day/vgimage/copier"
)

var sobelX = [3][3]int32 {
	{1,0,-1},
	{2,0,-2},
	{1,0,-1},
}
var sobelY = [3][3]int32 {
	{1,2,1},
	{0,0,0},
	{-1,-2,-1},
}
const sobelXLtr = int32(1)
const sobelYLtr = int32(1)


var prewittX = [3][3]int32 {
	{-1,0,1},
	{-1,0,1},
	{-1,0,1},
}
var prewittY = [3][3]int32 {
	{-1,-1,-1},
	{0,0,0},
	{1,1,1},
}
const prewittXLtr = int32(-1)
const prewittYLtr = int32(-1)


var scharrX = [3][3]int32 {
	{3 ,0,-3},
	{10,0,-10},
	{3 ,0,-3},
}
var scharrY = [3][3]int32 {
	{3,10,3},
	{0, 0,0},
	{-3,-10,-3},
}
const scharrXLtr = int32(1)
const scharrYLtr = int32(1)


type SourceFilter uint32

const (
	SF_Sobel3x3 = SourceFilter(iota)
	SF_Prewitt3x3
	SF_Scharr3x3
)

type job3x3 struct{
	rectangular
	src *image.Gray16
	dst NormalSink
	operatorX [3][3]int32
	operatorY [3][3]int32
	l2rX int32
	l2rY int32
	intensity float64 // $C/0xffff where $C is the scale factor.
}

func (self *job3x3) calculate(pt image.Point) (int32,int32) {
	pt.X--
	pt.Y--
	ix := int32(0)
	iy := int32(0)
	for x:=0; x<3; x++ {
		for y:=0; y<3; y++ {
			g := int32(self.src.Gray16At(self.clamp(pt.X+x,pt.Y+y)).Y)
			ix += self.operatorX[y][x]*g
			iy += self.operatorY[y][x]*g
		}
	}
	return ix*self.l2rX,iy*self.l2rY
}
func (self *job3x3) Operate(pt image.Point) {
	var n vgimage.Normal64
	sx,sy := self.calculate(pt)
	n.X = clampf(float64(sx)*self.intensity)
	n.Y = clampf(float64(sy)*self.intensity)
	n.Z = 0.0
	// (n.X²+n.Y²+n.Z²) should be == 1
	// n.Z² = 1-(n.X²+n.Y²)
	z2 := 1-((n.X*n.X)+(n.Y*n.Y))
	if z2>0.0 {
		n.Z = clampf(math.Sqrt(z2))
	}
	self.dst.SetNormal64(pt.X,pt.Y,n)
}

/*
Creates a Job, that generates a Normal-Map from a Height-Map using an operator such as Sobel. It returns a nil-Pointer
on failure, (if illegal arguments are given).

Use copier.Operate(Operator,image.Rectangle) to perform the job.

The parameter scaleFactor is the weight of the Height-Map. The default value should be 1.0 !
*/
func CreateGeneratorJob(dst NormalSink,src *image.Gray16, f SourceFilter,scaleFactor float64) copier.Operator {
	jb := new(job3x3)
	jb.rect = src.Bounds()
	jb.src = src
	jb.dst = dst
	switch f {
	case SF_Sobel3x3:
		jb.operatorX = sobelX
		jb.operatorY = sobelY
		jb.l2rX = sobelXLtr
		jb.l2rY = sobelYLtr
	case SF_Prewitt3x3:
		jb.operatorX = prewittX
		jb.operatorY = prewittY
		jb.l2rX = prewittXLtr
		jb.l2rY = prewittYLtr
	case SF_Scharr3x3:
		jb.operatorX = scharrX
		jb.operatorY = scharrY
		jb.l2rX = scharrXLtr
		jb.l2rY = scharrYLtr
	default: return nil
	}
	jb.intensity = scaleFactor/float64(0xffff)
	return jb
}

