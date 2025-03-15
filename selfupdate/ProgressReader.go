package selfupdate

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"time"
)

// ProgressReader 包装原始的 io.Reader，在每次读取时更新进度条
type ProgressReader struct {
	reader io.Reader
	bar    *progressbar.ProgressBar
}

func NewProgressReader(reader io.Reader, total int64) *ProgressReader {
	bar := progressbar.NewOptions64(
		total,
		progressbar.OptionSetWriter(os.Stderr), // 将进度条输出到 stderr
		progressbar.OptionSetDescription("Downloading"), // 设置进度条描述
		progressbar.OptionSetWidth(50),                  // 设置进度条宽度
		progressbar.OptionThrottle(65*time.Millisecond), // 设置更新频率
		progressbar.OptionShowBytes(true),               // 显示已下载的字节数
		progressbar.OptionShowCount(),                   // 显示总进度
		progressbar.OptionOnCompletion(func() {
			fmt.Fprintln(os.Stderr, "\nDownload complete!") // 下载完成后输出提示
		}),
		progressbar.OptionSpinnerType(14), // 设置进度条的旋转动画类型
	)

	return &ProgressReader{
		reader: reader,
		bar:    bar,
	}
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.bar.Add64(int64(n))
	return n, err
}
