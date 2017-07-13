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
An Image Copy algorithm. Supports both, single threaded and multithreaded operations.
*/
package copier


import (
	"image"
	//"image/color"
	"image/draw"
	"errors"
)

var ESizeMismatch = errors.New("Size-Mismatch")

func iterate(r image.Rectangle, pchan chan <- image.Point) {
	xb,yb,xe,ye := r.Min.X,r.Min.Y,r.Max.X,r.Max.Y
	for x := xb; x<xe; x++ {
		for y := yb; y<ye; y++ {
			pchan <- image.Point{x,y}
		}
	}
	close(pchan)
}
func copyf(dst draw.Image, src image.Image, pchan <- chan image.Point, qchan chan <- int) {
	for pt := range pchan {
		c := src.At(pt.X,pt.Y)
		dst.Set(pt.X,pt.Y,c)
	}
	qchan <- 1
}

/*
Copies an Image using one thread.
*/
func Copy(dst draw.Image, src image.Image, blind bool) error {
	dR,sR := dst.Bounds(),src.Bounds()
	if !(blind||dR.Eq(sR)) { return ESizeMismatch }
	pchan := make(chan image.Point,16)
	qchan := make(chan int,1)
	go iterate(sR,pchan)
	copyf(dst,src,pchan,qchan)
	return nil
}

/*
Copies an Image using multiple threads. The images should be thread-safe.
*/
func CopyMT(dst draw.Image, src image.Image, blind bool, nthreads int) error {
	dR,sR := dst.Bounds(),src.Bounds()
	if !(blind||dR.Eq(sR)) { return ESizeMismatch }
	pchan := make(chan image.Point,nthreads)
	qchan := make(chan int,nthreads)
	go iterate(sR,pchan)
	for i:=0; i<nthreads; i++ {
		go copyf(dst,src,pchan,qchan)
	}
	for i:=0; i<nthreads; i++ {
		<- qchan
	}
	return nil
}

type Operator interface{
	Operate(pt image.Point)
}
func operateF(o Operator, pchan <- chan image.Point, qchan chan <- int) {
	for pt := range pchan {
		o.Operate(pt)
	}
	qchan <- 1
}

/*
Performs the Task 'o' on every Pixel within the Rectangle 'r'.
*/
func Operate(o Operator, r image.Rectangle) {
	pchan := make(chan image.Point,16)
	qchan := make(chan int,1)
	go iterate(r,pchan)
	operateF(o,pchan,qchan)
}

/*
Performs the Task 'o' on every Pixel within the Rectangle 'r'.
This implementation uses multiple threads, so the Task should be threadsafe.
*/
func OperateMT(o Operator, r image.Rectangle, nthreads int) {
	pchan := make(chan image.Point,nthreads)
	qchan := make(chan int,nthreads)
	go iterate(r,pchan)
	for i:=0; i<nthreads; i++ {
		go operateF(o,pchan,qchan)
	}
	for i:=0; i<nthreads; i++ {
		<- qchan
	}
}

