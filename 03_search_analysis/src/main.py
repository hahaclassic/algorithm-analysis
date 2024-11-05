import histograms as hist
import list_funcs as lf

LENGTH = 1020 # Task number: 8086
MIN_VAL, MAX_VAL = -12345, 12345
OUT_OF_RANGE_VALUE = -123456

def linearSearchHistogram(arr: list[int]) -> None:
    comparisons, indexes = [], []

    for x in arr:
        idx, comp = lf.linearSearch(arr, x)
        comparisons.append(comp)
        indexes.append(idx)
    
    idx, comp = lf.linearSearch(arr, OUT_OF_RANGE_VALUE)
    comparisons.append(comp)
    indexes.append(idx)

    hist.createHistogram(comparisons, indexes, "Линейный поиск")

def binarySearchHistograms(arr: list[int]) -> None:
    arr.sort()
    comparisons, indexes = [], []
    
    for x in arr:
        idx, comp = lf.binarySearch(arr, x)
        comparisons.append(comp)
        indexes.append(idx)
    
    idx, comp = lf.binarySearch(arr, OUT_OF_RANGE_VALUE)
    comparisons.append(comp)
    indexes.append(idx)

    hist.createHistogram(comparisons, indexes, "Бинарный поиск")

    comparisons.sort()

    hist.createHistogram(comparisons, indexes, 
                         "Отсортированный результат бинарного поиска")

if __name__ == '__main__':
    arr = lf.generateRandom(LENGTH, MIN_VAL, MAX_VAL)

    linearSearchHistogram(arr)

    binarySearchHistograms(arr)
