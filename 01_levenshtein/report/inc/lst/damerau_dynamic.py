def DynamicDamerauLevenshtein(s1: str, s2: str) -> int:
    length1, length2 = len(s1), len(s2)
    if length1 == 0 or length2 == 0:
        return abs(length1 - length2)
    matrix = getInitialMatrix(length1, length2)

    for i in range(1, length1 + 1):
        for j in range(1, length2 + 1):
            cost = 0 if s1[i - 1] == s2[j - 1] else REPLACE_COST
            matrix[i][j] = min(
                matrix[i - 1][j] + DELETE_COST,
                matrix[i][j - 1] + INSERT_COST,
                matrix[i - 1][j - 1] + cost)

            if i > 1 and j > 1 and s1[i - 1] == s2[j - 2] \
                and s1[i - 2] == s2[j - 1]:
                matrix[i][j] = min(
                    matrix[i][j], 
                    matrix[i - 2][j - 2] + TRANSPOSITION_COST)

    return matrix[length1][length2]
