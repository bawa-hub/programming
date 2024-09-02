import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        // Input the string
        String s = scanner.nextLine();

        // Precompute the frequency of each character
        int[] hash = new int[256]; 
        for (int i = 0; i < s.length(); i++) {
            hash[s.charAt(i)]++;
        }

        // Input the number of queries
        int q = scanner.nextInt();
        scanner.nextLine();  // Consume the newline character

        // Process each query
        while (q-- > 0) {
            char c = scanner.nextLine().charAt(0);
            System.out.println(hash[c]);
        }
        
        scanner.close();
    }
}

