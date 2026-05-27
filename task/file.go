package task

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func (manager *TaskManager) SaveToFile(path string) error {
	file, err := os.Create(path)

	if err != nil {
		return fmt.Errorf("Failed to create file: %v", path)
	}

	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	// First row
	csvWriter.Write([]string{"id", "contents", "created"})

	for _, task := range manager.tasks {
		csvWriter.Write([]string{task.id, task.contents, fmt.Sprint(task.created)})
	}

	return nil
}

func (manager *TaskManager) ReadFromFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("Failed to read file: %v", path)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		return fmt.Errorf("Failed to read csv file: %v", path)
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		id := record[0]
		contents := record[1]
		created, err := strconv.Atoi(record[2])

		if err != nil {
			return fmt.Errorf("Failed conversion from string to int")
		}

		task := NewTask(id, contents, int64(created))

		manager.AddTask(task)
	}

	return nil
}
