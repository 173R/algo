/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// Суть такая: Идём по левым узлам в основном цикле, и левым узлам выставляем
// указатель Next на правый елемент parent.left.Next = parent.right
// а во внутреннему идём по уровню используя родителя
// и его Next указатель который был получен на прошлой итерации внешнего цикла
// Parent.Right.Next = Parent.Next.Left и  Parent.Next.Left.Next = Parent.Next.Right
// И в следующей итерации меняем parent в этом внутреннем цикле на parent.Next и по кругу

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	for parent := root; parent.Left != nil; parent = parent.Left {
		parent.Left.Next = parent.Right;

		for p := parent; p.Next != nil; p = p.Next {
			p.Right.Next = p.Next.Left
			p.Next.Left.Next = p.Next.Right
		}
	}

	return root
}