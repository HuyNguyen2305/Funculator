package parser

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

type Parser struct {
	input string
	pos   int
}

func Evaluate(input string) (float64, error) {
	if len(input) == 0 || len(input) > 500 {
		return 0, fmt.Errorf("input expression is empty or exceeds the maximum length of 500 characters")
	}

	p := &Parser{input: input}

	result, err := p.parseExpression()
	if err != nil {
		return 0, err
	}

	p.skipSpaces()

	if p.pos != len(p.input) {
		return 0, fmt.Errorf("unexpected character: %c", p.input[p.pos])
	}

	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, fmt.Errorf("result is not a finite number")
	}

	return result, nil
}

func (p *Parser) parseExpression() (float64, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}

	for {
		p.skipSpaces()

		if p.match('+') {
			right, err := p.parseTerm()
			if err != nil {
				return 0, err
			}
			left += right

		} else if p.match('-') {
			right, err := p.parseTerm()
			if err != nil {
				return 0, err
			}
			left -= right

		} else {
			break
		}
	}
	return left, nil
}

func (p *Parser) parseTerm() (float64, error) {
	left, err := p.parseUnary()
	if err != nil {
		return 0, err
	}

	for {
		p.skipSpaces()

		if p.match('*') {
			right, err := p.parseUnary()
			if err != nil {
				return 0, err
			}
			left *= right

		} else if p.match('/') {
			right, err := p.parseUnary()
			if err != nil {
				return 0, err
			}

			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}

			left /= right

		} else {
			break
		}
	}

	return left, nil
}

func (p *Parser) parseUnary() (float64, error) {
	p.skipSpaces()

	if p.match('+') {
		return p.parseUnary()
	}

	if p.match('-') {
		value, err := p.parseUnary()
		return -value, err
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() (float64, error) {
	p.skipSpaces()

	if p.match('(') {
		value, err := p.parseExpression()
		if err != nil {
			return 0, err
		}

		p.skipSpaces()

		if !p.match(')') {
			return 0, fmt.Errorf("Missing closing parenthesis")
		}

		return value, nil
	}

	return p.parseNumber()
}

func (p *Parser) parseNumber() (float64, error) {
	p.skipSpaces()

	start := p.pos
	hasDigit := false
	hasDot := false

	for p.pos < len(p.input) {
		ch := rune(p.input[p.pos])

		if unicode.IsDigit(ch) {
			hasDigit = true
			p.pos++
		} else if ch == '.' && !hasDot {
			hasDot = true
			p.pos++
		} else {
			break
		}
	}

	if !hasDigit {
		return 0, fmt.Errorf("expected a number")
	}

	return strconv.ParseFloat(p.input[start:p.pos], 64)
}

func (p *Parser) match(ch byte) bool {
	if p.pos < len(p.input) && p.input[p.pos] == ch {
		p.pos++
		return true
	}
	return false
}

func (p *Parser) skipSpaces() {
	for p.pos < len(p.input) &&
		(p.input[p.pos] == ' ' ||
			p.input[p.pos] == '\t' ||
			p.input[p.pos] == '\n' ||
			p.input[p.pos] == '\r') {
		p.pos++
	}
}
