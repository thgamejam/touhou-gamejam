#!/bin/bash
# shellcheck disable=SC2086
# shellcheck disable=SC2207
# shellcheck disable=SC2068

pwd=$(pwd)
template_api_path="$pwd/api/template/v1"
template_proto_path="$template_api_path/template.proto"
template_error_path="$template_api_path/error_reason.proto"
template_app_path="$pwd/app/template"

# 下划线转驼峰
function to_camel() {
  arr=$(echo $1 | tr '_' ' ')
  result=''
  for var in ${arr[@]}; do
    firstLetter=$(echo ${var:0:1} | awk '{print toupper($0)}')
    otherLetter=${var:1}
    result=$result$firstLetter$otherLetter
  done
  echo $result
}

echo -n "输入微服务的名称(下划线间隔单词): "
read -r app_name

app_name_camel=$(to_camel $app_name)

api_path="$pwd/api/$app_name/v1"
proto_path="$api_path/$app_name.proto"
error_path="$api_path/error_reason.proto"

mkdir -p $api_path

cp $template_proto_path $proto_path
cp $template_error_path $error_path

app_path="$pwd/app/$app_name"

mkdir -p $app_path
cp -r "$template_app_path/cmd" $app_path
cp -r "$template_app_path/configs" $app_path
cp -r "$template_app_path/internal" $app_path

# 获取目录下所有文件
function get_files() {
  files=()
  paths=$(ls "$1")
  for path in $paths; do
    if [ -d $1"/"$path ]; then
      temps=$(get_files $1"/"$path)
      files=("${files[*]}" "${temps[*]}")
    else
      local path=$1"/"$path
      files[${#files[@]}]=$path
    fi
  done
  echo ${files[*]}
}

files=($(get_files $app_path))
files[${#files[@]}]=$proto_path
files[${#files[@]}]=$error_path

for file in ${files[*]}; do
  echo $file
  # 替换文件关键字
  sed -i "s/template/$app_name/g" $file
  sed -i "s/Template/$app_name_camel/g" $file
  # 重名
  temp=${file##*/}
  if [ ${temp%.*} == "template" ]; then
    path=${file%/*}
    format=${file##*.}
    mv $file "$path/$app_name.$format"
  fi
done
