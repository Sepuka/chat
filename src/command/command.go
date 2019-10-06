package command

type Executor interface {
	Exec() error
}

type Preceptable interface {
	Precept() string

}