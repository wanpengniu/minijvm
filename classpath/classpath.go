package classpath

import (
	"os"
	"path/filepath"
)

// Classpath 三个字段分别存放三种类路径
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// ReadClass 依次从启动类路径、扩展类路径和用户类路径中搜索class文件
func (classpath *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := classpath.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := classpath.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return classpath.userClasspath.readClass(className)
}

func (classpath *Classpath) String() string {
	return classpath.userClasspath.String()
}

// 使用-Xjre选项解析启动类路径和扩展类路径，使用-classpath/-cp选项解析用户类路径
func (classpath *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	classpath.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	classpath.extClasspath = newWildcardEntry(jreExtPath)
}

// 如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径
func (classpath *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	classpath.userClasspath = newEntry(cpOption)
}

// 优先使用用户输入的-Xjre选项作为jre目录。如果没有输入该选项，则在当前目录下寻找jre目录。如果找不到，尝试使用JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
