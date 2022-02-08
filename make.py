#!/usr/bin/python3

import os
import sys
import fnmatch
import subprocess
from typing import List

PWD                 = os.getcwd()
NODE_MODULES_PATH   = os.popen('npm root -g').readlines()[0].strip()
PROTOC_GEN_TS_PATH  = os.path.join(NODE_MODULES_PATH, 'ts-protoc-gen', 'bin', 'protoc-gen-ts')
if sys.platform == 'win32':
    PROTOC_GEN_TS_PATH  = os.path.join(NODE_MODULES_PATH, 'ts-protoc-gen', 'bin', 'protoc-gen-ts.cmd')
VERSION             = os.popen('git describe --tags --always').readlines()[0].strip()

API_PROTO_PATH      = os.path.join(PWD, 'api')
SERVICE_PATH        = os.path.join(PWD, 'service')
WEB_PATH            = os.path.join(PWD, 'web')

PROJECT_NAME_MARK   = '@(project.name)'
JAVA_OUT_PATH       = os.path.join(SERVICE_PATH, PROJECT_NAME_MARK, 'src', 'main', 'java')
TS_OUT_PATH         = os.path.join(WEB_PATH, 'src', 'api', PROJECT_NAME_MARK)

# 获取 path 下的所有文件
def get_files(path=PWD) -> List[str]:
    files = []
    for item in os.listdir(path):
        file = os.path.join(path, item)
        if os.path.isfile(file):
            files.append(file)
        else:
            files += get_files(file)
    return files

API_PROTO_FILES     = fnmatch.filter(get_files(API_PROTO_PATH), "*.proto")

def init():
    print('初始化环境')
    cmds = [
        'npm install -g ts-protoc-gen',
        'PowerShell .\\testing-environment.ps1',
        ]
    error_count = 0
    for cmd in cmds:
        print(cmd)
        exitcode, output = subprocess.getstatusoutput(cmd)
        if exitcode == 1:
            error_count += 1
            print("error:\n", output)
    print('初始化环境完成')
    print(f'发现 {error_count} 个错误')
    pass

# 构建 proto api 文件
def api():
    print('构建 proto api 文件')
    error_count = 0
    for file in API_PROTO_FILES:
        proto_path = os.path.dirname(file)
        project_name = os.path.basename(proto_path)
        java_out = JAVA_OUT_PATH.replace(PROJECT_NAME_MARK, project_name)
        ts_out = TS_OUT_PATH.replace(PROJECT_NAME_MARK, project_name)

        for path in [java_out, ts_out]:
            if not os.path.exists(path):
                os.makedirs(path)

        # cmd = f'''
        #     protoc `
        #         --proto_path="{proto_path}" `
        #         --plugin=protoc-gen-ts="{PROTOC_GEN_TS_PATH}" `
        #         --java_out="{java_out}" `
        #         --js_out="import_style=commonjs,binary:{ts_out}" `
        #         --ts_out="{ts_out}" `
        #         {file}
        #     '''
        cmd = f'protoc --proto_path="{proto_path}" --plugin=protoc-gen-ts="{PROTOC_GEN_TS_PATH}" --java_out="{java_out}" --js_out="import_style=commonjs,binary:{ts_out}" --ts_out="{ts_out}" {file}'
        print(cmd)
        
        exitcode, output = subprocess.getstatusoutput(cmd)
        if exitcode == 1:
            error_count += 1
            print("error:\n", output)

    print('api 文件构建完成')
    print(f'发现 {error_count} 个错误')
    pass

if __name__ == "__main__":

    if len(sys.argv) <= 1:
        print('输入 python make.py help 获得帮助')
        exit(1)

    match sys.argv[1]:
        case "help":
            print('help \t  帮助')
            print('init \t  初始化环境')
            print('api  \t  构建 proto api 文件')
        case "init":
            init()
        case "api":
            api()
        case _:
            print('输入 python make.py help 获得帮助')

    pass