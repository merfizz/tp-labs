package main

type Building struct {
	GameObject
	build bool
}

func (b *Building) IsBuilt() bool {
	return b.build
}
