# NOTE: MicroPython

import utime
import urandom
from src import levenshtein as lvnsht
from micropython import const

MEASURES_PER_LEN = const(15)
MIN_LEN = const(1)
MAX_LEN = const(6)

RECURSIVE_LEV = const("Recur Lev      ")
RECURSIVE_CACHE_LEV = const("Recur Cache Lev")
DYNAMIC_LEV = const("Dynamic Lev    ")
DYNAMIC_DAMERAU = const("Dynamic Damerau")


def randString(length: int) -> str:
    alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    return ''.join(urandom.choice(alphabet) for _ in range(length))


def measureTime(func, s1: str, s2: str) -> int:
    start = utime.ticks_ms()
    func(s1, s2)
    end = utime.ticks_ms()

    return utime.ticks_diff(end, start)


def runBenchmark():
    algorithms = {
        RECURSIVE_LEV: lvnsht.RecursiveLevenshtein,
        RECURSIVE_CACHE_LEV: lvnsht.RecursiveCacheLevenshtein,
        DYNAMIC_LEV: lvnsht.DynamicLevenshtein,
        DYNAMIC_DAMERAU: lvnsht.DynamicDamerauLevenshtein
    }

    for length in range(MIN_LEN, MAX_LEN + 1):
        times = {
            RECURSIVE_LEV: [],
            RECURSIVE_CACHE_LEV: [],
            DYNAMIC_LEV: [],
            DYNAMIC_DAMERAU: []
        }

        for _ in range(10): 
            s1, s2 = randString(length), randString(length)

            for algoType in algorithms:
                times[algoType].append(measureTime(algorithms[algoType], s1, s2))

        for algoType in algorithms:
            print("| ", length, " | ", algoType, " | time = ", sum(times[algoType]) / len(times[algoType]), " ms", sep="")
        print()

if __name__ == "__main__":
    runBenchmark()
