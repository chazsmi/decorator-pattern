package client

import "net/http"

func Auth(token string) Decorator {
	return Header("Authoriszation", token)
}

func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}
