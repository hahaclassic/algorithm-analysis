import matplotlib.pyplot as plt
import numpy as np

lengths = np.array([1, 2, 3, 4, 5, 6])
levenshtein_matrix = np.array([0.100097, 0.176366, 0.340942, 0.512742, 0.741701, 1.012530])
damerau_levenshtein_matrix = np.array([0.097736, 0.204006, 0.394079, 0.631439, 0.947312, 1.341707])
levenshtein_recursive = np.array([0.095129, 0.399600, 1.860301, 9.813755, 52.949607, 275.854151])
levenshtein_recursive_cache = np.array([0.347810, 1.138200, 2.702778, 4.151481, 6.649858, 9.246883])

colors = {
    'levenshtein_matrix': '#1f77b4',  # Синий
    'damerau_levenshtein_matrix': '#ff7f0e',  # Оранжевый
    'levenshtein_recursive': '#2ca02c',  # Зеленый
    'levenshtein_recursive_cache': '#d62728'  # Красный
}

plt.figure(figsize=(12, 6))

plt.subplot(1, 2, 1)
plt.plot(lengths, levenshtein_matrix, 'o-', color=colors['levenshtein_matrix'], label='алг. поиска расст. (АПР) Левенштейна (матрица)', markersize=8, markerfacecolor='none')
plt.plot(lengths, damerau_levenshtein_matrix, 'x-', color=colors['damerau_levenshtein_matrix'], label='АПР Дамерау-Левенштейна (матрица)', markersize=8, markerfacecolor='none')
plt.plot(lengths, levenshtein_recursive, 's-', color=colors['levenshtein_recursive'], label='АПР Левенштейна (рекурсия)', markersize=8, markerfacecolor='none')
plt.plot(lengths, levenshtein_recursive_cache, 's-', color=colors['levenshtein_recursive_cache'], label='АПР Левенштейна (рекурсия с кэшем)', markersize=8, markerfacecolor='none')
plt.yscale('log')
plt.xlabel('Длина строки, кол-во символов')
plt.ylabel('Время, мс')
plt.title('Сравнение алгоритмов (логарифмическая шкала)')
plt.grid(True, which="both", linestyle='--', linewidth=0.5)
plt.legend()

plt.subplot(1, 2, 2)
plt.plot(lengths, levenshtein_matrix, 'o-', color=colors['levenshtein_matrix'], label='алг. поиска расст. (АПР) Левенштейна (матрица)', markersize=8, markerfacecolor='none')
plt.plot(lengths, damerau_levenshtein_matrix, 'x-', color=colors['damerau_levenshtein_matrix'], label='АПР Дамерау-Левенштейна (матрица)', markersize=8, markerfacecolor='none')
plt.plot(lengths, levenshtein_recursive, 's-', color=colors['levenshtein_recursive'], label='АПР Левенштейна (рекурсия)', markersize=8, markerfacecolor='none')
plt.plot(lengths, levenshtein_recursive_cache, 's-', color=colors['levenshtein_recursive_cache'], label='АПР Левенштейна (рекурсия с кэшем)', markersize=8, markerfacecolor='none')
plt.xlabel('Длина строки, кол-во символов')
plt.ylabel('Время, мс')
plt.title('Сравнение алгоритмов (обычная шкала)')
plt.grid(True, linestyle='--', linewidth=0.5)
plt.legend()

plt.tight_layout()
plt.show()
