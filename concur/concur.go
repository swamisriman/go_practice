package concur

type result struct {
	string
	bool
}

func CheckWebsites(websiteChecker func(string) bool, websites []string) map[string]bool {
	websiteResult := make(map[string]bool)
	resultChannel := make(chan result)

	for _, website := range websites {
		go func(url string) {
			resultChannel <- result{url, websiteChecker(url)}
		}(website)
	}

	for i := 0; i < len(websites); i++ {
		resultInstance := <-resultChannel
		websiteResult[resultInstance.string] = resultInstance.bool
	}

	return websiteResult
}
