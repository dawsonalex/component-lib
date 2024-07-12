package tpl

import (
	"encoding/xml"
	"fmt"
	"io"
)

type ComponentType interface {
	HeroConfig | TextConfig
}

type HeroConfig struct {
	HeroText string `xml:"hero-text"`
}

type TextConfig struct {
	Value string `xml:"value"`
}

type Component[t ComponentType] struct {
	// Name is the name of the component type
	Name string `xml:"-"`

	// ID is the identifier for the instance in the tree.
	Id         string `xml:"id,attr,omitempty"`
	GridRow    string `xml:"grid-row,attr"`
	GridColumn string `xml:"grid-column,attr"`
	//Style      string `xml:"style,attr"`
	CustomConfig t `xml:"custom-config"`
}

func implements[t ComponentType]() {
	var _ xml.Marshaler = &Component[t]{}
}

// This exists as an annoying hack because the type keyword can't be used in generic functions.
// It's used to prevent MarshalXML on Compoent[t] from recursing into a stack overflow.
type componentType[t ComponentType] Component[t]

func (c Component[t]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Use our silly little type to prevent a stack overflow, passing `c` to EncodeElement will cause a stack overflow.
	localComponent := componentType[t](c)
	start.Name.Local = c.Name
	if err := e.EncodeElement(localComponent, start); err != nil {
		return err
	}
	start.Name.Local = c.Name
	return nil
}

//func (c Component[t]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
//	//TODO implement me
//	panic("implement me")
//}

type ComponentTree struct {
	// XMLName is used in XML marshalling
	//XMLName    struct{} `xml:"components"`
	Components []any
}

// TODO: figure out wtf is going on with marshalling here.
func (t *ComponentTree) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "updated-name"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, component := range t.Components {
		err = e.Encode(&component)
		if err != nil {
			return err
		}
	}
	type componentTree ComponentTree

	return e.EncodeToken(start.End())
	//return e.EncodeElement(t, start)
}

// TODO: implement MarshalXML so element names are observed
func (t *ComponentTree) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		if startElement, ok := token.(xml.StartElement); ok {
			switch startElement.Name.Local {
			case "hero":
				var hero Component[HeroConfig]
				if err := d.DecodeElement(&hero, &startElement); err != nil {
					// TODO: This should wrap an error indicating line and cursor position.
					return err
				}
				hero.Name = "hero"
				t.Components = append(t.Components, hero)
			default:
				line, column := d.InputPos()
				return fmt.Errorf("unexpected element '%s' at position %d:%d", startElement.Name.Local, line, column)
			}
		}
	}

	return nil
}
