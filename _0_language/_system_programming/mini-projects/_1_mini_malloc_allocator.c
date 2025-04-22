// 🔥 Project: Build Your Own malloc() + free() (Mini Allocator)

// We’ll simulate a heap using a fixed-size array and manage memory ourselves using:
//     A custom allocator: my_malloc(size)
//     A custom deallocator: my_free(ptr)

// 🧱 How It Works (Behind the Scenes)
// We’ll divide a large static array (heap) into blocks:
//     Each block will have a header with:
//         size: how much space the block takes
//         is_free: if it’s available
//     You walk the heap like a linked list of blocks and find one that fits.

#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
#include <string.h>

#define HEAP_SIZE 1024  // Simulated heap size

char heap[HEAP_SIZE];  // Simulated raw memory

typedef struct block {
    size_t size;
    int is_free;
    struct block *next;
} block_t;

block_t *free_list = (block_t *)heap;  // Start of heap

void init_heap() {
    free_list->size = HEAP_SIZE - sizeof(block_t);
    free_list->is_free = 1;
    free_list->next = NULL;
}

void *my_malloc(size_t size) {
    block_t *curr = free_list;

    while (curr != NULL) {
        if (curr->is_free && curr->size >= size) {
            // Split the block if possible
            if (curr->size > size + sizeof(block_t)) {
                block_t *new_block = (block_t *)((char *)curr + sizeof(block_t) + size);
                new_block->size = curr->size - size - sizeof(block_t);
                new_block->is_free = 1;
                new_block->next = curr->next;

                curr->size = size;
                curr->next = new_block;
            }

            curr->is_free = 0;
            return (char *)curr + sizeof(block_t);  // Return pointer to usable memory
        }

        curr = curr->next;
    }

    return NULL;  // Out of memory
}

void my_free(void *ptr) {
    if (!ptr) return;

    block_t *block_ptr = (block_t *)((char *)ptr - sizeof(block_t));
    block_ptr->is_free = 1;

    // Simple coalescing (merge adjacent free blocks)
    block_t *curr = free_list;
    while (curr && curr->next) {
        if (curr->is_free && curr->next->is_free) {
            curr->size += sizeof(block_t) + curr->next->size;
            curr->next = curr->next->next;
        } else {
            curr = curr->next;
        }
    }
}

void *my_calloc(size_t num, size_t size) {
    size_t total_size = num * size;
    void *ptr = my_malloc(total_size);

    if (ptr != NULL) {
        memset(ptr, 0, total_size);
    }

    return ptr;
}
// 🔍 Uses your my_malloc(), then zeroes the memory just like real calloc.

void *my_realloc(void *ptr, size_t new_size) {
    if (ptr == NULL) {
        return my_malloc(new_size);
    }

    block_t *block_ptr = (block_t *)((char *)ptr - sizeof(block_t));
    if (block_ptr->size >= new_size) {
        return ptr;  // Already big enough
    }

    // Need new block
    void *new_ptr = my_malloc(new_size);
    if (new_ptr != NULL) {
        memcpy(new_ptr, ptr, block_ptr->size);  // Copy old content
        my_free(ptr);
    }

    return new_ptr;
}
// 🔍 Realloc copies old data to a new block and frees the old one if needed.

int main() {
    init_heap();

    printf("Allocating 100 bytes...\n");
    void *p1 = my_malloc(100);
    if (p1) {
        memset(p1, 0, 100);
        printf("Success! Pointer: %p\n", p1);
    }

    printf("Allocating 200 bytes...\n");
    void *p2 = my_malloc(200);
    if (p2) {
        memset(p2, 1, 200);
        printf("Success! Pointer: %p\n", p2);
    }

    printf("Freeing first block...\n");
    my_free(p1);

    printf("Freeing second block...\n");
    my_free(p2);

    return 0;
}



// 💪 Level-Up Path from Here:
//     Add realloc() support (resize a block)
//     Track block alignment and padding
//     Use sbrk() / mmap() for real memory pages (on Linux/macOS)
//     Integrate with C malloc hooks (overriding malloc/free)