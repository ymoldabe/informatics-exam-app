package pkg

import (
	"Zhantasov/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetTestFiles(filePaths []string) ([]models.Data, error) {

	var resData []models.Data

	for _, path := range filePaths {

		fileData, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("Ошибка чтения файла:", err)
			return nil, err
		}

		var questions models.Data

		// Разбор данных из JSON
		if err := json.Unmarshal(fileData, &questions); err != nil {
			fmt.Println("Ошибка разбора JSON:", err)
			return nil, err
		}
		resData = append(resData, questions)
	}
	return resData, nil

}
