package command

type Executor interface {
	Exec()
}

type Preceptable interface {
	Precept() string

}