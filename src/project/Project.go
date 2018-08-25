package Project

type Project struct{
	name string
	owner string
}

func (p Project) SetName(name string) { 
	p.name = name
}

func (p Project) SetOwner(owner string){
	p.owner = owner
}