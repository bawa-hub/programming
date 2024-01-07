public class DataTypes {
    public static void main(String[] args) {

        // There are 8 data types predefined in Java, known as primitive data types.
        // Note: In addition to primitive data types, there are also referenced types
        // (object type).
        //
        // Primitive data types:
        //
        // byte
        // short
        // int
        // long
        // char
        // float
        // double
        // boolean
        //
        // Referenced types:
        //
        // String
        // Class
        // Array
        // List
        // Map
        // Stack
        // Queue
        // HashSet
        // TreeSet
        // PriorityQueue
        // Hashtable

        // boolean
        boolean flag = true;
        System.out.println(flag);

        // byte
        // byte data type can have values from -128 to 127
        // (8-bit signed two's complement integer).
        byte b = 127;
        System.out.println(b);

        // short
        // short data type can have values from -32768 to 32767
        // (16-bit signed two's complement integer).
        short s = 32767;
        System.out.println(s);

        // int
        // int data type can have values from -2147483648 to 2147483647
        // (32-bit signed two's complement integer).
        int i = 2147483647;
        System.out.println(i);

        // long
        // long data type can have values from -9223372036854775808 to
        // 9223372036854775807
        // (64-bit signed two's complement integer).
        long l = 9223372036854775807L;
        System.out.println(l);

        // char
        // char data type can have values from -128 to 127
        // (8-bit signed two's complement integer).
        char c = 'A';
        System.out.println(c);

        // float
        // float data type can have values from -3.4028235E38 to 3.4028235E38
        // (32-bit IEEE 754 single precision floating point number).
        float f = 3.4028235E38f;
        System.out.println(f);

        // double
        // double data type can have values from -1.7976931348623157E308 to
        // 1.7976931348623157E308
        // (64-bit IEEE 754 double precision floating point number).
        double d = 1.7976931348623157E308d;
        System.out.println(d);

        // String
        String str = "Hello, World!";
        System.out.println(str);
    }
}
