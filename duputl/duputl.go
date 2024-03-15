package github.com/shockliu/tools/duputl

import (
  "time"
)

type DupNode struct {
  value string
  ctime int64
  next  *DupNode
}

type DeDuplicate struct {
  head *DupNode
  timeOut  int64
  max int
}

// 创建一个新的去重对象，设置超时时间为timeset,最大容量为max,max<1则不限制容量
func NewDeDup(timeset int64,max int) *DeDuplicate{
  return &DeDuplicate{timeOut:timeset,max:max}
}

// 返回true为重复，否则未重复，且放入缓存比较队列
func (dedup *DeDuplicate) DukChk(val string)  bool {
  // 检查队列是否有
  var pre *DupNode
  for cur := dedup.head; cur != nil; cur = cur.next {
    if time.Now().Unix()-cur.ctime > dedup.timeOut {
      // 超时 golang 队列释放内存
      if pre == nil {
        break
      } else {
        pre.next = nil
        break
      }
    }
    if cur.value == val {
      // logger.Debugf("重复节点%s\n", val)
      return  true
    }
    pre = cur
  }
  dedup.head=&DupNode{value:val,ctime:time.Now().Unix(), next: dedup.head}
  return false
}
