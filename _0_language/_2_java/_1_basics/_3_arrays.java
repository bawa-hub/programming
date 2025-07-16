// https://www.programiz.com/java-programming/copy-arrays

public class _3_arrays {

    public static void main(String[] args) {

        // declare an array
        // dataType[] arrayName;

        // declare an array
        double[] data;
        // allocate memory
        data = new double[10];

        // double[] data = new double[10]; //declare and allocate in single line

        // declare and initialize and array
        int[] age = { 12, 4, 5, 2, 5 };
        System.out.println("First Element: " + age[0]);

        // loop through the array

        // using for loop
        System.out.println("Using for Loop:");
        for (int i = 0; i < age.length; i++) {
            System.out.println(age[i]);
        }

        // Using the for-each Loop
        for(int a : age) {
            System.out.println(a);
          }
    }

}
