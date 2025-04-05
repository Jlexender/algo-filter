package internal

func handleCmd(command, argument string) {
	switch command {
	case "new":
		NewFile(argument)
	case "load":
		LoadFile(argument)
	}

}
