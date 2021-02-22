package usecase

type Nginx struct {
	Application       *Application
	MaxAllowedRequest int
	RateLimiter       map[string]int
}

func NewNginxServer() *Nginx {
	return &Nginx{
		Application:       &Application{},
		MaxAllowedRequest: 2,
		RateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.CheckRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.Application.HandleRequest(url, method)
}

func (n *Nginx) CheckRateLimiting(url string) bool {
	if n.RateLimiter[url] == 0 {
		n.RateLimiter[url] = 1
	}
	if n.RateLimiter[url] > n.MaxAllowedRequest {
		return false
	}
	n.RateLimiter[url] = n.RateLimiter[url] + 1
	return true
}
