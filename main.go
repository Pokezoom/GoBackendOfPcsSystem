package main

import (
	"GoDockerBuild/middleware"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

var c middleware.EGorm
var o middleware.EGorm
var z middleware.EGorm

type Data struct {
	AppId            string
	CommunityId      string
	Title            string //圈子名称
	AppName          string //店铺名称
	NewIndustryName  string //圈子行业类别
	UserCount        int    //圈子有效用户数
	FeedsCount       int    //圈子动态总数
	NoticeNum        int    //圈子置顶动态数
	NoticeRecordsNum int    //置顶动态浏览数
	NoticeZan        int    //置顶点赞数
	NoticeComment    int    //置顶评论数
}

const batchSize = 500

func main() {
	//go tool pprof http://localhost:8080/debug/pprof/heap 使用命令行工具抓取和查看堆剖析数据
	//导航到 http://localhost:8080/debug/pprof/。将看到多个剖析选项。点击 heap 链接
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	community := c.GDB("c").Table("t_community")
	baseInfo := o.GDB("o").Table("t_app_base_info")
	//feedsRecord := z.GDB("z").Table("t_community_feeds_record") //圈子浏览记录
	var totalData []Data

	// Fetch communities
	{
		err := community.
			Select("app_id, id as community_id, user_count, feeds_count, title").
			Where("created_at >= ? AND feeds_count > ?", "2023-01-01", 10).
			Find(&totalData).Error

		if err != nil {
			fmt.Println("Error fetching from t_community:", err)
			return
		}
		fmt.Println(totalData[0])
	}
	//批量 将appName和NewIndustryName字段写入
	for i := 0; i < len(totalData); i += batchSize {
		end := i + batchSize
		if end > len(totalData) {
			end = len(totalData)
		}
		batch := totalData[i:end]

		// Extract app IDs from the batch
		appIds := make([]string, len(batch))
		for j, data := range batch {
			appIds[j] = data.AppId
		}

		// Fetch appName and NewIndustryName for the batch of app_ids
		var infos []struct {
			AppId           string `gorm:"column:app_id"`
			AppName         string `gorm:"column:app_name"`
			NewIndustryName string `gorm:"column:new_industry_name"`
		}
		err := baseInfo.
			Select("app_id, app_name, new_industry_name").
			Where("app_id IN (?)", appIds).
			Find(&infos).Error

		if err != nil {
			fmt.Println("Error fetching batch from t_app_base_info:", err)
			continue
		}

		// Create a map for faster look-up
		infoMap := make(map[string]struct {
			AppName         string
			NewIndustryName string
		})
		for _, info := range infos {
			infoMap[info.AppId] = struct {
				AppName         string
				NewIndustryName string
			}{AppName: info.AppName, NewIndustryName: info.NewIndustryName}
		}

		// Update the Data struct slice with fetched values
		for j, data := range batch {
			if info, exists := infoMap[data.AppId]; exists {
				totalData[i+j].AppName = info.AppName
				totalData[i+j].NewIndustryName = info.NewIndustryName
			}
		}
	}
	fmt.Println(totalData[0])
}
