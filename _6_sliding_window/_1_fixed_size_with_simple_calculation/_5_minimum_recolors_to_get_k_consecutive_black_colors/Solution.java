public class Solution {
    public int minimumRecolors(String blocks, int k) {
        int i = 0, j = 0, len = blocks.length(), mini = 1000, cnt = 0;

        while (j < len) {
            if (blocks.charAt(j) == 'W') {
                cnt++;
            }

            if (j - i + 1 == k) {
                mini = Math.min(mini, cnt);
                if (blocks.charAt(i) == 'W') {
                    cnt--;
                }
                i++;
            }

            j++;
        }

        return mini;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        String blocks = "BBWWBBWW";
        int k = 3;
        int result = solution.minimumRecolors(blocks, k);
        System.out.println(result);
    }
}
