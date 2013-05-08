package script

import (
	"fmt"
	"io"
	"strconv"
)

// TODO: parse tree could keep line number references for runtime error reporting

type Parser struct {
	lexer
	world
}

func NewParser() *Parser {
	p := new(Parser)
	p.world.init()
	return p
}

func (p *Parser) parse(src io.Reader) (code []Expr, err error) {
	defer func() {
		panc := recover()
		if panc != nil {
			code = nil
			err = fmt.Errorf("%v", panc)
		}
	}()

	p.lexer.init(src)

	expr, err := p.parseLine()
	for err != io.EOF {
		if err != nil {
			return nil, err
		}
		code = append(code, expr)
		expr, err = p.parseLine()
	}
	return code, nil
}

func (p *Parser) parseLine() (ex Expr, err error) {

	p.advance()
	switch p.typ {
	case EOF:
		return nil, io.EOF // marks end of input
	case EOL:
		return &nop{}, nil // empty line
	default:
		Expr := p.parseExpr()
		p.advance()
		if p.typ == EOL || p.typ == EOF { // statement has to be terminated
			return Expr, nil
		} else {
			panic(p.unexpected())
		}
	}
}

func (p *Parser) parseIdent() Expr {
	switch p.peekTyp {
	case LPAREN:
		return p.parseCall()
	case ASSIGN:
		return p.parseAssign()
	default:
		return p.getvar(p.str)
	}
}

func (p *Parser) parseExpr() Expr {
	switch p.typ {
	case IDENT:
		return p.parseIdent()
	case NUM:
		return p.parseNum()
	case LPAREN:
		return List(p.parseArgs())
	default:
		panic(p.unexpected())
	}
}

func (p *Parser) parseCall() Expr {
	funcname := p.str
	p.advance()
	assert(p.typ == LPAREN)
	args := p.parseArgs()
	return p.newCall(funcname, args)
}

func (p *Parser) parseAssign() Expr {
	left := p.str
	p.advance()
	assert(p.typ == ASSIGN)
	p.advance()
	right := p.parseExpr()
	return p.newAssign(left, right)
}

func (p *Parser) parseNum() Expr {
	val, err := strconv.ParseFloat(p.str, 64)
	if err != nil {
		panic(err)
	}
	return num(val)
}

func (p *Parser) parseArgs() []Expr {
	var args []Expr
	p.advance()
	for {
		switch p.typ {
		case RPAREN:
			return args
		case NUM:
			args = append(args, p.parseNum())
		case IDENT:
			args = append(args, p.parseIdent())
		default:
			panic(p.unexpected())
		}
		p.advance()
		if p.typ != COMMA && p.typ != RPAREN {
			panic(fmt.Errorf(`%v: expected "," or ")"`, p.Position))
		}
		if p.typ == COMMA {
			p.advance()
			if p.typ == RPAREN {
				panic(p.unexpected())
			}
		}
	}
}

func assert(test bool) {
	if !test {
		panic("assertion failed")
	}
}
