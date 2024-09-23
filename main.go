package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Открытие файла для чтения gh
	file, err := os.Open("students.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
	}
	defer file.Close()

	// Словарь для хранения оценок студентов
	students := make(map[string][]int)

	// Чтение файла построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue // Игнорирование строк, содержащие кол-во данных в строке не равным двум
		}
		name := parts[0]

		// Преобразование оценки из строки в число
		score, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Ошибка при преобразовании оценки:", err)
			continue
		}
		if score < 0 {
			continue // Игнорирование строк с отрицательными оценками
		}

		// Добавление оценки в срез для данного студента
		students[name] = append(students[name], score)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Сортировка студентов по имени
	names := make([]string, 0, len(students))
	for name := range students {
		names = append(names, name)
	}
	sort.Strings(names)

	// Вывод результатов
	for _, name := range names {
		scores := students[name]
		fmt.Printf("%s \n", name)
		fmt.Printf("Scores: ")
		for _, score := range scores {
			fmt.Printf("%d, ", score)
		}
		// Рассчет средней оценки для каждого студента
		var total float32 = 0
		for _, score := range scores {
			total += float32(score)
		}
		average := total / float32(len(scores))
		fmt.Printf("\nСредняя оценка: %.2f\n\n", average)
	}
}
