// a string is a sequence of characters terminated with a null character \0.
// When the compiler encounters a sequence of characters enclosed in the double quotation marks,
// it appends a null character \0 at the end by default

// declaration
char s[5]; // declared a string of 5 characters

// initialize strings
char c[] = "abcd";
char c[50] = "abcd";
char c[] = {'a', 'b', 'c', 'd', '\0'};
char c[5] = {'a', 'b', 'c', 'd', '\0'};
char c[5] = "abcde"; // This is bad and you should never do this.
// we are trying to assign 6 characters (the last character is '\0') to a char array having 5 characters.

// Assigning Values to Strings
// Arrays and strings are second-class citizens in C;
// they do not support the assignment operator once it is declared
char c[100];
c = "C programming"; // Error! array type is not assignable.
                     // Use the strcpy() function to copy the string instead.