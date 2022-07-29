package types

type Command struct {
	Use    string                              `json:"use"`
	Action func(args []string) (string, error) `json:"action"`
}
