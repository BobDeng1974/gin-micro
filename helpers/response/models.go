package response

// 统一 json 结构体
type JsonObject struct {

	/** 状态码 */
	Code string          `json:"code"`

	/** 内容体 */
	Content interface{}  `json:"content"`

	/** 消息 */
	Message string        `json:"message"`
}

// 全局分页对象
type PageBean struct {

	/** 当前页  */
	Page      int          `json:"page"`

	/** 每页显示的最大行数 */
	PageSize  int          `json:"page_size"`

	/** 总记录数 */
	Total     int          `json:"total"`

	/** 每行的数据 */
	Rows      interface{}  `json:"rows"`
}


