package main

import (
    "fmt"
)

func main() {
    a := NewBinTreeNode(1)
    tree1 := NewBinaryTree(a)
    a.SetLChild(NewBinTreeNode(2))
    a.SetRChild(NewBinTreeNode(5))
    a.GetLChild().SetRChild(NewBinTreeNode(3))
    a.GetLChild().GetRChild().SetLChild(NewBinTreeNode(4))
    a.GetRChild().SetLChild(NewBinTreeNode(6))
    a.GetRChild().SetRChild(NewBinTreeNode(7))
    a.GetRChild().GetLChild().SetRChild(NewBinTreeNode(9))
    a.GetRChild().GetRChild().SetRChild(NewBinTreeNode(8))

    node2 := a.GetLChild()
    node9 := a.GetRChild().GetLChild().GetRChild()

    fmt.Println("结点2是叶子结点吗? ", node2.IsLeaf())
    fmt.Println("结点9是叶子结点吗? ", node9.IsLeaf())

    fmt.Println("这棵树的结点总数：", a.GetSize())

    l := tree1.InOrder()//中序遍历
    for e := l.Front(); e != nil; e = e.Next() {
        obj, _ := e.Value.(*BinTreeNode)
        fmt.Printf("data:%v\n", *obj)
    }

    //查询节点信息
    result := tree1.Find(6)
    fmt.Printf("结点6的父节点数据:%v\t结点6的右孩子节点数据:%v\t节点6的做孩子节点:%v\n", result.GetParent().GetData(),result.GetRChild().GetData(),result.GetLChild().GetData())
}


//         1
 //      /   \
 //     /     \
 //    2       5
 //     \    / \
 //      3   6   7
 //      /    \   \
 //      4     9   8

