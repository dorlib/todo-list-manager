package cmd

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getExistingArgs(argsMap map[string]interface{}) map[string]string {
	args := make(map[string]string)

	for k, v := range argsMap {
		if v != nil && v != "" {
			args[k] = v.(string)
		}
	}

	return args
}
