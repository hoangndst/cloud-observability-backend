package config

type Configuration struct {
	Server struct {
		KeepAlivePeriodSeconds int
		ListenAddr             string `default:""`
		Port                   int    `default:"8080"`
		ResponseHeaders        map[string]string
		Cors                   struct {
			AllowOrigins []string
			AllowMethods []string
			AllowHeaders []string
		}

		TrustedProxies []string
		SSL            struct {
			Enabled         *bool  `default:"false"`
			RedirectToHTTPS *bool  `default:"true"`
			ListenAddr      string `default:""`
			Port            int    `default:"443"`
			CertFile        string `default:""`
			CertKey         string `default:""`
			LetsEncrypt     struct {
				Enabled   *bool  `default:"false"`
				AcceptTOS *bool  `default:"false"`
				Cache     string `default:"data/certs"`
				Hosts     []string
			}
		}
	}
	Database struct {
		Dialect    string `default:"sqlite3"`
		Connection string `default:"data/gotify.db"`
	}
	DefaultUser struct {
		Name string `default:"admin"`
		Pass string `default:"admin"`
	}
}
