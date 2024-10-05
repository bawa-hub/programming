public class Solution {
    public int divisorSubstrings(int num, int k) {
        String s = Integer.toString(num);

        int i = 0, j = 0;
        int n = s.length();
        int cnt = 0;

        while (j < n) {
            if (j - i + 1 == k) {
                int curr = Integer.parseInt(s.substring(i, j + 1));
                if (curr != 0 && num % curr == 0) {
                    cnt++;
                }
                i++;
            }
            j++;
        }

        return cnt;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int num = 120, k = 2;
        int result = solution.divisorSubstrings(num, k);
        System.out.println(result);
    }
}
