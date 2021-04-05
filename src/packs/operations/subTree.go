package operations


type SubTree interface {
	GetFather()					(*SubTree, bool)
	GetRoot()					*WotoSerialized
	GetValue(_s string)			(string, bool)
	GetObjectType(_s string)	ObjectType
	IsValue(_s string)			bool
	IsSubtree(_s string)		bool
	IsTop()						bool
}

type subTree struct {

}

func (s subTree) GetFather() (*SubTree, bool) {
	panic("implement me")
}

func (s subTree) GetRoot() *WotoSerialized {
	panic("implement me")
}

func (s subTree) GetValue(_s string) (string, bool) {
	panic("implement me")
}

func (s subTree) GetObjectType(_s string) ObjectType {
	panic("implement me")
}

func (s subTree) IsValue(_s string) bool {
	panic("implement me")
}

func (s subTree) IsSubtree(_s string) bool {
	panic("implement me")
}

func (s subTree) IsTop() bool {
	panic("implement me")
}
