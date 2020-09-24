package main


import (
	"math"
	"fmt"
)


type ColorSpace interface {
	PrintValues()
}


type HSL struct {
	H, S, L float64
}


func ( hsl HSL ) PrintValues() {
	fmt.Println( hsl.H, hsl.S, hsl.L )
}


type RGB struct {
	R, G, B float64
}


func ( rgb RGB ) PrintValues() {
	fmt.Println( rgb.R, rgb.G, rgb.B )
}


func RGB_RED() RGB {
	return RGB{255, 0, 0}
}


func RGBtoHSL( rgb RGB ) HSL {

	r := rgb.R / 256.0
	g := rgb.G / 256.0
	b := rgb.B / 256.0

	maxCol := math.Max( r, math.Max( g, b ) )
	minCol := math.Min( r, math.Min( g, b ) )

	var h, s, l float64

	if ( r == g ) && ( r == b ) {
		h = 0.0
		s = 0.0
		l = b //can be any between r, g & b
	} else {
		l = ( minCol + maxCol ) / 2

		if l < 0.5 {
			s = ( maxCol - minCol ) / ( maxCol + minCol )
		} else {
			s = ( maxCol - minCol ) / ( 2.0 - maxCol - minCol )
		}

		if r == maxCol {
			h = ( g - b ) / ( maxCol - minCol )
		} else if g == maxCol {
			h = 2.0 + ( b - r ) / ( maxCol - minCol )
		} else {
			h = 4.0 + ( r - g ) / ( maxCol - minCol )
		}

		h /= 6

		if h < 0 {
			h += 1
		}
	}

	hsl := HSL{h, s, l}
	return hsl
}


func HSltoRGB( hsl HSL ) RGB {

	var r, g, b float64
	var temp1, temp2, tempr, tempg, tempb float64

	h := hsl.H / 256.0
	s := hsl.S / 256.0
	l := hsl.L / 256.0

	if s == 0 {
		r, g, b = l, l, l
	}

	if l < 0.5 {
		temp2 = l * ( 1 + s)
	} else {
		temp2 = ( l + s ) - ( l * s )
	}

	temp1 = 2 * l - temp2
    tempr = h + 1.0 / 3.0

    if tempr > 1 {
    	tempr -= 1
    }

    tempg = h
    tempb = h - 1.0 / 3.0

    if tempb < 0 {
    	tempb += 1
    }

    /**
    RED
    */
    if tempr < ( 1.0 / 6.0 ) {
    	r = temp1 + ( temp2 - temp1 ) * 6.0 * tempr
    } else if tempr < 0.5 {
    	r = temp2
    } else if tempr < 2.0 / 3.0 {
    	r = temp1 + ( temp2 - temp1 ) * ( (2.0 / 3.0) - tempr ) * 6.0
    } else {
    	r = temp1
    }

    /**
    GREEN
    */
    if tempg < ( 1.0 / 6.0 ) {
    	g = temp1 + ( temp2 - temp1 ) * 6.0 * tempg
    } else if tempg < 0.5 {
    	g = temp2
    } else if tempg < ( 2.0 / 3.0 ) {
    	g = temp1 + ( temp2 - temp1 ) * ( (2.0 / 3.0) - tempg ) * 6.0
    } else {
    	g = temp1
    }

    /**
    BLUE
    */
    if tempb < ( 1.0 / 6.0 ) {
    	b = temp1 + ( temp2 - temp1 ) * 6.0 * tempb
    } else if tempb < 0.5 {
    	b = temp2
    } else if tempb < ( 2.0 / 3.0 ) {
    	b = temp1 + ( temp2 - temp1 ) * ( (2.0 / 3.0) - tempb ) * 6.0
    } else {
    	b = temp1
    }

	rgb := RGB{r, g, b}
	return rgb
}