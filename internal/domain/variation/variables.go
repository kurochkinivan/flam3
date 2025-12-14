package variation

const (
	Linear       Name = "linear"
	Sinusoidal   Name = "sinusoidal"
	Spherical    Name = "spherical"
	Swirl        Name = "swirl"
	Horseshoe    Name = "horseshoe"
	Polar        Name = "polar"
	Handkerchief Name = "handkerchief"
	Heart        Name = "heart"
	Disk         Name = "disk"
	Spiral       Name = "spiral"
	Hyperbolic   Name = "hyperbolic"
	Diamond      Name = "diamond"
	Ex           Name = "ex"
	Bent         Name = "bent"
	Fisheye      Name = "fisheye"
	Eyefish      Name = "eyefish"
	Bubble       Name = "bubble"
	Cylinder     Name = "cylinder"
	Tangent      Name = "tangent"
	Cross        Name = "cross"
	Exponential  Name = "exponential"
	Power        Name = "power"
	Cosine       Name = "cosine"
)

var registry = map[Name]NamedFunction{
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
