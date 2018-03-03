package graph

// Problem represents the simmulated annealing inputs
type Problem struct {
	Graph               *Graph
	StartingTemperature float64
	LimitTemperature    float64
	NumberOfIterations  int
	Decay               DecayFunction
}
