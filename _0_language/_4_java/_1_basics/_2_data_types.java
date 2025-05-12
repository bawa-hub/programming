
// https://www.geeksforgeeks.org/data-types-in-java/
// https://www.programiz.com/java-programming/variables-primitive-data-types
// https://docs.oracle.com/javase/8/docs/api/java/lang/Number.html
// https://stackoverflow.com/questions/801117/whats-the-difference-between-a-single-precision-and-double-precision-floating-p
// https://www.programiz.com/java-programming/examples/int-to-char-conversion


public class _2_data_types {

    public static void main(String[] args) {


        /***
         *  8 Primitive Data Types
         *  Primitive Data Type: These are the basic building blocks that store simple values directly in memory
         * 
         *  1. boolean type : true/false
         *  2. byte type :  -128 to 127 (8-bit signed two's complement integer)
         *  3. short type : -32768 to 32767 (16-bit signed two's complement integer)
         *  4. int type : -231 to 231-1 (32-bit signed two's complement integer)
         *  5. long type : -263 to 263-1 (64-bit signed two's complement integer)
         *  6. double type : double-precision 64-bit floating-point (It should never be used for precise values such as currency)
         *  7. float type : single-precision 32-bit floating-point (It should never be used for precise values such as currency)
         *  8. char type : 16-bit Unicode character (minimum value of the char data type is '\u0000' (0) and the maximum value of the is '\uffff')
         * 
         *  Non-Primitive Data Types (Object Types): These are reference types that store memory addresses of objects.
         *  Examples of Non-primitive data types are String, Array, Class, Interface, and Object
         */



        //  Conversions

        // char to string and vice-versa
        char ch = 'c';
        String st = Character.toString(ch);  //String st = String.valueOf(ch);
        System.out.println("The string is: " + st);

        // char array to String
        char[] ch1 = {'a', 'e', 'i', 'o', 'u'};
        String st1 = String.valueOf(ch1);
        String st2 = new String(ch1);
        System.out.println(st1);
        System.out.println(st2);

        // String to char array
        String st3 = "This is great";
        char[] chars = st3.toCharArray();
        System.out.println(Arrays.toString(chars));

        // char to int
        char a = '5';
        char b = 'c';
        char c = '9';
        int num1 = a; // ASCII value of characters is assigned
        int num2 = b;
        int num3 = Character.getNumericValue(a);
        int num4 = Integer.parseInt(String.valueOf(a));
        int num5 = c - '0';
        System.out.println(num1); // 53
        System.out.println(num2); // 99
        System.out.println(num3); // 5
        System.out.println(num4); // 5
        System.out.println(num5); // 9

        // int to char
        int num6 = 80;
        int num7 = 81;
        int num8 = 1;
        char d = (char) num6;
        char e = (char) num7;
        char f = Character.forDigit(num8, 10);
        char g = (char)(num8 + '0');
        System.out.println(d); // P
        System.out.println(e); // Q
        System.out.println(f); // 1
        System.out.println(g); // 1

        // long to int
        // int to long
        // boolean to string
        // string to boolean
        // string to int

       
    }
    
}
