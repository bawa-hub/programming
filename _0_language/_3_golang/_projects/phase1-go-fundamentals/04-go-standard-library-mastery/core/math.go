package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("🚀 Go math Package Mastery Examples")
	fmt.Println("====================================")

	// 1. Mathematical Constants
	fmt.Println("\n1. Mathematical Constants:")
	fmt.Printf("Pi (π): %.15f\n", math.Pi)
	fmt.Printf("E (e): %.15f\n", math.E)
	fmt.Printf("Phi (φ): %.15f\n", math.Phi)
	fmt.Printf("Sqrt2 (√2): %.15f\n", math.Sqrt2)
	fmt.Printf("SqrtE (√e): %.15f\n", math.SqrtE)
	fmt.Printf("SqrtPi (√π): %.15f\n", math.SqrtPi)
	fmt.Printf("Ln2 (ln(2)): %.15f\n", math.Ln2)
	fmt.Printf("Ln10 (ln(10)): %.15f\n", math.Ln10)

	// 2. Basic Functions
	fmt.Println("\n2. Basic Functions:")
	
	// Absolute value
	fmt.Printf("Abs(-5.5): %.2f\n", math.Abs(-5.5))
	fmt.Printf("Abs(3.14): %.2f\n", math.Abs(3.14))
	
	// Maximum and minimum
	fmt.Printf("Max(10, 20): %.2f\n", math.Max(10, 20))
	fmt.Printf("Min(10, 20): %.2f\n", math.Min(10, 20))
	fmt.Printf("Max(-5, -10): %.2f\n", math.Max(-5, -10))
	
	// Sign bit
	fmt.Printf("Signbit(-3.14): %t\n", math.Signbit(-3.14))
	fmt.Printf("Signbit(3.14): %t\n", math.Signbit(3.14))
	
	// NaN and Infinity checks
	fmt.Printf("IsNaN(math.NaN()): %t\n", math.IsNaN(math.NaN()))
	fmt.Printf("IsNaN(3.14): %t\n", math.IsNaN(3.14))
	fmt.Printf("IsInf(math.Inf(1), 1): %t\n", math.IsInf(math.Inf(1), 1))
	fmt.Printf("IsInf(math.Inf(-1), -1): %t\n", math.IsInf(math.Inf(-1), -1))
	fmt.Printf("IsInf(3.14, 0): %t\n", math.IsInf(3.14, 0))

	// 3. Trigonometric Functions
	fmt.Println("\n3. Trigonometric Functions:")
	
	// Basic trig functions
	angle := math.Pi / 4 // 45 degrees
	fmt.Printf("Angle: %.4f radians (45 degrees)\n", angle)
	fmt.Printf("Sin(π/4): %.6f\n", math.Sin(angle))
	fmt.Printf("Cos(π/4): %.6f\n", math.Cos(angle))
	fmt.Printf("Tan(π/4): %.6f\n", math.Tan(angle))
	
	// Inverse trig functions
	fmt.Printf("Asin(0.707): %.6f radians\n", math.Asin(0.707))
	fmt.Printf("Acos(0.707): %.6f radians\n", math.Acos(0.707))
	fmt.Printf("Atan(1): %.6f radians\n", math.Atan(1))
	fmt.Printf("Atan2(1, 1): %.6f radians\n", math.Atan2(1, 1))
	
	// Hyperbolic functions
	fmt.Printf("Sinh(1): %.6f\n", math.Sinh(1))
	fmt.Printf("Cosh(1): %.6f\n", math.Cosh(1))
	fmt.Printf("Tanh(1): %.6f\n", math.Tanh(1))
	
	// Degree to radian conversion
	degrees := 90.0
	radians := degrees * math.Pi / 180
	fmt.Printf("%.0f degrees = %.6f radians\n", degrees, radians)
	fmt.Printf("Sin(90°): %.6f\n", math.Sin(radians))

	// 4. Exponential and Logarithmic Functions
	fmt.Println("\n4. Exponential and Logarithmic Functions:")
	
	// Exponential functions
	fmt.Printf("Exp(1): %.6f (e^1)\n", math.Exp(1))
	fmt.Printf("Exp(2): %.6f (e^2)\n", math.Exp(2))
	fmt.Printf("Exp2(3): %.6f (2^3)\n", math.Exp2(3))
	fmt.Printf("Expm1(1): %.6f (e^1 - 1)\n", math.Expm1(1))
	
	// Logarithmic functions
	fmt.Printf("Log(math.E): %.6f (ln(e))\n", math.Log(math.E))
	fmt.Printf("Log10(100): %.6f (log10(100))\n", math.Log10(100))
	fmt.Printf("Log2(8): %.6f (log2(8))\n", math.Log2(8))
	fmt.Printf("Log1p(1): %.6f (ln(1 + 1))\n", math.Log1p(1))
	
	// Power functions
	fmt.Printf("Pow(2, 3): %.6f (2^3)\n", math.Pow(2, 3))
	fmt.Printf("Pow(10, -2): %.6f (10^-2)\n", math.Pow(10, -2))
	fmt.Printf("Sqrt(16): %.6f (√16)\n", math.Sqrt(16))
	fmt.Printf("Cbrt(27): %.6f (∛27)\n", math.Cbrt(27))

	// 5. Rounding Functions
	fmt.Println("\n5. Rounding Functions:")
	
	// Test values
	values := []float64{3.2, 3.5, 3.7, -3.2, -3.5, -3.7}
	
	for _, val := range values {
		fmt.Printf("Value: %6.1f | Ceil: %6.1f | Floor: %6.1f | Round: %6.1f | Trunc: %6.1f\n",
			val, math.Ceil(val), math.Floor(val), math.Round(val), math.Trunc(val))
	}
	
	// Round to even
	fmt.Printf("RoundToEven(2.5): %.1f\n", math.RoundToEven(2.5))
	fmt.Printf("RoundToEven(3.5): %.1f\n", math.RoundToEven(3.5))

	// 6. Special Functions
	fmt.Println("\n6. Special Functions:")
	
	// Gamma function
	fmt.Printf("Gamma(5): %.6f (4! = 24)\n", math.Gamma(5))
	fmt.Printf("Gamma(0.5): %.6f (√π)\n", math.Gamma(0.5))
	
	// Error functions
	fmt.Printf("Erf(1): %.6f\n", math.Erf(1))
	fmt.Printf("Erfc(1): %.6f\n", math.Erfc(1))
	
	// Bessel functions
	fmt.Printf("J0(1): %.6f\n", math.J0(1))
	fmt.Printf("J1(1): %.6f\n", math.J1(1))
	fmt.Printf("Y0(1): %.6f\n", math.Y0(1))
	fmt.Printf("Y1(1): %.6f\n", math.Y1(1))

	// 7. Integer Functions
	fmt.Println("\n7. Integer Functions:")
	
	// Modulo operation
	fmt.Printf("Mod(10, 3): %.6f\n", math.Mod(10, 3))
	fmt.Printf("Mod(-10, 3): %.6f\n", math.Mod(-10, 3))
	
	// Split into integer and fractional parts
	integer, fractional := math.Modf(3.14159)
	fmt.Printf("Modf(3.14159): integer=%.0f, fractional=%.6f\n", integer, fractional)
	
	// Split into mantissa and exponent
	mantissa, exponent := math.Frexp(8.0)
	fmt.Printf("Frexp(8.0): mantissa=%.6f, exponent=%d\n", mantissa, exponent)
	
	// Multiply by 2^exp
	result := math.Ldexp(0.5, 3) // 0.5 * 2^3 = 4
	fmt.Printf("Ldexp(0.5, 3): %.6f\n", result)
	
	// Binary exponent
	fmt.Printf("Ilogb(8.0): %d\n", math.Ilogb(8.0))
	fmt.Printf("Logb(8.0): %.6f\n", math.Logb(8.0))

	// 8. Mathematical Series and Approximations
	fmt.Println("\n8. Mathematical Series and Approximations:")
	
	// Taylor series for e^x
	x := 1.0
	terms := 10
	exact := math.Exp(x)
	approximation := taylorExp(x, terms)
	fmt.Printf("e^%.1f (exact): %.10f\n", x, exact)
	fmt.Printf("e^%.1f (Taylor, %d terms): %.10f\n", x, terms, approximation)
	fmt.Printf("Error: %.2e\n", math.Abs(exact-approximation))
	
	// Taylor series for sin(x)
	x = math.Pi / 6 // 30 degrees
	exact = math.Sin(x)
	approximation = taylorSin(x, terms)
	fmt.Printf("sin(π/6) (exact): %.10f\n", exact)
	fmt.Printf("sin(π/6) (Taylor, %d terms): %.10f\n", x, terms, approximation)
	fmt.Printf("Error: %.2e\n", math.Abs(exact-approximation))

	// 9. Statistical Functions
	fmt.Println("\n9. Statistical Functions:")
	
	// Generate random data
	rand.Seed(time.Now().UnixNano())
	data := make([]float64, 100)
	for i := range data {
		data[i] = rand.Float64() * 100
	}
	
	// Calculate statistics
	mean := calculateMean(data)
	variance := calculateVariance(data, mean)
	stdDev := math.Sqrt(variance)
	
	fmt.Printf("Data points: %d\n", len(data))
	fmt.Printf("Mean: %.6f\n", mean)
	fmt.Printf("Variance: %.6f\n", variance)
	fmt.Printf("Standard deviation: %.6f\n", stdDev)

	// 10. Complex Number Operations
	fmt.Println("\n10. Complex Number Operations:")
	
	// Create complex numbers
	z1 := complex(3, 4) // 3 + 4i
	z2 := complex(1, 2) // 1 + 2i
	
	// Basic operations
	fmt.Printf("z1 = %.2f\n", z1)
	fmt.Printf("z2 = %.2f\n", z2)
	fmt.Printf("z1 + z2 = %.2f\n", z1+z2)
	fmt.Printf("z1 - z2 = %.2f\n", z1-z2)
	fmt.Printf("z1 * z2 = %.2f\n", z1*z2)
	fmt.Printf("z1 / z2 = %.2f\n", z1/z2)
	
	// Magnitude and phase
	magnitude := math.Sqrt(real(z1)*real(z1) + imag(z1)*imag(z1))
	phase := math.Atan2(imag(z1), real(z1))
	fmt.Printf("|z1| = %.6f\n", magnitude)
	fmt.Printf("arg(z1) = %.6f radians\n", phase)

	// 11. Numerical Integration
	fmt.Println("\n11. Numerical Integration:")
	
	// Integrate x^2 from 0 to 1 (exact result: 1/3)
	integral := simpsonsRule(func(x float64) float64 { return x * x }, 0, 1, 1000)
	exact := 1.0 / 3.0
	fmt.Printf("∫₀¹ x² dx (exact): %.10f\n", exact)
	fmt.Printf("∫₀¹ x² dx (Simpson's rule): %.10f\n", integral)
	fmt.Printf("Error: %.2e\n", math.Abs(exact-integral))

	// 12. Root Finding
	fmt.Println("\n12. Root Finding:")
	
	// Find root of x^2 - 2 = 0 (exact: √2)
	root := bisectionMethod(func(x float64) float64 { return x*x - 2 }, 1, 2, 1e-10)
	exactRoot := math.Sqrt2
	fmt.Printf("√2 (exact): %.10f\n", exactRoot)
	fmt.Printf("√2 (bisection): %.10f\n", root)
	fmt.Printf("Error: %.2e\n", math.Abs(exactRoot-root))

	// 13. Matrix Operations
	fmt.Println("\n13. Matrix Operations:")
	
	// 2x2 matrix multiplication
	a := [][]float64{{1, 2}, {3, 4}}
	b := [][]float64{{5, 6}, {7, 8}}
	c := multiplyMatrices(a, b)
	
	fmt.Println("Matrix A:")
	printMatrix(a)
	fmt.Println("Matrix B:")
	printMatrix(b)
	fmt.Println("A × B:")
	printMatrix(c)

	// 14. Distance and Geometry
	fmt.Println("\n14. Distance and Geometry:")
	
	// Euclidean distance
	p1 := Point{0, 0}
	p2 := Point{3, 4}
	distance := euclideanDistance(p1, p2)
	fmt.Printf("Distance between (0,0) and (3,4): %.6f\n", distance)
	
	// Area of triangle
	p3 := Point{1, 0}
	area := triangleArea(p1, p2, p3)
	fmt.Printf("Area of triangle: %.6f\n", area)

	// 15. Random Number Generation
	fmt.Println("\n15. Random Number Generation:")
	
	// Generate random numbers with different distributions
	fmt.Println("Uniform distribution [0, 1):")
	for i := 0; i < 5; i++ {
		fmt.Printf("%.6f ", rand.Float64())
	}
	fmt.Println()
	
	// Normal distribution (Box-Muller transform)
	fmt.Println("Normal distribution (μ=0, σ=1):")
	for i := 0; i < 5; i++ {
		fmt.Printf("%.6f ", boxMullerTransform())
	}
	fmt.Println()

	// 16. Performance Benchmarking
	fmt.Println("\n16. Performance Benchmarking:")
	
	// Benchmark different approaches
	iterations := 1000000
	
	// Benchmark math.Sqrt vs custom implementation
	start := time.Now()
	for i := 0; i < iterations; i++ {
		math.Sqrt(float64(i))
	}
	sqrtTime := time.Since(start)
	
	start = time.Now()
	for i := 0; i < iterations; i++ {
		customSqrt(float64(i))
	}
	customTime := time.Since(start)
	
	fmt.Printf("math.Sqrt: %v\n", sqrtTime)
	fmt.Printf("customSqrt: %v\n", customTime)
	fmt.Printf("Speedup: %.2fx\n", float64(customTime)/float64(sqrtTime))

	fmt.Println("\n🎉 math Package Mastery Complete!")
}

// Helper functions for mathematical operations

// Taylor series for e^x
func taylorExp(x float64, terms int) float64 {
	result := 1.0
	term := 1.0
	for i := 1; i < terms; i++ {
		term *= x / float64(i)
		result += term
	}
	return result
}

// Taylor series for sin(x)
func taylorSin(x float64, terms int) float64 {
	result := x
	term := x
	for i := 1; i < terms; i++ {
		term *= -x * x / float64((2*i)*(2*i+1))
		result += term
	}
	return result
}

// Calculate mean
func calculateMean(data []float64) float64 {
	sum := 0.0
	for _, x := range data {
		sum += x
	}
	return sum / float64(len(data))
}

// Calculate variance
func calculateVariance(data []float64, mean float64) float64 {
	sum := 0.0
	for _, x := range data {
		diff := x - mean
		sum += diff * diff
	}
	return sum / float64(len(data))
}

// Simpson's rule for numerical integration
func simpsonsRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := f(a) + f(b)
	
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}
	
	return sum * h / 3
}

// Bisection method for root finding
func bisectionMethod(f func(float64) float64, a, b, tolerance float64) float64 {
	for b-a > tolerance {
		c := (a + b) / 2
		if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return (a + b) / 2
}

// Point structure for geometry
type Point struct {
	X, Y float64
}

// Euclidean distance between two points
func euclideanDistance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Area of triangle using cross product
func triangleArea(p1, p2, p3 Point) float64 {
	return math.Abs((p2.X-p1.X)*(p3.Y-p1.Y)-(p3.X-p1.X)*(p2.Y-p1.Y)) / 2
}

// Matrix multiplication
func multiplyMatrices(a, b [][]float64) [][]float64 {
	rows := len(a)
	cols := len(b[0])
	result := make([][]float64, rows)
	
	for i := range result {
		result[i] = make([]float64, cols)
		for j := range result[i] {
			for k := range a[i] {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	
	return result
}

// Print matrix
func printMatrix(m [][]float64) {
	for _, row := range m {
		for _, val := range row {
			fmt.Printf("%8.2f ", val)
		}
		fmt.Println()
	}
}

// Box-Muller transform for normal distribution
func boxMullerTransform() float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	return math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
}

// Custom square root implementation (Newton's method)
func customSqrt(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}
	if x == 0 {
		return 0
	}
	
	guess := x / 2
	for i := 0; i < 10; i++ {
		guess = (guess + x/guess) / 2
	}
	return guess
}
