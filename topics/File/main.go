package main

import (
	"fmt"
	"os"
)

func fileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// File does not exist
		return false, nil
	} else if err != nil {
		// An error occurred while checking existence
		return false, err
	}

	// File exists
	return true, nil
}

func main() {
	filePath := "/var/opt/mssql/dbm_certificate.cer"
	exists, err := fileExists(filePath)
	if err != nil {
		fmt.Printf("Error checking file existence: %v\n", err)
		return
	}

	if exists {
		fmt.Printf("File %s exists.\n", filePath)
	} else {
		fmt.Printf("File %s does not exist.\n", filePath)
	}
}
