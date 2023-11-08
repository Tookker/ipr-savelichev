package task

type Task struct {
	Id          uint   `gorm:"column:id"`
	Descryption string `gorm:"column:descryption"`
	EmployeID   uint   `gorm:"column:employeid"`
	ToolID      uint   `gorm:"column:toolid"`
}

type EmployeTask struct {
	Id              uint   `gorm:"column:id"`
	DescryptionTask string `gorm:"column:descryptiontask"`
	Name            string `gorm:"column:name"`
	Age             string `gorm:"column:age"`
	Sex             string `gorm:"column:sex"`
	DescryptionTool string `gorm:"column:descryptiontool"`
}
