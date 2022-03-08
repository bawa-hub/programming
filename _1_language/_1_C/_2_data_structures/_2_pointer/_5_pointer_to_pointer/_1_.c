#include <stdio.h>

int main()
{
    int a = 10;
    int *p = &a;
    int **q = &p;
    int ***r = &q;
    int ****s = &r;
    int *****t = &s;
    int ******u = &t;

    printf("addresses of a, p, q, r, s, t are: %x %x %x %x %x %x\n", p, q, r, s, t, u);
    printf("values of a, p, q, r, s, t are: %d %x %x %x %x %x\n", *p, *q, *r, *s, *t, *u);
    printf("%d %d %d %d %d %d %d\n", a, *p, **q, ***r, ****s, *****t, ******u);
    printf("address of a is: %x", **r);
}