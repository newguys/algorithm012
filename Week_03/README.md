学习笔记
学习笔记
1.递归 阶乘
递归代码模版
（1）递归终结条件
（2）处理当前逻辑层
（3）下探到下一层
（4）清理当前层

void recursion(int level, int param) { 
  // recursion terminator
  if (level > MAX_LEVEL) { 
    // process result 
    return ; 
  }

  // process current logic 
  process(level, param);

  // drill down 
  recursion(level + 1, param);

  // reverse the current level status if needed
}


2.递归题目找最近重复性

