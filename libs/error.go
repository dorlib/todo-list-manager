package libs

import "strings"

type AppError struct {
	Code   int    `json:"-"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

const DefaultErrorMsg = "unexpected error"

func (a *AppError) Error() string {
	sb := strings.Builder{}
	appendDash := false

	if a.Title != "" {
		sb.WriteString(a.Title)
		appendDash = true
	}

	if a.Detail != "" {
		if appendDash {
			sb.WriteString(" - ")
		}

		sb.WriteString(a.Detail)
		appendDash = true
	}

	return sb.String()
}
