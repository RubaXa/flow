package flow

type Fn func() (err error)

func Go(queue ...Fn) (err error) {
	for _, fn := range queue {
		err = fn()

		if err != nil {
			break
		}
	}

	return
}


