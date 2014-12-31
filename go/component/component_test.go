package component
import "testing" 

func TestGetName(t *testing.T) {
  name := "abc"
  testComponent := BaseComponent{name: name}
  if testComponent.GetName() != name {
    t.Errorf("Expected to retrieve component name %s, found name %s.", name, testComponent.GetName())
  }
}

func TestLoadComponent(t *testing.T) {
  testComponent := BaseComponent{}
  testComponant.loadComponent()
  if self.loaded != true {
    t.Error("Expected a components loaded property to be true after loadComponent call.")
  }
}
