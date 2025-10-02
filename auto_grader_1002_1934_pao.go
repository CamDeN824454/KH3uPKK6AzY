// 代码生成时间: 2025-10-02 19:34:58
package main

import (
    "fmt"
    "strings"
    "beego"
    "encoding/json"
)

// 定义题目结构体
type Question struct {
    ID        int    "json:\"id\""
    Title     string "json:\"title\""
    Answer    string "json:\"answer\""
    InputType string "json:\"input_type\""
}

// 定义答案结构体
type Answer struct {
    ID     int    "json:\"id\""
    UserID int    "json:\"user_id\""
    Answer string "json:\"answer\""
}

// 定义批改结果结构体
type GradingResult struct {
    QuestionID int    "json:\"question_id\""
    IsCorrect  bool   "json:\"is_correct\""
    Points     int    "json:\"points\""
}

// GradingService 批改服务
type GradingService struct {
    // 可以添加更多字段，如题目库、用户信息等
}

// NewGradingService 创建批改服务实例
func NewGradingService() *GradingService {
    return &GradingService{}
}

// Grade 批改答案
func (s *GradingService) Grade(q *Question, a *Answer) (*GradingResult, error) {
    // 检查题目ID和用户ID是否匹配
    if q.ID != a.ID {
        return nil, fmt.Errorf("题目ID和答案ID不匹配")
    }

    // 将题目和答案都转换为小写，以便进行大小写不敏感的比较
    lowerAnswer := strings.ToLower(a.Answer)
    lowerCorrectAnswer := strings.ToLower(q.Answer)

    // 根据题目类型进行批改
    switch q.InputType {
    case "exact":
        // 完全匹配
        if lowerAnswer == lowerCorrectAnswer {
            return &GradingResult{
                QuestionID: q.ID,
                IsCorrect:  true,
                Points:     10, // 假设每题10分
            }, nil
        } else {
            return &GradingResult{
                QuestionID: q.ID,
                IsCorrect:  false,
                Points:     0,
            }, nil
        }
    case "regex":
        // 正则匹配
        if matched, err := regexp.MatchString(q.Answer, a.Answer); matched && err == nil {
            return &GradingResult{
                QuestionID: q.ID,
                IsCorrect:  true,
                Points:     10,
            }, nil
        } else {
            return &GradingResult{
                QuestionID: q.ID,
                IsCorrect:  false,
                Points:     0,
            }, nil
        }
    default:
        // 未知题目类型
        return nil, fmt.Errorf("未知题目类型: %s", q.InputType)
    }
}

// main 函数
func main() {
    beego.Router("/grade", &GradingController{})
    beego.Run()
}

// GradingController 批改控制器
type GradingController struct {
    beego.Controller
}

// Post 处理POST请求，进行批改
func (c *GradingController) Post() {
    var q Question
    var a Answer
    err := json.Unmarshal(c.Ctx.Input.RequestBody, &q)
    if err != nil {
        c.CustomAbort(400, fmt.Sprintf("解析题目失败: %v", err))
        return
    }

    err = json.Unmarshal(c.Ctx.Input.RequestBody, &a)
    if err != nil {
        c.CustomAbort(400, fmt.Sprintf("解析答案失败: %v", err))
        return
    }

    service := NewGradingService()
    result, err := service.Grade(&q, &a)
    if err != nil {
        c.CustomAbort(500, fmt.Sprintf("批改失败: %v", err))
        return
    }

    c.Data["json"] = result
    c.ServeJSON()
}
