package config

type Config struct {
	General     General     `toml:"general"`
	Servos      Servos      `toml:"servos"`
	Camera      Camera      `toml:"camera"`
	Targeting   Targeting   `toml:"targeting"`
	Calibration Calibration `toml:"calibration"`
}

type General struct {
	Show         bool `toml:"show"`
	Period       int  `toml:"period"`
	IdleDuration int  `toml:"idleDuration"`
}

type Servos struct {
	PinX int `toml:"pinX"`
	PinY int `toml:"pinY"`
}

type Frame struct {
	Width  int `toml:"width"`
	Height int `toml:"height"`
}

type AngleOfView struct {
	Horizontal float64 `toml:"horizontal"`
	Vertical   float64 `toml:"vertical"`
}

type Camera struct {
	Source      int         `toml:"source"`
	MaxFPS      int         `toml:"maxFPS"`
	Frame       Frame       `toml:"frame"`
	AngleOfView AngleOfView `toml:"angleOfView"`
}

type Target struct {
	Red       int `toml:"red"`
	Green     int `toml:"green"`
	Blue      int `toml:"blue"`
	Thickness int `toml:"thickness"`
}

type Other struct {
	Red       int `toml:"red"`
	Green     int `toml:"green"`
	Blue      int `toml:"blue"`
	Thickness int `toml:"thickness"`
}

type MidRect struct {
	Red       int `toml:"red"`
	Green     int `toml:"green"`
	Blue      int `toml:"blue"`
	Thickness int `toml:"thickness"`
}

type Color struct {
	Target  Target  `toml:"target"`
	Other   Other   `toml:"other"`
	MidRect MidRect `toml:"midRect"`
}

type Targeting struct {
	AimArea  float64  `toml:"aimArea"`
	Cascades []string `toml:"cascades"`
	Color    Color    `toml:"color"`
}

type Invert struct {
	X bool `toml:"x"`
	Y bool `toml:"y"`
}

type Coefficient struct {
	X float64 `toml:"x"`
	Y float64 `toml:"y"`
}

type Tolerate struct {
	X int `toml:"x"`
	Y int `toml:"y"`
}

type Calibration struct {
	CalibrateOnStart bool        `toml:"calibrateOnStart"`
	Invert           Invert      `toml:"invert"`
	Coefficient      Coefficient `toml:"coefficient"`
	Tolerate         Tolerate    `toml:"tolerate"`
}
