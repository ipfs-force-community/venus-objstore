package objstore

import (
	"fmt"
	"plugin"
)

const SymbolNameConstructor = "Constructor"

// Constructor
type Constructor = func(Config) (Store, error)

// loads the constructor from the given plugin
func LoadConstructor(path string) (Constructor, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open plugin %s: %w", path, err)
	}

	sym, err := plug.Lookup(SymbolNameConstructor)
	if err != nil {
		return nil, fmt.Errorf("lookup symbol %s: %w", SymbolNameConstructor, err)
	}

	constructor, ok := sym.(Constructor)
	if !ok {
		return nil, fmt.Errorf("unexpected constructor object of type %T", sym)
	}

	return constructor, nil
}
