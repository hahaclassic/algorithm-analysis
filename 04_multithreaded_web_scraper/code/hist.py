import matplotlib.pyplot as plt

# Data
threads = [1, 2, 4, 8, 16, 32, 48]
pages_per_sec = [1.10, 2.22, 3.43, 5.10, 6.31, 7.64, 6.92]

# Plotting
plt.figure(figsize=(10, 6))
plt.bar(threads, pages_per_sec, color='skyblue', edgecolor='black', zorder=2) 
plt.xlabel('Количество потоков, шт')
plt.ylabel('Скорость загрузки страниц, кол-во страниц/сек')
plt.title('Зависимость скорости загрузки страниц от количества потоков')
plt.xticks(threads)
# plt.xscale("log", base=2)
# Add grid on the background
plt.grid(axis='both', color='gray', linestyle='--', linewidth=0.5, zorder=1)

# Display the plot
plt.show()
