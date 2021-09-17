/*
 * @Autor: Ruic
 * @since: 2021-09-13 14:41:23
 * @LastEditTime: 2021-09-15 09:27:44
 */
package server

// NewHTTPServer new a HTTP server.
/*
func NewHTTPServer(c *conf.Server, account *service.AccountService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	return srv
}
*/
