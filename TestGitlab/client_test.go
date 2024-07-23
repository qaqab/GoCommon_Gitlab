package GoCommon_Gitlab

import (
	"fmt"
	"testing"

	"github.com/qaqab/GoCommon_DbManager"
)

func TestGitlabClient(t *testing.T) {
	clientAll := GoCommon_DbManager.ClientAll{ConfigSetting: struct {
		ConfigPath string
		ConfigName string
	}{ConfigPath: "/root/go-project/EasyCommon/YamlFile/", ConfigName: "test"}}
	clientAll.DbManagerClient("gitlab")
	git_Client := clientAll.GitClient

	// 打印Gitlab客户端信息
	fmt.Println(git_Client)
	fmt.Println(clientAll.GitlabSettingData.GithubUrl)
	fmt.Println(clientAll.GitlabSettingData.Token)
	fmt.Println(clientAll.GitlabSettingData.Username)
	fmt.Println(clientAll.GitlabSettingData.Password)

}
