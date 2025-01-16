#!/bin/bash

#!/bin/bash

echo "########## 正在检查 Go 版本 ##########"

read -p "⚠️ generateTool 目录不存在，是否自动初始化？[Y/N] " choice
echo $choice

# 查找 Go 命令路径
goCmd=$(which go)

# 检查 Go 是否安装
if [[ -z "$goCmd" ]]; then
    echo "❌ 出错了！Go 命令不存在。请安装 Go 1.20 或以上版本。"
    exit 1
fi

# 输出 Go 路径和版本
echo "✅ Go 命令路径：$goCmd"
$goCmd version

# 提取 Go 版本号
goVersion=$($goCmd version | awk '{print $3}' | sed 's/go//')

# 判断版本是否 >= 1.20
requiredVersion="1.20"
if [[ $(echo -e "$goVersion\n$requiredVersion" | sort -V | head -n1) != "$requiredVersion" ]]; then
    echo "❌ 出错了！Go 的版本不正确。当前版本：$goVersion，期待：$requiredVersion 或以上版本。"
    exit 1
fi

echo "✅ 检查通过！Go 版本正确：$goVersion"

# 获取当前脚本目录
dirName=$(dirname "$0")
echo "当前脚本路径：$dirName"

# 获取分支名称参数
branchName=$1

# 检查分支参数是否为空
if [ -z "$branchName" ]; then
    echo "❌ 请输入打包的分支名称！" >&2
    exit 1
fi

# 打印分支名称
echo "🔧 打包分支：$branchName"

# 检查 generateTool 目录是否存在
if [ ! -d generateTool1 ]; then
    read -p "⚠️ generateTool 目录不存在，是否自动初始化？[Y/N] " choice
    if [[ "$choice" == "y" || "$choice" == "Y" ]]; then
        echo "正在初始化项目..."
        bash init.bash
        if [[ $? -ne 0 ]]; then
            echo "❌ 初始化失败，请检查 init.bash 脚本！" >&2
            exit 1
        fi
    else
        echo "❌ 请先执行 bash init.bash 脚本初始化项目！" >&2
        exit 1
    fi
fi

# 生成日志文件名
LOG_FILE="build_$(date +%Y%m%d_%H%M%S).log"

# 执行生成脚本并记录日志
echo "🚀 正在执行生成脚本 ./generateTool/comon.bash"
/bin/bash ./generateTool/comon.bash > "$LOG_FILE" 2>&1

# 检查命令是否成功
if [[ $? -eq 0 ]]; then
    echo "🎉 恭喜！项目生成成功了！日志已保存到：$LOG_FILE"
else
    echo "❌ 生成失败！请查看日志：$LOG_FILE" >&2
    exit 1
fi
