package usecase

type Server interface {
    HandleRequest(string, string) (int, string)
}