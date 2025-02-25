package utils

// source: https://platform.openai.com/docs/models
func CheckValidModel(model *string) bool {
	models := []string{
		"gpt-4o",
		"gpt-4o-2024-05-13",

		"gpt-4-turbo",
		"gpt-4-turbo-2024-04-09",
		"gpt-4-turbo-preview",
		"gpt-4-0125-preview",
		"gpt-4-1106-preview",
		"gpt-4",
		"gpt-4-0613",
		"gpt-4-0314", // legacy

		"gpt-3.5-turbo-0125",
		"gpt-3.5-turbo",
		"gpt-3.5-turbo-1106",
		"gpt-3.5-turbo-instruct",
	}
	for _, m := range models {
		if *model == m {
			return true
		}
	}
	return false
}
