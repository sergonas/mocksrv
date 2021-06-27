package mocksrv

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Handler(config ConfigRoot) *gin.Engine {
	r := gin.Default()
	proxies := createProxies(config.Proxies)
	createRoots(r, config.Roots, proxies)

	return r
}

func createProxies(proxies []Proxy) map[string]*httputil.ReverseProxy {
	proxiesMap := make(map[string]*httputil.ReverseProxy)
	for _, proxy := range proxies {
		proxyUrl, err := url.Parse(proxy.Host)
		if err != nil {
			log.Fatal(err)
		}
		proxiesMap[proxy.Id] = httputil.NewSingleHostReverseProxy(proxyUrl)
	}

	return proxiesMap
}

func createRoots(r *gin.Engine, roots []Root, proxies map[string]*httputil.ReverseProxy) {
	for _, root := range roots {
		resp := root.Response
		r.Handle(root.Method, root.Path, func(c *gin.Context) {
			switch resp.Type {
			case ProxyType:
				proxy := proxies[resp.ProxyId]

				if resp.ProxyPath != "" {
					c.Request.URL.Path = resp.ProxyPath
				}
				if len(resp.Headers) > 0 {
					for k, v := range resp.Headers {
						c.Request.Header.Add(k, v)
					}
				}

				proxy.ServeHTTP(c.Writer, c.Request)
			case RespondType:
				for k, v := range root.Response.Headers {
					c.Header(k, v)
				}
				c.String(resp.Code, resp.Body)
			}
		})
	}
}
