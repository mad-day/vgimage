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

package normalmath

import "math"
import "github.com/mad-day/vgimage"

type Angle struct{
	X,Y,Z float64
}
func (a *Angle) FromNormal64(n vgimage.Normal64) {
	a.X = math.Atan2(n.Y,n.Z)
	z := math.Sqrt( (n.Y*n.Y) + (n.Z*n.Z) )
	a.Y = math.Atan2(n.X,z)
	a.Z = 0
}

func (a Angle) Normal64() vgimage.Normal64 {
	n := vgimage.Normal64{0,0,1}
	
	//Applying Z will cause nothing
	
	//Apply Y:
	n.X = math.Sin(a.Y)
	n.Z = math.Cos(a.Y)
	
	//Apply X:
	n.Y = math.Sin(a.X)*n.Z
	n.Z = math.Cos(a.X)*n.Z
	return n
}

