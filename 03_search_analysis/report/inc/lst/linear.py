def linearSearch(arr: list[int], elem: int) -> tuple[int, int]:
    idx, comparisons = -1, 0

    for i in range(len(arr)):
        comparisons += 1
        if arr[i] == elem:
            idx = i
            break
    
    return idx, comparisons