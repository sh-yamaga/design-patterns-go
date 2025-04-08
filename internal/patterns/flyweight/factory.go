package flyweight

// TreeFactory holds TreeTypes that are cached and reused.
type TreeFactory struct {
	treeType map[string]*TreeType
}

// NewTreeFactory initialize TreeFactory instance.
func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeType: make(map[string]*TreeType),
	}
}

// GetTreeType returns a TreeType that matches the given arguments from the cache;
// if it does not exist, it creates and caches a new TreeType and then returns it.
func (tf *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	key := name + "_" + color + "_" + texture
	if _, ok := tf.treeType[key]; !ok {
		tf.treeType[key] = &TreeType{
			Name:    name,
			Color:   color,
			Texture: texture,
		}
	}

	return tf.treeType[key]
}
