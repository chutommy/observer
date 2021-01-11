package config

// Config holds IoT specifications and customizable variables.
type Config struct {
	General     General     `toml:"general"`
	Servos      Servos      `toml:"servos"`
	Camera      Camera      `toml:"camera"`
	Targeting   Targeting   `toml:"targeting"`
	Calibration Calibration `toml:"calibration"`
}

// General is a configuration of the Observer's runtime.
type General struct {
	Show         bool `toml:"show"`
	Period       int  `toml:"period"`
	IdleDuration int  `toml:"idleDuration"`
}

// Servos stores Servo's GPIO pin numbers.
type Servos struct {
	PinX int `toml:"pinX"`
	PinY int `toml:"pinY"`
}

// Frame stores the frame resolution.
type Frame struct {
	Width  int `toml:"width"`
	Height int `toml:"height"`
}

// AngleOfView stores the angle of view of the camera.
type AngleOfView struct {
	Horizontal float64 `toml:"horizontal"`
	Vertical   float64 `toml:"vertical"`
}

// Camera stores the camera's specification.
type Camera struct {
	Source      int         `toml:"source"`
	MaxFPS      int         `toml:"maxFPS"`
	Frame       Frame       `toml:"frame"`
	AngleOfView AngleOfView `toml:"angleOfView"`
}

// Target stores specification of the target's rectangle.
type Target struct {
	Red       int `toml:"red"`
	Green     int `toml:"green"`
	Blue      int `toml:"blue"`
	Thickness int `toml:"thickness"`
}

// Target stores specification of the target's rectangle.

// Other stores specification of the rectangles of others.
type Other struct {
	Red       int `toml:"red"`
	Green     int `toml:"green"`
	Blue      int `toml:"blue"`
	Thickness int `toml:"thickness"`
}

// MidRect stores specification of the middle rectangle indicating the focus area.
type MidRect struct {
	Red       int `toml:"red"`
	Green     int `toml:"green"`
	Blue      int `toml:"blue"`
	Thickness int `toml:"thickness"`
}

// Color stores the customizable colors of rectangles.
type Color struct {
	Target  Target  `toml:"target"`
	Other   Other   `toml:"other"`
	MidRect MidRect `toml:"midRect"`
}

// Targeting specifies the targeting method, sensitivity and colors.
type Targeting struct {
	AimArea  float64  `toml:"aimArea"`
	Cascades []string `toml:"cascades"`
	Color    Color    `toml:"color"`
}

// Invert defines whether the movement of the Servos is inverted or not.
type Invert struct {
	X bool `toml:"x"`
	Y bool `toml:"y"`
}

// Coefficient defines the coefficient of the Servos movement.
type Coefficient struct {
	X float64 `toml:"x"`
	Y float64 `toml:"y"`
}

// Tolerate modifies the focus area.
type Tolerate struct {
	X int `toml:"x"`
	Y int `toml:"y"`
}

// Calibration stores the settings of the coefficients for optimized and smooth movements.
type Calibration struct {
	CalibrateOnStart bool        `toml:"calibrateOnStart"`
	Invert           Invert      `toml:"invert"`
	Coefficient      Coefficient `toml:"coefficient"`
	Tolerate         Tolerate    `toml:"tolerate"`
}
