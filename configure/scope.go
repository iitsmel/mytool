package configure

type Scope struct {
	HTTP    HTTPscope
	General Generalscope
	Plan    Planscope
	Attack  AttackFeature
}

type AttackFeature struct {
	XSS    bool
	SQL    bool
	Target string
}

type Generalscope struct {
	Exit       string
	Help       bool
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

type Planscope struct {
	Target string
}
