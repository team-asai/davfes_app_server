package main

import (
    "davfes_app/app"
    "os"
    "syscall"
)

func main() {
  // daemon実行
//   errcd := daemon(0, 0)
  
//   if errcd != 0 {
    // panic(errcd
//     os.Exit(1)
//   }

  app.Start()
}

func daemon(nochdir, noclose int) int {
	var ret uintptr
	var err syscall.Errno
	
	// バックグラウンドプロセスにする
	// 子プロセスを生成し，親プロセスを終了する
	ret, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
	 return -1
	}
	switch ret {
	 case 0:
	  // 子プロセスが生成できたらそのまま処理を続ける
	  break
	 default:
	  // 親プロセスだとここで終了する
	  os.Exit(0)
	}
	
	 
	// 新しいセッションを生成(子プロセスがセッションリーダになる)
	pid, _ := syscall.Setsid()
	if pid == -1 {
	 return -1
	}
	 
	if nochdir == 0 {
	 // カレントディレクトリの再設定。ここではルートにしておく
	 os.Chdir("/")
	}
	 
	// ファイルのパーミッションを再設定(必須ではない。オプション)
	syscall.Umask(0)
	
	if noclose == 0 {
	 // 標準入出力先を/dev/nullファイルに変更して、すべて破棄する
	 f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
	 if e == nil {
	  fd := int(f.Fd())
	  syscall.Dup2(fd, int(os.Stdin.Fd()))
	  syscall.Dup2(fd, int(os.Stdout.Fd()))
	  syscall.Dup2(fd, int(os.Stderr.Fd()))
	 }
	}
	
	return 0
}