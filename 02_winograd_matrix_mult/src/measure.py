import matrix
import time 
import random
from reader import get_value
from tabulate import tabulate

DEFAULT_NUM_OF_REPEATS = 5
DEFAULT_START = 50
DEFAULT_END = 300
DEFAULT_STEP = 50

def measure_time() -> None:
    algorithms = {
        "Standard Mult": matrix.standard_mult,
        "Winograd Mult": matrix.winograd_mult,
        "Optimized Winograd Mult": matrix.winograd_optimized_mult,
    }

    params = get_measurement_parameters()
    table = [["Size", "Standard Mult", "Winograd Mult", "Optimized Winograd Mult"]]

    for size in range(params['start'], params['end'] + 1, params['step']):
        print(f"size = {size}: measuring...")
        row = [size]
        for _, func in algorithms.items():
            elapsed_time = 0
            for _ in range(params['repeats']):
                m1 = generate_matrix(size)
                m2 = generate_matrix(size)

                start_time = time.process_time()
                func(m1, m2)
                end_time = time.process_time()
                elapsed_time += end_time - start_time

            avg_time = elapsed_time / params['repeats']
            row.append(f"{avg_time:.4f} seconds")

        print(f"size = {size}: done.")
        table.append(row)

    print(tabulate(table, headers="firstrow", tablefmt="grid"))


def generate_matrix(size: int) -> list[list[int]]:
    return [[random.randint(0, 10) for _ in range(size)] for _ in range(size)]


def get_measurement_parameters() -> dict[str, int]:
    info = "(Press enter for default value = {})"
    return {
        "start": get_value("Enter start", 
                           DEFAULT_START, lambda x: x > 0),
        "end": get_value("Enter end", 
                         DEFAULT_END, lambda x: x > 0),
        "step": get_value("Enter step", 
                          DEFAULT_STEP, lambda x: x > 0),
        "repeats": get_value("Enter number of repeats per one size", 
                             DEFAULT_NUM_OF_REPEATS, lambda x: x > 0)
    }
