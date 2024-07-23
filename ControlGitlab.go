package GoCommon_Gitlab

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	gitlab "github.com/xanzy/go-gitlab"
)

type Project struct {
	WebURL string `json:"web_url"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}
type GitlabSetting struct {
	Token     string         `json:"token"`
	GithubUrl string         `json:"url"`
	GitClient *gitlab.Client `json:"client"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
}

func (gitlabSetting GitlabSetting) GetAllProject() ([]*Project, error) {
	// 创建Gitlab分页查询项目列表的选项对象
	lbo := &gitlab.ListProjectsOptions{ListOptions: gitlab.ListOptions{Page: 1, PerPage: 100}}
	var pro []*Project
	for {
		// 使用Gitlab客户端的Projects方法分页查询项目列表
		p, _, err := gitlabSetting.GitClient.Projects.ListProjects(lbo)
		if err != nil {
			fmt.Printf("list projects failed:%v\n", err)
			return nil, err
		}
		// 遍历查询到的项目列表，并将项目信息转换为Project对象后添加到pro切片中
		for _, v := range p {
			pro = append(pro, &Project{WebURL: v.WebURL, Name: v.Name, ID: v.ID})
		}

		// 如果查询到的项目数量小于50，则跳出循环
		if len(p) < 50 {
			break
		}
		// 更新分页查询选项的页码，以便下一轮查询
		lbo.ListOptions.Page++
	}
	// 返回查询到的项目列表
	return pro, nil
}

func (gitlabSetting GitlabSetting) GetAllBranch(projectId int) []string {
	// 调用GitlabClient的Branches.ListBranches方法获取项目下的所有分支
	branches, _, err := gitlabSetting.GitClient.Branches.ListBranches(projectId, &gitlab.ListBranchesOptions{
		ListOptions: gitlab.ListOptions{PerPage: 100},
	})
	if err != nil {
		panic(err)
	}

	// 创建一个空字符串切片用于存储分支名称
	projectName_list := []string{}

	// 遍历每个分支
	// 为每个分支创建一个克隆
	for _, branch := range branches {
		// 将分支名称添加到切片中
		projectName_list = append(projectName_list, branch.Name)
	}

	// 返回分支名称切片
	return projectName_list
}
func (gitlabSetting GitlabSetting) CloneRepoBranch(repoURL, branchName, localPath string) {
	fmt.Println("开始克隆仓库", repoURL, "分支", branchName, "到", localPath)
	// 克隆特定的分支
	// Clone the specific branch
	_, err := git.PlainClone(localPath, false, &git.CloneOptions{
		// 仓库地址
		URL: repoURL,
		// 分支名称
		// Fixed: use plumbing.NewBranchReferenceName instead of reference.NewBranch
		ReferenceName: plumbing.NewBranchReferenceName(branchName), // 使用 plumbing.NewBranchReferenceName 替代 reference.NewBranch
		// 只克隆单个分支
		SingleBranch: true,
		// 认证信息
		Auth: &http.BasicAuth{
			// 用户名
			Username: gitlabSetting.Username,
			// 密码
			Password: gitlabSetting.Password,
		},
	})

	if err != nil {
		// 如果克隆失败，则抛出异常
		panic(err)
	} else {
		fmt.Println("克隆仓库成功")
	}
}
