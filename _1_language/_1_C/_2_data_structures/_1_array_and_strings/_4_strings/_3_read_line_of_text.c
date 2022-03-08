#include <stdio.h>
int main()
{
    char name[30];
    printf("Enter name: ");
    fgets(name, sizeof(name), stdin); // read string
    printf("Name: ");
    puts(name); // display string
    return 0;
}

// Output
// Enter name: Tom Hanks
// Name: Tom Hanks

// Note:
// The gets() function can also be to take input from the user.
// However, it is removed from the C standard.
// It's because gets() allows you to input any length of characters.
// Hence, there might be a buffer overflow.