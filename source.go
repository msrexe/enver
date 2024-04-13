package enver

type Source interface {
	// Type returns the type of the source.
	Type() string
	// Get returns the key-value pairs of the source.
	Get() (map[string]any, error)
}
