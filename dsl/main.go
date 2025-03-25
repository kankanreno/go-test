package main

import (
	"fmt"
	"github.com/kankanreno/go-test/dsl/dsl"
)

// ****************** 示例使用 ******************
/*
现在解析器可以正确处理以下语法结构：
// 复杂赋值
x = (a + b) * (c - d) / 2

// 嵌套逻辑表达式
if (age > 18 || hasPermission) && !isExpired {
    // ...
}

// 多运算符表达式
return a > 0 && b < 10 || c == "test"
*/
func main() {
	// 示例脚本
	//script := `
	//	for i = 0; i < 10; i = i + 1 {
	//	  if i > 5 {
	//	    return "big"
	//	  } else {
	//	    return "small"
	//	  }
	//	}
	//`
	//script := `
	//	if (a2 >= 1.23) {
	//		return "high"
	//	} else {
	//		return "low"
	//	}
	//`
	//script := `B=days if(B=="1"){return "GW_PAR1"}if(B=="2"){return "A"}if(B=="3"){return "AAA"}else{return "FILL"}`
	script := `if(days>=3){return "B"}else{return "END"}`

	// 词法分析
	tokenizer := dsl.NewTokenizer(script)
	tokens := tokenizer.Tokenize()
	for _, token := range tokens {
		fmt.Println(token.Value)
	}

	// 语法解析
	parser := dsl.NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	// 打印AST结构
	fmt.Println("Parsed AST:")
	dsl.PrintAST(ast, 0)

	// 创建执行器
	executor := dsl.NewExecutor()

	// 设置执行器环境变量(参数)
	env := dsl.NewEnvironment(map[string]any{
		"days": "5",
	})
	executor.SetEnvironment(env)

	// 执行脚本
	result, err := executor.Execute(ast.Statements)
	if err != nil {
		fmt.Println("Execute error:", err)
	}
	fmt.Println(result)
}
