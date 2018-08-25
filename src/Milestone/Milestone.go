package milestone

import "Web/src/Project"

type milestone struct { 
	parent Project.Project
	name string
}

func (m milestone) SetName(name string) { 
	m.name = name
}

func (m milestone) SetParent(parent Project.Project){
	m.parent = parent
}