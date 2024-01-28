public class Main {
    public static void main(String[] args) {

        // syntax
        // dataType[] arrayName;

        // declare an array
        double[] data;

        // allocate memory
        data = new double[10];

        // declare and initialize and array
        int[] age = { 12, 4, 5, 2, 5 };

        System.out.println("Using for Loop:");
        for (int i = 0; i < age.length; i++) {
            System.out.println(age[i]);
        }

        System.out.println("Using for-each Loop:");
        for (int a : age) {
            System.out.println(a);
        }

        // multi-dimensional array
        int[][] arr = new int[2][3];

        // initialize the array
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr[i].length; j++) {
                arr[i][j] = i * 10 + j;
            }
        }

        System.out.println("Using for Loop:");
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr[i].length; j++) {
                System.out.print(arr[i][j] + " ");
            }
            System.out.println();
        }

        // create a 3d array
        int[][][] test = {
                {
                        { 1, -2, 3 },
                        { 2, 3, 4 }
                },
                {
                        { -4, -5, 6, 9 },
                        { 1 },
                        { 2, 3 }
                }
        };

        // for..each loop to iterate through elements of 3d array
        for (int[][] array2D : test) {
            for (int[] array1D : array2D) {
                for (int item : array1D) {
                    System.out.println(item);
                }
            }
        }

    }
}
