# Leetcode

[![Leetcode Testing Github Actions](https://github.com/sunist-c/LeetCode/actions/workflows/testing.yml/badge.svg)](https://github.com/sunist-c/LeetCode/actions/workflows/testing.yml)

[![Current Commit](https://img.shields.io/badge/2023.8.29-last_commit-blue)](https://studio.sunist.work/sunist-c/leetcode)
[![Total Problems](https://img.shields.io/badge/52+_problems-8A2BE2)](https://studio.sunist.work/sunist-c/leetcode)

[![Average Difficulty](https://img.shields.io/badge/difficulty-1.7750-lightyellow)](https://studio.sunist.work/sunist-c/leetcode)
[![Average Cpu Usage](https://img.shields.io/badge/cpu_usage_rank-83.20%25-lightgreen)](https://studio.sunist.work/sunist-c/leetcode)
[![Average Memory Usage](https://img.shields.io/badge/memory_usage_rank-67.72%25-lightgreen)](https://studio.sunist.work/sunist-c/leetcode)

这里是`sunist-c`的力扣刷题仓库，用于记录并提交刷题的过程。

刷题规约：

1. 每周完成至少三道题
2. 每周至少完成一道中等难度题
3. 每周完成至少一篇题解

使用Gitea-Actions完成自动化测试，刷题日志和链接放到changelog.json中

json格式如下：

```json
[
	{
		"date": "2023-06-30",
		"name": "2490. 回环句",
		"difficulty": "easy",
		"link": "https://leetcode.cn/problems/circular-sentence",
		"cpu_usage_rank": "100%",
		"memory_usage_rank": "7.41%"
	}
]
```

## 如果你也想用

1. Fork或者Clone本仓库
2. 删除`changelog.json`中的内容(而不是删除文件)以及`problems`文件夹中的内容(删除文件)并修改`readme.md`中的内容
3. 在`problems`文件夹中创建题目，建议的目录如下：
    ```text
    .
    └── problems
        ├──202x-x
        |   ├── problem-name
        |   |   ├── 114.go
        |   |   ├── 114_test.go
        |-- ...
    ```
    其中`xxx.go`是题目文件，`xxx_test.go`是测试文件，`xxx`是题目的编号，`problem-name`是题目的名称，`202x.x`是题目的日期
4. 完成题目与测试文件的编写
5. 执行测试
6. 更新`changelog.json`和`readme.md`中的内容

### 示例

我们需要完成的题目是1+1，题目的编号是114，题目的名称是`add-two-numbers`，做题的时间是2023.6.1，题目的难度是`medium`，题目的链接是`https://leetcode-cn.com/problems/add-two-numbers/`

1. 在`problems`文件夹中创建`2023-06`文件夹
2. 在`2023-06`文件夹中创建`add-two-numbers`文件夹
3. 在`add-two-numbers`文件夹中创建`114.go`,`114_test.go`和`114_pre.go`文件
4. 编写题目代码：
    在`114_pre.go`中，这里是用于定义题目给出的结构的地方(虽然此处并没有用到)：
    ```go
    package add_two_numbers
    
    type ListNode struct {
        Val  int
        Next *ListNode
    }
    ```
    
    在`114.go`中，实现`a+b`的题目要求函数：
    ```go
    package add_two_numbers
    
    func addTwoNumbers(a, b int) int {
        return a + b
    }
    ```
    
    在`114_test.go`中，要注意这里的`addTwoNumbersInput`和`originFunction`作为适配器的设计思想：
    
    ```go
    package add_two_numbers
    
    import (
        "leetcode/cases"
        "testing"
    )
   
    type addTwoNumbersInput struct {
        a int
        b int
    }
    
    func TestAddTwoNumbers(t *testing.T) {
        testCases := []cases.CommonJudgeTestCase[addTwoNumbersInput, int]{
            {
                Input:  addTwoNumbersInput{a: 1, b: 1},
                WantOutput: 2,
            },
            // ...cases
        }
   
        originFunction := func(input addTwoNumbersInput) int {
            return addTwoNumbers(input.a, input.b)
        }
        
        cases.NewCommonTestCases("AddTowNumbers", originFunction, testCases...).Run(t)
    }
    ```
5. 执行`recorder.go`，这个步骤会自动统计并更新`readme.md`和`changelog.json`：
    ```shell
    go run utils/recorder.go \
      #最好使用绝对路径 
      -f ${LeetCode仓库根目录} \ 
      # 0-Easy, 1-Medium, 2-Difficult
      -d ${难度} \ 
      # leetcode的链接的题目名称部份
      -l 'add-tow-numbers' \ 
      # 题目名称,就是题目的标题
      -n '114. 1+1' \ 
      # 从leetcode上获取的排名, 请使用0-1的小数表示, 保留四位
      -c ${CPU使用排名} -m ${内存使用排名} 
    ```
6. 添加Commit并Push到你自己的仓库 