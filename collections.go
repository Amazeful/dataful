package dataful

type Collection string

const (
	CollectionUser    Collection = "user"
	CollectionChannel Collection = "channel"
	CollectionCommand Collection = "command"
	CollectionFilters Collection = "filters"
	CollectionPurge   Collection = "purge"
	CollectionAlerts  Collection = "alerts"
)
