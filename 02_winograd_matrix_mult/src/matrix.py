
def standard_mult(m1: list[list[int]], m2: list[list[int]]) -> list[list[int]]:
    m1_rows, m1_cols, m2_cols = len(m1), len(m1[0]), len(m2[0])
    if m1_cols != len(m2):
        return None
    
    result = [[0] * m2_cols for _ in range(m1_rows)]
    
    for i in range(m1_rows):
        for j in range(m2_cols):
            for k in range(m1_cols):
                result[i][j] += m1[i][k] * m2[k][j]
    
    return result


def winograd_mult(m1: list[list[int]], m2: list[list[int]]) -> list[list[int]]:
    m1_rows, m1_cols, m2_cols = len(m1), len(m1[0]), len(m2[0])
    if m1_cols != len(m2):
        return None
    
    result = [[0] * m2_cols for _ in range(m1_rows)]
    mul_rows, mul_cols = [0] * m1_rows, [0] * m2_cols

    for i in range(m1_rows):
        for j in range(m1_cols // 2):
            mul_rows[i] = mul_rows[i] + m1[i][2 * j] * m1[i][2 * j + 1]
    for i in range(m2_cols):
        for j in range(m1_cols // 2):
            mul_cols[i] = mul_cols[i] + m2[2 * j][i] * m2[2 * j + 1][i]

    for i in range(m1_rows):
        for j in range(m2_cols):
            result[i][j] = -mul_rows[i] - mul_cols[j]
            for k in range(m1_cols // 2):
                result[i][j] = result[i][j] + (m1[i][2 * k] + m2[2 * k + 1][j]) \
                    * (m1[i][2 * k + 1] + m2[2 * k][j])
    
    if (m1_cols % 2 == 1):
        for i in range(m1_rows):
            for j in range(m2_cols):
                result[i][j] = result[i][j] + m1[i][m1_cols - 1] * m2[m1_cols - 1][j]

    return result


def winograd_optimized_mult(m1: list[list[int]], m2: list[list[int]]) -> list[list[int]]:
    m1_rows, m1_cols, m2_cols = len(m1), len(m1[0]), len(m2[0])
    if m1_cols != len(m2):
        return None
    
    result = [[0] * m2_cols for _ in range(m1_rows)]
    mul_rows, mul_cols = [0] * m1_rows, [0] * m2_cols

    m1_half_cols = m1_cols >> 1
    is_odd = (m1_cols & 1 == 1)

    for i in range(m1_rows):
        for j in range(m1_half_cols):
            mul_rows[i] += m1[i][j << 1] * m1[i][(j << 1) + 1]

    for i in range(m2_cols):
        for j in range(m1_half_cols):
            mul_cols[i] += m2[j << 1][i] * m2[(j << 1) + 1][i]

    for i in range(m1_rows):
        for j in range(m2_cols):
            result[i][j] -= mul_rows[i] + mul_cols[j]
            for k in range(m1_half_cols):
                result[i][j] += (m1[i][k << 1] + m2[(k << 1) + 1][j]) \
                    * (m1[i][(k << 1) + 1] + m2[k << 1][j])
            if is_odd:
                result[i][j] += m1[i][m1_cols - 1] * m2[m1_cols - 1][j]

    return result
