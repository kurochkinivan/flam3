package variation

const (
	Linear       VariationName = "linear"
	Sinusoidal   VariationName = "sinusoidal"
	Spherical    VariationName = "spherical"
	Swirl        VariationName = "swirl"
	Horseshoe    VariationName = "horseshoe"
	Polar        VariationName = "polar"
	Handkerchief VariationName = "handkerchief"
	Heart        VariationName = "heart"
	Disk         VariationName = "disk"
	Spiral       VariationName = "spiral"
	Hyperbolic   VariationName = "hyperbolic"
	Diamond      VariationName = "diamond"
	Ex           VariationName = "ex"
	Bent         VariationName = "bent"
	Fisheye      VariationName = "fisheye"
	Eyefish      VariationName = "eyefish"
	Bubble       VariationName = "bubble"
	Cylinder     VariationName = "cylinder"
	Tangent      VariationName = "tangent"
	Cross        VariationName = "cross"
	Exponential  VariationName = "exponential"
	Power        VariationName = "power"
	Cosine       VariationName = "cosine"
)

var registry = map[VariationName]NamedFunction{
	Linear:       NewNamedFunction(Linear, linear),
	Sinusoidal:   NewNamedFunction(Sinusoidal, sinusoidal),
	Spherical:    NewNamedFunction(Spherical, spherical),
	Swirl:        NewNamedFunction(Swirl, swirl),
	Horseshoe:    NewNamedFunction(Horseshoe, horseshoe),
	Polar:        NewNamedFunction(Polar, polar),
	Handkerchief: NewNamedFunction(Handkerchief, handkerchief),
	Heart:        NewNamedFunction(Heart, heart),
	Disk:         NewNamedFunction(Disk, disk),
	Spiral:       NewNamedFunction(Spiral, spiral),
	Hyperbolic:   NewNamedFunction(Hyperbolic, hyperbolic),
	Diamond:      NewNamedFunction(Diamond, diamond),
	Ex:           NewNamedFunction(Ex, ex),
	Bent:         NewNamedFunction(Bent, bent),
	Fisheye:      NewNamedFunction(Fisheye, fisheye),
	Eyefish:      NewNamedFunction(Eyefish, eyefish),
	Bubble:       NewNamedFunction(Bubble, bubble),
	Cylinder:     NewNamedFunction(Cylinder, cylinder),
	Tangent:      NewNamedFunction(Tangent, tangent),
	Cross:        NewNamedFunction(Cross, cross),
	Exponential:  NewNamedFunction(Exponential, exponential),
	Power:        NewNamedFunction(Power, power),
	Cosine:       NewNamedFunction(Cosine, cosine),
}
