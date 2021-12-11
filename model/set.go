package model

// hash256,文件指纹，
//在不同的文件路径下，会有内容相同的文件，用来保证，相同文件只会被存储一次
//string > 文件hash值，
//fiarr > 文件信息
type HashStore map[string]FiArr

var HashStores HashStore

type FiArr map[FileDto]struct{}

func (a *FiArr) Add(finfo FileDto) bool {
	_, ok := (*a)[finfo]
	if !ok {
		(*a)[finfo] = struct{}{}
		return true
	}
	return false
}
