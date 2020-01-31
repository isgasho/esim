package new

func ModelInit() {
	fc1 := &FileContent{
		FileName: "user.go",
		Dir:      "internal/domain/entity",
		Content: `package entity

type User struct {

	//ID
	ID int {{!}}gorm:"column:id;primary_key"{{!}}

	//优惠券号码
	UserNmae string {{!}}gorm:"column:user_name"{{!}}

	//密码
	PassWord string {{!}}gorm:"column:pass_word"{{!}}
}
`,
	}

	fc2 := &FileContent{
		FileName: "README.md",
		Dir:      "internal/domain/service",
		Content:  `domain service`,
	}

	Files = append(Files, fc1, fc2)
}
