#include <stdio.h>
int main()
{
    int var = 5;
    int var1 = 7;
    printf("var: %d\n", var);

    // Notice the use of & before var
    printf("address of var: %p in memory\n", &var);
    printf("address of var1: %p in memory", &var1);
    return 0;
}