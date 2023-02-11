package dal

import (
	"context"
	"douyin/common/conf"
	"fmt"
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func testInit() {
	conf.InitConfig()
	DSN := conf.Database.DSN()
	var err error
	DB, err = gorm.Open(mysql.Open(DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Fatal(err)
	}
	DB = DB.Debug()
}

func TestCreateVideo(t *testing.T) {
	testInit()
	err := CreateVideo(context.Background(), 3, "3_title", "3_playUrl", "3_coverUrl")
	if err != nil {
		panic(err)
	}
}

func TestMGetVideoByUserId(t *testing.T) {
	testInit()
	userId := int64(23)
	videoInfo, err := MGetVideoByUserID(context.Background(), userId)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(videoInfo); i++ {
		fmt.Println(*videoInfo[i])
	}
}

func TestGetLikeCount(t *testing.T) {
	testInit()
	// 需要先验证是否已点赞，如果点过赞不应执行插入操作
	vid := int64(11)
	cnt, err := GetLikeCount(context.Background(), vid)
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
}

func TestGetCommentCount(t *testing.T) {
	testInit()
	vid := int64(11)
	cnt, err := GetCommentCount(context.Background(), vid)
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
}

func TestIsFavorite(t *testing.T) {
	testInit()
	uid := int64(23)
	vid := int64(11)
	flag, err := IsFavorite(context.Background(), vid, uid)
	if err != nil {
		panic(err)
	}
	fmt.Println(flag)
}

func TestMGetVideoByTime(t *testing.T) {
	testInit()
	lastTime := time.Now()
	videos, nextTime, err := MGetVideoByTime(context.Background(), lastTime, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(nextTime)
	for i := 0; i < len(videos); i++ {
		fmt.Println(*videos[i])
	}
}

func TestLikeVideo(t *testing.T) {
	testInit()
	userId := int64(1)
	videoId := int64(11)
	if err := LikeVideo(context.Background(), userId, videoId); err != nil {
		panic(err)
	}
}

func TestUnLikeVideo(t *testing.T) {
	testInit()
	userId := int64(23)
	videoId := int64(10)
	if err := UnLikeVideo(context.Background(), userId, videoId); err != nil {
		panic(err)
	}
}

func TestMGetLikeList(t *testing.T) {
	testInit()
	userId := int64(1)
	favorites, err := MGetLikeList(context.Background(), userId)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(favorites); i++ {
		fmt.Printf("%#v\n", *favorites[i])
	}
}
