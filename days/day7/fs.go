package main

type elType uint8

const (
	elTypeFile elType = iota
	elTypeDir
)

type fs struct {
	name    string
	t       elType
	content map[string]*fs
	size    uint
}

func newFs() *fs {
	f := &fs{
		name:    "/",
		t:       elTypeDir,
		content: make(map[string]*fs),
		size:    0,
	}
	f.content[".."] = f
	f.content["."] = f
	f.content["/"] = f
	return f
}

func (f *fs) addDir(name string) {
	newFs := &fs{
		name:    name,
		t:       elTypeDir,
		content: make(map[string]*fs),
		size:    0,
	}
	newFs.content[".."] = f
	newFs.content["."] = newFs
	f.content[name] = newFs
}

func (f *fs) addFile(name string, size uint) {
	f.content[name] = &fs{
		name: name,
		t:    elTypeFile,
		size: size,
	}
	f.incrementSize(size)
}

func (f *fs) incrementSize(size uint) {
	currentFs := f

	for ; currentFs.name != "/"; currentFs = currentFs.content[".."] {
		currentFs.size += size
	}

	currentFs.size += size
}
