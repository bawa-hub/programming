# 50 Easy String Problems (with solution recipes) — Fluency Builder

Use this as a daily drill list. For each problem, your goal is to implement the **recipe** quickly in C++ and test edge cases.

Legend:
- **Patt**: main pattern
- **Recipe**: minimal steps to implement
- **TC/SC**: time/space

---

## A) Basics / Traversal / Indexing (10)

1. **(E) Reverse String (344)**  
   - **Patt**: two pointers  
   - **Recipe**: `l=0,r=n-1` swap while `l<r`.  
   - **TC/SC**: $O(n)$ / $O(1)$

2. **(E) Reverse Words in a String III (557)**  
   - **Patt**: scan + reverse subranges  
   - **Recipe**: iterate; when hit space/end reverse `[start,i-1]`.  
   - **TC/SC**: $O(n)$ / $O(1)$

3. **(E) Valid Palindrome (125)**  
   - **Patt**: two pointers + filtering  
   - **Recipe**: `l/r`; skip non-alnum; compare `tolower`.  
   - **TC/SC**: $O(n)$ / $O(1)$

4. **(E) Valid Palindrome II (680)**  
   - **Patt**: two pointers + one deletion  
   - **Recipe**: on mismatch, check `isPal(l+1,r)` or `isPal(l,r-1)`.  
   - **TC/SC**: $O(n)$ / $O(1)$

5. **(E) Detect Capital (520)**  
   - **Patt**: counting  
   - **Recipe**: count uppercase; valid if all, none, or only first.  
   - **TC/SC**: $O(n)$ / $O(1)$

6. **(E) To Lower Case (709)**  
   - **Patt**: char transform  
   - **Recipe**: for each char: if 'A'..'Z' add 32 (or `tolower`).  
   - **TC/SC**: $O(n)$ / $O(1)$

7. **(E) Length of Last Word (58)**  
   - **Patt**: reverse scan  
   - **Recipe**: skip trailing spaces, then count until space/end.  
   - **TC/SC**: $O(n)$ / $O(1)$

8. **(E) Implement strStr() (28)**  
   - **Patt**: brute force (easy)  
   - **Recipe**: try each start i; match needle sequentially.  
   - **TC/SC**: $O(nm)$ / $O(1)$

9. **(E) Excel Sheet Column Title (168)**  
   - **Patt**: base conversion  
   - **Recipe**: while n>0: n--; push char('A'+n%26); n/=26; reverse.  
   - **TC/SC**: $O(log n)$ / $O(1)$

10. **(E) Excel Sheet Column Number (171)**  
   - **Patt**: base conversion  
   - **Recipe**: ans=0; for c: ans=ans*26 + (c-'A'+1).  
   - **TC/SC**: $O(n)$ / $O(1)$

---

## B) Frequency / Hashing (12)

11. **(E) Valid Anagram (242)**  
   - **Patt**: frequency array  
   - **Recipe**: count 26 for s; subtract for t; all zeros.  
   - **TC/SC**: $O(n)$ / $O(1)$

12. **(E) Ransom Note (383)**  
   - **Patt**: frequency  
   - **Recipe**: count magazine; consume for note; fail if negative.  
   - **TC/SC**: $O(n)$ / $O(1)$

13. **(E) First Unique Character in a String (387)**  
   - **Patt**: freq + scan  
   - **Recipe**: count 26; scan s for first with freq==1.  
   - **TC/SC**: $O(n)$ / $O(1)$

14. **(E) Find the Difference (389)**  
   - **Patt**: XOR  
   - **Recipe**: xor all chars in s and t; result char is answer.  
   - **TC/SC**: $O(n)$ / $O(1)$

15. **(E) Isomorphic Strings (205)**  
   - **Patt**: bijection mapping  
   - **Recipe**: map s->t and t->s; verify consistency.  
   - **TC/SC**: $O(n)$ / $O(Σ)$

16. **(E) Word Pattern (290)**  
   - **Patt**: bijection mapping  
   - **Recipe**: split words; map char->word and word->char.  
   - **TC/SC**: $O(n)$ / $O(n)$

17. **(E) Check if One String Swap Can Make Strings Equal (1790)**  
   - **Patt**: mismatch positions  
   - **Recipe**: collect indices where s[i]!=t[i]; must be 0 or 2 and swap matches.  
   - **TC/SC**: $O(n)$ / $O(1)$

18. **(E) Determine if String Halves Are Alike (1704)**  
   - **Patt**: counting  
   - **Recipe**: count vowels in first half and second half.  
   - **TC/SC**: $O(n)$ / $O(1)$

19. **(E) Check if All Characters Have Equal Number of Occurrences (1941)**  
   - **Patt**: frequency  
   - **Recipe**: count; ensure all non-zero counts equal.  
   - **TC/SC**: $O(n)$ / $O(1)$

20. **(E) Find Words That Can Be Formed by Characters (1160)**  
   - **Patt**: freq compare  
   - **Recipe**: count chars; for each word count and compare <=.  
   - **TC/SC**: $O(total chars)$ / $O(1)$

21. **(E) Jewels and Stones (771)**  
   - **Patt**: set membership  
   - **Recipe**: put jewels in set; count stones in set.  
   - **TC/SC**: $O(n)$ / $O(1)$

22. **(E) Unique Morse Code Words (804)**  
   - **Patt**: mapping + set  
   - **Recipe**: transform each word to morse; insert in set; return size.  
   - **TC/SC**: $O(total chars)$ / $O(n)$

---

## C) Two Pointers / Simple Greedy (10)

23. **(E) Backspace String Compare (844)**  
   - **Patt**: two pointers from end  
   - **Recipe**: scan backward with skip counters; compare next valid chars.  
   - **TC/SC**: $O(n)$ / $O(1)$

24. **(E) Merge Strings Alternately (1768)**  
   - **Patt**: two pointers  
   - **Recipe**: append from each string if available.  
   - **TC/SC**: $O(n+m)$ / $O(1)$ extra

25. **(E) Reverse Prefix of Word (2000)**  
   - **Patt**: find index + reverse  
   - **Recipe**: find first occurrence of ch; reverse s[0..idx].  
   - **TC/SC**: $O(n)$ / $O(1)$

26. **(E) Rotated Digits (788)**  
   - **Patt**: digit check  
   - **Recipe**: for each number, validate digits; must contain at least one {2,5,6,9}.  
   - **TC/SC**: $O(n log n)$ / $O(1)$

27. **(E) Long Pressed Name (925)**  
   - **Patt**: two pointers  
   - **Recipe**: advance in typed; match name chars; allow repeats of typed prev.  
   - **TC/SC**: $O(n+m)$ / $O(1)$

28. **(E) Check If Two String Arrays are Equivalent (1662)**  
   - **Patt**: streaming comparison  
   - **Recipe**: pointers on arrays + indices in strings; compare current chars without concatenating.  
   - **TC/SC**: $O(total chars)$ / $O(1)$

29. **(E) Valid Parentheses (20)**  
   - **Patt**: stack  
   - **Recipe**: push opens; on close check top matches; end stack empty.  
   - **TC/SC**: $O(n)$ / $O(n)$

30. **(E) Remove All Adjacent Duplicates In String (1047)**  
   - **Patt**: stack simulation  
   - **Recipe**: use string as stack; if top==c pop else push.  
   - **TC/SC**: $O(n)$ / $O(n)$

31. **(E) Make The String Great (1544)**  
   - **Patt**: stack  
   - **Recipe**: cancel if abs(a-b)==32 (case pair); else push.  
   - **TC/SC**: $O(n)$ / $O(n)$

32. **(E) Check If a Word Occurs As a Prefix of Any Word in a Sentence (1455)**  
   - **Patt**: split + startsWith  
   - **Recipe**: parse words; return first index where word starts with searchWord.  
   - **TC/SC**: $O(n)$ / $O(1)$

---

## D) String Building / Parsing (10)

33. **(E) Add Strings (415)**  
   - **Patt**: digit simulation  
   - **Recipe**: add from end with carry; reverse result.  
   - **TC/SC**: $O(n)$ / $O(1)$ extra

34. **(E) Add Binary (67)**  
   - **Patt**: digit simulation  
   - **Recipe**: add bits with carry from end.  
   - **TC/SC**: $O(n)$ / $O(1)$

35. **(E) Roman to Integer (13)**  
   - **Patt**: mapping + greedy scan  
   - **Recipe**: sum values; if current < next, subtract current else add.  
   - **TC/SC**: $O(n)$ / $O(1)$

36. **(E) Integer to Roman (12)**  
   - **Patt**: greedy mapping  
   - **Recipe**: iterate value-symbol pairs; while n>=value append symbol and subtract.  
   - **TC/SC**: $O(1)$ / $O(1)$

37. **(E) Valid Word Abbreviation (408)**  
   - **Patt**: two pointers parse number  
   - **Recipe**: scan abbr; if digit parse len (no leading 0); jump in word; else match char.  
   - **TC/SC**: $O(n)$ / $O(1)$

38. **(E) Number of Segments in a String (434)**  
   - **Patt**: scan  
   - **Recipe**: count transitions from space->non-space.  
   - **TC/SC**: $O(n)$ / $O(1)$

39. **(E) Replace All Digits with Characters (1844)**  
   - **Patt**: build/modify string  
   - **Recipe**: for odd i: s[i]=s[i-1]+(s[i]-'0').  
   - **TC/SC**: $O(n)$ / $O(1)$

40. **(E) Check if the Sentence Is Pangram (1832)**  
   - **Patt**: bitset/freq  
   - **Recipe**: mark 26 letters; verify all seen.  
   - **TC/SC**: $O(n)$ / $O(1)$

41. **(E) Maximum Number of Words Found in Sentences (2114)**  
   - **Patt**: count spaces  
   - **Recipe**: words = spaces+1 for non-empty; take max.  
   - **TC/SC**: $O(total chars)$ / $O(1)$

42. **(E) Goal Parser Interpretation (1678)**  
   - **Patt**: parsing  
   - **Recipe**: scan; "G"->G, "()"->o, "(al)"->al.  
   - **TC/SC**: $O(n)$ / $O(1)$

---

## E) Small Pattern Matching / Simple Logic (8)

43. **(E) Repeated Substring Pattern (459)**  
   - **Patt**: string trick  
   - **Recipe**: if `s` is in `(s+s).substr(1,2n-2)` then true.  
   - **TC/SC**: $O(n)$ / $O(n)$

44. **(E) Longest Common Prefix (14)**  
   - **Patt**: vertical scan  
   - **Recipe**: take first string as prefix; shrink while mismatch with each word.  
   - **TC/SC**: $O(total chars)$ / $O(1)$

45. **(E) Valid Word Square (422)**  
   - **Patt**: index bounds  
   - **Recipe**: for i,j: char at words[i][j] must equal words[j][i] (check bounds).  
   - **TC/SC**: $O(n^2)$ / $O(1)$

46. **(E) String Matching in an Array (1408)**  
   - **Patt**: brute force contains  
   - **Recipe**: sort by length; for each i check if it’s substring of any j!=i.  
   - **TC/SC**: $O(n^2 * L^2)$ / $O(1)$

47. **(E) Find the Index of the First Occurrence in a String (28)**  
   - **Patt**: brute / KMP optional  
   - **Recipe**: same as #8; write clean loops and bounds.  
   - **TC/SC**: $O(nm)$ / $O(1)$

48. **(E) Check if a String Is an Acronym of Words (2828)**  
   - **Patt**: simple build  
   - **Recipe**: build acronym from first char of each word; compare.  
   - **TC/SC**: $O(total chars)$ / $O(1)$

49. **(E) Shuffle String (1528)**  
   - **Patt**: index placement  
   - **Recipe**: create res size n; res[indices[i]] = s[i].  
   - **TC/SC**: $O(n)$ / $O(n)$

50. **(E) Defanging an IP Address (1108)**  
   - **Patt**: build string  
   - **Recipe**: append chars; if '.' append "[.]".  
   - **TC/SC**: $O(n)$ / $O(n)$

---

## Daily fluency routine (10–20 min)

- Implement 3 problems/day from different sections.
- Always test:
  - empty string, 1-char, all same chars
  - punctuation/spaces (where applicable)
  - uppercase/lowercase differences

