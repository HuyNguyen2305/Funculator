package services

import "funculator/parser"

func Calculate(expression string) (float64, error) {
	return parser.Evaluate(expression)
}
