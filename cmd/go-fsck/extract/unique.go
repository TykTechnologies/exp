package extract

import "github.com/TykTechnologies/exp/cmd/go-fsck/model"

func unique(defs []*model.Definition) []*model.Definition {
	result := make([]*model.Definition, 0, len(defs))

	for _, def := range defs {
		var match bool
		for _, res := range result {
			if res.Package.Equal(def.Package) {
				match = true
				res.Merge(def)
				break
			}
		}

		if !match {
			result = append(result, def)
		}
	}

	return result
}
