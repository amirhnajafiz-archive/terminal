package command

type Command struct {
	Use    string `json:"use"`
	Action func(args []string) error
}
