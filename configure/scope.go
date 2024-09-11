package configure

type Scope struct {
	HTTP    HTTPscope
	General Generalscope
	Pentest Penetrationscope
}

type Generalscope struct {
	Exit string
	Help bool
	Mode string

	FlagGarage []string
	Detail     bool
	Ignore     bool
	Delete     bool
}

type HTTPscope struct {
	URL          string
	Domain       string
	Port         int
	ResponseCode string
	Cookie       int
}

type Penetrationscope struct {
	XSS          bool
	SQLInjection bool
}
