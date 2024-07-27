package gite

type Group struct {
	app    *Server
	prefix string
}

func (g Group) Get(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{prefix: g.prefix, path: path}
	g.app.Get(p.buildGroup(), handlers...)
}

func (g Group) Post(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{prefix: g.prefix, path: path}
	g.app.Post(p.buildGroup(), handlers...)
}

func (g Group) Put(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{prefix: g.prefix, path: path}
	g.app.Put(p.buildGroup(), handlers...)
}

func (g Group) Patch(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{prefix: g.prefix, path: path}
	g.app.Patch(p.buildGroup(), handlers...)
}

func (g Group) Delete(path string, handlers ...HanlderFunc) {
	p := &buildPathConfig{prefix: g.prefix, path: path}
	g.app.Delete(p.buildGroup(), handlers...)
}

func (g Group) Group(prefix string) Router {
	newGroup := &Group{app: g.app, prefix: prefix}
	return newGroup
}
