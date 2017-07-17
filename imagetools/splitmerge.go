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


package imagetools

import (
	"image"
	"image/color"
	"github.com/mad-day/vgimage/copier"
)



type jobYuvSplitAndMerge struct{
	in1 ,in2  YuvInput
	out1,out2 YuvOutput
}

func (self *jobYuvSplitAndMerge) Operate(pt image.Point) {
	c1 := self.in1.YCbCrAt(pt.X,pt.Y)
	c2 := self.in2.YCbCrAt(pt.X,pt.Y)
	c1.Y,c2.Y = c2.Y,c1.Y
	self.out1.SetYCbCr(pt.X,pt.Y,c1)
	self.out2.SetYCbCr(pt.X,pt.Y,c2)
}

func SplitImage(dstColor image.Image, dstLuma image.Image, src image.Image,neutral color.Color) copier.Operator {
	j := &jobYuvSplitAndMerge {
		NewYuvInput(src),yuvUniform{color.YCbCrModel.Convert(neutral).(color.YCbCr)},
		ToYuvOutput(dstColor),ToYuvOutput(dstLuma),
	}
	return j
}
func MergeImage(dst image.Image, srcColor image.Image, srcLuma image.Image) copier.Operator {
	j := &jobYuvSplitAndMerge {
		NewYuvInput(srcColor),NewYuvInput(srcLuma),
		ToYuvOutput(dst),nilYuvOutput{},
	}
	return j
}

