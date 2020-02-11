package helper

import (
	"encoding/json"
	"testing"

	"github.com/spf13/cast"
)

func TestLookup(t *testing.T) {
	var f interface{}

	out, err := Lookup("../test-fixtures/hiera.yaml", "deep", "is_utc", "", map[string]interface{}{"service": "api", "environment": "live", "facts": "{timezone=>'CET'}"})
	if err != nil {
		t.Errorf("Error lookup: %s", err)
	}

	err = json.Unmarshal(out, &f)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %s", err)

	}

	v := cast.ToString(f)

	if v != "false" {
		t.Errorf("aws_cloudwatch_enable is %s; want %s", v, "true")
	}

	out2, err2 := Lookup("../doesnt_exists/hiera.yaml", "deep", "is_utc", "", map[string]interface{}{"service": "api", "environment": "live", "facts": "{timezone=>'CET'}"})
	if err2 == nil {
		t.Errorf("Error invalid config should not return: %s", out2)
	}
}