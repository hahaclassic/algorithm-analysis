def binarySearch(arr: list[int], elem: int) -> tuple[int, int]:
    idx, comparisons = -1, 0
    left, right = 0, len(arr) - 1

    while left <= right:
        comparisons += 1
        mid = (left + right) // 2            
        if arr[mid] == elem:
            idx = mid
            break
        elif arr[mid] < elem:
            left = mid + 1
        else:
            right = mid - 1

    return idx, comparisons
