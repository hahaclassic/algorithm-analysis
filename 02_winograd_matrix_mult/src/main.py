from enum import Enum
import typing
import matrix
import measure
from reader import get_value

class Operation(Enum):
    MULT = 1
    MEASURE = 2
    EXIT = 0


def menu():
    print("----------------------------------")
    print("1. Single multiplication")
    print("2. Time measurement")
    print("0. Exit")
    print("----------------------------------")


def get_operation() -> Operation:
    menu()
    inputMsg = "Enter the option number or 0 to exit the program: "
    operation = None

    while operation is None:
        try:
            operation = Operation(int(input(inputMsg)))
        except ValueError:
            print("\nInvalid data. Enter number from 0 to 2.")

    return operation


def start():
    operation = get_operation()

    functions: dict[Operation, typing.Callable[[None], None]] = {
        Operation.MULT: single_matrix_multiplication,
        Operation.MEASURE: measure.measure_time
    }

    while operation != Operation.EXIT:
        functions[operation]()
        operation = get_operation()


def single_matrix_multiplication():
    m1_rows = get_value("Enter number of rows of the 1st matrix: ", check=lambda x: x > 0)
    m1_cols = get_value("Enter number of columns of the 1st matrix: ", check=lambda x: x > 0)
    m2_rows = get_value("Enter number of rows of the 2nd matrix: ", check=lambda x: x > 0)
    m2_cols = get_value("Enter number of columns of the 2nd matrix: ", check=lambda x: x > 0)

    print("\nEnter elements of the 1st matrix:")
    m1 = [[get_value(f"Elem [{i}][{j}]: ") for j in range(m1_cols)] for i in range(m1_rows)]

    print("\nEnter elements of the 2nd matrix:")
    m2 = [[get_value(f"Elem [{i}][{j}]: ") for j in range(m2_cols)] for i in range(m2_rows)]

    algorithms = {
        "Standard": matrix.standard_mult,
        "Winograd": matrix.winograd_mult,
        "Winograd Optimized": matrix.winograd_optimized_mult
    }

    for name in algorithms:
        print(f"\n{name}")
        result = algorithms[name](m1, m2)
        if result is None:
            print("Error. Invalid input")
            break

        for row in result:
            print(row)


if __name__ == "__main__":
    start()

