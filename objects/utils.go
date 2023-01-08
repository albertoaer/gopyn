package objects

import "github.com/albertoaer/gopyn/common"

func ObjectsToRefs(objects ...PyObject) []common.PyObjectRef {
	refs := make([]common.PyObjectRef, len(objects))
	for i := range objects {
		refs[i] = objects[i].Ref()
	}
	return refs
}

func RefsToObjects(refs ...common.PyObjectRef) []PyObject {
	objects := make([]PyObject, len(refs))
	for i := range refs {
		objects[i] = Wrap(refs[i])
	}
	return objects
}
