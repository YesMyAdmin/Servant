package utils
import (
    "archive/zip"
    "io"
    "os"
    "path/filepath"
)

// ZipFolder 压缩文件夹为 ZIP 文件
func ZipFolder(sourceDir, targetZip string) error {
    // 创建 ZIP 文件
    zipFile, err := os.Create(targetZip)
    if err != nil {
        return err
    }
    defer zipFile.Close()

    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close()

    // 遍历目录
    err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // 跳过目录本身（但需要在 ZIP 中保留目录结构）
        relPath, err := filepath.Rel(sourceDir, path)
        if err != nil {
            return err
        }
        if relPath == "." {
            return nil // 不添加根目录自身
        }

        // 如果是目录，只需在 ZIP 中创建条目（可选）
        if info.IsDir() {
            // 在 ZIP 中创建目录条目（通常不需要，但有些工具需要）
            // 可直接 return nil，因为 zip.Writer 会在添加文件时自动创建父目录
            return nil
        }

        // 打开源文件
        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close()

        // 创建 ZIP 中的文件头
        header := &zip.FileHeader{
            Name:   relPath,
            Method: zip.Deflate, // 压缩算法
        }
        header.SetMode(info.Mode())

        writer, err := zipWriter.CreateHeader(header)
        if err != nil {
            return err
        }

        _, err = io.Copy(writer, file)
        return err
    })

    return err
}