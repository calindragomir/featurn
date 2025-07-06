package flags

type Rule struct {
	Attribute    string      `json:"attribute"`
	Operator     string      `json:"operator"`
	MatchValue   interface{} `json:"match_value"`
	ReturnValue  interface{} `json:"return_value"`
}

type Flag struct {
	Key          string `json:"key"`
	DefaultValue interface{} `json:"default"`
	Rules        []Rule `json:"rules"`
}

var testFlags = []Flag{
	{
		Key:          "new-homepage",
		DefaultValue: false,
		Rules: []Rule{
			{Attribute: "country", Operator: "eq", MatchValue: "US", ReturnValue: true},
		},
	},
}

func GetFlagByKey(key string) (Flag, bool) {
	for _, f := range testFlags {
		if f.Key == key {
			return f, true
		}
	}
	return Flag{}, false
}
