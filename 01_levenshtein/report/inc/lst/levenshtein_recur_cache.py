def RecursiveCacheLevenshtein(s1: str, s2: str, memo: dict = None) -> int:
    if memo is None:
        memo = {}
        
    length1, length2 = len(s1), len(s2)
    key = (length1, length2)
    if key in memo:
        return memo[key]

    if length1 == 0 or length2 == 0:
        return abs(length1 - length2)
    if s1[-1] == s2[-1]:
        distance = RecursiveCacheLevenshtein(s1[:-1], s2[:-1])
    else:
        distance = min(
            RecursiveCacheLevenshtein(s1, s2[:-1], memo) + INSERT_COST,
            RecursiveCacheLevenshtein(s1[:-1], s2, memo) + DELETE_COST,
            RecursiveCacheLevenshtein(s1[:-1], s2[:-1], memo) + REPLACE_COST)
        
    memo[key] = distance
    return distance 