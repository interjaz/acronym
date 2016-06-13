package main

import "testing"

func TestEnglishWikiParser_SingleEntry(t *testing.T) {

	// Aggregate
	wikiParser := NewEnglishWikiParser()
	content := "<ul><li>...<a href=\"Url\" title=\"Acronym Title\">Acronym</a>Description</li></ul>"

	// Act
	acronyms, err := wikiParser.Parse(content)

	// Assert
	if err != nil {
		t.Error("Parse should succeed - should not contain any errors")
	}

	if acronyms == nil {
		t.Error("Parse should succeed - should have at least one acronym")
	}

	if len(acronyms) != 1 {
		t.Errorf("Parse should succeed - should have exactly one acronym, has: %d", len(acronyms))
	}

	if acronyms[0].Url != "https://en.wikipedia.orgUrl" {
		t.Errorf("Wrong Url field")
	}

	if acronyms[0].Acronym != "Acronym" {
		t.Errorf("Wrong Acronym field")
	}

	if acronyms[0].Description != "Description" {
		t.Errorf("Wrong Decription field")
	}

}

func TestEnglishWikiParser_RealSimpleEntry(t *testing.T) {

	// Aggregate
	wikiParser := NewEnglishWikiParser()
	content :=
		`
		<li><a href="/wiki/BA_(disambiguation)" class="mw-redirect mw-disambig" title="BA (disambiguation)">BA</a>&nbsp;– (i) <a href="/wiki/Bachelor_of_Arts" title="Bachelor of Arts">Bachelor of Arts</a>&nbsp;– (s) <a href="/wiki/Bahrain" title="Bahrain">Bahrain</a> (FIPS 10-4 country code)&nbsp;– <a href="/wiki/Bosnia_and_Herzegovina" title="Bosnia and Herzegovina">Bosnia and Herzegovina</a> (ISO 3166 digram)&nbsp;– (i) <a href="/wiki/British_Airways" title="British Airways">British Airways</a>&nbsp;– <a href="/wiki/British_Army" title="British Army">British Army</a>&nbsp;– <a href="/wiki/Bank_of_America" title="Bank of America">Bank of America</a></li>
		`

	// Act
	acronyms, err := wikiParser.Parse(content)

	// Assert
	if err != nil {
		t.Error("Parse should succeed - should not contain any errors")
	}

	if acronyms == nil {
		t.Error("Parse should succeed - should have at least one acronym")
	}

	if len(acronyms) != 1 {
		t.Errorf("Parse should succeed - should have exactly one acronym, has: %d.Found:\n%v", len(acronyms), acronyms)
	}

	if acronyms[0].Url != "https://en.wikipedia.org/wiki/BA_(disambiguation)" {
		t.Errorf("Wrong Url field")
	}

	if acronyms[0].Acronym != "BA" {
		t.Errorf("Wrong Acronym field")
	}

	if acronyms[0].Description != " – (i) Bachelor of Arts – (s) Bahrain (FIPS 10-4 country code) – Bosnia and Herzegovina (ISO 3166 digram) – (i) British Airways – British Army – Bank of America" {
		t.Errorf("Wrong Decription field: %s", acronyms[0].Description)
	}
}

func TestEnglishWikiParser_RealComplexEntry(t *testing.T) {

	// Aggregate
	wikiParser := NewEnglishWikiParser()
	content :=
		`
<li><a href="/wiki/BBC_(disambiguation)" class="mw-disambig" title="BBC (disambiguation)">BBC</a>
	<ul>
		<li>(i) <a href="/w/index.php?title=Breeding_Bird_Census&amp;action=edit&amp;redlink=1" class="new" title="Breeding Bird Census (page does not exist)">Breeding Bird Census</a></li>
		<li><a href="/wiki/BBC" title="BBC">British Broadcasting Corporation</a></li>
		<li>(i) Bumper to Back of Cab (trucking measurement)</li>
	</ul>
</li>
`

	// Act
	acronyms, err := wikiParser.Parse(content)

	// Assert
	if err != nil {
		t.Error("Parse should succeed - should not contain any errors")
	}

	if acronyms == nil {
		t.Error("Parse should succeed - should have at least one acronym")
	}

	if len(acronyms) != 1 {
		t.Errorf("Parse should succeed - should have exactly three acronym, has: %d.Found:\n%v", len(acronyms), acronyms)
	}
}

func TestEnglishWikiParser_MultipleEntry(t *testing.T) {

	// Aggregate
	wikiParser := NewEnglishWikiParser()
	content := "<ul><li>...<a href=\"Url\" title=\"Acronym\">Description Inner</a>Description Outer</li><li>...<a href=\"Url\" title=\"Acronym\">Description Inner</a>Description Outer</li></ul>"

	// Act
	acronyms, err := wikiParser.Parse(content)

	// Assert
	if err != nil {
		t.Error("Parse should succeed - should not contain any errors")
	}

	if acronyms == nil {
		t.Error("Parse should succeed - should have at least one acronym")
	}

	if len(acronyms) != 2 {
		t.Errorf("Parse should succeed - should have exactly one acronym, has: %d", len(acronyms))
	}

}

func TestEnglishFetcher_UpdateAll(t *testing.T) {

	// Aggregate
	lettersInAlphabet := 26 + 1
	reader := newFakeReader()
	parser := newFakeParser()
	parser.acronymsToCreate = 5
	acronymRepository := newFakeAcronymRepository()
	englishFetcher := NewEnglishFetcher(reader, parser, acronymRepository)

	// Act
	err := englishFetcher.UpdateAll()

	// Assert
	if err != nil {
		t.Error("UpdateAll should succeed")
	}

	if reader.readCalled != lettersInAlphabet {
		t.Errorf("Should call Read: %d, Actual: %d", lettersInAlphabet, reader.readCalled)
	}

	if parser.parseCalled != lettersInAlphabet {
		t.Errorf("Should call Parse: %d, Actual: %d", lettersInAlphabet, reader.readCalled)
	}

	if acronymRepository.openCalled != 1 &&
		acronymRepository.closeCalled != 1 {
		t.Error("Database left open")
	}

	if acronymRepository.insertCalled != lettersInAlphabet*parser.acronymsToCreate {
		t.Errorf("Not all acronyms have been inserted. Should: %d Actual: %d", lettersInAlphabet*parser.acronymsToCreate, acronymRepository.insertCalled)
	}

}

// -- Fakes

// Fake Acronym
type fakeAcronymRepository struct {
	openCalled      int
	closeCalled     int
	insertCalled    int
	deleteAllCalled int
}

func newFakeAcronymRepository() *fakeAcronymRepository {
	repository := new(fakeAcronymRepository)
	return repository
}

func (f *fakeAcronymRepository) Insert(a *Acronym) error {
	f.insertCalled++
	return nil
}

func (f *fakeAcronymRepository) Open() error {
	f.openCalled++
	return nil
}

func (f *fakeAcronymRepository) Close() error {
	f.closeCalled++
	return nil
}

func (f *fakeAcronymRepository) DeleteAll() error {
	f.deleteAllCalled++
	return nil
}

// Fake Reader
type fakeReader struct {
	readCalled int
}

func newFakeReader() *fakeReader {
	reader := new(fakeReader)
	return reader
}

func (f *fakeReader) Read(string) (string, error) {
	f.readCalled++
	return "", nil
}

// Fake Parser
type fakeParser struct {
	parseCalled      int
	acronymsToCreate int
}

func newFakeParser() *fakeParser {
	parser := new(fakeParser)
	return parser
}

func (f *fakeParser) Parse(string) ([]AcronymWikiResult, error) {
	f.parseCalled++
	acronyms := make([]AcronymWikiResult, 0)
	acronym := NewAcronymWikiResult()
	acronym.Acronym = "test"

	for i := 0; i < f.acronymsToCreate; i++ {
		acronyms = append(acronyms, *acronym)
	}

	return acronyms, nil
}
