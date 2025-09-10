# Process Monitor - Project 5

## Learning Objectives
- Master system programming with Go
- Understand process management
- Learn system call interfaces
- Practice real-time monitoring

## Features to Implement
1. **Process Listing**: Display running processes with details
2. **Resource Monitoring**: CPU, memory, I/O usage per process
3. **Process Tree**: Show parent-child relationships
4. **Signal Handling**: Send signals to processes
5. **Real-time Updates**: Live monitoring with refresh
6. **Filtering & Search**: Find processes by name, PID, user

## Technical Concepts
- `/proc` filesystem for process information
- `os` package for process operations
- `syscall` package for system calls
- `gopsutil` library for system metrics
- Signal handling with `os/signal`
- Terminal UI with `termui` or `tview`

## Implementation Steps
1. Read process information from `/proc`
2. Parse and display process data
3. Add real-time monitoring
4. Implement process tree visualization
5. Add signal sending capabilities
6. Create filtering and search
7. Add terminal UI for better UX
8. Write comprehensive tests

## Expected Learning Outcomes
- Deep understanding of process management
- Experience with system programming
- Understanding of Linux internals
- Practice with real-time data processing
