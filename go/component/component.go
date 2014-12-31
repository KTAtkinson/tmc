package component

type Component interface {
  GetId() string
  GetName() string
  GetChildren() []Component

  SetName(string)
  AddChild(Component)

  IsDecendentOf(Component) bool
}

type BaseComponent struct {
  id, name string
  loaded bool
  children []Component
}

func (c BaseComponent) GetId() string {
  if c.loaded != true {
    c.loadComponent()
  }
  return c.id
}

func (c BaseComponent) GetName() string {
  if c.loaded != true {
    c.loadComponent()
  }
  return c.name
}


func (c BaseComponent) GetChildren() []Component {
  if c.loaded != true {
    c.loadComponent()
  }

  return c.children
}

func (c BaseComponent) SetName(newName string) {
  c.name = newName
}

func (c BaseComponent) AddChild(child Component) {
  c.children = append(c.children, child)
}

func (c BaseComponent) IsDecendentOf(decent Component) bool {
  children := decent.GetChildren()
  for i := 0; i < len(children); i++ {
    child := children[i]

    if child.GetId() == c.GetId() {
      return true
    } else if c.IsDecendentOf(child) {
      return true
    }
  }

  return false
}

func (c BaseComponent) loadComponent() {
  c.loaded = true
}


