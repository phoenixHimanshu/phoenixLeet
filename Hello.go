package main

type astronaut struct {
	name string
	age  int
}

type spaceship struct {
	name   string
	weight int
}

func main() {
	m := make(map[astronaut]spaceship)
	Rot := astronaut{"Rot", 35}
	Mike := astronaut{"Mike", 25}
	Rob := astronaut{"Rob", 35}

	m[Rot] = spaceship{"ISV", 1234}
	m[Mike] = spaceship{"ISLV", 2929}
	m[Rob] = spaceship{"NASA", 7909}

	delete(m, Rot)
	val, ok := m[Rob]
	println(ok)
	println(val.name)
	//new()

}
