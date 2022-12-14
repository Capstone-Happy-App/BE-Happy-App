package community

import (
	"time"
)

type CoreCommunity struct {
	ID           uint
	Title        string
	Role         string
	Descriptions string
	Logo         string
	Members      int64
	Feeds        []CoreFeed
}

type CoreFeed struct {
	ID          uint
	Name        string
	Text        string
	UserID      uint
	CommunityID uint
	Date        time.Time
	Comments    []CoreComment
}

type CoreComment struct {
	ID     uint
	Name   string
	Text   string
	UserID uint
	FeedID uint
	Date   time.Time
}

type DataInterface interface {
	Insert(userid int, core CoreCommunity) (string, error)
	SelectList() ([]CoreCommunity, string, error)
	SelectMembers(communityid int) ([]string, string, error)
	Delete(userid, communityid int) (int64, string, error)
	UpdateCommunity(communityid int, core CoreCommunity) (string, error)
	GetUserRole(Userid, communityid int) (string, error)
	InsertToJoin(userid, communityid int) (string, error)
	CheckJoin(userid, communityid int) (string, error)
	SelectCommunityIdWithFeed(feedid int) (int, error)
	InsertFeed(CoreFeed) (string, error)
	SelectCommunity(userid, communityid int) (CoreCommunity, string, error)
	SelectFeed(feedid int) (CoreFeed, string, error)
	InsertComment(CoreComment) (string, error)
	DeleteCommunity(communityid int) (string, error)
	ChangeAdmin(communityid int) (string, string, error)
	SelectListCommunityWithParam(param string) ([]CoreCommunity, string, error)
}

type UsecaseInterface interface {
	AddNewCommunity(userid int, core CoreCommunity) (string, error)
	GetListCommunity() ([]CoreCommunity, string, error)
	GetMembers(communityid int) ([]string, string, error)
	Leave(userid, communityid int) (string, error)
	UpdateCommunity(userid int, core CoreCommunity) (string, error)
	JoinCommunity(userid, communityid int) (string, error)
	PostFeed(CoreFeed) (string, error)
	GetCommunityFeed(userid, communityid int) (CoreCommunity, string, error)
	GetDetailFeed(feedid int) (CoreFeed, string, error)
	AddComment(CoreComment) (string, error)
	GetListCommunityWithParam(param string) ([]CoreCommunity, string, error)
}
