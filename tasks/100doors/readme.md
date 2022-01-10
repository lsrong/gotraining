# 100 doors
There are 100 doors in a row that are all initially closed.   
一排有 100 扇门最初都是关闭的。

You make 100 passes by the doors.  
你通过门 100 次。

The first time through, visit every door and  toggle  the door  (if the door is closed,  open it;   if it is open,  close it).  
第一次通过，访问每扇门并拨动门（如果门关闭，打开它；如果它是打开的，关闭它）。

The second time, only visit every 2nd door   (door #2, #4, #6, ...),   and toggle it.  
第二次，只访问每个第二个门（门 2、4、6，...），然后切换它。

The third time, visit every 3rd door   (door #3, #6, #9, ...), etc,   until you only visit the 100th door.  
第三次，访问每 3 个门（门 3、6、9，...）等，直到您只访问第 100 个门。

## Task
Answer the question:   what state are the doors in after the last pass?   Which are open, which are closed?  
回答问题：最后一次通过后门处于什么状态？哪些是开放的，哪些是封闭的？


## Notice 
唯一保持打开的门是那些数字是完美正方形(n^2)的门.   
```
打印的结果: 
1 0 0 1 0 0 0 0 1 0
0 0 0 0 0 1 0 0 0 0
0 0 0 0 1 0 0 0 0 0
0 0 0 0 0 1 0 0 0 0
0 0 0 0 0 0 0 0 1 0
0 0 0 0 0 0 0 0 0 0
0 0 0 1 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
1 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 1

唯一打开的门是整数的完美平方：1、4、9、16、25、36、49、64、81 和 100: 
1 0 0 4 0 0 0 0 9 0  
0 0 0 0 0 16 0 0 0 0  
0 0 0 0 25 0 0 0 0 0  
0 0 0 0 0 36 0 0 0 0  
0 0 0 0 0 0 0 0 49 0  
0 0 0 0 0 0 0 0 0 0  
0 0 0 64 0 0 0 0 0 0  
0 0 0 0 0 0 0 0 0 0  
81 0 0 0 0 0 0 0 0 0  
0 0 0 0 0 0 0 0 0 100  
```

