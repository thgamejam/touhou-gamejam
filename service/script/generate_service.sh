#!/usr/bin/env bash
# shellcheck disable=SC2086
# shellcheck disable=SC2207
# shellcheck disable=SC2068

pwd=$(pwd)
api_path="$pwd/api"

# 获取目录下所有服务proto文件
function get_services() {
  files=()
  paths=$(ls "$1")
  for file in $paths; do
    if [ -d $1"/"$file ]; then
      temps=$(get_services $1"/"$file)
      files=("${files[*]}" "${temps[*]}")
    else
      local path=$1"/"$file
      format=${file##*.}
      if [[ $file != "error_reason.proto" && $format == "proto" ]]; then
        files[${#files[@]}]=$path
      fi
    fi
  done
  echo ${files[*]}
}

files=($(get_services $api_path))

echo "选择需要生成的service.proto文件"
echo "index" $'\t' "file"
i=1
for file in ${files[*]}; do
  echo $i $'\t' $file
  (( i++ ))
done

echo -n "输入序号: "
read -r i

if [[ i -lt 1 || i -gt ${#files[@]} ]]; then
  echo "ERROR 不存在的序号！"
  exit 1
fi
i=$i-1

file=${files[$i]}
path=${file%/*}
path=${path%/*}
app_name=${path##*/}
app_path="$pwd/app/$app_name"
service_path="$app_path/internal/service"

kratos proto server $file -t $service_path
