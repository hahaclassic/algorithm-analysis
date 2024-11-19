def RecursiveLevenshtein(s1: str, s2: str) -> int:
    length1, length2 = len(s1), len(s2)
    if length1 == 0 or length2 == 0:
        return abs(length1 - length2)
    if s1[-1] == s2[-1]:
        return RecursiveLevenshtein(s1[:-1], s2[:-1])
    
    return min(
        RecursiveLevenshtein(s1, s2[:-1]) + INSERT_COST,
        RecursiveLevenshtein(s1[:-1], s2) + DELETE_COST,
        RecursiveLevenshtein(s1[:-1], s2[:-1]) + REPLACE_COST)