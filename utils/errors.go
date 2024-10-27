package utils

import "fmt"

func TaskError(message string) error {
	return fmt.Errorf("Task Error: %s", message);
}