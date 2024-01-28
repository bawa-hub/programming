// There are two types of operators used for accessing members of a structure.
// . - Member operator
// -> - Structure pointer operator

#include <stdio.h>
#include <string.h>

// create struct with person1 variable
struct Person
{
    char name[50];
    int citNo;
    float salary;
} person1;

int main()
{

    // assign value to name of person1
    strcpy(person1.name, "George Orwell");
    // assign values to other person1 variables
    person1.citNo = 1984;
    person1.salary = 2500;
    // print struct variables
    printf("Name: %s\n", person1.name);
    printf("Citizenship No.: %d\n", person1.citNo);
    printf("Salary: %.2f", person1.salary);
    return 0;
}
// we have used strcpy() function to assign the value to person1.name.
// This is because name is a char array (C-string) and
// we cannot use the assignment operator = with it after we have declared the string.