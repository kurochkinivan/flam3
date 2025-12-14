package variation

import "math"

func linear(x, y float64) (newX, newY float64) {
	newX, newY = x, y
	return newX, newY
}

func sinusoidal(x, y float64) (newX, newY float64) {
	newX = math.Sin(x)
	newY = math.Sin(y)

	return newX, newY
}

func spherical(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	r2 := r * r

	newX = x / r2
	newY = y / r2

	return newX, newY
}

func swirl(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	r2 := r * r
	sinR2 := math.Sin(r2)
	cosR2 := math.Cos(r2)

	newX = x*sinR2 - y*cosR2
	newY = x*cosR2 + y*sinR2

	return newX, newY
}

func horseshoe(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)

	newX = (x - y) * (x + y) / r
	newY = 2 * x * y / r

	return newX, newY
}

func polar(x, y float64) (newX, newY float64) {
	theta := math.Atan2(y, x)
	r := math.Sqrt(x*x + y*y)
	newX = theta / math.Pi
	newY = r - 1.0
	return newX, newY
}

func handkerchief(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(y, x)
	newX = r * math.Sin(theta+r)
	newY = r * math.Cos(theta-r)
	return newX, newY
}

func heart(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(y, x)
	newX = r * math.Sin(theta*r)
	newY = -r * math.Cos(theta*r)
	return newX, newY
}

func disk(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x+y*y) * math.Pi
	theta := math.Atan2(y, x) / math.Pi
	newX = theta * math.Sin(r)
	newY = theta * math.Cos(r)
	return newX, newY
}

func spiral(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(y, x)
	newX = (1.0 / r) * (math.Cos(theta) + math.Sin(r))
	newY = (1.0 / r) * (math.Sin(theta) - math.Cos(r))
	return newX, newY
}

func hyperbolic(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(y, x)
	newX = math.Sin(theta) / r
	newY = r * math.Cos(theta)
	return newX, newY
}

func diamond(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(y, x)
	newX = math.Sin(theta) * math.Cos(r)
	newY = math.Cos(theta) * math.Sin(r)
	return newX, newY
}

func ex(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(y, x)

	p0 := math.Pow(math.Sin(theta+r), 3)
	p1 := math.Pow(math.Cos(theta-r), 3)

	newX = r * (p0 + p1)
	newY = r * (p0 - p1)
	return newX, newY
}

func bent(x, y float64) (newX, newY float64) {
	if x >= 0 && y >= 0 {
		newX, newY = x, y
	} else if x < 0 && y >= 0 {
		newX, newY = 2*x, y
	} else if x >= 0 && y < 0 {
		newX, newY = x, 0.5*y
	} else {
		newX, newY = 2*x, 0.5*y
	}
	return newX, newY
}

func fisheye(x, y float64) (newX, newY float64) {
	r := 2.0 / (1.0 + math.Sqrt(x*x+y*y))
	newX = r * y
	newY = r * x
	return newX, newY
}

func eyefish(x, y float64) (newX, newY float64) {
	r := 2.0 / (1.0 + math.Sqrt(x*x+y*y))
	newX = r * x
	newY = r * y
	return newX, newY
}

func bubble(x, y float64) (newX, newY float64) {
	r := 4 + x*x + y*y
	newX = (4 * x) / r
	newY = (4 * y) / r
	return newX, newY
}

func cylinder(x, y float64) (newX, newY float64) {
	newX = math.Sin(x)
	newY = y
	return newX, newY
}

func tangent(x, y float64) (newX, newY float64) {
	newX = math.Sin(x) / math.Cos(y)
	newY = math.Tan(y)
	return newX, newY
}

func cross(x, y float64) (newX, newY float64) {
	r := math.Sqrt(1.0 / ((x*x - y*y) * (x*x - y*y)))
	newX = x * r
	newY = y * r
	return newX, newY
}

func exponential(x, y float64) (newX, newY float64) {
	newX = math.Exp(x-1.0) * math.Cos(math.Pi*y)
	newY = math.Exp(x-1.0) * math.Sin(math.Pi*y)
	return newX, newY
}

func power(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(y, x)
	p := math.Pow(r, math.Sin(theta))
	newX = p * math.Cos(theta)
	newY = p * math.Sin(theta)
	return newX, newY
}

func cosine(x, y float64) (newX, newY float64) {
	newX = math.Cos(math.Pi*x) * math.Cosh(y)
	newY = -math.Sin(math.Pi*x) * math.Sinh(y)
	return newX, newY
}
