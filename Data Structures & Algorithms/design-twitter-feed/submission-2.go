
type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].time > h[j].time } // max-heap
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x any)        { *h = append(*h, x.(Tweet)) }
func (h *TweetHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Tweet struct {
	id   int
	time int
}
type Twitter struct {
	time    int
	tweets  map[int][]Tweet
	follows map[int]map[int]struct{}
}

func Constructor() Twitter {
	return Twitter{
		tweets:  make(map[int][]Tweet),
		follows: make(map[int]map[int]struct{}),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.time++
	this.tweets[userId] = append(this.tweets[userId], Tweet{tweetId, this.time})
	fmt.Printf("tweets: %v - follows: %v\n", this.tweets, this.follows)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &TweetHeap{}
	// get my tweet
	for _, tweet := range this.tweets[userId] {
		*h = append(*h, tweet)
	}
	// get tweet from followee
	if _, exist := this.follows[userId]; exist {
		for followeeId, _ := range this.follows[userId] {
			for _, tweet := range this.tweets[followeeId] {
				*h = append(*h, tweet)
			}
		}
	}
	heap.Init(h)
	posts := []int{}
	count := 0
	for count < 10 && h.Len() > 0 {
		posts = append(posts, heap.Pop(h).(Tweet).id)
		count++
	}
	return posts
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {return}
	if _, exist := this.follows[followerId]; !exist {
		this.follows[followerId] = make(map[int]struct{})
	}
	this.follows[followerId][followeeId] = struct{}{}
	fmt.Printf("tweets: %v - follows: %v\n", this.tweets, this.follows)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, exist := this.follows[followerId]; exist {
		delete(this.follows[followerId], followeeId)
	}
	fmt.Printf("tweets: %v - follows: %v\n", this.tweets, this.follows)
}