package userService

// RegisterParams TODO 验证带完善
type RegisterParams struct {
	VerificationType int64  `json:"type" binding:"required,numeric,gt=0" msg:"注册类型错误"`
	AreaCode         int64  `json:"area_code" binding:"numeric,gt=0" msg:"请选择正确的区域代码"`
	Mobile           int64  `json:"mobile" binding:"numeric,gt=0" msg:"请输入正确的手机号"`
	Email            string `json:"email" binding:"-" msg:"邮箱名不能为空"`
	Password         string `json:"password" binding:"-" msg:"请输入正确格式的密码"`
	Code             int64  `json:"code" binding:"required,numeric,gt=0" msg:"请输入正确的验证码"`
}

// VerificationCodeParams TODO 验证带完善
type VerificationCodeParams struct {
	VerificationType int64  `json:"type" binding:"required,numeric,gt=0" msg:"错误的验证码类型"`
	AreaCode         int64  `json:"area_code" binding:"numeric,gt=0" msg:"请选择正确的区域代码"`
	Mobile           int64  `json:"mobile" binding:"numeric,gt=0" msg:"请输入正确的电话号码"`
	Email            string `json:"email" binding:"email" msg:"请输入正确的邮箱地址"`
}

type LoginCodeParams struct {
	VerificationType int64  `json:"type" binding:"required,numeric,gt=0" msg:"错误的登录类型"`
	AreaCode         int64  `json:"area_code" binding:"numeric,gt=0" msg:"请选择正确的区域代码"`
	Mobile           int64  `json:"mobile" binding:"numeric,gt=0" msg:"请输入正确的电话号码"`
	Email            string `json:"email" binding:"email" msg:"请输入正确的邮箱地址"`
}

type LoginParams struct {
	VerificationType int64  `json:"type" binding:"required,numeric,gt=0" msg:"错误的登录类型"`
	AreaCode         int64  `json:"area_code" binding:"numeric,gt=0" msg:"请选择正确的区域代码"`
	Mobile           int64  `json:"mobile" binding:"numeric,gt=0" msg:"请输入正确的电话号码"`
	Email            string `json:"email" binding:"email" msg:"请输入正确的邮箱地址"`
	Code             int64  `json:"code" binding:"required,numeric,gt=0" msg:"请输入正确的验证码"`
}

type JWTClaims struct {
	ID       int64
	AreaCode int64
	Mobile   int64
	Email    string
}
