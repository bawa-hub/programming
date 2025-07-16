#include <stdio.h>
int main()
{
    int a = 10, b = 3;
    float c, d;
    c = a / b;
    d = (float)a / b;
    printf("Result C: %.2f\n", c);
    printf("Result D:%.2f", d);
    return 0;
}