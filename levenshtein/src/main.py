import levenshtein as lvnst
from enum import Enum
import typing

class Operation(Enum):
    RECURSIVE_LEVENSHTEIN = 1
    RECURSIVE_CACHE_LEVENSHTEIN = 2
    DYNAMIC_LEVENSHTEIN = 3
    DYNAMIC_DAMERAU_LEVENSHTEIN = 4
    EXIT = 0

def Menu():
    print("----------------------------------")
    print("1. Recursive Levenshtein")
    print("2. Recursive Cache Levenshtein")
    print("3. Dynamic Levenshtein")
    print("4. Dynamic Damerau Levenshtein")
    print("0. Exit")
    print("----------------------------------")

def Start():
    Menu()
    inputMsg = "Enter the algorithm number or 0 to exit the program: "
    operation = Operation(int(input(inputMsg)))

    algorithm: dict[Operation, typing.Callable[[str, str], int]] = {
        Operation.RECURSIVE_LEVENSHTEIN: lvnst.RecursiveLevenshtein,
        Operation.RECURSIVE_CACHE_LEVENSHTEIN: lvnst.RecursiveCacheLevenshtein,
        Operation.DYNAMIC_LEVENSHTEIN: lvnst.DynamicLevenshtein,
        Operation.DYNAMIC_DAMERAU_LEVENSHTEIN: lvnst.DynamicDamerauLevenshtein
    }

    while operation != Operation.EXIT:
        if Operation.RECURSIVE_LEVENSHTEIN.value <= operation.value \
            <= Operation.DYNAMIC_DAMERAU_LEVENSHTEIN.value:
        
            s1 = input("Enter first string: ")
            s2 = input("Enter second string: ")
        
            result = algorithm[operation](s1, s2)
            print("\nResult: ", result)
        else:
            print("Wrong number. Try again.")

        Menu()
        operation = Operation(int(input(inputMsg)))

if __name__ == "__main__":
    Start()