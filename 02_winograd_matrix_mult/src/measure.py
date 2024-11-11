import matrix
import time 
import random

DEFAULT_NUM_OF_REPEATS = 5
DEFAULT_START = 50
DEFAULT_END = 301
DEFAULT_STEP = 50

def measure_time() -> None:
    algorithms = {
        "Standard Mult": matrix.standard_mult,
        "Winograd Mult": matrix.winograd_mult,
        "Optimized Winograd Mult": matrix.winograd_optimized_mult,
        "NIKITA": matrix.dotprod_winograd,
    }

    params = get_measurement_parameters()

    for size in range(params['start'], params['end'], params['step']):
        print("----------------------------------")
        print(f"SIZE: {size}")
        print("----------------------------------")
        for name, func in algorithms.items():
            elapsed_time = 0
            for _ in range(params['repeats']):
                m1 = generate_matrix(size)
                m2 = generate_matrix(size)

                start_time = time.process_time()
                func(m1, m2)
                end_time = time.process_time()
                elapsed_time += end_time - start_time

            print(f"{name}: {elapsed_time / params['repeats']:.4f} seconds")


def generate_matrix(size: int) -> list[list[int]]:
    return [[random.randint(0, 10) for _ in range(size)] for _ in range(size)]


def get_measurement_parameters() -> dict[str, int]:
    info = "(Press enter for default value)"
    return {
        "start": get_value(f"Enter start {info}: ", DEFAULT_START),
        "end": get_value(f"Enter end {info}: ", DEFAULT_END),
        "step": get_value(f"Enter step {info}: ", DEFAULT_STEP),
        "repeats": get_value(f"Enter number of repeats per one size {info}: ", DEFAULT_NUM_OF_REPEATS)
    }


def get_value_default(msg: str, default: int) -> int:
    val = None 
    try: 
        val = int(input(msg))
    except ValueError:
        val = default

    if val < 1:
        val = default
    
    return val
