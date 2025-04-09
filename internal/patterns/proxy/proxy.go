package proxy

import "fmt"

// ProxyDatabase is a struct that acts as a proxy to the RealDatabase
type ProxyDatabase struct {
	realDatabase *RealDatabase
	lastQuery    string
	lastResult   string
	cacheEnabled bool
}

func NewProxyDatabase(realDB *RealDatabase) *ProxyDatabase {
	return &ProxyDatabase{
		realDatabase: realDB,
		cacheEnabled: true,
	}
}

// Query executes the given query, incorporating caching mechanisms
func (pdb *ProxyDatabase) Query(query string) (string, error) {
	// Use cache if it is enabled and the query is the same as the last one.
	if pdb.cacheEnabled && query == pdb.lastQuery && pdb.lastResult != "" {
		fmt.Println("Proxy: using cache.")

		return pdb.lastResult, nil
	}

	// Querying the RealDatabase
	result, err := pdb.realDatabase.Query(query)
	if err != nil {
		return "", err
	}

	// Cache the result
	pdb.lastQuery = query
	pdb.lastResult = result

	return result, nil
}
