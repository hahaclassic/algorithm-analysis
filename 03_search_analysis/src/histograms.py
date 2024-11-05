import matplotlib.pyplot as plt

def createHistogram(comparisons: list[int], indexes: list[int], title: str) -> None:
    plt.bar(indexes, comparisons, zorder=2)

    plt.title(title)
    plt.xlabel("Индекс элемента")
    plt.ylabel("Количество сравнений")

    plt.grid(axis='y', zorder=1)
    plt.show()

