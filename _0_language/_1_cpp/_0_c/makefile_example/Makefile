# Makefile for a simple C program

# Compiler and flags
CC = gcc
CFLAGS = -Wall -g

# Targets and dependencies
all: myprogram

myprogram: main.o
	$(CC) $(CFLAGS) -o myprogram main.o

main.o: main.c
	$(CC) $(CFLAGS) -c main.c

clean:
	rm -f myprogram *.o
