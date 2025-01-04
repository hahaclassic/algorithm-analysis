import pandas as pd
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches

df = pd.read_csv('log.csv')

tasks = ["Чтение", "Извлечение данных", "Сохранение", "Ожидание"]
colors = [
    (0.2, 0.7, 0.8), 
    (0.3, 0.45, 0.8),   
    (0.25, 0.65, 0.25),  
    (0.77, 0.77, 0.77)  
]

for col in df.columns[1:]:  
    df[col] = df[col] / 1000

fig, ax = plt.subplots(figsize=(10, 6))

start = 0

for i, row in df.iterrows():
    if i == 0:
        start = row['readFileContentQueued']

    ax.barh(row['TaskID'], row['readFileContentStarted'] - row['readFileContentQueued'], 
            left=row['readFileContentQueued'] - start, height=0.6, color=colors[3], zorder=2, edgecolor='black')
    ax.barh(row['TaskID'], row['readFileContentFinished'] - row['readFileContentStarted'], 
            left=row['readFileContentStarted'] - start, height=0.6, color=colors[0], zorder=2, edgecolor='black')
    
    ax.barh(row['TaskID'], row['parseRecipeStarted'] - row['parseRecipeQueued'], 
            left=row['parseRecipeQueued'] - start, height=0.6, color=colors[3], zorder=2, edgecolor='black')
    ax.barh(row['TaskID'], row['parseRecipeFinished'] - row['parseRecipeStarted'], 
            left=row['parseRecipeStarted'] - start, height=0.6, color=colors[1], zorder=2, edgecolor='black')
    
    ax.barh(row['TaskID'], row['saveRecipeStarted'] - row['saveRecipeQueued'], 
            left=row['saveRecipeQueued'] - start, height=0.6, color=colors[3], zorder=2, edgecolor='black')
    ax.barh(row['TaskID'], row['saveRecipeFinished'] - row['saveRecipeStarted'], 
            left=row['saveRecipeStarted'] - start, height=0.6, color=colors[2], zorder=2, edgecolor='black')

ax.set_xlabel('Время, мc', fontsize=12)
ax.set_ylabel('Идентификатор задачи (taskID)', fontsize=12)
ax.set_title('Диаграмма Ганта', fontsize=14)
plt.xticks(rotation=45)

handles = [
    mpatches.Patch(color=colors[0], label=tasks[0]),  # Чтение
    mpatches.Patch(color=colors[1], label=tasks[1]),  # Извлечение данных
    mpatches.Patch(color=colors[2], label=tasks[2]),  # Сохранение
    mpatches.Patch(color=colors[3], label=tasks[3]),  # Ожидание
]

ax.legend(handles=handles, loc='upper left', bbox_to_anchor=(1, 1))
# ax.legend(handles=handles, loc='lower right')

ax.grid(True, axis='x', linestyle='-', color='gray', alpha=0.7, zorder=0) 

plt.tight_layout()
plt.show()
