package expr

import "testing"

var exprExampleList = []struct {
	control map[string]interface{}
	expr    string
	result  bool
}{
	{
		map[string]interface{}{"pType": 66, "pid": 317},
		`pType > 36 && pid > 310`,
		true,
	},

	{
		map[string]interface{}{"pType": 66, "pid": 317},
		`pType < 90 && pid == 310`,
		false,
	},

	{
		map[string]interface{}{"version": "11.0.1", "pid": 317},
		`ver_compare(version, ">", "10.1.1")`,
		true,
	},

	{
		map[string]interface{}{"version": "11.0.1", "pid": 317},
		`ver_compare(version, "<", "10.1.1")`,
		false,
	},

	{
		map[string]interface{}{"pid": 317},
		`in_array(pid, []int{317, 318, 319})`,
		true,
	},

	{
		map[string]interface{}{"pid": "317"},
		`in_array(pid, []string{"317", "318", "319"})`,
		true,
	},

}

func TestNewEngine(t *testing.T) {
	for id, v := range exprExampleList {
		engine, err := NewEngine(v.expr)
		if err != nil {
			t.Error(err)
		}
		result, err := engine.RunRule(v.control)
		if err != nil {
			t.Error(err)
		}
		if v.result != result {
			t.Errorf("id %d failed test %v , result = %v", id, v, result)
		}
	}
}
