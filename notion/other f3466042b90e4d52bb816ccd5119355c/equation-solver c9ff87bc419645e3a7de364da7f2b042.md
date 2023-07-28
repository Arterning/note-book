# equation-solver

# Description

Equation solver

# 方案 1

使用 Apache Commons Math 中的 Expression Evaluator

## 存在的问题

在循环中多次执行 evaluate 函数可能导致性能问题，因为在每次执行 evaluate 函数时，需要创建表达式解析器，进行词法分析，语法分析，符号解析等。如果表达式比较复杂，这些操作可能需要很多时间。因此，如果需要计算多个表达式，最好将所有表达式编译为字节代码，并在循环中使用已编译的字节代码来执行计算。这可以在大大提高运算速度的同时避免性能问题。

# 方案 2

并行计算

# 方案 3

是否可以考虑使用 spark