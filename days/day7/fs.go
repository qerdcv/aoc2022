package main

type fType uint8

const (
	file fType = iota
	dir
)

const (
	rootDir = "/"
	prevDir = ".."
	curDir  = "."
)

type fs struct {
	name    string
	t       fType
	content map[string]*fs
	size    uint
}

func newFs() *fs {
	f := &fs{
		name:    "/",
		t:       dir,
		content: make(map[string]*fs),
		size:    0,
	}
	f.content[prevDir] = f
	f.content[curDir] = f
	f.content[rootDir] = f
	return f
}

func (f *fs) addDir(name string) {
	newFs := &fs{
		name:    name,
		t:       dir,
		content: make(map[string]*fs),
		size:    0,
	}
	newFs.content[prevDir] = f
	newFs.content[curDir] = newFs
	f.content[name] = newFs
}

func (f *fs) addFile(name string, size uint) {
	f.content[name] = &fs{
		name: name,
		t:    file,
		size: size,
	}
	f.incrementSize(size)
}

func (f *fs) incrementSize(size uint) {
	currentFs := f

	for ; currentFs.name != rootDir; currentFs = currentFs.content[".."] {
		currentFs.size += size
	}

	currentFs.size += size
}

func isNameReserved(name string) bool {
	return name == rootDir || name == curDir || name == prevDir
}
