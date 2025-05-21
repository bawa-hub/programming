package _2_dsa_java._3_techniques._1_hash_table;
import java.util.Scanner;

public class _2_char_hashing {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        String s = sc.nextLine();

        // Precompute character frequencies
        int[] hash = new int[256];
        for (int i = 0; i < s.length(); i++) {
            hash[s.charAt(i)]++;
        }

        int q = sc.nextInt();
        sc.nextLine(); // consume newline

        while (q-- > 0) {
            char c = sc.nextLine().charAt(0);
            System.out.println(hash[c]);
        }

        sc.close();
    }
}
