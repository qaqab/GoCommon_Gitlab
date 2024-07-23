package GoCommon_Gitlab

import (
	"testing"

	"github.com/qaqab/GoCommon_DbManager"
	"github.com/qaqab/GoCommon_Gitlab"
)

func TestGitlabClone(t *testing.T) {
	clientAll := GoCommon_DbManager.ClientAll{ConfigSetting: struct {
		ConfigPath string
		ConfigName string
	}{ConfigPath: "/root/go-project/EasyCommon/YamlFile/", ConfigName: "test"}}

	clientAll.DbManagerClient("gitlab")
	git_Client := clientAll.GitClient

	// 创建一个GitlabSetting对象，并设置Token和Url
	gitlabSetting := GoCommon_Gitlab.GitlabSetting{Token: clientAll.GitlabSettingData.Token, GithubUrl: clientAll.GitlabSettingData.GithubUrl, Username: clientAll.GitlabSettingData.Username, Password: clientAll.GitlabSettingData.Password}

	// 设置 Gitlab 客户端
	gitlabSetting.GitClient = git_Client

	// 设置本地路径
	localPath := "/root/gitlab"

	// 获取所有项目
	// 获取项目
	project, _ := gitlabSetting.GetAllProject()

	// 遍历项目列表
	for _, v := range project {
		// 获取项目的所有分支
		branch_list := gitlabSetting.GetAllBranch(v.ID)

		// 遍历分支列表
		for _, branchName := range branch_list {
			// 克隆项目的指定分支到本地路径
			gitlabSetting.CloneRepoBranch(v.WebURL, branchName, localPath+"/"+v.Name+"/"+branchName)
		}
	}
}
