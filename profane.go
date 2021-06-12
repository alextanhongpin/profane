package profane

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var vowels *regexp.Regexp
var nonconsonants *regexp.Regexp

//go:embed profane.txt
var raw string

func init() {
	// Regex to match non-consonants.
	nonconsonants = regexp.MustCompile("(?i)[^b-df-hj-np-tv-z]+")

	// Regex to match only vowels.
	vowels = regexp.MustCompile("(?i)[aeiou]+")
}

type Profane struct {
	words map[string]bool
	re    *regexp.Regexp
}

func New() *Profane {
	words := newSet(buildWords())
	return &Profane{
		words: words,
		re:    newExactMatch(newSlice(words)),
	}
}

// Has checks if the word is in the list.
func (p *Profane) Has(word string) bool {
	return p.words[word]
}

// Add adds a new words the profanity list.
func (p *Profane) Add(words ...string) {
	for _, word := range words {
		p.words[newWord(word)] = true
	}
	p.re = newExactMatch(newSlice(p.words))
}

// Remove remove a list of words from the profanity list.
func (p *Profane) Remove(words ...string) {
	for _, word := range words {
		delete(p.words, newWord(word))
	}
	p.re = newExactMatch(newSlice(p.words))
}

func (p *Profane) Match(word string) bool {
	word = normalize(word)
	return p.re.MatchString(word)
}

func (p *Profane) Regexp() *regexp.Regexp {
	r := *p.re
	return &r
}

// ReplaceGarbled replaces profane words with $@!#%.
func (p *Profane) ReplaceGarbled(s string) string {
	s = normalize(s)
	return p.re.ReplaceAllString(s, "$@!#%")
}

// ReplaceStars replaces profane words with '*' up to the word's length.
func (p *Profane) ReplaceStars(s string) string {
	s = normalize(s)
	return p.re.ReplaceAllStringFunc(s, func(s string) string {
		if len(s) < 3 {
			return strings.Repeat("*", len(s))
		}
		head := s[:1]
		body := strings.Repeat("*", len(s)-2)
		tail := s[len(s)-1:]
		return fmt.Sprintf("%s%s%s", head, body, tail)
	})
}

// ReplaceVowels replaces the vowels in the profane word with '*'.
func (p *Profane) ReplaceVowels(s string) string {
	s = normalize(s)
	return p.re.ReplaceAllStringFunc(s, func(s string) string {
		return vowels.ReplaceAllString(s, "*")
	})
}

// ReplaceCustom replaces the profane word with the custom string.
func (p *Profane) ReplaceCustom(s, t string) string {
	s = normalize(s)
	return p.re.ReplaceAllString(s, t)
}

// Say we have a word "hell", and "hello", and we test against "hello", it will match against "hell" first.
func sortByLen(s []string) []string {
	result := make([]string, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) > len(result[j])
	})
	return result
}

func newSet(in []string) map[string]bool {
	result := make(map[string]bool)
	for _, each := range in {
		// Skip empty string.
		if len(each) == 0 {
			continue
		}
		result[each] = true
	}
	return result
}

func newSlice(in map[string]bool) []string {
	var result []string
	for each := range in {
		result = append(result, each)
	}
	return result
}

func newExactMatch(in []string) *regexp.Regexp {
	// Sort in descending len.
	// Always attempt to match the longest words first.
	// If there's AA and A, match AA first.
	sorted := sortByLen(in)
	expr := fmt.Sprintf(`(?i)\b(%s)\b`, strings.Join(sorted, "|"))
	return regexp.MustCompile(expr)
}

var replacer = strings.NewReplacer(
	`0`, `o`,
	`1`, `i`,
	`2`, `z`,
	`3`, `e`,
	`4`, `a`,
	`5`, `s`,
	`6`, `b`,
	`7`, `t`,
	`8`, `b`,
	`9`, `g`,
	`@`, `a`,
	`.`, `\.`, // Exact punctiation.
)

func normalize(word string) string {
	return replacer.Replace(word)
}

func newWord(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	word = normalize(word)
	return word
}

func newWords(words []string) []string {
	result := make([]string, len(words))
	for i, word := range words {
		result[i] = newWord(word)
	}
	return result
}

func buildWords() []string {
	words := strings.Split(raw, "\n")
	return newWords(words)
}
