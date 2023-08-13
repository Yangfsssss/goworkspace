package parser

import (
	"engine"
	"model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄:</span>([\d]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高:</span>([\d]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入:</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重:</span>([\d]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别:</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座:</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚姻状况:</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历:</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业: </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯:</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件:</span>([^<]+)</td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车:</span>([^<]+)</td>`)

func ParseProfile(contents []byte) engine.ParserResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extraString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extraString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extraString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Marriage = extraString(contents, marriageRe)
	profile.Income = extraString(contents, incomeRe)
	profile.Xinzuo = extraString(contents, xinzuoRe)
	profile.Education = extraString(contents, educationRe)
	profile.Occupation = extraString(contents, occupationRe)
	profile.Hokou = extraString(contents, hokouRe)
	profile.House = extraString(contents, houseRe)
	profile.Car = extraString(contents, carRe)
	profile.Gender = extraString(contents, genderRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

	return result
}

func extraString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
