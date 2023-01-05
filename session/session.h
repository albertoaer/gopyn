void initialize();
void finalize();
void* mainSession(void);
void* newSession(void);
void useSession(void* session);
void endSession(void* session);
void runString(const char* code);
void runFile(const char* filename);