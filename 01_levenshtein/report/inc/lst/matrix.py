def createMatrix(rows: int, cols: int) -> list[list[int]]:
    return [[0 for _ in range(cols)] for _ in range(rows)]

def getInitialMatrix(rows: int, cols: int) -> list[list[int]]:
    matrix = createMatrix(rows+1, cols+1)
    for i in range(1, rows + 1):
        matrix[i][0] = matrix[i - 1][0] + DELETE_COST
    for j in range(1, cols + 1):
        matrix[0][j] = matrix[0][j - 1] + INSERT_COST

    return matrix