package godx

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSkeleton(t *testing.T) {
	b, err := ioutil.ReadFile("./main/cmp.xml")
	if err != nil {
		t.Error("File cmp.xml not found.")
		return
	}

	x2go := New(b)

	bones := x2go.Skeleton()

	expect := map[string][]string{
		"":                    []string{"soap:Envelope"},
		"soap:Envelope":       []string{"soap:Header", "soap:Body", "xmlns:soap,attr", "xmlns:soap1,attr", "xmlns:xsd,attr"},
		"soap:Body":           []string{"soap1:executeBatch"},
		"soap1:executeBatch":  []string{"soap1:sessionID", "soap1:commands"},
		"soap1:commands":      []string{"xsd:audienceID", "xsd:audienceLevel", "xsd:debug", "xsd:eventParameters", "xsd:interactiveChannel", "xsd:methodIdentifier", "xsd:relyOnExistingSession"},
		"xsd:eventParameters": []string{"xsd:name", "xsd:valueAsString", "xsd:valueDataType", "xsd:valueAsNumeric"},
		"xsd:audienceID":      []string{"xsd:name", "xsd:valueAsString", "xsd:valueDataType"},
	}

	for k, v := range bones {
		if len(v) != len(expect[k]) {
			t.Error("Something went wrong.")
			t.Errorf("%# v:%# v\n!=\n%# v:%# v", k, v, k, expect[k])
		}
	}
	// arrange("", bones.(map[string][]string))
}

func TestIdentifyType(t *testing.T) {
	bones := map[string][]string{
		"":                    []string{"soap:Envelope"},
		"soap:Envelope":       []string{"soap:Header", "soap:Body", "xmlns:soap,attr", "xmlns:soap1,attr", "xmlns:xsd,attr"},
		"soap:Body":           []string{"soap1:executeBatch"},
		"soap1:executeBatch":  []string{"soap1:sessionID", "soap1:commands"},
		"soap1:commands":      []string{"xsd:audienceID", "xsd:audienceLevel", "xsd:debug", "xsd:eventParameters", "xsd:interactiveChannel", "xsd:methodIdentifier", "xsd:relyOnExistingSession"},
		"xsd:eventParameters": []string{"xsd:name", "xsd:valueAsString", "xsd:valueDataType", "xsd:valueAsNumeric"},
		"xsd:audienceID":      []string{"xsd:name", "xsd:valueAsString", "xsd:valueDataType"},
	}

	id := Identify(bones)

	expect := map[string]map[string]string{
		"":                map[string]string{"soap:Envelope": "Envelope"},
		"Envelope":        map[string]string{"soap:Header": "Header", "soap:Body": "Body", "xmlns:soap,attr": "string", "xmlns:soap1,attr": "string", "xmlns:xsd,attr": "string"},
		"Body":            map[string]string{"soap1:executeBatch": "ExecuteBatch"},
		"executeBatch":    map[string]string{"soap1:sessionID": "string", "soap1:commands": "Commands"},
		"commands":        map[string]string{"xsd:audienceID": "string", "xsd:audienceLevel": "string", "xsd:debug": "string", "xsd:eventParameters": "EventParameters", "xsd:interactiveChannel": "string", "xsd:methodIdentifier": "string", "xsd:relyOnExistingSession": "string"},
		"eventParameters": map[string]string{"xsd:name": "string", "xsd:valueAsString": "string", "xsd:valueDataType": "string", "xsd:valueAsNumeric": "string"},
		"audienceID":      map[string]string{"xsd:name": "string", "xsd:valueAsString": "string", "xsd:valueDataType": "string"},
	}

	for k, v := range expect {
		if len(v) != len(id[k]) {
			t.Error("It might be wrong.")
			t.Errorf("%# v:%# v\n!=\n%# v:%# v", k, v, k, id[k])
		}
	}
}

func arrange(key string, bones map[string][]string) {
	for i := range bones[key] {
		fmt.Println(key + ":" + bones[key][i])
		arrange(bones[key][i], bones)
	}
}
