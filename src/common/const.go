package common

const (
	Uid1LikeUid2    = 1 << 0
	Uid1DisLikeUid2 = 1 << 2

	Uid2LikeUid1    = 1 << 1
	Uid2DisLikeUid1 = 1 << 3
)

const (
	Like    = "liked"
	DisLike = "disliked"
)
