package lc

type Twitter struct {
	UserId       []int
	TweetID      []int
	FollowRelate map[int][]int
}

func ConstructorTwitter() Twitter {
	var userId []int
	var tweetId []int
	followRelate := make(map[int][]int)
	return Twitter{UserId: userId, TweetID: tweetId, FollowRelate: followRelate}
}
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.UserId = append(t.UserId, userId)
	t.TweetID = append(t.TweetID, tweetId)
	return
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	//If contain key, ok == true
	feedUser := make(map[int]int)
	var feedList []int
	if _, ok := t.FollowRelate[userId]; ok {
		for _, i2 := range t.FollowRelate[userId] {
			if _, ok := feedUser[i2]; !ok {
				feedUser[i2] = 0
			}
		}
	}
	//map{userid/follower => 1,2,3,4}
	feedUser[userId] = 0
	for i, i2 := range t.UserId {
		//Exist key
		if _, ok := feedUser[i2]; ok {
			if len(feedList) >= 10 {
				feedList = feedList[:len(feedList)-1]
			}
			feedList = append([]int{t.TweetID[i]}, feedList...)
		}
	}
	return feedList
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	//Check if is followed
	if _, ok := t.FollowRelate[followerId]; ok {
		//Check
		var isExist bool = false
		for _, i2 := range t.FollowRelate[followerId] {
			if i2 == followeeId {
				isExist = true
				break
			}
		}
		if !isExist {
			t.FollowRelate[followerId] = append(t.FollowRelate[followerId], followeeId)
		}

	} else {
		t.FollowRelate[followerId] = []int{followeeId}
	}
	return
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := t.FollowRelate[followerId]; ok {
		var isExist bool = false
		for _, i2 := range t.FollowRelate[followerId] {
			if i2 == followeeId {
				isExist = true
				break
			}
		}
		if isExist {
			var temp []int
			for _, v := range t.FollowRelate[followerId] {
				if v != followeeId {
					temp = append(temp, v)
				}
			}
			t.FollowRelate[followerId] = temp
		}
	}
	return
}
