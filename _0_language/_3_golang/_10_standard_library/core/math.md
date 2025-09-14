# math Package - Mathematical Functions 🔢

The `math` package provides mathematical constants and functions. It's essential for numerical computations, scientific calculations, and mathematical operations.

## 🎯 Key Concepts

### 1. **Mathematical Constants**
- `Pi` - π (3.141592653589793)
- `E` - e (2.718281828459045)
- `Phi` - φ (1.618033988749895)
- `Sqrt2` - √2 (1.4142135623730951)
- `SqrtE` - √e (1.6487212707001282)
- `SqrtPi` - √π (1.7724538509055159)
- `Ln2` - ln(2) (0.6931471805599453)
- `Ln10` - ln(10) (2.302585092994046)

### 2. **Basic Functions**
- `Abs()` - Absolute value
- `Max()` - Maximum of two values
- `Min()` - Minimum of two values
- `Signbit()` - Check if sign bit is set
- `IsNaN()` - Check if value is NaN
- `IsInf()` - Check if value is infinity

### 3. **Trigonometric Functions**
- `Sin()` - Sine
- `Cos()` - Cosine
- `Tan()` - Tangent
- `Asin()` - Arc sine
- `Acos()` - Arc cosine
- `Atan()` - Arc tangent
- `Atan2()` - Arc tangent of y/x
- `Sinh()` - Hyperbolic sine
- `Cosh()` - Hyperbolic cosine
- `Tanh()` - Hyperbolic tangent

### 4. **Exponential and Logarithmic Functions**
- `Exp()` - e^x
- `Exp2()` - 2^x
- `Expm1()` - e^x - 1
- `Log()` - Natural logarithm
- `Log10()` - Base-10 logarithm
- `Log2()` - Base-2 logarithm
- `Log1p()` - ln(1 + x)
- `Pow()` - x^y
- `Sqrt()` - Square root
- `Cbrt()` - Cube root

### 5. **Rounding Functions**
- `Ceil()` - Round up
- `Floor()` - Round down
- `Round()` - Round to nearest
- `RoundToEven()` - Round to even
- `Trunc()` - Truncate

### 6. **Special Functions**
- `Gamma()` - Gamma function
- `Lgamma()` - Log gamma function
- `Erf()` - Error function
- `Erfc()` - Complementary error function
- `J0()` - Bessel function of first kind
- `J1()` - Bessel function of first kind
- `Y0()` - Bessel function of second kind
- `Y1()` - Bessel function of second kind

### 7. **Integer Functions**
- `Mod()` - Modulo operation
- `Modf()` - Split into integer and fractional parts
- `Frexp()` - Split into mantissa and exponent
- `Ldexp()` - Multiply by 2^exp
- `Ilogb()` - Binary exponent
- `Logb()` - Binary exponent

## 🚀 Common Patterns

### Basic Math
```go
result := math.Sqrt(16) // 4
max := math.Max(10, 20) // 20
abs := math.Abs(-5)     // 5
```

### Trigonometric Calculations
```go
angle := math.Pi / 4
sin := math.Sin(angle)
cos := math.Cos(angle)
```

### Exponential Functions
```go
exp := math.Exp(1)      // e^1
log := math.Log(math.E) // ln(e) = 1
pow := math.Pow(2, 3)   // 2^3 = 8
```

## ⚠️ Common Pitfalls

1. **NaN values** - Always check for NaN results
2. **Infinity values** - Handle infinity cases
3. **Precision loss** - Be aware of floating-point precision
4. **Domain errors** - Check input ranges for functions
5. **Overflow/underflow** - Handle extreme values

## 🎯 Best Practices

1. **Check for special values** - Always check for NaN and Inf
2. **Use appropriate precision** - Choose float32 vs float64
3. **Handle edge cases** - Check input ranges
4. **Use constants** - Use math constants instead of hardcoded values
5. **Consider performance** - Some functions are expensive

## 🔍 Advanced Features

### Custom Math Functions
```go
func CustomSqrt(x float64) float64 {
    if x < 0 {
        return math.NaN()
    }
    return math.Sqrt(x)
}
```

### Mathematical Series
```go
func TaylorSeries(x float64, terms int) float64 {
    result := 1.0
    term := 1.0
    for i := 1; i < terms; i++ {
        term *= x / float64(i)
        result += term
    }
    return result
}
```

## 📚 Real-world Applications

1. **Scientific Computing** - Physics, chemistry, biology
2. **Graphics Programming** - 3D graphics, animations
3. **Signal Processing** - Audio, image processing
4. **Statistics** - Data analysis, machine learning
5. **Cryptography** - Random number generation

## 🧠 Memory Tips

- **math** = **M**athematical **A**lgorithms **T**oolkit **H**elper
- **Abs** = **A**bsolute value
- **Max/Min** = **M**aximum/**M**inimum
- **Sin/Cos/Tan** = **S**ine/**C**osine/**T**angent
- **Log** = **L**ogarithm
- **Exp** = **E**xponential
- **Sqrt** = **S**quare **R**oot
- **Pow** = **P**ower

Remember: The math package is your gateway to mathematical operations in Go! 🎯
