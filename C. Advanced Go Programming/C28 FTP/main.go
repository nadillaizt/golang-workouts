package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
)

func getStringEntryType(t ftp.EntryType) string {
	switch t {
	case ftp.EntryTypeFile:
		return "(file)"
	case ftp.EntryTypeFolder:
		return "(folder)"
	case ftp.EntryTypeLink:
		return "(link)"
	default:
		return ""
	}
}

func main() {
	const FTP_ADDR = "0.0.0.0:2121"
	const FTP_USERNAME = "admin"
	const FTP_PASSWORD = "123456"

	conn, err := ftp.Dial(FTP_ADDR)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = conn.Login(FTP_USERNAME, FTP_PASSWORD)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("======= PATH ./")

	entries, err := conn.List(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, entry := range entries {
		fmt.Println(" ->", entry.Name, getStringEntryType(entry.Type))
	}

	//----------------------------------------------------------------
	fmt.Println("======= PATH ./somefolder")

	err = conn.ChangeDir("./somefolder")
	if err != nil {
		log.Fatal(err.Error())
	}

	//----------------------------------------------------------------
	entries, err = conn.List(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, entry := range entries {
		fmt.Println(" ->", entry.Name, getStringEntryType(entry.Type))
	}

	//----------------------------------------------------------------
	err = conn.ChangeDirToParent()
	if err != nil {
		log.Fatal(err.Error())
	}

	//----------------------------------------------------------------
	fmt.Println("======= SHOW CONTENT OF FILES")

	fileTest1Path := "test1.txt"
	fileTest1, err := conn.Retr(fileTest1Path)
	if err != nil {
		log.Fatal(err.Error())
	}

	test1ContentInBytes, err := io.ReadAll(fileTest1)
	fileTest1.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(" ->", fileTest1Path, "->", string(test1ContentInBytes))

	fileTest2Path := "somefolder/test3.txt"
	fileTest2, err := conn.Retr(fileTest2Path)
	if err != nil {
		log.Fatal(err.Error())
	}

	test2ContentInBytes, err := io.ReadAll(fileTest2)
	fileTest2.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(" ->", fileTest2Path, "->", string(test2ContentInBytes))

	//----------------------------------------------------------------
	fmt.Println("======= DOWNLOAD FILE TO LOCAL")

	fileMoviePath := "movie.mp4"
	fileMovie, err := conn.Retr(fileMoviePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	destinationMoviePath := "/Users/novalagung/Desktop/downloaded-movie.mp4"
	f, err := os.Create(destinationMoviePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = io.Copy(f, fileMovie)
	if err != nil {
		log.Fatal(err.Error())
	}
	fileMovie.Close()
	f.Close()

	fmt.Println("File downloaded to", destinationMoviePath)

	//----------------------------------------------------------------
	fmt.Println("======= UPLOAD FILE TO FTP SERVER")

	sourceFile := "/Users/novalagung/Desktop/Go-Logo_Aqua.png"
	f, err = os.Open(sourceFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	destinationFile := "./somefolder/logo.png"
	err = conn.Stor(destinationFile, f)
	if err != nil {
		log.Fatal(err.Error())
	}
	f.Close()

	fmt.Println("File", sourceFile, "uploaded to", destinationFile)

	entries, err = conn.List("./somefolder")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, entry := range entries {
		fmt.Println(" ->", entry.Name, getStringEntryType(entry.Type))
	}
}
