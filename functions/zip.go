package functions

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type ZipFiles struct {
	directory string
	zipName   string
	files     []string
}

func NewZipFiles(directory string) *ZipFiles {
	return &ZipFiles{
		directory: directory,
		zipName:   "trimestral.zip",
		files:     []string{},
	}
}

func (receiver *ZipFiles) appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(receiver.getPathDirectory(filename))

	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}

	defer file.Close()

	wr, err := zipw.Create(filename)

	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}

	return nil
}

func (receiver *ZipFiles) AddFiles(files []string) *ZipFiles {
	receiver.files = files
	return receiver
}

func (receiver *ZipFiles) AddFile(file string) *ZipFiles {
	receiver.files = append(receiver.files, file)
	return receiver
}

func (receiver *ZipFiles) GetFilesDirectory() (filesRs []string) {
	files, err := ioutil.ReadDir(receiver.directory)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filesRs = append(filesRs, file.Name())
	}

	return
}

func (receiver *ZipFiles) Generate() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC

	file, err := os.OpenFile(receiver.GetZip(), flags, 0644)

	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}

	defer file.Close()

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	for _, filename := range receiver.files {
		if err := receiver.appendFiles(filename, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", fmt.Sprintf("%s/%s", receiver.directory, filename), err)
		}
	}
}

func (receiver *ZipFiles) getPathDirectory(filename string) string {
	return fmt.Sprintf("%s/%s", receiver.directory, filename)
}

func (receiver *ZipFiles) GetZip() string {
	return fmt.Sprintf("%s/%s", receiver.directory, receiver.zipName)
}

func (receiver *ZipFiles) GetDirectory() string {
	return receiver.directory
}
