// Package handlebars provides handlebars evaluation
package handlebars

import (
	"reflect"
	"sync"
)

type Handlebars struct {
	// helpers stores all globally registered helpers
	helpers map[string]reflect.Value

	// protects global helpers
	helpersMutex sync.RWMutex

	// partials stores all global partials
	partials map[string]*partial

	// protects global partials
	partialsMutex sync.RWMutex
}

func New() *Handlebars {
	h := Handlebars{
		helpers:  make(map[string]reflect.Value),
		partials: make(map[string]*partial),
	}

	// register builtin helpers
	h.RegisterHelper("if", ifHelper)
	h.RegisterHelper("unless", unlessHelper)
	h.RegisterHelper("with", withHelper)
	h.RegisterHelper("each", eachHelper)
	h.RegisterHelper("log", logHelper)
	h.RegisterHelper("lookup", lookupHelper)
	h.RegisterHelper("equal", equalHelper)

	return &h
}

// Render parses a template and evaluates it with given context
//
// Note that this function call is not optimal as your template is parsed everytime you call it. You should use Parse() function instead.
func (h *Handlebars) Render(source string, ctx interface{}) (string, error) {
	// parse template
	tpl, err := h.Parse(source)
	if err != nil {
		return "", err
	}

	// renders template
	str, err := tpl.Exec(ctx)
	if err != nil {
		return "", err
	}

	return str, nil
}

// MustRender parses a template and evaluates it with given context. It panics on error.
//
// Note that this function call is not optimal as your template is parsed everytime you call it. You should use Parse() function instead.
func (h *Handlebars) MustRender(source string, ctx interface{}) string {
	return h.MustParse(source).MustExec(ctx)
}
