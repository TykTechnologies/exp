package model

// GetOrder returns a list of type names in the order of declaration.
func (p DeclarationList) GetOrder(root string) []string {
	// Step 1: Gather all type declarations from the root type info
	var (
		rootTypeDeclarations []string
		rootTypeInfo         *TypeInfo
	)
	for _, decl := range p {
		for _, t := range decl.Types {
			rootTypeDeclarations = append(rootTypeDeclarations, t.Name)
			if root != "" && t.Name == root {
				rootTypeInfo = t
			}
		}
	}

	// Step 2: If root type info is not found, return rootTypeDeclarations as is
	if root == "" || rootTypeInfo == nil {
		return rootTypeDeclarations
	}

	// Build lookup array for valid declarations
	rootTypeDeclarationsMap := make(map[string]bool)
	for _, v := range rootTypeDeclarations {
		rootTypeDeclarationsMap[v] = true
	}

	// Create a map to track the visited type declarations
	visited := make(map[string]bool)
	visited[root] = true

	// Initialize the output list with the root type name
	outputList := []string{rootTypeInfo.Name}

	// Traverse the fields in the root type info
	for _, field := range rootTypeInfo.Fields {
		if visited[field.Type] {
			continue
		}
		ok, _ := rootTypeDeclarationsMap[field.Type]
		if !ok {
			continue
		}

		// Add the field's type to the output list
		outputList = append(outputList, field.Type)
		visited[field.Type] = true
	}

	// Traverse the whole DeclarationList to add any unreferenced type declarations
	for _, decl := range p {
		for _, t := range decl.Types {
			if visited[t.Name] {
				continue
			}
			outputList = append(outputList, t.Name)
			visited[t.Name] = true
		}
	}

	return outputList
}
