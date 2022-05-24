#include <stdio.h>
int main()
{
    int testInteger;
    float num1;
    double num2;

    printf("Enter an integer: ");
    scanf("%d", &testInteger);

    printf("Enter a float: ");
    scanf("%f", &num1);

    printf("Enter another double: ");
    scanf("%lf", &num2);

    printf("Number = %d\n", testInteger);
    printf("num1 = %f\n", num1);
    printf("num2 = %lf\n", num2);

    return 0;
}