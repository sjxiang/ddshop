package requests


type RegisterUsingEmail struct {
	Email    string  `json:"email"       binding:"required"`  
	Mobile   string  `json:"mobile"      binding:"required"`  
	NickName string  `json:"nick_name"   binding:"required"`  
	Password string  `json:"password"    binding:"required,min=3,max=20"` 
}