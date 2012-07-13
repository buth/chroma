package white

import (
	"fmt"
	"testing"
)

func TestTemp(t *testing.T) {
	var x, y, z, rmul, gmul, bmul, temp, rOut, gOut, bOut float64
	green := 1.0
	r, g, b := 0.8, 0.8, 0.8

	fmt.Println("test of Flash6500 wb")

	x, y, z = illum2xyz(27)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, ":   ", rOut, gOut, bOut)

	/// Tests of white-balance at different temperatures
	fmt.Println("/// Tests of white-balance at different temperatures")
	fmt.Println("/// R-G-B values should be roughly equal, near 1.0")

	//24000 K:  0.3961, 0.5123, 1.0000
	fmt.Println("test of 12000 K wb")
	r, g, b = 0.3961, 0.5123, 1.0000
	temp = 24000.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	//12000 K:  0.5431, 0.6389, 1.0000
	fmt.Println("test of 12000 K wb")
	r, g, b = 0.5431, 0.6389, 1.0000
	temp = 12000.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	//10000 K:  0.6268, 0.7039, 1.0000
	fmt.Println("test of 10000 K wb")
	r, g, b = 0.6268, 0.7039, 1.0000
	temp = 10000.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	//8000 K:  0.7874, 0.8187, 1.0000
	fmt.Println("test of 8000 K wb")
	r, g, b = 0.7874, 0.8187, 1.0000
	temp = 8000.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	//6500 K:  1.0000, 0.9436, 0.9621
	fmt.Println("test of 6500K wb")
	r, g, b = 1.0000, 0.9436, 0.9621
	temp = 6500.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	//3500 K:  1.0000, 0.5809, 0.2433
	fmt.Println("test of 3500K wb")
	r, g, b = 1.0000, 0.5809, 0.2433
	temp = 3500.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	//3000 K:  1.0000, 0.4870, 0.1411
	fmt.Println("test of 3000K wb")
	r, g, b = 1.0000, 0.4870, 0.1411
	temp = 3000.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	//2500 K:  1.0000, 0.3814, 0.0588
	fmt.Println("test of 2500K wb")
	r, g, b = 1.0000, 0.3814, 0.0588
	temp = 2500.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	//2100 K:  1.0000, 0.2889, 0.0120
	fmt.Println("test of 2100K wb")
	r, g, b = 1.0000, 0.2889, 0.0120
	temp = 2100.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGB:", rOut, gOut, bOut)

	// Test against Rawtherapee
	r, g, b = 0.8, 0.8, 0.8
	r2, g2, b2 := 0.690196, 0.780392, 0.952941
	fmt.Println("Test against Rawtherapee: 4500 K")
	temp = 4500.0
	x, y, z = temp2xyz(temp)
	rmul, gmul, bmul = xyz2mul(x, y, z, green, false)
	rOut, gOut, bOut = mul2RGB(r, g, b, rmul, gmul, bmul)
	fmt.Println(temp, " K:   RGBref:", r2, g2, b2 , "RGBout:", rOut, gOut, bOut)



}
