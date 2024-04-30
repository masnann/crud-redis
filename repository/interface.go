package repository

type PingRepositoryInterface interface {
	Ping() (string, error)
}
