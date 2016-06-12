package AcronymServerFetcher

func main() {

	reader := NewHttpClient()
	parser := NewEnglishWikiParser()
	repository := NewSqliteAcronymRepository("./AcronymDb.sqlite")
	englishFetcher := NewEnglishFetcher(reader, parser, repository)
	englishFetcher.UpdateAll()

}
