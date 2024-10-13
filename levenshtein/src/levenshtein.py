DELETE_COST = 1
INSERT_COST = 1
REPLACE_COST = 1
TRANSPOSITION_COST = 1

def createMatrix(rows: int, cols: int) -> list[list[int]]:
    return [[0 for _ in range(cols)] for _ in range(rows)]

def RecursiveLevenshtein(s1: str, s2: str) -> int:
    if len(s1) == 0 or len(s2) == 0:
        return abs(len(s1) - len(s2))
   
    if s1[0] == s2[0]:
        return RecursiveLevenshtein(s1[1:], s2[1:])
    
    return min(
        RecursiveLevenshtein(s1, s2[1:]) + INSERT_COST,
        RecursiveLevenshtein(s1[1:], s2) + DELETE_COST,
        RecursiveLevenshtein(s1[1:], s2[1:]) + REPLACE_COST
    )

def RecursiveCacheLevenshtein(s1: str, s2: str, memo: dict = None) -> int:
    if memo is None:
        memo = {}

    key = (len(s1), len(s2))
    if key in memo:
        return memo[key]

    if len(s1) == 0 or len(s2) == 0:
        return abs(len(s1) - len(s2))
    if s1[0] == s2[0]:
        return RecursiveCacheLevenshtein(s1[1:], s2[1:])

    distance = min(
        RecursiveCacheLevenshtein(s1, s2[1:], memo) + INSERT_COST,
        RecursiveCacheLevenshtein(s1[1:], s2, memo) + DELETE_COST,
        RecursiveCacheLevenshtein(s1[1:], s2[1:], memo) + REPLACE_COST
    )
    memo[key] = distance

    return distance 


def DynamicLevenshtein(s1: str, s2: str) -> int:
    length1, length2 = len(s1), len(s2)

    if len(s1) == 0 or len(s2) == 0:
        return abs(length1 - length2)
    
    matrix = createMatrix(length1+1, length2+1)
    for i in range(1, length1 + 1):
        matrix[i][0] = matrix[i - 1][0] + DELETE_COST

    for j in range(1, length2 + 1):
        matrix[0][j] = matrix[0][j - 1] + INSERT_COST

    for i in range(1, length1 + 1):
        for j in range(1, length2 + 1): 
            cost = 0 if s1[i - 1] == s2[j - 1] else REPLACE_COST
           
            matrix[i][j] = min(
                matrix[i - 1][j] + DELETE_COST, 
                matrix[i][j - 1] + INSERT_COST,
                matrix[i - 1][j - 1] + cost
            )

    return matrix[length1][length2]

def DynamicDamerauLevenshtein(s1: str, s2: str) -> int:
    length1, length2 = len(s1), len(s2)

    if len(s1) == 0 or len(s2) == 0:
        return abs(length1 - length2)

    matrix = createMatrix(length1+1, length2+1)

    for i in range(1, length1 + 1):
        for j in range(1, length2 + 1):
            cost = 0 if s1[i - 1] == s2[j - 1] else REPLACE_COST

            matrix[i][j] = min(
                matrix[i - 1][j] + DELETE_COST,
                matrix[i][j - 1] + INSERT_COST,
                matrix[i - 1][j - 1] + cost
            )

            if i > 1 and j > 1 and s1[i - 1] == s2[j - 2] \
                and s1[i - 2] == s2[j - 1]:
        
                matrix[i][j] = min(
                    matrix[i][j], 
                    matrix[i - 2][j - 2] + TRANSPOSITION_COST
                )

    return matrix[length1][length2]