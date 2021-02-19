package usecase

type Button struct {
    Command command
}

func (b *Button) Press() {
    b.Command.Execute()
}