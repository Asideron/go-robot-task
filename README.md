# go-robot-task

It's a toy Golang script for my student to compare dynamic and recursive solutions of the robot-task.   

## The robot task

**Description:** There is a field of (N x M) cells. You are a robot standing in the lower left corner of the field at the coordinate of (1, 1) which can only move up and to the right. 

**The task:** How many ways do you have to get to the cell (i, j)?  

**Constraints:**    
    - N, M > 0 where N, M are integers.
    - i < N + 1 and j < M + 1

## Demonstration

```
Dynamic solution
Steps:  817190    
Calls:  299       

Recursive solution
Steps:  817190    
Calls:  1634379 
```