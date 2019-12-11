package domain

const DefaultHostsLimit = 2

type ClientProperties struct {
	ClientId   uint64  `sql:",pk"`
	HostsLimit uint8   `pg:"default:2"`
	Client     *Client `sql:"fk:id,notnull"`
}
