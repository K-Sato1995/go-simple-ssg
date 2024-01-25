package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a project",
	Long:  "Initialize a project",
	Run:   initialize,
}

func init() {
	RootCmd.AddCommand(initCmd)
}

const BASE_PROJECT_PATH = "example"
const REPO_URL = "https://github.com/K-Sato1995/go-simple-ssg"

func initialize(cmd *cobra.Command, args []string) {
	fmt.Println("Initializing project...")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a directory name for the project (default is 'myProject'): ")
	dirName, _ := reader.ReadString('\n')
	dirName = strings.TrimSpace(dirName)
	if dirName == "" {
		dirName = "myProject" // default directory name
	}

	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating project directory:", err)
			return
		}
	}
	tempDir, err := ioutil.TempDir("", "repoClone")
	if err != nil {
		fmt.Println("Error creating a temporary directory:", err)
		return
	}
	defer os.RemoveAll(tempDir) // Clean up

	err = gitClone(REPO_URL, tempDir)
	if err != nil {
		fmt.Println("Error cloning repository:", err)
		return
	}

	baseProjectPath := filepath.Join(tempDir, BASE_PROJECT_PATH)
	projectDir := filepath.Join(".", dirName)

	err = copyDir(baseProjectPath, projectDir)
	if err != nil {
		fmt.Println("Error copying example:", err)
		return
	}

	fmt.Println("Project initialized successfully in directory:", dirName)
}

func gitClone(repoURL, directoryName string) error {
	cmd := exec.Command("git", "clone", repoURL, directoryName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyDir(srcPath, dstPath string) error {
	return filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relativePath, err := filepath.Rel(srcPath, path)
		if err != nil {
			return err
		}
		dst := filepath.Join(dstPath, relativePath)

		if info.IsDir() {
			return os.MkdirAll(dst, info.Mode())
		} else {
			if _, err := os.Stat(dst); err == nil {
				return nil
			}
			return copyFile(path, dst)
		}
	})
}

func copyFile(srcFile, dstFile string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
