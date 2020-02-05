// Code generated by go generate; DO NOT EDIT.
package nominal_catalog

import (
	"github.com/docker/oscalkit/types/oscal/validation_root"
)

// NOT TO BE USED FOR A BASE METASCHEMA ONLY FOR A MODULE
type NominalCatalog struct {

	// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
	Param *Param `xml:"param,omitempty" json:"param,omitempty"`
	// A partition or component of a control or part
	Part *Part `xml:"part,omitempty" json:"part,omitempty"`
}

// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
type Param struct {

	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`
	// Another parameter invoking this one
	DependsOn string `xml:"depends-on,attr,omitempty" json:"dependsOn,omitempty"`

	// A short name for the parameter.
	Label Label `xml:"label,omitempty" json:"label,omitempty"`
	// Indicates and explains the purpose and use of a parameter
	Descriptions []Usage `xml:"usage,omitempty" json:"descriptions,omitempty"`
	// A formal or informal expression of a constraint or test
	Constraints []Constraint `xml:"constraint,omitempty" json:"constraints,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// A prose statement that provides a recommendation for the use of a parameter.
	Guidance []Guideline `xml:"guideline,omitempty" json:"guidance,omitempty"`
	// A recommended parameter value or set of values.
	Value Value `xml:"value,omitempty" json:"value,omitempty"`
	// A set of parameter value choices, that may be picked from to set the parameter value.
	Select *Select `xml:"select,omitempty" json:"select,omitempty"`
}

// A prose statement that provides a recommendation for the use of a parameter.
type Guideline struct {

	// Prose permits multiple paragraphs, lists, tables etc.
	Prose *Prose `xml:"prose,omitempty" json:"prose,omitempty"`
}

// Presenting a choice among alternatives
type Select struct {

	// When selecting, a requirement such as one or more
	HowMany string `xml:"how-many,attr,omitempty" json:"howMany,omitempty"`

	// A value selection among several such options
	Alternatives []Choice `xml:"choice,omitempty" json:"alternatives,omitempty"`
}

// A partition or component of a control or part
type Part struct {

	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Identifying the purpose and intended use of the property, part or other object.
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
	// A namespace qualifying the name.
	Ns string `xml:"ns,attr,omitempty" json:"ns,omitempty"`
	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// Prose permits multiple paragraphs, lists, tables etc.
	Prose *Prose `xml:"prose,omitempty" json:"prose,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// A partition or component of a control or part
	Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
}

// A placeholder for a missing value, in display.

type Label string

// Indicates and explains the purpose and use of a parameter
type Usage struct {
	// Unique identifier of the containing object
	Id    string `xml:"id,attr,omitempty" json:"id,omitempty"`
	Value string `xml:",chardata" json:"value,omitempty"`
}

// A formal or informal expression of a constraint or test
type Constraint struct {
	// A formal (executable) expression of a constraint
	Test  string `xml:"test,attr,omitempty" json:"test,omitempty"`
	Value string `xml:",chardata" json:"value,omitempty"`
}

// Indicates a permissible value for a parameter or property

type Value string

// A value selection among several such options

type Choice string

// Prose permits multiple paragraphs, lists, tables etc.

type Prose = Markup

type Link = validation_root.Link

type Prop = validation_root.Prop

type Title = validation_root.Title
