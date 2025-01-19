package main

func ExecuteHandler(flagUsed Flag, attr *FlagAttr) error {
	handlerMap := make(map[Flag]func() error)

	f, ok := handlerMap[flagUsed]

	if !ok {
		return ErrWrongArgument
	}

	return f()

}
