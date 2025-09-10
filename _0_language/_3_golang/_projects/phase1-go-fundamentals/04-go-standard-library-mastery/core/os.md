# os Package - Operating System Interface 🖥️

The `os` package provides a platform-independent interface to operating system functionality. It's essential for file operations, environment variables, process management, and system calls.

## 🎯 Key Concepts

### 1. **File Operations**
- `Open()` - Open file for reading
- `Create()` - Create or truncate file
- `OpenFile()` - Open file with specific flags
- `Remove()` - Remove file or directory
- `Rename()` - Rename file or directory
- `Mkdir()` - Create directory
- `MkdirAll()` - Create directory tree
- `RemoveAll()` - Remove directory tree

### 2. **File Information**
- `Stat()` - Get file info
- `Lstat()` - Get file info (don't follow symlinks)
- `FileInfo` - Interface for file information
- `FileMode` - File permissions and type

### 3. **Environment Variables**
- `Getenv()` - Get environment variable
- `Setenv()` - Set environment variable
- `Unsetenv()` - Unset environment variable
- `Environ()` - Get all environment variables
- `LookupEnv()` - Get environment variable with existence check

### 4. **Process Management**
- `Getpid()` - Get process ID
- `Getppid()` - Get parent process ID
- `Getuid()` - Get user ID
- `Getgid()` - Get group ID
- `Exit()` - Exit program
- `Args` - Command line arguments

### 5. **Standard Streams**
- `Stdin` - Standard input
- `Stdout` - Standard output
- `Stderr` - Standard error

### 6. **File Permissions**
- `ModeDir` - Directory
- `ModeAppend` - Append-only
- `ModeExclusive` - Exclusive use
- `ModeTemporary` - Temporary file
- `ModeSymlink` - Symbolic link
- `ModeDevice` - Device file
- `ModeNamedPipe` - Named pipe
- `ModeSocket` - Unix domain socket
- `ModeSetuid` - Setuid
- `ModeSetgid` - Setgid
- `ModeCharDevice` - Character device
- `ModeSticky` - Sticky bit

## 🚀 Common Patterns

### File Reading
```go
file, err := os.Open("filename.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

### File Writing
```go
file, err := os.Create("output.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

### Environment Variables
```go
home := os.Getenv("HOME")
if home == "" {
    log.Fatal("HOME not set")
}
```

### Command Line Arguments
```go
if len(os.Args) < 2 {
    log.Fatal("Usage: program <filename>")
}
filename := os.Args[1]
```

## ⚠️ Common Pitfalls

1. **Not closing files** - Always use `defer file.Close()`
2. **Ignoring errors** - Always check error returns
3. **Race conditions** - Be careful with concurrent file access
4. **Path separators** - Use `filepath` package for cross-platform paths
5. **Permission errors** - Check file permissions before operations

## 🎯 Best Practices

1. **Always close files** - Use defer statements
2. **Check errors** - Handle all error returns
3. **Use filepath** - For cross-platform path operations
4. **Check permissions** - Before file operations
5. **Use appropriate flags** - For file opening operations

## 🔍 Advanced Features

### File Flags
```go
file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
```

### File Modes
```go
info, err := os.Stat("file.txt")
if err != nil {
    log.Fatal(err)
}
mode := info.Mode()
if mode.IsDir() {
    fmt.Println("It's a directory")
}
```

### Process Information
```go
pid := os.Getpid()
uid := os.Getuid()
gid := os.Getgid()
```

## 📚 Real-world Applications

1. **File Management** - Reading, writing, organizing files
2. **Configuration** - Reading environment variables
3. **Logging** - Writing to log files
4. **CLI Tools** - Processing command line arguments
5. **System Administration** - Managing processes and files

## 🧠 Memory Tips

- **os** = **O**perating **S**ystem
- **Open** = **O**pen for **R**eading
- **Create** = **C**reate **N**ew file
- **Stat** = **S**tatus of file
- **Getenv** = **G**et **E**nvironment **V**ariable
- **Args** = **A**rguments from command line
- **Stdin/Stdout/Stderr** = **S**tandard **I**nput/**O**utput/**E**rror

Remember: The os package is your gateway to the operating system! 🎯
