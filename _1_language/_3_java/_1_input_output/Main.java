import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        // java output
        // System.out.println();
        // System.out.print();
        // System.out.printf();

        // print() - It prints string inside the quotes.
        // println() - It prints string inside the quotes similar like print() method.
        // Then the cursor moves to the beginning of the next line.
        // printf() - It provides string formatting (similar to printf in C/C++
        // programming).

        Double number = -10.6;
        System.out.println("I am " + "awesome.");
        System.out.println("Number = " + number);

        // java input
        // Java provides different ways to get input from the user.

        Scanner input = new Scanner(System.in);

        // Getting integer input
        System.out.print("Enter an integer: ");
        int num = input.nextInt();
        System.out.println("You entered " + num);

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

        input.close();

    }
}
