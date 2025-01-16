#!/bin/bash

function runCmd() {
  pwd
  cmdFilePath=$1
  tips=$2

  echo ""
  echo "【${tips}】开始！！！"

  echo "即将执行" /bin/bash "${cmdFilePath}"

  /bin/bash "${cmdFilePath}"
  if [[ $? != 0 ]];then
    echo "【运行出错了出错啦!】执行命令：${cmdFilePath}:${tips}"
    exit 1;
  fi

  printf "\nᕙ( ＾‿ゝ＾ c)【Done!】${cmdFilePath}:${tips}\n\n"
}

runCmd "./ac.bash" "设置 go env 相关的配置"
if [[ $? != 0 ]];then exit 1; fi
