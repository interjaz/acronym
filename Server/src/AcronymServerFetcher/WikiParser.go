package main

import (
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

	aExpr := "<a[a-zA-Z0-9 =\"\\-]*href=\"([a-zA-Z0-9\\/_()+&%?=.;]*)\"[a-zA-Z0-9 =z\\-\" ;?=.]*title=\"([a-zA-Z0-9 ()+-\\/]*)\"[a-zA-Z0-9 =\"\\-]*>(.+)<\\/a>(.+)"

	aRegex, _ := regexp.Compile(aExpr)

	acronyms = make([]AcronymWikiResult, 0)

	for {
		data, isDone := readTag(tokenizer, "li")

		if isDone {
			return acronyms, nil
		}

		found := aRegex.FindStringSubmatch(data)

		if found != nil {

			acronym := NewAcronymWikiResult()
			acronym.Url = englishWikiUrl + toHumanText(found[1])
			acronym.Acronym = toHumanText(found[2])
			description := found[3] + " " + found[4]
			acronym.Description = toHumanText(description)

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

var commonStripTags = []*regexp.Regexp{
	commonTagBegin("a"), commonTagEnd("a"),
	// It is possible to have multiple nesting of <a> tags...
	commonTagBegin("a"), commonTagEnd("a"),
	commonTagBegin("b"), commonTagEnd("b"),
	commonTagBegin("i"), commonTagEnd("i"),
	commonTagBegin("u"), commonTagEnd("u"),
	commonTagBegin("strong"), commonTagEnd("strong"),
	commonTagBegin("li"), commonTagEnd("li"),
	commonTagBegin("ul"), commonTagEnd("ul"),
	commonTagBegin("sup"), commonTagEnd("sup"),
	toRegexp("\\r"), toRegexp("\\n"), toRegexp("\t"),
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
	r, _ := regexp.Compile("<\\/" + tag + ">")
	return r
}

func toHumanText(content string) string {

	for i := 1; i < len(commonStripTags); i++ {
		content = commonStripTags[i].ReplaceAllString(content, emptyString)
	}

	content = strings.TrimSpace(content)
	content = strings.Replace(content, "( ", "(", -1)
	content = strings.Replace(content, " )", ")", -1)
	content = strings.Replace(content, "  ", " ", -1)
	content = html.UnescapeString(content)

	return content
}
