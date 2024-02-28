package plugin

type Plugin interface {
	Name() string
	Description() string
	Version() string

	Run(payload map[string]any) error
}
