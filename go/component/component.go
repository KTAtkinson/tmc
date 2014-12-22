type Component interface {
  GetId() string
  GetName() string

  SetName(string)
  AddChild(Component)

  IsDecendentOf(Component) bool
}

type BaseComponent struct {
  id, name string
  children []Component
}

func (c BaseComponent) GetId() string {
  return c.id
}

func (c BaseComponent) GetName() string {
  return c.name
}

func (c BaseComponent) SetName(newName string) {
  c.name = newName
}

func (c BaseComponent) AddChild(child Component) {
  c.children = append(c.children, child)
}

func (c BaseComponent) IsDecendentOf(decent Component) bool {
  for i := 0; i < len(self.children); i++ {
    child := decent.children[i]

    if child.GetId() == c.GetId() {
      return true
    }
    else if c.IsDecendentOf(child) {
      return true
    }
  }

  return false
}

func BaseComponentFromJson(json string) Component, nil {
  b := byte(json)

  var c BaseComponent
  err := json.Unmarshal(b, &c)

  if err != nil {
    return nil, err
  }

  return c, nil
}

func NewBaseComponent(id, name string) BaseComponent {
  return BaseComponent{
      id: id,
      name: name,
  }
}


