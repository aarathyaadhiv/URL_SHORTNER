package interfaceRepo

type Repository interface {
	Save(url string, short string) error
	IsShortExist(short string)(bool,error)
	GetURL(short string) (string, error)
	UpdateCount(short string) error
	GetCount(short string) (int, error)
}
