package flow

func Defer(fn func(Stream)) Stream {
	return DeferMany(1, fn)
}

func DeferMany(size int, fn func(Stream)) Stream {
	ch := make(Stream, size)
	go fn(ch)
	close(ch)
	return ch
}
