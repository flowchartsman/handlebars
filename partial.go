package handlebars

import (
	"fmt"
)

// partial represents a partial template
type partial struct {
	handlebars *Handlebars
	name       string
	source     string
	tpl        *Template
}

// newPartial instanciates a new partial
func newPartial(h *Handlebars, name string, source string, tpl *Template) *partial {
	return &partial{
		handlebars: h,
		name:       name,
		source:     source,
		tpl:        tpl,
	}
}

// RegisterPartial registers a global partial. That partial will be available to all templates.
func (h *Handlebars) RegisterPartial(name string, source string) {
	h.partialsMutex.Lock()
	defer h.partialsMutex.Unlock()

	if h.partials[name] != nil {
		panic(fmt.Errorf("Partial already registered: %s", name)) // TODO don't panic
	}

	h.partials[name] = newPartial(h, name, source, nil)
}

// RegisterPartials registers several global partials. Those partials will be available to all templates.
func (h *Handlebars) RegisterPartials(partials map[string]string) {
	for name, p := range partials {
		h.RegisterPartial(name, p)
	}
}

// RegisterPartialTemplate registers a global partial with given parsed template. That partial will be available to all templates.
func (h *Handlebars) RegisterPartialTemplate(name string, tpl *Template) {
	h.partialsMutex.Lock()
	defer h.partialsMutex.Unlock()

	if h.partials[name] != nil {
		panic(fmt.Errorf("Partial already registered: %s", name)) // TODO don't panic
	}

	h.partials[name] = newPartial(h, name, "", tpl)
}

// RemovePartial removes the partial registered under the given name. The partial will not be available globally anymore. This does not affect partials registered on a specific template.
func (h *Handlebars) RemovePartial(name string) {
	h.partialsMutex.Lock()
	defer h.partialsMutex.Unlock()

	delete(h.partials, name)
}

// RemoveAllPartials removes all globally registered partials. This does not affect partials registered on a specific template.
func (h *Handlebars) RemoveAllPartials() {
	h.partialsMutex.Lock()
	defer h.partialsMutex.Unlock()

	h.partials = make(map[string]*partial)
}

// findPartial finds a registered global partial
func (h *Handlebars) findPartial(name string) *partial {
	h.partialsMutex.RLock()
	defer h.partialsMutex.RUnlock()

	return h.partials[name]
}

// template returns parsed partial template
func (p *partial) template() (*Template, error) {
	if p.tpl == nil {
		var err error

		p.tpl, err = p.handlebars.Parse(p.source)
		if err != nil {
			return nil, err
		}
	}

	return p.tpl, nil
}
