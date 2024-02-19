package interfaceUsecase

type Usecase interface {
	Shorten(url string) (string, error)
	Redirect(shorten string) (string, error)
	Count(shorten string) (int, error)
}
