# Simmulated Annealing

*"Simulated annealing (SA) is a probabilistic technique for approximating the global optimum of a given function." - [Wikipedia](https://en.wikipedia.org/wiki/Simulated_annealing)*

## Introduction

During the semester I took an Introduction to Artificial Intellegence course, one of the exercises was to apply the Simmulated Annealing *algorithm* to the [Travelling Salesman Problem](https://en.wikipedia.org/wiki/Travelling_salesman_problem)

Me and my colleagues implemented it in **Java**, since it is pretty much the *lingua franca* of the course.
But I wasn't satisfied. During the vacations I wanted to do it in Go, this wonderfully simple C-like language, and so I did. This is the result.

---

## Usage

The program takes a text file as input with the following format

| Node | X | Y |
|------|---|---|
|Lisbon| 50 | 40 |
|Porto| 140 | 0 |
|London | 300 | 300|
|*(and so on)*|...|...|

After compilling you can run the program simply by typing:

`./simm <input>.txt`

You can call for help by typing:

`./simm -h`