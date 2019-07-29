package models

import (
	"time"

	"github.com/astaxie/beego"
	// beego的orm模块
	"github.com/astaxie/beego/orm"
	// Go语言的sql驱动,否则无法操作数据库
	_ "github.com/go-sql-driver/mysql"

	"go-microservice/utils"
)

/* 用户 table_name = user */
type User struct {
	Id int `json:"user_id"`
	// 用户昵称
	Name string `orm:"size(32)"  json:"name"`
	// 用户密码加密
	Password_hash string `orm:"size(128)" json:"password"`
	// 手机号做登录
	Mobile string `orm:"size(11);unique"  json:"mobile"`
	// 真实姓名 实名认证
	Real_name string `orm:"size(32)" json:"real_name"`
	// 身份证号 实名认证
	Id_card string `orm:"size(20)" json:"id_card"`
	// 用户头像路径 通过fastdfs进行图片存储
	Avatar_url string `orm:"size(256)" json:"avatar_url"`
	// 1:N 用户发布的房屋信息 一个人多套房
	Houses []*House `orm:"reverse(many)" json:"houses"`
	// 1:N 用户下的订单 一个人多次订单
	Orders []*OrderHouse `orm:"reverse(many)" json:"orders"`
}

/* 房屋信息 table_name = house */
type House struct {
	// 房屋编号
	Id int `json:"house_id"`
	// 房屋主人的用户编号 与用户进行关联
	User *User `orm:"rel(fk)" json:"user_id"`
	// 归属地的区域编号 和 地区表进行关联
	Area *Area `orm:"rel(fk)" json:"area_id"`
	// 房屋标题
	Title string `orm:"size(64)" json:"title"`
	// 单价,单位:分 每次的价格要乘以100
	// TODO - 金融一般分为单位,元为单位 浮点数 数值转换丢失精度
	Price int `orm:"default(0)" json:"price"`
	// 地址
	Address string `orm:"size(512)" orm:"default("")" json:"address"`
	// 房间数目
	Room_count int `orm:"default(1)" json:"room_count"`
	// 房屋总面积
	Acreage int `orm:"default(0)" json:"acreage"`
	// 房屋单元,如 几室几厅
	Unit string `orm:"size(32)" orm:"default("")" json:"unit"`
	// 房屋容纳的总人数
	Capacity int `orm:"default(1)" json:"capacity"`
	// 房屋床铺的配置
	Beds string `orm:"size(64)" orm:"default("")" json:"beds"`
	// 押金
	Deposit int `orm:"default(0)" json:"deposit"`
	// 最少入住的天数
	Min_days int `orm:"default(1)" json:"min_days"`
	// 最多入住的天数 0表示不限制
	Max_days int `orm:"default(0)" json:"max_days"`
	// 预定完成的该房屋的订单数
	Order_count int `orm:"default(0)" json:"order_count"`
	// 房屋主图片路径
	Index_image_url string `orm:"size(256)" orm:"default("")" json:"index_image_url"`
	// 1:N 房屋设施 与 设施表进行关联
	Facilities []*Facility `orm:"reverse(many)" json:"facilities"`
	// 1:N 房屋的图片 除主要图片之外的其他图片地址
	Images []*HouseImage `orm:"reverse(many)" json:"img_urls"`
	// 1:N 房屋的订单 与 房订单表进行管理 历史订单记录
	Orders []*OrderHouse `orm:"reverse(many)" json:"orders"`
	Ctime  time.Time     `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

// 首页最多展示的房屋数量
var HOME_PAGE_MAX_HOUSES int = 5

// 房屋列表页面每页显示条目数
var HOUSE_LIST_PAGE_CAPACITY int = 2

// 处理房子信息
func (this *House) To_house_info() interface{} {
	house_info := map[string]interface{}{
		"house_id":    this.Id,
		"title":       this.Title,
		"price":       this.Price,
		"area_name":   this.Area.Name,
		"img_url":     utils.AddDomain2Url(this.Index_image_url),
		"room_count":  this.Room_count,
		"order_count": this.Order_count,
		"address":     this.Address,
		"user_avatar": utils.AddDomain2Url(this.User.Avatar_url),
		"ctime":       this.Ctime.Format("2006-01-02 15:04:05"),
	}

	return house_info
}

// 处理1个房子的全部信息
func (this *House) To_one_house_desc() interface{} {
	house_desc := map[string]interface{}{
		"hid":         this.Id,
		"user_id":     this.User.Id,
		"user_name":   this.User.Name,
		"user_avatar": utils.AddDomain2Url(this.User.Avatar_url),
		"title":       this.Title,
		"price":       this.Price,
		"address":     this.Address,
		"room_count":  this.Room_count,
		"acreage":     this.Acreage,
		"unit":        this.Unit,
		"capacity":    this.Capacity,
		"beds":        this.Beds,
		"deposit":     this.Deposit,
		"min_days":    this.Min_days,
		"max_days":    this.Max_days,
	}

	// 房屋图片
	img_urls := []string{}
	for _, img_url := range this.Images {
		img_urls = append(img_urls, utils.AddDomain2Url(img_url.Url))
	}
	house_desc["img_urls"] = img_urls

	// 房屋设施
	facilities := []int{}
	for _, facility := range this.Facilities {
		facilities = append(facilities, facility.Id)
	}
	house_desc["facilities"] = facilities

	// 评论信息
	comments := []interface{}{}
	orders := []OrderHouse{}
	o := orm.NewOrm()
	order_num, err := o.QueryTable("order_house").
		Filter("house__id", this.Id).
		Filter("status", ORDER_STATUS_COMPLETE).
		OrderBy("-ctime").
		Limit(10).
		All(&orders)
	if err != nil {
		beego.Error("select orders comments error, err =", err, "house id = ", this.Id)
	}
	for i := 0; i < int(order_num); i++ {
		o.LoadRelated(&orders[i], "User")
		var username string
		if orders[i].User.Name == "" {
			username = "匿名用户"
		} else {
			username = orders[i].User.Name
		}

		comment := map[string]string{
			"comment":   orders[i].Comment,
			"user_name": username,
			"ctime":     orders[i].Ctime.Format("2006-01-02 15:04:05"),
		}
		comments = append(comments, comment)
	}
	house_desc["comments"] = comments

	return house_desc
}

/* 区域信息 table_name = area */
// TODO - 区域信息是需要我们手动添加到数据库中的
type Area struct {
	// 区域编号     1    2
	Id int `json:"aid"`
	// 区域名字     昌平 海淀
	Name string `orm:"size(32)" json:"aname"`
	// 区域所有的房屋 与 房屋表进行关联
	Houses []*House `orm:"reverse(many)" json:"houses"`
}

/* 设施信息 table_name = "facility"*/
// TODO - 设施信息 需要我们提前手动添加的
type Facility struct {
	// 设施编号
	Id int `json:"fid"`
	// 设施名字
	Name string `orm:"size(32)"`
	// 都有哪些房屋有此设施 与 房屋表进行关联
	Houses []*House `orm:"rel(m2m)"`
}

/* 房屋图片 table_name = "house_image"*/
type HouseImage struct {
	// 图片id
	Id int `json:"house_image_id"`
	// 图片url 存放我们房屋的图片
	Url string `orm:"size(256)" json:"url"`
	// 图片所属房屋编号
	House *House `orm:"rel(fk)" json:"house_id"`
}

const (
	ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"  //待接单
	ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT" //待支付
	ORDER_STATUS_PAID         = "PAID"         //已支付
	ORDER_STATUS_WAIT_COMMENT = "WAIT_COMMENT" //待评价
	ORDER_STATUS_COMPLETE     = "COMPLETE"     //已完成
	ORDER_STATUS_CANCELED     = "CONCELED"     //已取消
	ORDER_STATUS_REJECTED     = "REJECTED"     //已拒单
)

/* 订单 table_name = order */
type OrderHouse struct {
	// 订单编号
	Id int `json:"order_id"`
	// 下单的用户编号 与 用户表进行关联
	User *User `orm:"rel(fk)" json:"user_id"`
	// 预定的房间编号 与 房屋信息进行关联
	House *House `orm:"rel(fk)" json:"house_id"`
	// 预定的起始时间
	Begin_date time.Time `orm:"type(datetime)"`
	// 预定的结束时间
	End_date time.Time `orm:"type(datetime)"`
	// 预定总天数
	Days int
	// 房屋的单价
	House_price int
	// 订单总金额
	Amount int
	// 订单状态
	Status string `orm:"default(WAIT_ACCEPT)"`
	// 订单评论
	Comment string `orm:"size(512)"`
	// TODO - 每次更新此表都会更新这个字段
	Ctime time.Time `orm:"auto_now;type(datetime)" json:"ctime"`
	// 表示个人征信情况 true表示良好
	Credit bool
}

// 处理订单信息
func (this *OrderHouse) To_order_info() interface{} {
	order_info := map[string]interface{}{
		"order_id":   this.Id,
		"title":      this.House.Title,
		"img_url":    utils.AddDomain2Url(this.House.Index_image_url),
		"start_date": this.Begin_date.Format("2006-01-02 15:04:05"),
		"end_date":   this.End_date.Format("2006-01-02 15:04:05"),
		"ctime":      this.Ctime.Format("2006-01-02 15:04:05"),
		"days":       this.Days,
		"amount":     this.Amount,
		"status":     this.Status,
		"comment":    this.Comment,
		"credit":     this.Credit,
	}

	return order_info
}

// 数据库的初始化
func init() {
	// 调用什么驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	// 连接数据(默认参数,mysql数据库,"数据库用户名:数据库密码@tcp("+数据库地址+":"+数据库端口+")/库名?格式",默认参数）
	orm.RegisterDataBase("default",
		"mysql",
		"root:root@tcp("+utils.G_mysql_addr+":"+utils.G_mysql_port+")/rent_house?charset=utf8",
		30)

	// 注册 model 建表
	// 用户表 房屋表 地区表 设施表 房屋图片表 房屋订单表
	orm.RegisterModel(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))

	// create table
	// 参数1 别名,默认参数
	// 参数2 是否强制替换模块
	// TODO - 如果表变更,把false改为true,重新执行这段代码,就把数据库中的表重新创建了,之前数据丢失
	// 参数3 如果没有则同步或创建
	// 如原有3张表,现在是6张表,会把没有的3张表创建出来
	orm.RunSyncdb("default", false, true)
}
