package serv

type servOption struct {
	host     string
	port     int16
	isSecure bool
}

type OptFn func(opt *servOption)

func defaultOpt() *servOption {
	return &servOption{
		host:     "0.0.0.0",
		port:     8012,
		isSecure: false,
	}
}

func Host(host string) OptFn {
	return func(opt *servOption) {
		opt.host = host
	}
}

func Port(port int16) OptFn {
	return func(opt *servOption) {
		opt.port = port
	}
}

func Secure(isSecure bool) OptFn {
	return func(opt *servOption) {
		opt.isSecure = isSecure
	}
}

func RunServ(opts ...OptFn) {
	opt := defaultOpt()
	for _, optFn := range opts {
		optFn(opt)
	}
	// TODO
}
