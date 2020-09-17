import os, sys

prefix = "lc"
testSuffix = "_test"
suffix = ".go"

def formatFileName(idx, isTestFile):
    _suffix = testSuffix if isTestFile else suffix
    if (idx < 10):
        return prefix+"00"+str(idx)+ _suffix
    elif (idx < 100):
        return prefix+"0"+str(idx)+_suffix
    
    return prefix+str(idx)+_suffix

def formatFolderName(idx):
    if (idx < 10):
        return prefix+"00"+str(idx)
    elif (idx < 100):
        return prefix+"0"+str(idx)
    
    return prefix+str(idx)

def initIdx(idx):
    folderName = formatFolderName(idx)
    if (os.path.exists(folderName) == False):
        # 创建文件夹
        os.mkdir(folderName)
        # 创建solution.go
        f = open(folderName+"/solution.go", 'w')
        f.write('''package solution
''')
        f.close()

        # 创建solution_test.go
        ft = open(folderName+"/solution_test.go", 'w')
        ft.write('''package solution

import (
	"testing"
)

func TestX(t *testing.T) {

	if false {
		t.Fail()
	}
}''')
        ft.close()

if (len(sys.argv) == 2):
    try:
        idx = int(sys.argv[1])
    except:
        print("wrong parameter")
        exit()
    if (idx < 1 or idx > 900):
        print("wrong parameter")
        exit()

    initIdx(idx)
    print("successfully generate files")

elif (len(sys.argv) == 3):
    try:
        left = int(sys.argv[1])
        right = int(sys.argv[2])
    except:
        print("wrong parameter")
        exit()
    if (left < 1 or right > 900 or left > right):
        print("wrong parameter")
        exit()

    for i in range(left, right+1):
        initIdx(i)
    print("successfully generate files")

else:
    print("parameters: 0 [200], if only one parameter, we will generate the No.parameter file, if achieve two parameter, we will generate files in that range")