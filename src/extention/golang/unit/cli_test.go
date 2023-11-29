package gunit

// func testCommand(t *testing.T, cmd string) {

// 	tmp := t.TempDir()

// 	cfg := config.GetDefaultConfiguration()
// 	cmdline := cli.New(cfg)
// 	client := cmdline.GetClient()
// 	cc := cli.Bind(nil)

// 	task := &cli.Task{
// 		Context: cc,
// 		Command: cmd,
// 		WD:      tmp,
// 	}
// 	err := client.Run(task)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestCommandHelp(t *testing.T) {
// 	testCommand(t, "help")
// 	testCommand(t, "help pwd")
// }

// func TestCommandCD(t *testing.T) {
// 	testCommand(t, "cd ..")
// }

// func TestCommandLS(t *testing.T) {
// 	testCommand(t, "ls")
// }

// func TestCommandMkdir(t *testing.T) {
// 	testCommand(t, "mkdir foo")
// }

// func TestCommandPwd(t *testing.T) {
// 	testCommand(t, "pwd")
// }

// func TestCommandSleep(t *testing.T) {
// 	t0 := time.Now()
// 	testCommand(t, "sleep 3000")
// 	t1 := time.Now()
// 	fmt.Println("t0=", t0)
// 	fmt.Println("t1=", t1)
// }
