package model

import (
	"golang.org/x/exp/slices"
)

// TypeInfo returns a `*TypeInfo` for a type by name.
func (p DeclarationList) TypeInfo(name string) *TypeInfo {
	for _, decl := range p {
		for _, t := range decl.Types {
			if t.Name == name {
				return t
			}
		}
	}
	return nil
}

// TypeDeclarations returns a default order of type names and a lookup map.
func (p DeclarationList) TypeDeclarations() ([]string, map[string]bool) {
	result := []string{}
	resultMap := map[string]bool{}
	for _, decl := range p {
		for _, t := range decl.Types {
			result = append(result, t.Name)
			resultMap[t.Name] = true
		}
	}
	return result, resultMap
}

// GetOrder returns a list of type names in the order of declaration.
func (p DeclarationList) GetOrder(root string) []string {
	typeOrder, typeOrderMap := p.TypeDeclarations()

	// If root element can't be found, return default order
	if !slices.Contains(typeOrder, root) {
		return typeOrder
	}

	wantOrder := []string{root}
	wantOrderIndex := 0

	// getTypes ranges over type fields and gets the slice of
	// type names that are referenced.
	var getTypes func(info *TypeInfo) []string
	getTypes = func(info *TypeInfo) []string {
		result := []string{}

		// Log `B` from `type A []*B`
		if typeRef := info.TypeRef(); typeRef != "" {
			if valid, _ := typeOrderMap[typeRef]; valid {
				result = append(result, typeRef)
			}
		}

		// Struct fields
		for _, field := range info.Fields {
			typeRef := field.TypeRef()

			// Skip seen type declarations
			if slices.Contains(wantOrder, typeRef) {
				continue
			}

			if valid, _ := typeOrderMap[typeRef]; valid {
				result = append(result, typeRef)
			}
		}

		return result
	}

	for {
		if wantOrderIndex < len(wantOrder) {
			elem := wantOrder[wantOrderIndex]
			names := getTypes(p.TypeInfo(elem))
			if len(names) > 0 {
				wantOrder = append(wantOrder, names...)
			}
			wantOrderIndex++
			continue
		}
		break
	}

	for _, typeName := range typeOrder {
		if !slices.Contains(wantOrder, typeName) {
			wantOrder = append(wantOrder, typeName)
		}
	}

	return wantOrder
}
