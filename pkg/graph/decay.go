package graph

type DecayFunction func(float64) float64

func Decay(decay string, variance float64) DecayFunction {
	switch decay {
	case "geometric":
		return GeometricDecay(variance)

	case "gradual":
		return GradualDecay(variance)

	default:
		return GeometricDecay(variance)
	}
}

// T[k] = T[k-1] * alfa
func GeometricDecay(alfa float64) DecayFunction {
	return func(temp float64) float64 {
		return temp * alfa
	}
}

// T[k] = T[k-1] / (1 + b T[k-1])
func GradualDecay(beta float64) DecayFunction {
	return func(temp float64) float64 {
		return temp / (1 + beta*temp)
	}
}
