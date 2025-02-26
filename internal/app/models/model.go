package models

func GetModelsToMigrate() []interface{} {
	return []interface{}{
		&Users{},
		&Academic{},
		&Achievement{},
		&Activity{},
		&Thesis{},
		&Predicate{},
		&Course{},
	}
}
