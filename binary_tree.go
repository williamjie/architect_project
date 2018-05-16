package main

import "fmt"

//二叉树结构体
//如果每个节点有两个指针，分别用来指向左子树和右子树，我们把这样的结构叫做二叉树
type Student struct {
    Name  string
    Age   int
    Score float32
    left  *Student
    right *Student
}

func trans(root *Student) {
    if root == nil {
        return
    }
    //fmt.Println(root) //前序遍历
    trans(root.left)
    //fmt.Println(root) //中序遍历
    trans(root.right)
    fmt.Println(root) //后序遍历
}

func main() {
    var root *Student = new(Student)

    root.Name = "stu01"
    root.Age = 11
    root.Score = 100

    var left1 *Student = new(Student)
    left1.Name = "stu02"
    left1.Age = 12
    left1.Score = 100

    root.left = left1

    var right1 *Student = new(Student)
    right1.Name = "Stu04"
    right1.Age = 13
    right1.Score = 100

    root.right = right1

    var left2 *Student = new(Student)
    left2.Name = "Stu03"
    left2.Age = 25
    left2.Score = 100

    left1.left = left2
    trans(root)
}