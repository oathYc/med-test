package builder

// builder标准
type IBuilder interface {
	// 添加
	Add(value interface{}) error
	// 更新
	Update(attrs ...interface{}) error
	// 删除
	Delete(value interface{}, where ...interface{}) error
	// 查询单条
	Find(out interface{}, where ...interface{}) error
	// 查询多条
	Get(out interface{}, where ...interface{}) error
	// 统计
	Count(value interface{}) error
	// 查询且条件
	Where(query interface{}, args ...interface{}) IBuilder
	// 查询或条件
	Or(query interface{}, args ...interface{}) IBuilder
	// 查询定位
	Offset(offset interface{}) IBuilder
	// 查询区间
	Limit(limit interface{}) IBuilder
	// 排序
	Order(value interface{}, reorder ...bool) IBuilder
	// 事务开始
	Begin() (IBuilder, error)
	// 事务回滚
	Rollback() error
	// 事务提交
	Commit() error
	// 模型
	Model(value interface{}) IBuilder
	// 检查数据是否为空
	IsEmpty(e error) bool
	// 检查是否为写库
	IsWrite() bool
	// 日志
	SetLogger(log logger) IBuilder
}

// 日志标准
type logger interface {
	Print(v ...interface{})
}
