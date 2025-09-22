package main

import "fmt"

type CompressionStrategy interface {
	Compress(input string) string
}

type ZipCompression struct{}

type GzipCompression struct{}

func (ZipCompression) Compress(input string) string {
	return fmt.Sprintf("[zip]%s[/zip]", input)
}

func (GzipCompression) Compress(input string) string {
	return fmt.Sprintf("[gzip]%s[/gzip]", input)
}

type Compressor struct {
	strategy CompressionStrategy
}

func NewCompressor(strategy CompressionStrategy) *Compressor {
	return &Compressor{strategy: strategy}
}

func (c *Compressor) SetStrategy(strategy CompressionStrategy) {
	c.strategy = strategy
}

func (c *Compressor) Do(input string) string {
	return c.strategy.Compress(input)
}

func main() {
	compressor := NewCompressor(ZipCompression{})
	fmt.Println(compressor.Do("hello"))

	compressor.SetStrategy(GzipCompression{})
	fmt.Println(compressor.Do("world"))
}
