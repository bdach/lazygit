package pull

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var SSHWithU2FAuthenticator = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Pulls changes from an origin repository, using an SSH key backed by an U2F authenticator",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.CreateFile("myfile1", "myfile1 content")
		shell.GitAddAll().Commit("commit 1")
		shell.CreateFile("myfile2", "myfile2 content")
		shell.GitAddAll().Commit("commit 2")
		shell.CreateFile("myfile3", "myfile3 content")
		shell.GitAddAll().Commit("commit 3")

		shell.SetLocalConfig("core.askpass", "../../mock_u2f_pin_prompt.sh")
		shell.RunCommand("git clone --bare . ../origin")

		shell.HardReset("HEAD~2")
		shell.AddRemote("origin", "../origin")
		shell.Fetch("origin")
		shell.SetUpstream("master", "origin/master")
	},
	Run: func(shell *Shell, input *Input, assert *Assert, keys config.KeybindingConfig) {
		assert.CommitCount(1)

		input.PressKeys(keys.Universal.PullFiles)
		assert.CommitCount(3)
		assert.MatchHeadCommitMessage(Equals("commit 3"))
	},
})
