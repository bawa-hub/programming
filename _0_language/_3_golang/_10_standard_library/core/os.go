package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	fmt.Println("🚀 Go os Package Mastery Examples")
	fmt.Println("==================================")

	// 1. Command Line Arguments
	fmt.Println("\n1. Command Line Arguments:")
	fmt.Printf("Program name: %s\n", os.Args[0])
	fmt.Printf("Number of arguments: %d\n", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("  [%d] %s\n", i, arg)
	}

	// 2. Environment Variables
	fmt.Println("\n2. Environment Variables:")
	
	// Get specific environment variable
	home := os.Getenv("HOME")
	if home != "" {
		fmt.Printf("HOME: %s\n", home)
	} else {
		fmt.Println("HOME not set")
	}

	// Get all environment variables
	fmt.Println("\nAll environment variables:")
	envVars := os.Environ()
	for i, env := range envVars {
		if i >= 5 { // Show only first 5
			fmt.Printf("... and %d more\n", len(envVars)-5)
			break
		}
		fmt.Printf("  %s\n", env)
	}

	// Set and get environment variable
	os.Setenv("MY_VAR", "Hello World")
	value, exists := os.LookupEnv("MY_VAR")
	if exists {
		fmt.Printf("MY_VAR: %s\n", value)
	}

	// 3. Process Information
	fmt.Println("\n3. Process Information:")
	fmt.Printf("Process ID: %d\n", os.Getpid())
	fmt.Printf("Parent Process ID: %d\n", os.Getppid())
	fmt.Printf("User ID: %d\n", os.Getuid())
	fmt.Printf("Group ID: %d\n", os.Getgid())

	// 4. Working Directory
	fmt.Println("\n4. Working Directory:")
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %v", err)
	} else {
		fmt.Printf("Current directory: %s\n", wd)
	}

	// Change working directory
	originalWd, _ := os.Getwd()
	tempDir := "/tmp"
	err = os.Chdir(tempDir)
	if err != nil {
		log.Printf("Error changing directory: %v", err)
	} else {
		fmt.Printf("Changed to: %s\n", tempDir)
		// Change back
		os.Chdir(originalWd)
	}

	// 5. File Operations
	fmt.Println("\n5. File Operations:")
	
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "go_os_example_*.txt")
	if err != nil {
		log.Printf("Error creating temp file: %v", err)
	} else {
		fmt.Printf("Created temp file: %s\n", tempFile.Name())
		
		// Write to file
		content := "Hello from Go os package!\nThis is a test file.\n"
		_, err = tempFile.WriteString(content)
		if err != nil {
			log.Printf("Error writing to file: %v", err)
		}
		
		// Close file
		tempFile.Close()
		
		// Read from file
		readFile, err := os.Open(tempFile.Name())
		if err != nil {
			log.Printf("Error opening file: %v", err)
		} else {
			scanner := bufio.NewScanner(readFile)
			fmt.Println("File content:")
			for scanner.Scan() {
				fmt.Printf("  %s\n", scanner.Text())
			}
			readFile.Close()
		}
		
		// Clean up
		os.Remove(tempFile.Name())
		fmt.Println("Temp file removed")
	}

	// 6. File Information
	fmt.Println("\n6. File Information:")
	
	// Get current file info
	_, err = os.Stat("os.go")
	if err != nil {
		log.Printf("Error getting file info: %v", err)
	} else {
		info, _ := os.Stat("os.go")
		fmt.Printf("File: %s\n", info.Name())
		fmt.Printf("Size: %d bytes\n", info.Size())
		fmt.Printf("Mode: %s\n", info.Mode())
		fmt.Printf("ModTime: %s\n", info.ModTime())
		fmt.Printf("IsDir: %t\n", info.IsDir())
		fmt.Printf("Sys: %v\n", info.Sys())
	}

	// 7. Directory Operations
	fmt.Println("\n7. Directory Operations:")
	
	// Create test directory
	testDir := "test_dir"
	err = os.Mkdir(testDir, 0755)
	if err != nil {
		log.Printf("Error creating directory: %v", err)
	} else {
		fmt.Printf("Created directory: %s\n", testDir)
		
		// Create nested directories
		nestedDir := filepath.Join(testDir, "nested", "deep")
		err = os.MkdirAll(nestedDir, 0755)
		if err != nil {
			log.Printf("Error creating nested directories: %v", err)
		} else {
			fmt.Printf("Created nested directories: %s\n", nestedDir)
		}
		
		// List directory contents
		entries, err := os.ReadDir(testDir)
		if err != nil {
			log.Printf("Error reading directory: %v", err)
		} else {
			fmt.Printf("Directory contents of %s:\n", testDir)
			for _, entry := range entries {
				fmt.Printf("  %s (isDir: %t)\n", entry.Name(), entry.IsDir())
			}
		}
		
		// Clean up
		err = os.RemoveAll(testDir)
		if err != nil {
			log.Printf("Error removing directory: %v", err)
		} else {
			fmt.Println("Directory removed")
		}
	}

	// 8. File Permissions
	fmt.Println("\n8. File Permissions:")
	
	// Create file with specific permissions
	permFile := "perm_test.txt"
	file, err := os.OpenFile(permFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error creating file: %v", err)
	} else {
		file.WriteString("Permission test file\n")
		file.Close()
		
		// Get file info
		info, err := os.Stat(permFile)
		if err != nil {
			log.Printf("Error getting file info: %v", err)
		} else {
			mode := info.Mode()
			fmt.Printf("File permissions: %s\n", mode)
			fmt.Printf("Is regular file: %t\n", mode.IsRegular())
			fmt.Printf("Is directory: %t\n", mode.IsDir())
			fmt.Printf("Is symbolic link: %t\n", mode&os.ModeSymlink != 0)
			fmt.Printf("Permission bits: %o\n", mode.Perm())
		}
		
		// Clean up
		os.Remove(permFile)
	}

	// 9. Standard Streams
	fmt.Println("\n9. Standard Streams:")
	fmt.Printf("Stdin: %v\n", os.Stdin)
	fmt.Printf("Stdout: %v\n", os.Stdout)
	fmt.Printf("Stderr: %v\n", os.Stderr)
	
	// Write to stderr
	fmt.Fprintf(os.Stderr, "This is written to stderr\n")

	// 10. File Copying
	fmt.Println("\n10. File Copying:")
	
	// Create source file
	srcFile := "source.txt"
	src, err := os.Create(srcFile)
	if err != nil {
		log.Printf("Error creating source file: %v", err)
	} else {
		src.WriteString("This is the source file content\n")
		src.Close()
		
		// Copy file
		dstFile := "destination.txt"
		err = copyFile(srcFile, dstFile)
		if err != nil {
			log.Printf("Error copying file: %v", err)
		} else {
			fmt.Printf("File copied from %s to %s\n", srcFile, dstFile)
			
			// Verify copy
			dst, err := os.Open(dstFile)
			if err != nil {
				log.Printf("Error opening destination file: %v", err)
			} else {
				content, err := io.ReadAll(dst)
				if err != nil {
					log.Printf("Error reading destination file: %v", err)
				} else {
					fmt.Printf("Copied content: %s", string(content))
				}
				dst.Close()
			}
		}
		
		// Clean up
		os.Remove(srcFile)
		os.Remove(dstFile)
	}

	// 11. File Walking
	fmt.Println("\n11. File Walking:")
	
	// Create test directory structure
	walkDir := "walk_test"
	os.MkdirAll(walkDir, 0755)
	os.MkdirAll(filepath.Join(walkDir, "subdir1"), 0755)
	os.MkdirAll(filepath.Join(walkDir, "subdir2"), 0755)
	
	// Create files
	os.WriteFile(filepath.Join(walkDir, "file1.txt"), []byte("File 1"), 0644)
	os.WriteFile(filepath.Join(walkDir, "subdir1", "file2.txt"), []byte("File 2"), 0644)
	os.WriteFile(filepath.Join(walkDir, "subdir2", "file3.txt"), []byte("File 3"), 0644)
	
	// Walk directory
	fmt.Printf("Walking directory: %s\n", walkDir)
	err = filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(walkDir, path)
		if relPath == "." {
			relPath = "."
		}
		fmt.Printf("  %s (%s)\n", relPath, info.Mode())
		return nil
	})
	if err != nil {
		log.Printf("Error walking directory: %v", err)
	}
	
	// Clean up
	os.RemoveAll(walkDir)

	// 12. System Calls
	fmt.Println("\n12. System Calls:")
	
	// Get system information
	var uname syscall.Utsname
	err = syscall.Uname(&uname)
	if err != nil {
		log.Printf("Error getting system info: %v", err)
	} else {
		fmt.Printf("System: %s\n", string(uname.Sysname[:]))
		fmt.Printf("Node: %s\n", string(uname.Nodename[:]))
		fmt.Printf("Release: %s\n", string(uname.Release[:]))
		fmt.Printf("Version: %s\n", string(uname.Version[:]))
		fmt.Printf("Machine: %s\n", string(uname.Machine[:]))
	}

	// 13. File Truncation
	fmt.Println("\n13. File Truncation:")
	
	// Create file with content
	truncFile := "truncate_test.txt"
	file, err = os.Create(truncFile)
	if err != nil {
		log.Printf("Error creating file: %v", err)
	} else {
		file.WriteString("This is a long line of text that will be truncated")
		file.Close()
		
		// Get original size
		info, _ := os.Stat(truncFile)
		fmt.Printf("Original size: %d bytes\n", info.Size())
		
		// Truncate file
		err = os.Truncate(truncFile, 10)
		if err != nil {
			log.Printf("Error truncating file: %v", err)
		} else {
			info, _ = os.Stat(truncFile)
			fmt.Printf("Truncated size: %d bytes\n", info.Size())
			
			// Read truncated content
			content, _ := os.ReadFile(truncFile)
			fmt.Printf("Truncated content: %q\n", string(content))
		}
		
		// Clean up
		os.Remove(truncFile)
	}

	// 14. File Renaming
	fmt.Println("\n14. File Renaming:")
	
	// Create file to rename
	oldName := "old_name.txt"
	newName := "new_name.txt"
	file, err = os.Create(oldName)
	if err != nil {
		log.Printf("Error creating file: %v", err)
	} else {
		file.WriteString("This file will be renamed")
		file.Close()
		
		// Rename file
		err = os.Rename(oldName, newName)
		if err != nil {
			log.Printf("Error renaming file: %v", err)
		} else {
			fmt.Printf("File renamed from %s to %s\n", oldName, newName)
			
			// Verify rename
			_, err = os.Stat(oldName)
			if os.IsNotExist(err) {
				fmt.Println("Old file no longer exists")
			}
			
			_, err = os.Stat(newName)
			if err == nil {
				fmt.Println("New file exists")
			}
		}
		
		// Clean up
		os.Remove(newName)
	}

	// 15. Process Exit
	fmt.Println("\n15. Process Exit:")
	fmt.Println("Process will exit in 2 seconds...")
	time.Sleep(2 * time.Second)
	
	// Note: We don't actually exit here to allow the program to complete
	// os.Exit(0) would terminate the program immediately

	fmt.Println("\n🎉 os Package Mastery Complete!")
}

// Helper function to copy file
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
