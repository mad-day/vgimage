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
 *
 *---------------------------------------------------------------------------
 *
 * SPECIAL THANKS: to the Irrlicht Engine team for their Implementation of the
 *                 quaternion. This code has been inspired by that.
 */


package normalmath

import "math"
import "github.com/mad-day/vgimage"

type Quaternion struct{
	X,Y,Z,W float64
}

func (q *Quaternion) Normalize() {
	n := math.Sqrt( (q.X*q.X) + (q.Y*q.Y) + (q.Z*q.Z) + (q.W*q.W) )
	if math.IsNaN(n) || math.IsInf(n,0) { return }
	q.X/=n
	q.Y/=n
	q.Z/=n
}
func (q *Quaternion) FromAngle(a Angle) {
	angle := a.X * 0.5;
	sr := math.Sin(angle)
	cr := math.Cos(angle)
	
	angle = a.Y * 0.5;
	sp := math.Sin(angle)
	cp := math.Cos(angle)
	
	angle = a.Z * 0.5;
	sy := math.Sin(angle)
	cy := math.Cos(angle)
	
	cpcy := cp * cy
	spcy := sp * cy
	cpsy := cp * sy
	spsy := sp * sy
	
	q.X = (sr*cpcy) - (cr*spsy)
	q.Y = (cr*spcy) + (sr*cpsy)
	q.Z = (cr*cpsy) - (sr*spcy)
	q.W = (cr*cpcy) + (sr*spsy)
	q.Normalize()
}
func (q *Quaternion) ToAngle(a *Angle) {
	test := ((q.Y*q.W) - (q.X*q.Z))*2.0
	if equalsf(test,1.0,0.000001) {
		a.Z = math.Atan2(q.X,q.W)* -2.0
		a.X = 0
		a.Y = math.Pi/2.0
	} else if equalsf(test,-1.0,0.000001) {
		a.Z = math.Atan2(q.X,q.W)*2.0
		a.X = 0
		a.Y = math.Pi/ -2.0
	} else {
		if test>1.0 { test = 1.0 } else if test< -1.0 { test = -1.0 }
		sx := q.X*q.X
		sy := q.Y*q.Y
		sz := q.Z*q.Z
		sw := q.W*q.W
		a.Z = math.Atan2(2.0 * ((q.X*q.Y)+(q.Z*q.W)), sx-(sy+sz)+sw)
		a.X = math.Atan2(2.0 * ((q.Y*q.Z)+(q.X*q.W)), -(sx+sy)+sz+sw)
		a.Y = math.Asin(test)
	}
}

func (q Quaternion) Multiply(other Quaternion) Quaternion {
	res := Quaternion{}
	res.W = (other.W * q.W) - (other.X * q.X) - (other.Y * q.Y) - (other.Z * q.Z)
	res.X = (other.W * q.X) + (other.X * q.W) + (other.Y * q.Z) - (other.Z * q.Y)
	res.Y = (other.W * q.Y) + (other.Y * q.W) + (other.Z * q.X) - (other.X * q.Z)
	res.Z = (other.W * q.Z) + (other.Z * q.W) + (other.X * q.Y) - (other.Y * q.X)
	return res
}

// ---------------------------------------------------------------------------------
// From here, we have optimized conversions from Quaternion to Normal-Vector and
// vice versa, bypassing the step over the Angle value.


/*
Produces the same result as Quaternion.ToAngle() followed by Angle.Normal64(), but faster.
*/
func (q *Quaternion) Normal64() vgimage.Normal64 {
	n := vgimage.Normal64{0,0,1}
	test := ((q.Y*q.W) - (q.X*q.Z))*2.0
	
	//Applying Z will cause nothing, so we ignore it completely.
	if equalsf(test,1.0,0.000001) {
		//Apply Y: Y = math.Pi/2.0
		n.X = 1
		n.Z = 0
		
		//Apply X: X = 0
		n.Y = 0
	} else if equalsf(test,-1.0,0.000001) {
		//Apply Y: Y = math.Pi/ -2.0
		n.X = -1
		n.Z = 0
		
		//Apply X: X = 0
		n.Y = 0
	} else {
		if test>1.0 { test = 1.0 } else if test< -1.0 { test = -1.0 }
		sx := q.X*q.X
		sy := q.Y*q.Y
		sz := q.Z*q.Z
		sw := q.W*q.W
		at := math.Atan2(2.0 * ((q.Y*q.Z)+(q.X*q.W)), -(sx+sy)+sz+sw)
		
		//Apply Y: Y = math.Asin(test)
		n.X = test
		n.Z = math.Cos(math.Asin(test))
		
		//Apply X: X = math.Atan2(2.0 * ((q.Y*q.Z)+(q.X*q.W)), -(sx+sy)+sz+sw)
		n.Y = math.Sin(at)*n.Z
		n.Z = math.Cos(at)*n.Z
	}
	
	return n
}

/*
Produces the same result as Angle.FromNormal64() followed by Quaternion.FromAngle(), but faster.
*/
func (q *Quaternion) FromNormal64(n vgimage.Normal64) {
	z := math.Sqrt( (n.Y*n.Y) + (n.Z*n.Z) )
	
	angle := math.Atan2(n.Y,n.Z) * 0.5;
	sr := math.Sin(angle)
	cr := math.Cos(angle)
	
	angle = math.Atan2(n.X,z) * 0.5;
	sp := math.Sin(angle)
	cp := math.Cos(angle)
	
	q.X = (sr*cp)
	q.Y = (cr*sp)
	q.Z = -(sr*sp)
	q.W = (cr*cp)
	q.Normalize()
}

