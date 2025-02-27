package models

import (
	"reflect"
	"testing"
)

func TestGetModelsToMigrate(t *testing.T) {
	expectedModels := []interface{}{
		&Users{},
		&Academic{},
		&Achievement{},
		&Activity{},
		&Thesis{},
		&Predicate{},
		&Course{},
	}

	models := GetModelsToMigrate()

	if !reflect.DeepEqual(models, expectedModels) {
		t.Errorf("Expected %v, but got %v", expectedModels, models)
	}
}
