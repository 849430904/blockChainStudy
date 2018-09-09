package main
import (
	"fmt"
	"bufio"
	"os" 
)
/**
https://studygolang.com/static/pkgdoc/pkg/os.htm#OpenFile
	
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。
	它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。
	如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
	flag:文件操作的模式
	FileMode：一般用于uninx或linux

flag int:https://studygolang.com/static/pkgdoc/pkg/os.htm#pkg-constants
	const (
		O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
		O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
		O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
		O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
		O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
		O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
		O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
		O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件 <使用这个要特别小心，可能把文件清了>
	)

perm FileMode:https://studygolang.com/static/pkgdoc/pkg/os.htm#FileMode

	const (
    // 单字符是被String方法用于格式化的属性缩写。
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
    ModeAppend                                     // a: 只能写入，且只能写入到末尾
    ModeExclusive                                  // l: 用于执行
    ModeTemporary                                  // T: 临时文件（非备份文件）
    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
    ModeDevice                                     // D: 设备
    ModeNamedPipe                                  // p: 命名管道（FIFO）
    ModeSocket                                     // S: Unix域socket
    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
    ModeSticky                                     // t: 只有root/创建者能删除/移动文件
    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
)
**/

//打开一个存在的文件中，将原来的内容覆盖成新的内容10句 "你好，尚硅谷!"
func main() {

	//创建一个新文件，写入内容 5句 "hello, Gardon"
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return 
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入5句 "你好,尚硅谷!"
	str := "你好,尚硅谷!\r\n" // \r\n 表示换行
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()

}