package backend

import "github.com/a-clap/distillation-gui/backend/parameters"

// Models allows us to create models.ts in frontend with needed structures
type Models struct {
}

func (*Models) Temperature() parameters.Temperature {
	return parameters.Temperature{}
}
