package templates

import "regexp"

func FormatURL(url string, urlTemplate string) string {
	var re = regexp.MustCompile(`({URL})`)
	s := re.ReplaceAllString(urlTemplate, url)

	return s
}
