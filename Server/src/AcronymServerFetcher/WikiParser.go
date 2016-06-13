package main

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// WikiParser is an interface used by WikiParsers
type WikiParser interface {
	Parse(content string) (acronyms []AcronymWikiResult, err error)
}

// AcronymWikiResult structure for represending data returned by the WikiParser
type AcronymWikiResult struct {
	Acronym     string
	Description string
	Url         string
}

// NewAcronymWikiResult return new AcronymWikiResult
func NewAcronymWikiResult() *AcronymWikiResult {
	acronym := new(AcronymWikiResult)
	return acronym
}

// EnglishWikiParser is WikiParser suitable for parsing english wikipedia entries
type EnglishWikiParser struct {
}

// NewEnglishWikiParser returns new EnglishWikiParser
func NewEnglishWikiParser() *EnglishWikiParser {
	parser := new(EnglishWikiParser)
	return parser
}

// Parse func parses the english wikipedia's html body
func (proc *EnglishWikiParser) Parse(content string) (acronyms []AcronymWikiResult, err error) {

	contentReader := strings.NewReader(content)
	tokenizer := html.NewTokenizer(contentReader)

	aSimpleExpr := `<li>[^<]*<a href="([^"]+)"[^>]*>([^<]+)<\/a>(.+)<\/li>`
	aComplexExpr := `<li>[^<]*<a href="([^"]+)"[^>]*>([\S ]+)<\/a>([\s\S]+)<\/ul>`

	aSimpleRegex, err := regexp.Compile(aSimpleExpr)
	if err != nil {
		panic(err)
	}

	aComplexRegex, err := regexp.Compile(aComplexExpr)
	if err != nil {
		panic(err)
	}

	acronyms = make([]AcronymWikiResult, 0)
	for {
		data, isDone := readTag(tokenizer, "li")

		if isDone {
			return acronyms, nil
		}

		found := aComplexRegex.FindStringSubmatch(data)
		if found != nil {
			acronym := NewAcronymWikiResult()
			acronym.Url = englishWikiUrl + toHumanText(found[1])
			acronym.Acronym = toHumanText(found[2])
			acronym.Description = toHumanText(found[3])

			acronyms = append(acronyms, *acronym)
			continue
		}

		found = aSimpleRegex.FindStringSubmatch(data)
		if found != nil {
			acronym := NewAcronymWikiResult()
			acronym.Url = englishWikiUrl + toHumanText(found[1])
			acronym.Acronym = toHumanText(found[2])
			acronym.Description = toHumanText(found[3])

			acronyms = append(acronyms, *acronym)
		}

	}
}

func readTag(tokenizer *html.Tokenizer, tagName string) (data string, isFinished bool) {
	isFound := false
	depth := 0
	raw := ""

	for {
		tokenType := tokenizer.Next()
		tokenNameBytes, _ := tokenizer.TagName()
		tokenName := string(tokenNameBytes)

		if tokenizer.Err() != nil {
			return "", true
		}

		switch tokenType {
		case html.StartTagToken:
			if isFound {
				depth++
				raw += string(tokenizer.Raw())
			}
			if !isFound && tagName == tokenName {
				isFound = true
			}
		case html.EndTagToken:
			if depth == 0 {
				raw = fmt.Sprintf("<%s>%s</%s>", tagName, raw, tagName)
				return raw, false
			}

			if isFound {
				raw += string(tokenizer.Raw())
				depth--
			}
		case html.TextToken:
			if isFound {
				raw += string(tokenizer.Raw())
			}
		}
	}

}

var commonStripTagsEnumeration = []*regexp.Regexp{
	commonTagBegin("li"),
}
var commonStripTags = []*regexp.Regexp{
	commonTagEnd("li"),
	commonTagBegin("a"), commonTagEnd("a"),
	commonTagBegin("b"), commonTagEnd("b"),
	commonTagBegin("i"), commonTagEnd("i"),
	commonTagBegin("u"), commonTagEnd("u"),
	commonTagBegin("strong"), commonTagEnd("strong"),
	commonTagBegin("sup"), commonTagEnd("sup"),
	commonTagBegin("ul"), commonTagEnd("ul"),
	toRegexp("\\r"), toRegexp("\\t"),
}

var englishWikiUrl = "https://en.wikipedia.org"

func toRegexp(str string) *regexp.Regexp {
	r, _ := regexp.Compile(str)
	return r
}

func commonTagBegin(tag string) *regexp.Regexp {
	r, _ := regexp.Compile("<" + tag + "[^>]*>")
	return r
}

func commonTagEnd(tag string) *regexp.Regexp {
	r, _ := regexp.Compile("</" + tag + ">")
	return r
}

func toHumanText(content string) string {

	// Wikipedia has quite inconsistent formatting
	// lets try make best out of it

	for i := 0; i < len(commonStripTagsEnumeration); i++ {
		content = commonStripTagsEnumeration[i].ReplaceAllString(content, "\n– ")
	}

	for i := 0; i < len(commonStripTags); i++ {
		content = commonStripTags[i].ReplaceAllString(content, emptyString)
	}

	content = strings.TrimSpace(content)
	content = strings.Replace(content, "\n\n", "\n", -1)
	content = strings.Replace(content, "( ", "(", -1)
	content = strings.Replace(content, " )", ")", -1)
	content = strings.Replace(content, " – ", "\n– ", -1)
	content = strings.Replace(content, "  ", " ", -1)
	content = html.UnescapeString(content)

	content = strings.Replace(content, "  ", " ", -1)
	content = strings.Replace(content, "  ", " ", -1)

	return content
}
