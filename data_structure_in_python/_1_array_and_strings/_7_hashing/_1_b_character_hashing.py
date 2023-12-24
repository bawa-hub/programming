if __name__ == "__main__":
    s = input()

    # Precompute
    hash_map = {}
    for char in s:
        hash_map[char] = hash_map.get(char, 0) + 1

    q = int(input())
    while q > 0:
        c = input()
        # Fetch
        print(hash_map.get(c, 0))
        q -= 1

