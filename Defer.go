package flow

func Defer(fn func(Stream)) Stream {
	return DeferMany(1, fn)
}

func DeferMany(size int, fn func(Stream)) Stream {
	ret := make(Stream, size)
	go fn(ret)
	return ret
}
