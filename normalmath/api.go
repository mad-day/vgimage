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

import "github.com/mad-day/vgimage"

func CompbineNormals64(norms ...vgimage.Normal64) vgimage.Normal64{
	if len(norms)==0 { return vgimage.Normal64{0,0,1} }
	if len(norms)==1 { return norms[0] }
	T := Angle{}
	A := Quaternion{}
	N := Quaternion{}
	
	T.FromNormal64(norms[0])
	A.FromAngle(T)
	for _,e := range norms[1:] {
		T.FromNormal64(e)
		N.FromAngle(T)
		A = A.Multiply(N)
	}
	A.ToAngle(&T)
	return T.Normal64()
}