# Вариант: 
# двоичный сдвиг вместо умножения на 2; 
# использование инкремента (+=); 
# вынос начальной итерации из каждого внешнего цикла;

def standard_mult(m1: list[list[int]], m2: list[list[int]]) -> list[list[int]]:
    m1_rows, m2_rows = len(m1), len(m2)
    if m1_rows < 1 or m2_rows < 1:
        return None
    m1_cols, m2_cols = len(m1[0]), len(m2[0])
    if m1_cols < 1 or m2_cols < 1 or m1_cols != m2_rows:
        return None
    
    result = create_matrix(m1_rows, m2_cols)
    for i in range(m1_rows):
        for j in range(m2_cols):
            for k in range(m1_cols):
                result[i][j] += m1[i][k] * m2[k][j]
    
    return result


def winograd_mult(m1: list[list[int]], m2: list[list[int]]) -> list[list[int]]:
    m1_rows, m2_rows = len(m1), len(m2)
    if m1_rows < 1 or m2_rows < 1:
        return None
    m1_cols, m2_cols = len(m1[0]), len(m2[0])
    if m1_cols < 1 or m2_cols < 1 or m1_cols != m2_rows:
        return None
    
    result = create_matrix(m1_rows, m2_cols)
    mul_rows, mul_cols = [0] * m1_rows, [0] * m2_cols

    for i in range(m1_rows):
        for j in range(m1_cols // 2):
            mul_rows[i] = mul_rows[i] + m1[i][2 * j] * m1[i][2 * j + 1]
    for i in range(m2_cols):
        for j in range(m2_rows // 2):
            mul_cols[i] = mul_cols[i] + m2[2 * j][i] * m2[2 * j + 1][i]

    for i in range(m1_rows):
        for j in range(m2_cols):
            result[i][j] = -mul_rows[i] - mul_cols[j]
            for k in range(m1_cols // 2):
                result[i][j] = result[i][j] + (m1[i][2 * k] + m2[2 * k + 1][j]) \
                    * (m1[i][2 * k + 1] + m2[2 * k][j])
    
    if m1_cols % 2 == 1:
        for i in range(m1_rows):
            for j in range(m2_cols):
                result[i][j] = result[i][j] + m1[i][m1_cols - 1] * m2[m1_cols - 1][j]

    return result


def winograd_optimized_mult(m1: list[list[int]], m2: list[list[int]]) -> list[list[int]]:
    m1_rows, m2_rows = len(m1), len(m2)
    if m1_rows < 1 or m2_rows < 1:
        return None
    m1_cols, m2_cols = len(m1[0]), len(m2[0])
    if m1_cols < 1 or m2_cols < 1 or m1_cols != m2_rows:
        return None
    
    result = create_matrix(m1_rows, m2_cols)
    mul_rows, mul_cols = create_precomp_row_opt(m1), create_precomp_column_opt(m2)

    m1_half_cols = m1_cols >> 1

    result[0][0] -= mul_rows[0] + mul_cols[0]
    for k in range(m1_half_cols):
        result[0][0] += (m1[0][k << 1] + m2[(k << 1) + 1][0]) \
            * (m1[0][(k << 1) + 1] + m2[k << 1][0])
        
    for j in range(1, m2_cols):
        result[0][j] -= mul_rows[0] + mul_cols[j]
        for k in range(m1_half_cols):
            result[0][j] += (m1[0][k << 1] + m2[(k << 1) + 1][j]) \
                * (m1[0][(k << 1) + 1] + m2[k << 1][j])
    for i in range(1, m1_rows):
        for j in range(m2_cols):
            result[i][j] -= mul_rows[i] + mul_cols[j]
            for k in range(m1_half_cols):
                result[i][j] += (m1[i][k << 1] + m2[(k << 1) + 1][j]) \
                    * (m1[i][(k << 1) + 1] + m2[k << 1][j])
    
    if m1_cols % 2 == 1:
        for j in range(m2_cols):
                result[0][j] += m1[0][m1_cols - 1] * m2[m1_cols - 1][j]
        for i in range(1, m1_rows):
            for j in range(m2_cols):
                result[i][j] += m1[i][m1_cols - 1] * m2[m1_cols - 1][j]

    return result


def create_matrix(rows: int, cols: int) -> list[list[int]]:
    return [[0] * cols for _ in range(rows)]


def create_precomp_row_opt(m1: list[list[int]]) -> list[int]:
    m1_rows, m1_half_cols = len(m1), len(m1[0]) >> 1
    mul_rows = [0] * m1_rows

    for j in range(m1_half_cols):
        mul_rows[0] += m1[0][j << 1] * m1[0][(j << 1) + 1]
    for i in range(1, m1_rows):
        for j in range(m1_half_cols):
            mul_rows[i] += m1[i][j << 1] * m1[i][(j << 1) + 1]

    return mul_rows


def create_precomp_column_opt(m2: list[list[int]]) -> list[int]:
    m2_cols, m2_half_rows = len(m2[0]), len(m2) >> 1
    mul_cols = [0] * m2_cols

    for j in range(m2_half_rows):
        mul_cols[0] += m2[j << 1][0] * m2[(j << 1) + 1][0]
    for i in range(1, m2_cols):
        for j in range(m2_half_rows):
            mul_cols[i] += m2[j << 1][i] * m2[(j << 1) + 1][i]
    
    return mul_cols
