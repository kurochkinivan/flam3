package input_config

const (
	DefaultWidth           = 1920
	DefaultHeight          = 1080
	DefaultSeed            = 5.0
	DefaultIterations      = 2500
	DefaultOutputPath      = "fractal.png"
	DefaultThreads         = 1
	DefaultGammaCorrection = false
	DefaultGamma           = 2.2
	DefaultSymmetryLevel   = 1
)

// ApplyDefaults sets default values for unset fields.
func (c *Config) ApplyDefaults() {
	if c.Size.Width == 0 {
		c.Size.Width = DefaultWidth
	}
	if c.Size.Height == 0 {
		c.Size.Height = DefaultHeight
	}

	if c.Seed == 0 {
		c.Seed = DefaultSeed
	}

	if c.Iterations == 0 {
		c.Iterations = DefaultIterations
	}

	if c.Output == "" {
		c.Output = DefaultOutputPath
	}

	if c.Threads == 0 {
		c.Threads = DefaultThreads
	}

	if c.Gamma == 0 {
		c.Gamma = DefaultGamma
	}

	if c.SymmetryLevel == 0 {
		c.SymmetryLevel = DefaultSymmetryLevel
	}
}
