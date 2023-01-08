void* getAttr(void*, const char*);

const char* toRepr(void*);
const char* toStr(void*);
const char* asStr(void*);
unsigned int asBool(void*, int*);
long asInt(void*, int*);
double asDouble(void*, int*);

void* listFromSlice(size_t, void**);
void* tupleFromSlice(size_t, void**);