
import java.util.*;

// brute force
public class Main {
    public static int[] moveZeros(int n, int[] a) {
        // temporary array:
        ArrayList<Integer> temp = new ArrayList<>();
        // copy non-zero elements
        // from original -> temp array:
        for (int i = 0; i < n; i++) {
            if (a[i] != 0)
                temp.add(a[i]);
        }

        // number of non-zero elements.
        int nz = temp.size();

        // copy elements from temp
        // fill first nz fields of
        // original array:
        for (int i = 0; i < nz; i++) {
            a[i] = temp.get(i);
        }

        // fill rest of the cells with 0:
        for (int i = nz; i < n; i++) {
            a[i] = 0;
        }
        return a;
    }

    public static void main(String[] args) {
        int[] arr = { 1, 0, 2, 3, 2, 0, 0, 4, 5, 1 };
        int n = 10;
        int[] ans = moveZeros(n, arr);
        for (int i = 0; i < n; i++) {
            System.out.print(ans[i] + " ");
        }
        System.out.println("");
    }
}

// two pointer
class tUf {
    public static int[] moveZeros(int n, int[] a) {
        int j = -1;
        // place the pointer j:
        for (int i = 0; i < n; i++) {
            if (a[i] == 0) {
                j = i;
                break;
            }
        }

        // no non-zero elements:
        if (j == -1)
            return a;

        // Move the pointers i and j
        // and swap accordingly:
        for (int i = j + 1; i < n; i++) {
            if (a[i] != 0) {
                // swap a[i] & a[j]:
                int tmp = a[i];
                a[i] = a[j];
                a[j] = tmp;
                j++;
            }
        }
        return a;
    }

    public static void main(String[] args) {
        int[] arr = { 1, 0, 2, 3, 2, 0, 0, 4, 5, 1 };
        int n = 10;
        int[] ans = moveZeros(n, arr);
        for (int i = 0; i < n; i++) {
            System.out.print(ans[i] + " ");
        }
        System.out.println("");
    }
}
