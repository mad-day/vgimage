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
	"github.com/mad-day/vgimage/copier"
	"github.com/mad-day/vgimage/normalmath"
)

type jobStacking struct{
	src  []NormalSource
	dst  NormalSink
	nSrc int
}

func (self *jobStacking) Operate(pt image.Point) {
	A := normalmath.Quaternion{}
	C := normalmath.Quaternion{}
	
	A.FromNormal64(self.src[0].Normal64At(pt.X,pt.Y))
	for i,n := 1,self.nSrc; i<n; i++ {
		C.FromNormal64(self.src[i].Normal64At(pt.X,pt.Y))
		A = A.Multiply(C)
	}
	self.dst.SetNormal64(pt.X,pt.Y,A.Normal64())
}

/*
Creates a Job, that stacks multiple Normal-Maps above each other. The number of
sources must be at least 1, otherwise, the function fails and a nil-pointer is returned.

Use copier.Operate(Operator,image.Rectangle) to perform the job.
*/
func NewStackingJob(dst NormalSink,src ...NormalSource) copier.Operator {
	if len(src)==0 { return nil }
	js := new(jobStacking)
	js.src = src
	js.dst = dst
	js.nSrc = len(src)
	return js
}

