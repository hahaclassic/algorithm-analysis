import levenshtein as lvnst

str1 = input("Введите первую строку: ")
str2 = input("Введите вторую строку: ")

print("Левенштейн рекурсивный: ", lvnst.RecursiveLevenshtein(str1, str2))
print("Левенштейн рекурсивный с мемоизацией: ", lvnst.RecursiveCacheLevenshtein(str1, str2))
print("Левенштейн динамический: ", lvnst.DynamicLevenshtein(str1, str2))
print("Дамерау-Левенштейн динамический: ", lvnst.DynamicDamerauLevenshtein(str1, str2))