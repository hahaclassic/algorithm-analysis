def measure_time():
    sizes = [100, 200, 300, 400]  # Размеры квадратных матриц для замера времени
    algorithms = {
        "Standard Multiplication": standard_mult,
        "Winograd Multiplication": winograd_mult,
        "Optimized Winograd Multiplication": winograd_optimized_mult
    }

    print("Запуск замера времени для различных реализаций умножения матриц:")
    for size in sizes:
        print(f"\nРазмер матриц: {size}x{size}")
        m1 = generate_matrix(size)
        m2 = generate_matrix(size)

        for name, func in algorithms.items():
            elapsed_time = 0
            count = 1
            for _ in range(count):
                start_time = time.process_time()
                func(m1, m2)
                end_time = time.process_time()
                elapsed_time += end_time - start_time

            print(f"{name}: {elapsed_time / count:.4f} секунд")


def main():
    parser = argparse.ArgumentParser(description="CLI для умножения матриц")
    parser.add_argument(
        "mode",
        choices=["single", "benchmark"],
        help="Режим работы: single - одиночное умножение, benchmark - замер времени"
    )
    args = parser.parse_args()

    if args.mode == "single":
        single_matrix_multiplication()
    elif args.mode == "benchmark":
        time_measurement()
def generate_matrix(size: int) -> list[list[int]]:
    """Генерирует случайную квадратную матрицу заданного размера."""
    return [[random.randint(0, 10) for _ in range(size)] for _ in range(size)]