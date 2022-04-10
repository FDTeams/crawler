package parser

import (
	"regexp"
	"strconv"

	"github.com/seaung/crawler/pkg/engine"
	"github.com/seaung/crawler/pkg/models"
)

var (
	ageRe      = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d+])岁<div>`)
	heightRe   = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)cm</div>`)
	weightRe   = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)kg</div>`)
	incomeRe   = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([^<]+)</div>`)
	marriageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
	addressRe  = regexp.MustCompile(`<div class="m-btn purple" [^>]*>工作地:([^<]+)</div>`)
)

func ParseProfile(content []byte, name, gender string) engine.ParseResult {
	profile := models.Profile{}

	profile.Name = name
	profile.Gender = gender

	if age, err := strconv.Atoi(extractString(content, ageRe)); err != nil {
		profile.Age = age
	}

	if height, err := strconv.Atoi(extractString(content, heightRe)); err != nil {
		profile.Height = height
	}

	if weight, err := strconv.Atoi(extractString(content, weightRe)); err != nil {
		profile.Weight = weight
	}

	profile.Income = extractString(content, incomeRe)
	profile.Marriage = extractString(content, marriageRe)
	profile.Addresses = extractString(content, addressRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(content []byte, re *regexp.Regexp) string {
	subMatch := re.FindSubmatch(content)
	if len(subMatch) >= 2 {
		return string(subMatch[1])
	}
	return ""
}
