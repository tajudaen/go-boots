defer is use as to delay the execution of a statement till the end of the calling function before it returns.
for multiple defer statements; they are executed by Last In First Out(LIFO)

Panic is used when an application gets into a situation when it can't continue at all

Recover is used to recover from panics. It is only useful in deferred functions