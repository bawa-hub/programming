#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// ====== Block Structure for Allocator ======
typedef struct block {
    size_t size;
    int free;
    struct block *next;
} block_t;

#define HEAP_SIZE 1024 * 1024  // 1MB Heap
static char heap[HEAP_SIZE];
static block_t *free_list = NULL;

// ====== Leak Tracker Structure ======
typedef struct LeakEntry {
    void *ptr;
    size_t size;
    int is_freed;
    struct LeakEntry *next;
} LeakEntry;

LeakEntry *leak_list = NULL;

// ====== Leak Tracker Helpers ======
void add_leak_entry(void *ptr, size_t size) {
    LeakEntry *entry = (LeakEntry *)malloc(sizeof(LeakEntry));
    entry->ptr = ptr;
    entry->size = size;
    entry->is_freed = 0;
    entry->next = leak_list;
    leak_list = entry;
}

void mark_freed(void *ptr) {
    LeakEntry *curr = leak_list;
    while (curr) {
        if (curr->ptr == ptr) {
            curr->is_freed = 1;
            return;
        }
        curr = curr->next;
    }
}

void report_leaks() {
    LeakEntry *curr = leak_list;
    printf("\n==== MEMORY LEAK REPORT ====\n");
    int leaks = 0;
    while (curr) {
        if (!curr->is_freed) {
            printf("Leaked: ptr=%p, size=%zu bytes\n", curr->ptr, curr->size);
            leaks++;
        }
        curr = curr->next;
    }

    if (leaks == 0) {
        printf("No memory leaks detected ✅\n");
    } else {
        printf("%d leak(s) detected ❌\n", leaks);
    }
}

// ====== Allocator Helpers ======
void init_heap() {
    free_list = (block_t *)heap;
    free_list->size = HEAP_SIZE - sizeof(block_t);
    free_list->free = 1;
    free_list->next = NULL;
}

void split_block(block_t *block, size_t size) {
    block_t *new_block = (block_t *)((char *)block + sizeof(block_t) + size);
    new_block->size = block->size - size - sizeof(block_t);
    new_block->free = 1;
    new_block->next = block->next;

    block->size = size;
    block->next = new_block;
}

void *my_malloc(size_t size) {
    block_t *curr = free_list;

    while (curr) {
        if (curr->free && curr->size >= size) {
            if (curr->size > size + sizeof(block_t)) {
                split_block(curr, size);
            }
            curr->free = 0;
            void *user_ptr = (char *)curr + sizeof(block_t);
            add_leak_entry(user_ptr, size);
            return user_ptr;
        }
        curr = curr->next;
    }

    return NULL; // Not enough space
}

void my_free(void *ptr) {
    if (!ptr) return;

    block_t *block_ptr = (block_t *)((char *)ptr - sizeof(block_t));
    block_ptr->free = 1;
    mark_freed(ptr);
}

// ====== calloc and realloc ======
void *my_calloc(size_t num, size_t size) {
    size_t total_size = num * size;
    void *ptr = my_malloc(total_size);
    if (ptr) {
        memset(ptr, 0, total_size);
    }
    return ptr;
}

void *my_realloc(void *ptr, size_t new_size) {
    if (!ptr) return my_malloc(new_size);

    block_t *block_ptr = (block_t *)((char *)ptr - sizeof(block_t));
    if (block_ptr->size >= new_size) {
        return ptr;
    }

    void *new_ptr = my_malloc(new_size);
    if (new_ptr) {
        memcpy(new_ptr, ptr, block_ptr->size);
        my_free(ptr);
    }
    return new_ptr;
}

// ====== Main Test ======
int main() {
    init_heap();
    atexit(report_leaks);  // Report leaks on exit

    void *a = my_malloc(128);
    void *b = my_malloc(256);
    my_free(a);  // Free 'a'
    // b is not freed – intentional leak

    void *c = my_calloc(10, 32);  // 320 bytes
    void *d = my_realloc(c, 500); // reallocates to 500 bytes
    my_free(d);                   // free after realloc

    return 0;
}
