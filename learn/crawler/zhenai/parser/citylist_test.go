package parser

import (
	"os"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := os.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	exceptedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}

	exceptedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result.Requests should have %d elements, but got %d", resultSize, len(result.Requests))
	}

	for i, url := range exceptedUrls {
		if url != result.Requests[i].Url {
			t.Errorf("exceptedUrls[%d] is %s, but got %s", i, exceptedUrls[i], result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result.Items should have %d elements, but got %d", resultSize, len(result.Items))
	}

	for i, city := range exceptedCities {
		if city != result.Items[i] {
			t.Errorf("exceptedCities[%d] is %s, but got %s", i, exceptedCities[i], result.Items[i])
		}
	}
}
