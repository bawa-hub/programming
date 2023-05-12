#include <iostream>
using namespace std;

int minimumRecolors(string blocks, int k) {
        int i=0;
        int j=0;
        int len = blocks.size();
        int mini = 1000;
        int cnt = 0;

        while(j<len) {
            if(blocks[j]=='W') {
                cnt++;
            }

            if(j-i+1<k) j++;

            if(j-i+1==k) {
              mini = min(mini, cnt);
              if(blocks[i]=='W') cnt--;
              i++;j++;
            }
        }

        return mini;

    }

int main(){
     cout << "min: " << minimumRecolors("WWBBBWBBBBBWWBWWWB", 16) << endl;
}