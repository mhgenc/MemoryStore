package config

import (
	"path"
)

var (
	API_PORT          = "8080"
	LOG_FILE          = "log.txt"
	STORE_FILE        = "-data.txt"
	SAVE_INTERVAL     = "10"
	DIR_SAVED_FILES   = "./savedfiles"
	DIR_ARCHIVE_FILES = path.Join(DIR_SAVED_FILES, "archive")
)

/*

// GetEnv Return env value if its exists, else return default value
func GetEnv(key string, configVariable *string) {
	val, isExists := os.LookupEnv(key)
	if isExists {
		*configVariable = val
	}
}

//ReadConfigFile reads config file and sets parameters
func ReadConfigFile() {
	i := 0
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i++
		arr := strings.Split(scanner.Text(), "=")

		if len(arr) == 2 && len(arr[0]) > 0 && len(arr[1]) > 0 {
			err := os.Setenv(strings.Trim(arr[0], " "), strings.Trim(arr[1], " "))
			if err != nil{
				logger.Error("ReadConfigFile - ", err)
			}
		} else {
			log.Println("Config could not be processed. Line ", i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

*/
