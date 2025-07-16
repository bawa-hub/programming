import java.util.Scanner;

public class _1_input_output {

    public static void main(String[] args) {

        // Output

        // print() - It prints string inside the quotes.
        // println() - It prints string inside the quotes similar like print() method.
        //             Then the cursor moves to the beginning of the next line.
        // printf() - It provides string formatting (similar to printf in C/C++ programming).

        System.out.println("sadf");
        System.out.print("Hello");


        // Input

        Scanner input = new Scanner(System.in);
        System.out.print("Enter an integer: ");
        int number = input.nextInt(); // int input
        System.out.println("You entered " + number);
        input.close(); // closing the scanner object

        // use nextLong(), nextFloat(), nextDouble(), and next() methods to get long, float, double, and string input respectively from the user.

        // Getting float input
        System.out.print("Enter float: ");
        float myFloat = input.nextFloat();
        System.out.println("Float entered = " + myFloat);
    	
        // Getting double input
        System.out.print("Enter double: ");
        double myDouble = input.nextDouble();
        System.out.println("Double entered = " + myDouble);
    	
        // Getting String input
        System.out.print("Enter text: ");
        String myString = input.next();
        System.out.println("Text entered = " + myString);



    }
}
