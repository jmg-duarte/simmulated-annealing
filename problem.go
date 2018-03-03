package main

import "github.com/jmg-duarte/simmulated-annealing/graph"

// Problem represents the simmulated annealing inputs
type Problem struct {
	Graph               *graph.Graph
	StartingTemperature float64
	LimitTemperature    float64
	NumberOfIterations  int
	Decay               DecayFunction
}
