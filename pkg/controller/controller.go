package controller

import (
	CON "social-network-go/pkg/config"
)

func Get(id interface{}, what string) string {
	db := CON.DB()
	var RET string
	db.QueryRow("SELECT "+what+" AS RET FROM user1 WHERE id=?", id).Scan(&RET)
	return RET
}

func IsFollowing(by string, to string) bool {
	db := CON.DB()
	var followCount int
	db.QueryRow("SELECT COUNT(followID) AS followCount FROM follow WHERE followBy=? AND followTo=? LIMIT 1", by, to).Scan(&followCount)
	if followCount == 0 {
		return false
	}
	return true
}

func UsernameDecider(user int, session string) string {
	username := Get(user, "username")
	sesUsername := Get(session, "username")
	if username == sesUsername {
		return "You"
	}
	return username
}

func NoOfFollowers(user int) int {
	db := CON.DB()
	var followersCount int
	db.QueryRow("SELECT COUNT(followID) AS followersCount FROM follow WHERE followTo=?", user).Scan(&followersCount)
	return followersCount
}

func LikedOrNot(post int, user interface{}) bool {
	db := CON.DB()
	var likeCount int
	db.QueryRow("SELECT COUNT(likeID) AS likeCount FROM likes WHERE likeBy=? AND postID=?", user, post).Scan(&likeCount)
	if likeCount == 0 {
		return false
	}
	return true
}
