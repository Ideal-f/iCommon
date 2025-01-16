#!/bin/bash

# 导入环境变量
echo "dev" > .ENV

# 克隆拉取项目
if [ ! -d generateTool ];then git clone https://gogs.fabanyun.cn/devOpsGolang/fbyTool.git generateTool;fi

# 执行生成命令
/bin/bash ./generateTool/comon.bash

if [[ $? == 0 ]];then echo " ### 恭喜！项目生成成功了！###";fi