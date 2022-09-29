package community

type CoreCommunity struct {
	ID           uint
	Title        string
	Descriptions string
	Logo         string
	Members      int64
}

type DataInterface interface {
	Insert(userid int, core CoreCommunity) (string, error)
	SelectList() ([]CoreCommunity, string, error)
	SelectMembers(communityid int) ([]string, string, error)
	Delete(userid, communityid int) (string, error)
}

type UsecaseInterface interface {
	AddNewCommunity(userid int, core CoreCommunity) (string, error)
	GetListCommunity() ([]CoreCommunity, string, error)
	GetMembers(communityid int) ([]string, string, error)
	Leave(userid, communityid int) (string, error)
}
