public class DataTypes {
    public static void main(String[] args) {
        // There are 8 primitive data types in java

        // 1. boolean
        boolean flag = true;
        System.out.println(flag);

        // 2. byte
        // byte data type can have values from -128 to 127 (8-bit signed two's complement integer).
        byte range = 124;
        System.out.println(range);

        // 3. short
        // short data type in Java can have values from -32768 to 32767 (16-bit signed two's complement integer).
        short temperature = -200;
        System.out.println(temperature);

        // 4. int
        // int data type can have values from -231 to 231-1 (32-bit signed two's complement integer).
        // using Java 8 or later, you can use an unsigned 32-bit integer. This will have a minimum value of 0 and a maximum value of 232-1.
        int num = -4250000;
        System.out.println(num); 

        // 5. long
        // long data type can have values from -263 to 263-1 (64-bit signed two's complement integer).
        // using Java 8 or later, you can use an unsigned 64-bit integer with a minimum value of 0 and a maximum value of 264-1.
        long max = -42332200000L;
        System.out.println(max); 

        // 6. double
        // double data type is a double-precision 64-bit floating-point.
        // It should never be used for precise values such as currency.
        double number = -42.3;
        System.out.println(number);

        // 7. float
        // float data type is a single-precision 32-bit floating-point.
        float f = -42.3f;
        System.out.println(f); 

        // 8. char
        // It's a 16-bit Unicode character.
        // minimum value of the char data type is '\u0000' (0) and the maximum value of the is '\uffff'.
        char letter = '\u0051';
        System.out.println(letter);
        char letter1 = '9';
        System.out.println(letter1);
        char letter2 = 65;
        System.out.println(letter2);


    }
}