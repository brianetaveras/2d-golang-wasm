package gaia

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

type Component interface {
	onUpdate() error
	onDraw(screen *ebiten.Image) error
}

type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool
	Tag        string
	Components []Component
}

func (element *Element) Update() error {
	for _, component := range element.Components {
		err := component.onUpdate()

		if err != nil {
			return err
		}
	}

	return nil
}

func (element *Element) Draw(screen *ebiten.Image) error {
	for _, component := range element.Components {
		err := component.onDraw(screen)

		if err != nil {
			return err
		}
	}

	return nil
}

func (element *Element) AddComponent(newComponent Component) {
	for _, existing := range element.Components {
		if reflect.TypeOf(newComponent) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("%v exists within this element already", reflect.TypeOf((newComponent))))
		}
	}

	element.Components = append(element.Components, newComponent)
}

func (element *Element) GetComponentType(componentType Component) Component {
	typeToGet := reflect.TypeOf(componentType)

	for _, component := range element.Components {
		if reflect.TypeOf(component) == typeToGet {
			return component
		}
	}

	return nil
}
