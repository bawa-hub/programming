import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        String s = scanner.nextLine();

        int[] hash = new int[256]; 
        for (int i = 0; i < s.length(); i++) {
            hash[s.charAt(i)]++;
        }

        int q = scanner.nextInt();
        scanner.nextLine();  

        while (q-- > 0) {
            char c = scanner.nextLine().charAt(0);
            System.out.println(hash[c]);
        }
        
        scanner.close();
    }
}

