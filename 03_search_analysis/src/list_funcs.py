import random

def linearSearch(arr: list[int], elem: int) -> tuple[int, int]:
    """
        Returns index of the elem and number of comparisons. 
        If the elem is not in the list, returns -1.
    """
    idx, comparisons = -1, 0

    for i in range(len(arr)):
        comparisons += 1
        if arr[i] == elem:
            idx = i
            break
    
    return idx, comparisons


def binarySearch(arr: list[int], elem: int) -> tuple[int, int]:
    """
        Returns index of the elem and number of comparisons. 
        If the elem is not in the list, returns -1.
    """
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


def generateRandom(length: int, min_val: int, max_val: int) -> list[int]:
    if min_val > max_val:
        min_val, max_val = max_val, min_val

    arr: list[int] = []
    while length > 0:
        val = random.randint(min_val, max_val)
        if val not in arr:
            arr.append(val)
            length -= 1

    return arr
