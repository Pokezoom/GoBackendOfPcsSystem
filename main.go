package main

import (
	"GoDockerBuild/middleware"
	"fmt"
	"github.com/xuri/excelize/v2"
	"net/http"
	_ "net/http/pprof"
	"reflect"
	"strings"
)

var c middleware.EGorm
var o middleware.EGorm
var z middleware.EGorm

// Data 最终要筛选出来的表的列名
type Data struct {
	AppId            string `xlsx:"店铺id"`
	CommunityId      string `xlsx:"圈子id"`
	Title            string `xlsx:"圈子名称"`
	AppName          string `xlsx:"店铺名称"`
	NewIndustryName  string `xlsx:"圈子行业类别"`
	UserCount        int    `xlsx:"圈子有效用户数"`
	FeedsCount       int    `xlsx:"圈子动态总数"`
	NoticeNum        int    `xlsx:"圈子置顶动态数"`
	NoticeRecordsNum int    `xlsx:"置顶动态浏览数"`
	NoticeZan        int    `xlsx:"置顶点赞数"`
	NoticeComment    int    `xlsx:"置顶评论数"`
}

const batchSize = 500 // 单次最高查询量，若要提高效率

func main() {
	//go tool pprof http://localhost:8080/debug/pprof/heap 使用命令行工具抓取和查看堆剖析数据
	//导航到 http://localhost:8080/debug/pprof/。将看到多个剖析选项。点击 heap 链接
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	community := c.GDB("c").Table("t_community")
	communityFeeds := c.GDB("c").Table("t_community_feeds")
	baseInfo := o.GDB("o").Table("t_app_base_info")
	feedsRecord := z.GDB("z").Table("t_community_feeds_record") //圈子浏览记录
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

	// 写入每条圈子的NoticeNum, NoticeZan, NoticeComment，并收集所有置顶动态的feeds_id
	allTopFeedsIds := make([]string, 0)
	for i := 0; i < len(totalData); i += batchSize {
		end := i + batchSize
		if end > len(totalData) {
			end = len(totalData)
		}
		batch := totalData[i:end]

		// Extract community IDs from the batch
		communityIds := make([]string, len(batch))
		for j, data := range batch {
			communityIds[j] = data.CommunityId
		}

		var noticeData []struct {
			CommunityId   string `gorm:"column:community_id"`
			NoticeNum     int    `gorm:"column:notice_num"`
			NoticeZan     int    `gorm:"column:notice_zan"`
			NoticeComment int    `gorm:"column:notice_comment"`
			FeedsIds      string `gorm:"column:feeds_ids"` // Changed FeedsIds to string
		}

		err := communityFeeds.
			Select("community_id, GROUP_CONCAT(id) as feeds_ids, COUNT(id) as notice_num, SUM(zan_num) as notice_zan, SUM(comment_count) as notice_comment").
			Where("community_id IN (?) AND is_notice = 1 AND feeds_state = 0", communityIds).
			Group("community_id").
			Find(&noticeData).Error

		if err != nil {
			fmt.Println("Error fetching batch notice data from t_community_feeds:", err)
			continue
		}

		// Create a map for faster look-up
		noticeMap := make(map[string]struct {
			NoticeNum     int
			NoticeZan     int
			NoticeComment int
			FeedsIds      []string
		})
		for _, notice := range noticeData {
			ids := strings.Split(notice.FeedsIds, ",") // Split comma-separated string
			noticeMap[notice.CommunityId] = struct {
				NoticeNum     int
				NoticeZan     int
				NoticeComment int
				FeedsIds      []string
			}{NoticeNum: notice.NoticeNum, NoticeZan: notice.NoticeZan, NoticeComment: notice.NoticeComment, FeedsIds: ids}
			allTopFeedsIds = append(allTopFeedsIds, ids...)
		}

		// Update the Data struct slice with fetched values
		for j, data := range batch {
			if notice, exists := noticeMap[data.CommunityId]; exists {
				totalData[i+j].NoticeNum = notice.NoticeNum
				totalData[i+j].NoticeZan = notice.NoticeZan
				totalData[i+j].NoticeComment = notice.NoticeComment
			}
		}
	}

	fmt.Println(totalData[0])
	// 写入每个圈子的置顶动态浏览总数至NoticeRecordsNum字段
	var recordData []struct {
		CommunityId string `gorm:"column:community_id"`
		RecordsNum  int    `gorm:"column:records_num"`
	}
	err := feedsRecord.
		Select("community_id, COUNT(id) as records_num").
		Where("feeds_id IN (?)", allTopFeedsIds).
		Group("community_id").
		Find(&recordData).Error
	if err != nil {
		fmt.Println("Error fetching records data from t_community_feeds_record:", err)
	}
	recordsMap := make(map[string]int)
	for _, record := range recordData {
		recordsMap[record.CommunityId] = record.RecordsNum
	}
	for _, data := range totalData {
		if recordsNum, exists := recordsMap[data.CommunityId]; exists {
			data.NoticeRecordsNum = recordsNum
		}
	}

	fmt.Println(totalData[0])
	//开始构建xlsx文件
	excle(totalData)
}

func getExcelHeadersFromStruct(data interface{}) []string {
	var headers []string
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		header := field.Tag.Get("xlsx")
		if header != "" {
			headers = append(headers, header)
		}
	}
	return headers
}

func excle(totalData []Data) {

	f := excelize.NewFile()

	// 写列名
	headers := getExcelHeadersFromStruct(Data{})
	for i, header := range headers {
		col := toAlphaString(i) + "1"
		f.SetCellValue("Sheet1", col, header)
	}

	// Add data to the Excel file
	for i, data := range totalData {
		row := i + 2 // Start data at row 2, since row 1 has the column names
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), data.AppId)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), data.CommunityId)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), data.Title)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), data.AppName)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), data.NewIndustryName)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), data.UserCount)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), data.FeedsCount)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), data.NoticeNum)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", row), data.NoticeRecordsNum)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", row), data.NoticeZan)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", row), data.NoticeComment)
	}

	// Save the Excel file
	if err := f.SaveAs("totalData.xlsx"); err != nil {
		fmt.Println(err)
	}
}
func toAlphaString(i int) string {
	return string(rune('A' + i))
}
