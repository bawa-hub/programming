const arr = [1,2,3,4,4,4,5,6]

const freq = {}
let ans = 0;

for(let i=0;i<arr.length;i++) {
  freq[arr[i]] = (freq[arr[i]] || 0) + 1;
//   ans = Math.max(ans, freq[arr[i]]);
if(freq[arr[i]] > ans) ans = arr[i]
}

console.log(freq)
console.log('====================================');
console.log(ans);
console.log('====================================');