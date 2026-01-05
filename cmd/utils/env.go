package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetEnvString(name, defaultvalue string)string{
	value := os.Getenv("name")
	if strings.Trim(value, "") == ""{
		return defaultvalue
	}
	
	return value
}

func GetEnvInt(name string, defaultvalue int)int{
	value := os.Getenv("name")
	if strings.Trim(value, "") == ""{
		return defaultvalue
	}
	
	intValue,_ := strconv.Atoi(value)
	
	return intValue
}