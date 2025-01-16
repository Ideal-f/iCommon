#!/bin/bash

# 判断是否可以执行
if [ ! -d generateTool ];then echo "请先执行 bash init.bash 脚本初始化项目！";exit 1;fi

# 执行生成命令
/bin/bash ./generateTool/comon.bash

if [[ $? == 0 ]];then echo " ### 恭喜！项目生成成功了！###";fi