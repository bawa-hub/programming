#include <stdio.h>
void swap(int n1, int n2);

int main()
{
    int arr[3];
    for (int i = 0; i < 3; i++)
    {
        printf(" %p\n", &arr[i]);
    }
    printf("Address of array is %p", arr + 1);
    return 0;
}
