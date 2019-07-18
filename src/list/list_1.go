package main

import (
	"fmt"
)

const MAXSIZE = 20 			// 存储空间初始化存储量
const OK = 1
const ERROR = 0
const TRUE = 1
const FALSE = 0

type ElemType int 			// ElemType根据实际情况而定，这里假调为int
type Status int

// 顺序存储结构
type SqList struct {
	Data 	[MAXSIZE]ElemType		// 数组存储数据元素，最大值为MAXSIZE
	Length 	int						// 线性表当前长度
}

func NewSqList() SqList {
	return SqList{}
}

// Status是函数类型，其值是返回状态如OK等
// 初始化条件：顺序线性表已经存在，1<=i<=ListLength(L)
// 操作结果：用e返回L中第i个数据元素的值
func GetElem(L SqList, i int, elemType *ElemType) Status {
	if L.Length == 0 || i < 1 || i > L.Length {
		return ERROR
	}
	*elemType = L.Data[i-1]
	return OK
}

// 初始条件，顺序线性表L已存在,1<=i<=ListLength(L)
// 操作结果: 在L表中第i个位置之前插入新的元素e,L表长度加1
func ListInsert(L *SqList, i int, e ElemType) Status {
	if L.Length == MAXSIZE { // 顺序表已满
		return ERROR
	}
	if i < 1 || i > L.Length + 1 {	//i不在范围内
		return ERROR
	}
	var k int
	if i <= L.Length { // i 不在表尾
		for k=L.Length; k>=i-1; k-- {	// 将要插入位置后元素向后移动一位
			L.Data[k+1] = L.Data[k]
		}
	}
	L.Data[i-1] = e	//插入新元素
	L.Length++	// 长度加1
	return OK
}

// 初始条件，顺序线性表L已存在,1<=i<=ListLength(L)
// 操作结果: 删除L的第i个元素，并用e返回值,L表长度减1
func ListDelete(L *SqList, i int, elemType *ElemType) Status {
	if i < 1 || i > L.Length + 1 {	//i不在范围内
		return ERROR
	}
	*elemType = L.Data[i-1]	// 取出L中i位置的元素
	var k int
	for k=i-1; k<=L.Length; k++ {	// 将i后元素向前移动一位
		L.Data[k] = L.Data[k+1]
	}
	L.Length-- // 长度减1
	return OK
}

// 线性表，零个或多个数据元素的有限序列(在较复杂的线性表中，一个元素可以由多个数据项组成)
// 线性表的元素个数定义为经性表的长度，0为空表
// 直接前驱元素 a-1
// 直接后驱元素 a+1
func main()  {

	str := "线性表的抽象数据类型定义如下：\n"
	str += "ADT 线性表 List\n"
	str += "Data\n"
	str += "	线性表的数据对象集合为{a1,a2,...an}，每个元素的类型块均为DateType.\n"
	str += "Operation\n"
	str += "	InitList(*L): 初始化操作，建立一个空的线性表\n"
	str += "	InitEmpty(L): 或线必表为空，返回true，否由返回false\n"
	str += "	ClearList(*L): 将线性表清空\n"
	str += "	GetElem(L, i, *e): 将线性中L中第i个值返回线*e\n"
	str += "	LocateElem(L, e): 在线性表L中查找与给定值e相等的元素，如果查找成功，返回该元素在表中序号表示成功；否则，返回0表示失败\n"
	str += "	ListInsert(*L, i, e): 在线性表中第i个元素插入新元素e\n"
	str += "	ListDelete(*L, i, *e): 删除线性表L中第i个元素，并用e返回其值\n"
	str += "	Length(L): 返回性性表L的元素个数\n"
	str += "endADT\n"
	fmt.Println(str)

	L := NewSqList()

	var a1, a2, a3, a4, a5 ElemType = 1, 2, 3, 4, 5
	ListInsert(&L, 1, a1)
	ListInsert(&L, 2, a2)
	ListInsert(&L, 3, a3)
	ListInsert(&L, 4, a4)
	fmt.Println(L)
	fmt.Println("---")

	// 插入i
	status := ListInsert(&L, 2, a5)
	if status == OK {
		fmt.Println("insert ok", status)
	} else {
		fmt.Println("insert err", status)
	}
	fmt.Println(L)
	fmt.Println("---")

	// 查找
	var data ElemType
	status = GetElem(L, 2, &data)
	if status == OK {
		fmt.Println("get ok", status, "data:", data)
	} else {
		fmt.Println("get error", status)
	}
	fmt.Println(L)
	fmt.Println("---")

	// 删除
	var e ElemType
	status = ListDelete(&L, 2, &e)
	if status == OK {
		fmt.Println("delete ok", status, "e:", &e, e)
	} else {
		fmt.Println("delete err", status)
	}
	fmt.Println(L)
	fmt.Println("---")
}




