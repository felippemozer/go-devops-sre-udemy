package main

import (
	"cryptoserver/utils"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	f, err := os.OpenFile("cryptoserver.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Starting crypto server")

	router := mux.NewRouter()
	router.HandleFunc("/encrypt", Encrypt).Methods("POST")
	router.HandleFunc("/decrypt", Decrypt).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", router))
}

func Encrypt(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		log.Println("Error on get file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	log.Println("Arquivo carregado com sucesso")

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error on read file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("Arquivo lido com sucesso")

	filename := uuid.NewString()
	tempFile, err := os.CreateTemp("./", filename)
	if err != nil {
		log.Println("Error on create temporary file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
	log.Println("Arquivo temporário criado com sucesso")

	tempFileName := uuid.NewString()
	tempFile.Write(fileBytes)
	utils.EncryptLargeFiles(tempFile.Name(), tempFileName, []byte(os.Getenv("KEY")))

	returnFileBytes, err := os.ReadFile(tempFileName)
	if err != nil {
		log.Println("Error on encrypt file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer os.Remove(tempFileName)
	log.Println("Arquivo criptografado com sucesso")

	binFile, err := os.Create("./tmp/" + filename + ".bin")
	if err != nil {
		log.Println("Error on create binary file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer binFile.Close()
	_, err = binFile.Write(returnFileBytes)
	if err != nil {
		log.Println("Error on save data to binary file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Binário gerado com sucesso")

	res.WriteHeader(http.StatusOK)
}

func Decrypt(res http.ResponseWriter, req *http.Request) {
	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		log.Println("Error on get file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	log.Println("Arquivo carregado com sucesso")

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error on read file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("Arquivo lido com sucesso")

	filename := uuid.NewString()
	tempFile, err := os.CreateTemp("./", filename)
	if err != nil {
		log.Println("Error on create temporary file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
	log.Println("Arquivo temporário criado com sucesso")

	tempFileName := uuid.NewString()
	tempFile.Write(fileBytes)
	utils.DecryptLargeFiles(tempFile.Name(), tempFileName, []byte(os.Getenv("KEY")))
	returnFileBytes, err := os.ReadFile(tempFileName)
	if err != nil {
		log.Println("Error on decrypt file: ", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer os.Remove(tempFileName)
	log.Println("Arquivo descriptografado com sucesso")
	res.WriteHeader(http.StatusOK)
	// res.Header().Set("Content-Type", "application/octet-stream")
	// res.Write(returnFileBytes)
	outFile, _ := os.Create("./tmp/" + strings.Split(fileHeader.Filename, ".")[0] + "-dec")
	outFile.Write(returnFileBytes)
}
