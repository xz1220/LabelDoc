package model

type Task struct {
	TaskID      int64  `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:task_id" form:"task_id" json:"taskId"`
	TaskName    string `gorm:"type:varchar(50);column:task_name" form:"task_name" json:"taskName"`
	TaskDesc    string `gorm:"type:varchar(1024);column:task_desc" form:"task_desc" json:"taskDesc"`
	ImageNumber int64  `gorm:"type:int(64);column:image_number" form:"image_number" json:"imageNumber"`
	IsCreated   int64  `gorm:"column:is_created" form:"is_created"`
	TaskType    int64  `gorm:"column:task_type" form:"task_type"`
}

//TableName reset the Name field
func (Task) TableName() string {
	return "task"
}

type TaskList struct {
	TaskID      int64  `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:task_id" form:"task_id" json:"taskId"`
	TaskName    string `gorm:"type:varchar(50);column:task_name" form:"task_name" json:"taskName"`
	TaskDesc    string `gorm:"type:varchar(1024);column:task_desc" form:"task_desc" json:"taskDesc"`
	ImageNumber int64  `gorm:"type:int(64);column:image_number" form:"image_number" json:"imageNumber"`
	TaskType    int64  `gorm:"column:task_type" form:"task_type" json:"taskType"`
	Finish      int64  `gorm:"column:finish" form:"finish" json:"finish"`
}

type Tasklabelinfo struct {
	TaskID  int64 `gorm:"column:task_id" form:"task_id"`
	LabelID int64 `gorm:"column:label_id" form:"label_id"`
}

//TableName reset the Name field
func (Tasklabelinfo) TableName() string {
	return "tasklabelinfo"
}

type TaskReviewerInfo struct {
	TaskID     int64 `gorm:"column:task_id" form:"task_id"`
	ReviewerID int64 `gorm:"column:reviewer_id" form:"reviewer_id"`
}

//TableName reset the Name field
func (TaskReviewerInfo) TableName() string {
	return "taskreviewerinfo"
}

type TaskUserInfo struct {
	TaskID int64 `gorm:"column:task_id" form:"task_id"`
	UserID int64 `gorm:"column:user_id" form:"user_id"`
}

//TableName reset the Name field
func (TaskUserInfo) TableName() string {
	return "taskuserinfo"
}

type TaskAdminInfo struct {
	TaskID int64 `gorm:"column:task_id" form:"task_id"`
	UserID int64 `gorm:"column:admin_id" form:"admin_id"`
}

//TableName reset the Name field
func (TaskAdminInfo) TableName() string {
	return "taskadmininfo"
}

type Test struct {
	ID       int    `gorm:"AUTO_INCREMENT;primary_key;column:id"`
	Test     string `gorm:"type:varchar(255);column:test"`
	TestLong string `gorm:"type:text;column:test_long"`
}

//TableName reset the Name field
func (Test) TableName() string {
	return "test"
}
