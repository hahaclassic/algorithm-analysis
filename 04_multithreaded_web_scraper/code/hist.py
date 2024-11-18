import matplotlib.pyplot as plt

# Data
threads = [1, 2, 4, 8, 16, 32, 48]
pages_per_sec = [1.10, 1.50, 3.00, 4.50, 7.65, 12.15, 9.95]

# Plotting
plt.figure(figsize=(10, 6))
plt.bar(threads, pages_per_sec, color='skyblue', edgecolor='black', width=0.7)  # Увеличиваем ширину столбцов
plt.xlabel('Number of Threads')
plt.ylabel('Pages per Second')
plt.title('Performance vs. Number of Threads')
plt.xticks(threads)

# Add grid on the background
plt.grid(axis='y', color='gray', linestyle='--', linewidth=0.5)

# Display the plot
plt.show()
