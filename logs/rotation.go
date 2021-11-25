package logs

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func zipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

func rotation(logDirectory string) {
	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		panic(err)
	}

	startDate := time.Now().In(loc).Format("02-01-2006")

	for {
		dateNow := time.Now().In(loc).Format("02-01-2006")

		if startDate == dateNow {
			time.Sleep(10 * time.Second)
			continue
		}

		logFiles, _ := ioutil.ReadDir(logDirectory)

		var logFilesName []string

		for _, f := range logFiles {
			logFilesName = append(logFilesName, fmt.Sprintf("%s/%s", logDirectory, f.Name()))
		}

		if len(logFilesName) > 0 {
			err := zipFiles(fmt.Sprintf("%s-%s.zip", logDirectory, startDate), logFilesName); if err != nil {
				panic(err)
			}

			for _, f := range logFilesName {
				_ = os.Remove(f)
			}

			fmt.Println(fmt.Sprintf("Comppress Log %s Success....", startDate))
		}

		startDate = dateNow
		time.Sleep(5 * time.Second)
	}
}