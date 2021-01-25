package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekSeconds = 7 * 86400
	scorePerVote   = 432 // 每票占多少分
)

var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postID int64) error {
	_, err := rdb.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()

	_, err = rdb.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断用户投票的限制
	// getRedisKey(KeyPostTimeZSet) = bluebell:post:time
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrVoteTimeExpire
	}

	// 2. 更新帖子的分数
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val() // 源来的值

	if value == ov {
		return ErrVoteRepeated
	}

	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID).Result()
	if ErrVoteTimeExpire != nil {
		return err
	}
	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		_, err = rdb.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID).Result()
	} else {
		_, err = rdb.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		}).Result()
	}
	return err
}
