package tpl

import (
	"encoding/xml"
	"log"
	"testing"
)

func TestComponentUnmarshal(t *testing.T) {
	xmlStr := `
		<hero grid-row="1" grid-column="5">
			<custom-config>
				<hero-text>This is the hero</hero-text>
			</custom-config>
		</hero>`

	var heroComponent Component[HeroConfig]
	err := xml.Unmarshal([]byte(xmlStr), &heroComponent)
	if err != nil {
		t.Fatal(err)
	}

	expectedHeroText := "This is the hero"
	if heroComponent.CustomConfig.HeroText != expectedHeroText {
		t.Errorf("expected: %s but got: %s", expectedHeroText, heroComponent.CustomConfig.HeroText)
	}
}

func TestComponentTreeUnmarshal(t *testing.T) {
	xmlStr := `
		<components>
			<hero grid-row="1" grid-column="5">
				<custom-config>
					<hero-text>This is the hero</hero-text>
				</custom-config>
			</hero>
		</components>`

	var componentTree ComponentTree
	err := xml.Unmarshal([]byte(xmlStr), &componentTree)
	if err != nil {
		t.Fatal(err)
	}
}

func TestComponentTreeMarshal(t *testing.T) {
	var componentTree ComponentTree
	componentTree.Components = append(componentTree.Components, Component[HeroConfig]{
		Name:         "hero",
		Id:           "",
		GridRow:      "1",
		GridColumn:   "1",
		CustomConfig: HeroConfig{HeroText: "Hero Text Here"},
	})

	v, err := xml.Marshal(&componentTree)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%s", string(v))
}
