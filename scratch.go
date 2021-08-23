package main


// home := filepath.Join(system.home, ".*")

// 	dotfiles, err := filepath.Glob(home)
// 	if err != nil {
// 		return err
// 	}
// 	if len(dotfiles) == 0 {
// 		return nil
// 	}
// 	t := time.Now()
// 	timeStamp := t.Format("2006-01-02_03:04:05")
// 	folder := "backup-" + timeStamp
// 	backup := filepath.Join(system.backup, folder)
// 	err = os.MkdirAll(backup, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}
// 	err = os.Chown(backup, system.uid, system.gid)
// 	if err != nil {
// 		return err
// 	}
// 	for _, dotfile := range dotfiles {
// 		fmt.Println(dotfile)
// 	}
//  if _, err := os.Stat("/path/to/your-file"); os.IsNotExist(err) {
//
//  	os.MkdirAll("/path/to", 0700) // Create your file
//
//  }
//  src := filepath.Join(homeDir, filename)
//  dst := filepath.Join(backupDir, filename)
//  err = os.Rename(src, dst)
//  err = os.Remove(sourcePath)
//  err = os.Chmod("temp.txt", 0777)
//  config := filepath.Join(pwd, "config")
//  err = os.Symlink(config, "/home/ctw/.config")
//  statDotfiles, err := os.Stat(dirDotfiles)
//  fmt.Printf("Permission File Before: %s\n", stats.Mode())
//  err = os.Chmod("new.txt", 0700)
//  err = os.Chown("test.txt", os.Getuid(), os.Getgid())
//  if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
//  err := os.Mkdir(path, os.ModePerm)
//  path := "sample/deeply/nested/directory"
//  err := os.MkdirAll(path, os.ModePerm)
//  defer os.RemoveAll("subdir")
//  err = os.Chdir("subdir/parent/child")
//  if _, err := os.Stat(path); os.IsNotExist(err) {
//  os.Mkdir(path, 0755)
//  fmt.Println("Directory created")
//  } else {
//  fmt.Println("Directory already exists")
//  }
//  }
//  oldpath := "tmp"
//  newpath := "tmp2"
//  err := os.Rename(oldpath, newpath)
//  err := os.Remove("tmp")
//  err := os.RemoveAll("tmp")
//  case "nvim":
//  	nvimPath := filepath.Join(configPath, name)
//  	err = os.Symlink(nvimPath, zshSymlink)
//  case "git":
//  	gitPath := filepath.Join(configPath, name)
//  	err = os.Symlink(gitPath, zshSymlink)
//  for _, config := range configs {
//  if config.IsDir() {
//  _, err := exec.LookPath(config.Name())
//  fmt.Println(config.Name(), "is not installed")
//  cmd := exec.Command("firefox")
//  err := cmd.Run()
//  out, err := exec.Command("nvim").Output()
//  out, err := exec.Command("ls", "-l").Output()
//  cmd.Stdout = os.Stdout
//  cmd.Stderr = os.Stderr
//  cmd := exec.Command("ls")
//  err := cmd.Run()

//  Meraki Key
//  f1d58ffb63c5532bfbabd0950a45a2434669128b
//  Commands to get OS Info
//  cat /etc/os-release
//  lsb_release -a
//  hostnamectl requires systemd
//  uname -r | grep wsl
//  Install fira code font
//  sudo apt install fonts-firacode
//  Install Spacechip Font
//  Start the ssh-agent in the background.
//  $ eval "$(ssh-agent -s)"
//  > Agent pid 59566
//  Depending on your environment, you may need to use a different command. For example, you may need to use root access by running sudo -s -H before starting the ssh-agent, or you may need to use exec ssh-agent bash or exec ssh-agent zsh to run the ssh-agent.
//  Add your SSH private key to the ssh-agent. If you created your key with a different name, or if you are adding an existing key that has a different name, replace id_ed25519 in the command with the name of your private key file.
//  $ ssh-add ~/.ssh/id_ed25519
