package storage

import "sync"

// Create a map where key is shortened URL and the value the redirect URL
var urlStore = make(map[string]string);
var mutex = &sync.Mutex{};

func Save(shortCode, url string) {
	mutex.Lock();
	defer mutex.Unlock();

	urlStore[shortCode] = url;
}

func Get(shortCode string) (string, bool) {
	mutex.Lock();
	defer mutex.Unlock();

	url, exists := urlStore[shortCode];

	return url, exists;
}