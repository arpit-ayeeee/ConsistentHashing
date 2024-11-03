package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
	"os"
)

// StorageNode represents a storage node with a name and host address.
type StorageNode struct {
	Name string
	Host string
}

// fetchFile retrieves the content of a file from a storage node.
func (node *StorageNode) fetchFile(path string) (string, error) {
	url := fmt.Sprintf("https://%s:1231/%s", node.Host, path)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching file: %v", err)
	}
	defer resp.Body.Close()

	var body bytes.Buffer
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return body.String(), nil
}

// putFile uploads a file's content to a storage node.
func (node *StorageNode) putFile(path string) (string, error) {
	content, err := os.ReadFile(path) // Using os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	url := fmt.Sprintf("https://%s:1231/%s", node.Host, path)
	resp, err := http.Post(url, "application/text", bytes.NewBuffer(content))
	if err != nil {
		return "", fmt.Errorf("error posting file content: %v", err)
	}
	defer resp.Body.Close()

	var responseBody bytes.Buffer
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return responseBody.String(), nil
}

// storageNodes holds instances of the storage node objects.
var storageNodes = []StorageNode{
	{Name: "A", Host: ""},
	{Name: "B", Host: ""},
	{Name: "C", Host: ""},
	{Name: "D", Host: ""},
	{Name: "E", Host: ""},
}

// hashFn computes a hash for the key and returns an index for a storage node.
func hashFn(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % len(storageNodes)
}

// upload determines which storage node to store the file on and uploads it.
func upload(path string) (string, error) {
	index := hashFn(path)
	node := &storageNodes[index]
	return node.putFile(path)
}

// fetch determines which storage node the file is on and retrieves it.
func fetch(path string) (string, error) {
	index := hashFn(path)
	node := &storageNodes[index]
	return node.fetchFile(path)
}

func main() {
	// Simulate locating the files across storage nodes
	files := []string{"f1.txt", "f2.txt", "f3.txt", "f4.txt", "f5.txt"}
	for _, file := range files {
		nodeName := storageNodes[hashFn(file)].Name
		fmt.Printf("File %s resides on node %s\n", file, nodeName)
	}

	// Example: Upload and fetch a file
	filePath := "f1.txt"
	if response, err := upload(filePath); err != nil {
		log.Printf("Failed to upload %s: %v", filePath, err)
	} else {
		fmt.Printf("Upload response: %s\n", response)
	}

	if content, err := fetch(filePath); err != nil {
		log.Printf("Failed to fetch %s: %v", filePath, err)
	} else {
		fmt.Printf("Fetched content from %s: %s\n", filePath, content)
	}
}
