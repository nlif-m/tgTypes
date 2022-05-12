package tgTypes

type User struct {
	Id                          int64  `json:"id"`
	Is_bot                      bool   `json:"is_bot"`
	First_name                  string `json:"first_name"`
	Last_name                   string `json:"last_name"`
	Username                    string `json:"username"`
	Language_code               string `json:"language_code"`
	Can_join_groups             bool   `json:"can_join_groups"`
	Can_read_all_group_messages bool   `json:"can_read_all_group_messages"`
	Supports_inline_queries     bool   `json:"supports_inline_queries"`
}
