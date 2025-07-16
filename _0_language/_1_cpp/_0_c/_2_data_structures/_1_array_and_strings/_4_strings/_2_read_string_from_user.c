// scanf() function reads the sequence of characters until it encounters whitespace (space, newline, tab, etc.).
#include <stdio.h>
int main()
{
    char name[20];
    printf("Enter name: ");
    scanf("%s", name);
    printf("Your name is %s.", name);
    return 0;
}

// output
// Enter name: Dennis Ritchie
// Your name is Dennis.