package main

type RWMutex struct {
	w Mutex
	writerSem uint32
	readerSem uint32
	readerCount int32
	readerWait int32
}

//从RWMutex的实现来看，RWMutex类型其实组合了Mutex