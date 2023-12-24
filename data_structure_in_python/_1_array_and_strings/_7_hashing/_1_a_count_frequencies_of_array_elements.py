#  https://practice.geeksforgeeks.org/problems/frequency-of-array-elements-1587115620/0

def frequency(arr):
    # Using a dictionary to store frequencies
    freq_map = {}

    for i in arr:
        freq_map[i] = freq_map.get(i, 0) + 1

    # Traverse through the dictionary and print frequencies
    for key, value in freq_map.items():
        print(key, value)

if __name__ == "__main__":
    arr = [10, 5, 10, 15, 10, 5]
    frequency(arr)
