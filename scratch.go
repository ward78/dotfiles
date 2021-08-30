package main

// 	system, err := newSystem()
// 	if err != nil {
// 		return err
// 	}
// 	err = backup(system)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func newSystem() (*system, error) {
// 	if runtime.GOOS != "linux" {
// 		return nil, errors.New("script requires linux OS")
// 	}
// 	user, err := user.Current()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if user.Name != "root" {
// 		return nil, errors.New("root/sudo access is required")
// 	}
// 	dirScript, err := os.Getwd()
// 	if err != nil {
// 		return nil, err
// 	}
// 	sudoUser := os.Getenv("SUDO_USER")
// 	dirHome := filepath.Join("/home", sudoUser)
// 	if !strings.HasPrefix(dirScript, dirHome) {
// 		return nil, errors.New(dirScript + " doesn't reside in user's home directory " + dirHome)
// 	}
// 	sudoUID, err := strconv.Atoi(os.Getenv("SUDO_UID"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	sudoGID, err := strconv.Atoi(os.Getenv("SUDO_GID"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	s := &system{
// 		user: sudoUser,
// 		uid:  sudoUID,
// 		gid:  sudoGID,
// 		dir: map[string]string{
// 			"home":     dirHome,
// 			"dotfiles": dirScript,
// 			"backup":   filepath.Join(dirScript, "backup"),
// 		},
// 	}
// 	return s, nil
// }

// func backup(system *system) error {
// 	dirHome := system.dir["home"]
// 	h, err := os.Open(dirHome)
// 	if err != nil {
// 		return err
// 	}
// 	defer h.Close()
// 	files, err := h.Readdir(-1)
// 	if err != nil {
// 		return err
// 	}
// 	var dotfiles []string
// 	for _, file := range files {
// 		filename := file.Name()
// 		if filename[0:1] == "." {
// 			dotfiles = append(dotfiles, filename)
// 		}
// 	}
// 	if len(dotfiles) != 0 {
// 		dirBackup := system.dir["backup"]
// 		_, err := os.Stat(dirBackup)
// 		if errors.Is(err, os.ErrNotExist) {
// 			err = os.Mkdir(dirBackup, os.ModePerm)
// 			if err != nil {
// 				return err
// 			}
// 			err = os.Chown(dirBackup, system.uid, system.gid)
// 			if err != nil {
// 				return err
// 			}

// 		}
// 		timeStamp := time.Now().Format("2006-01-02_030405")
// 		dirTemp := filepath.Join(dirBackup, timeStamp)
// 		err = os.Mkdir(dirTemp, os.ModePerm)
// 		if err != nil {
// 			return err
// 		}
// 		err = os.Chown(dirTemp, system.uid, system.gid)
// 		if err != nil {
// 			return err
// 		}
// 		src := filepath.Join(dirHome, filename)
// 		dst := filepath.Join(dirTemp, filename)
// 		err = os.Rename(src, dst)
// 	}
// 	return nil
// }
//  err = os.Remove(sourcePath)
//  defer os.RemoveAll("subdir")
//  err = os.Chmod("temp.txt", 0777)

//  err = os.Symlink(config, "/home/ctw/.config")
//  statDotfiles, err := os.Stat(dirDotfiles)
//  err = os.Chdir("subdir/parent/child")

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
