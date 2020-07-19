package definition

// New constructor
func New(version string) Definition {
	return Definition{version}
}

// Definition def
type Definition struct {
	version string
}

// Version def
func (d Definition) Version() string {
	return d.version
}
