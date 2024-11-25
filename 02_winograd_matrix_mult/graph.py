import matplotlib.pyplot as plt

# Данные
sizes = [21, 41, 61, 81, 101, 121, 141, 161, 181, 201, 221, 241, 261, 281, 301, 321, 341]
standard_mult = [0.0013, 0.0077, 0.0248, 0.0579, 0.1118, 0.1924, 0.3040, 0.4522, 0.6419, 0.8778, 
                 1.1668, 1.5149, 1.9474, 2.4148, 2.9777, 3.6667, 4.3503]
winograd_mult = [0.0013, 0.0090, 0.0283, 0.0660, 0.1268, 0.2174, 0.3433, 0.5086, 0.7194, 0.9856, 
                 1.3059, 1.6916, 2.1496, 2.7247, 3.3824, 4.1631, 5.0158]
optimized_winograd_mult = [0.0014, 0.0100, 0.0318, 0.0746, 0.1436, 0.2462, 0.3882, 0.5765, 0.8139, 1.1183, 
                           1.4779, 1.9216, 2.4351, 3.0461, 3.7617, 4.5926, 5.5051]

# Создание графика
plt.figure(figsize=(10, 6))
plt.plot(sizes, standard_mult, label='Стандартный алгоритм', marker='o', linestyle='-', color='b')
plt.plot(sizes, winograd_mult, label='Оптимизированный алгоритм Винограда', marker='s', linestyle='--', color='g')
plt.plot(sizes, optimized_winograd_mult, label='Алгоритм Винограда', marker='^', linestyle='-.', color='r')

# Настройки графика
plt.title('Зависимость процессорного времени работы ')
plt.xlabel('Размер матрицы')
plt.ylabel('Процессорное время, с')
plt.grid(True, linestyle='--', alpha=0.6)
plt.xticks(sizes, rotation=45)
plt.yticks([i/10 for i in range(0, 60, 5)])
plt.legend()
plt.tight_layout()

plt.show()
