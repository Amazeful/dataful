package dataful

//ModelList defines an interface used for array of models.
//All model list classes must implement ModelList interface.
type ModelList interface {
	SetLoaded()
	GetList() interface{}
	Loaded() bool
}
