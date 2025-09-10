package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Custom Reader implementation
type CustomReader struct {
	data []byte
	pos  int
}

func NewCustomReader(data string) *CustomReader {
	return &CustomReader{
		data: []byte(data),
		pos:  0,
	}
}

func (r *CustomReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

// Custom Writer implementation
type CustomWriter struct {
	buffer []byte
}

func NewCustomWriter() *CustomWriter {
	return &CustomWriter{
		buffer: make([]byte, 0),
	}
}

func (w *CustomWriter) Write(p []byte) (n int, err error) {
	w.buffer = append(w.buffer, p...)
	return len(p), nil
}

func (w *CustomWriter) String() string {
	return string(w.buffer)
}

// Custom ReadWriter implementation
type CustomReadWriter struct {
	data   []byte
	pos    int
	closed bool
}

func NewCustomReadWriter() *CustomReadWriter {
	return &CustomReadWriter{
		data:   make([]byte, 0),
		pos:    0,
		closed: false,
	}
}

func (rw *CustomReadWriter) Read(p []byte) (n int, err error) {
	if rw.closed {
		return 0, io.ErrClosedPipe
	}
	if rw.pos >= len(rw.data) {
		return 0, io.EOF
	}
	n = copy(p, rw.data[rw.pos:])
	rw.pos += n
	return n, nil
}

func (rw *CustomReadWriter) Write(p []byte) (n int, err error) {
	if rw.closed {
		return 0, io.ErrClosedPipe
	}
	rw.data = append(rw.data, p...)
	return len(p), nil
}

func (rw *CustomReadWriter) Close() error {
	rw.closed = true
	return nil
}

func (rw *CustomReadWriter) Seek(offset int64, whence int) (int64, error) {
	if rw.closed {
		return 0, io.ErrClosedPipe
	}
	
	var newPos int64
	switch whence {
	case io.SeekStart:
		newPos = offset
	case io.SeekCurrent:
		newPos = int64(rw.pos) + offset
	case io.SeekEnd:
		newPos = int64(len(rw.data)) + offset
	default:
		return 0, fmt.Errorf("invalid whence: %d", whence)
	}
	
	if newPos < 0 || newPos > int64(len(rw.data)) {
		return 0, fmt.Errorf("invalid offset: %d", newPos)
	}
	
	rw.pos = int(newPos)
	return newPos, nil
}

func main() {
	fmt.Println("🚀 Go io Package Mastery Examples")
	fmt.Println("==================================")

	// 1. Basic Reader Interface
	fmt.Println("\n1. Basic Reader Interface:")
	reader := strings.NewReader("Hello, World!")
	buffer := make([]byte, 5)
	
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(buffer))
	}

	// 2. Basic Writer Interface
	fmt.Println("\n2. Basic Writer Interface:")
	var writer bytes.Buffer
	_, err = writer.Write([]byte("Hello, "))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Writer content: %s\n", writer.String())
	}

	// 3. ReadAll Function
	fmt.Println("\n3. ReadAll Function:")
	reader = strings.NewReader("This is a test string for ReadAll")
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("ReadAll result: %s\n", string(data))
	}

	// 4. WriteString Function
	fmt.Println("\n4. WriteString Function:")
	writer.Reset()
	_, err = io.WriteString(&writer, "Hello from WriteString!")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("WriteString result: %s\n", writer.String())
	}

	// 5. Copy Function
	fmt.Println("\n5. Copy Function:")
	src := strings.NewReader("Source data for copying")
	dst := &bytes.Buffer{}
	
	bytesCopied, err := io.Copy(dst, src)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Copied %d bytes: %s\n", bytesCopied, dst.String())
	}

	// 6. CopyN Function
	fmt.Println("\n6. CopyN Function:")
	src = strings.NewReader("This is a long string for CopyN")
	dst.Reset()
	
	bytesCopied, err = io.CopyN(dst, src, 10)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Copied %d bytes: %s\n", bytesCopied, dst.String())
	}

	// 7. ReadAtLeast Function
	fmt.Println("\n7. ReadAtLeast Function:")
	reader = strings.NewReader("Short")
	buffer = make([]byte, 10)
	
	n, err = io.ReadAtLeast(reader, buffer, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("ReadAtLeast %d bytes: %s\n", n, string(buffer[:n]))
	}

	// 8. ReadFull Function
	fmt.Println("\n8. ReadFull Function:")
	reader = strings.NewReader("Full")
	buffer = make([]byte, 4)
	
	n, err = io.ReadFull(reader, buffer)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("ReadFull %d bytes: %s\n", n, string(buffer))
	}

	// 9. LimitReader
	fmt.Println("\n9. LimitReader:")
	reader = strings.NewReader("This is a very long string that will be limited")
	limitedReader := io.LimitReader(reader, 20)
	
	data, err = io.ReadAll(limitedReader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Limited read: %s\n", string(data))
	}

	// 10. MultiReader
	fmt.Println("\n10. MultiReader:")
	reader1 := strings.NewReader("First part. ")
	reader2 := strings.NewReader("Second part. ")
	reader3 := strings.NewReader("Third part.")
	
	multiReader := io.MultiReader(reader1, reader2, reader3)
	data, err = io.ReadAll(multiReader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("MultiReader result: %s\n", string(data))
	}

	// 11. MultiWriter
	fmt.Println("\n11. MultiWriter:")
	writer1 := &bytes.Buffer{}
	writer2 := &bytes.Buffer{}
	writer3 := &bytes.Buffer{}
	
	multiWriter := io.MultiWriter(writer1, writer2, writer3)
	_, err = io.WriteString(multiWriter, "Written to all writers")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Writer1: %s\n", writer1.String())
		fmt.Printf("Writer2: %s\n", writer2.String())
		fmt.Printf("Writer3: %s\n", writer3.String())
	}

	// 12. TeeReader
	fmt.Println("\n12. TeeReader:")
	reader = strings.NewReader("Data for TeeReader")
	writer = &bytes.Buffer{}
	teeReader := io.TeeReader(reader, writer)
	
	data, err = io.ReadAll(teeReader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("TeeReader read: %s\n", string(data))
		fmt.Printf("TeeReader wrote: %s\n", writer.String())
	}

	// 13. Pipe
	fmt.Println("\n13. Pipe:")
	reader, writer := io.Pipe()
	
	// Write to pipe in goroutine
	go func() {
		defer writer.Close()
		writer.Write([]byte("Data written to pipe"))
	}()
	
	// Read from pipe
	data, err = io.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Pipe data: %s\n", string(data))
	}

	// 14. Custom Reader
	fmt.Println("\n14. Custom Reader:")
	customReader := NewCustomReader("Custom reader data")
	buffer = make([]byte, 10)
	
	n, err = customReader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Custom reader read %d bytes: %s\n", n, string(buffer[:n]))
	}

	// 15. Custom Writer
	fmt.Println("\n15. Custom Writer:")
	customWriter := NewCustomWriter()
	_, err = customWriter.Write([]byte("Custom writer data"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Custom writer: %s\n", customWriter.String())
	}

	// 16. Custom ReadWriter
	fmt.Println("\n16. Custom ReadWriter:")
	customRW := NewCustomReadWriter()
	
	// Write data
	_, err = customRW.Write([]byte("ReadWriter test data"))
	if err != nil {
		fmt.Printf("Error writing: %v\n", err)
	} else {
		fmt.Printf("Written to ReadWriter\n")
	}
	
	// Read data
	buffer = make([]byte, 20)
	n, err = customRW.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading: %v\n", err)
	} else {
		fmt.Printf("Read from ReadWriter: %s\n", string(buffer[:n]))
	}

	// 17. Seek Operations
	fmt.Println("\n17. Seek Operations:")
	customRW = NewCustomReadWriter()
	customRW.Write([]byte("0123456789"))
	
	// Seek to beginning
	pos, err := customRW.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Printf("Error seeking: %v\n", err)
	} else {
		fmt.Printf("Seeked to position: %d\n", pos)
	}
	
	// Read from beginning
	buffer = make([]byte, 5)
	n, err = customRW.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading: %v\n", err)
	} else {
		fmt.Printf("Read from beginning: %s\n", string(buffer[:n]))
	}
	
	// Seek to middle
	pos, err = customRW.Seek(3, io.SeekStart)
	if err != nil {
		fmt.Printf("Error seeking: %v\n", err)
	} else {
		fmt.Printf("Seeked to position: %d\n", pos)
	}
	
	// Read from middle
	n, err = customRW.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading: %v\n", err)
	} else {
		fmt.Printf("Read from middle: %s\n", string(buffer[:n]))
	}

	// 18. SectionReader
	fmt.Println("\n18. SectionReader:")
	reader = strings.NewReader("This is a long string for SectionReader")
	sectionReader := io.NewSectionReader(reader, 5, 10)
	
	data, err = io.ReadAll(sectionReader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("SectionReader result: %s\n", string(data))
	}

	// 19. Error Handling
	fmt.Println("\n19. Error Handling:")
	reader = strings.NewReader("")
	buffer = make([]byte, 10)
	
	n, err = reader.Read(buffer)
	if err == io.EOF {
		fmt.Println("Reached end of file (EOF)")
	} else if err != nil {
		fmt.Printf("Other error: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes\n", n)
	}

	// 20. File I/O Example
	fmt.Println("\n20. File I/O Example:")
	
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "io_example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
	} else {
		defer os.Remove(tempFile.Name())
		defer tempFile.Close()
		
		// Write to file
		_, err = io.WriteString(tempFile, "Hello from io package!")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		} else {
			fmt.Printf("Written to file: %s\n", tempFile.Name())
		}
		
		// Read from file
		tempFile.Seek(0, io.SeekStart)
		data, err = io.ReadAll(tempFile)
		if err != nil {
			fmt.Printf("Error reading from file: %v\n", err)
		} else {
			fmt.Printf("Read from file: %s\n", string(data))
		}
	}

	// 21. Buffered I/O
	fmt.Println("\n21. Buffered I/O:")
	reader = strings.NewReader("Line 1\nLine 2\nLine 3\n")
	scanner := bufio.NewScanner(reader)
	
	for scanner.Scan() {
		fmt.Printf("Scanned line: %s\n", scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}

	// 22. NopCloser
	fmt.Println("\n22. NopCloser:")
	reader = strings.NewReader("Data for NopCloser")
	closer := io.NopCloser(reader)
	
	data, err = io.ReadAll(closer)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("NopCloser result: %s\n", string(data))
	}
	
	// Close (does nothing)
	closer.Close()
	fmt.Println("NopCloser closed (no-op)")

	fmt.Println("\n🎉 io Package Mastery Complete!")
}
