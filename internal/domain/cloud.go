package domain

type RemoteCmd string

type Cloud interface {
	Run(pool *Pool, cmd RemoteCmd) ([]byte, error)
}
