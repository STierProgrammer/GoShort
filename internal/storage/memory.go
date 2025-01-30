package storage

import (
	"encoding/json"
	"os"
	"sync"
)

var filePath string;

var mutex = &sync.Mutex{};

var urlStore = make(map[string]string);

func LoadData(path string) error {
	mutex.Lock();
	defer mutex.Unlock();

	filePath = path;

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644);

	if err != nil {
		return err;
	}

	defer file.Close();

	// Decode the JSON data into the map
	decoder := json.NewDecoder(file);
	err = decoder.Decode(&urlStore);

	if err != nil && err.Error() != "EOF" {
		return err;
	}

	return nil;
}

func Save(shortCode, url string) error {
	mutex.Lock();
	defer mutex.Unlock();

	urlStore[shortCode] = url;

	return saveData();
}

func Get(shortCode string) (string, bool) {
	mutex.Lock();
	defer mutex.Unlock();

	url, exists := urlStore[shortCode];
	
	return url, exists;
}

func saveData() error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644);
	
	if err != nil {
		return err;
	}
	
	defer file.Close();

	encoder := json.NewEncoder(file);
	return encoder.Encode(urlStore);
}
