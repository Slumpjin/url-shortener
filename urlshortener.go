package urlshortener

import "strconv"

func shortenUrlPath(path string) (int64, error) {
	id, err := strconv.ParseInt(path, 36, 0)
	if err == nil {
		return id, nil
	}
	return 0, err
}
