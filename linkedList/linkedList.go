package linkedlist

type LinkedList struct {
	Root *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(newValue int) {
	newNode := NewNode(newValue, nil)
	if l.Root == nil {
		l.Root = newNode
		return
	}

	currentNode := l.Root

	for {
		if currentNode.Next == nil {
			currentNode.Next = newNode
			break
		}
		currentNode = currentNode.Next
	}
}

func (l *LinkedList) RemoveSingle(valueToRemove int) {
	if l.Root == nil {
		return
	}

	if l.Root.Value == valueToRemove {
		l.Root = l.Root.Next
		return
	}

	// si llega aqui es porque en el root no estaba el numero
	currentNode := l.Root

	for {
		if currentNode == nil {
			break
		}
		if currentNode.Next.Value == valueToRemove {
			currentNode.Next = currentNode.Next.Next
			break
		}
		currentNode = currentNode.Next
	}
}

func (l *LinkedList) RemoveAll(valueToRemove int) {
	//  *******@@@@
	// importantisimo ver que aqui es un FPR para eliminar todos las
	// ocurrencias de la raiz!!!
	// ^^^^^^^
	for l.Root != nil && l.Root.Value == valueToRemove {
		l.Root = l.Root.Next
	}

	if l.Root == nil {
		return
	}

	currentNode := l.Root

	for {
		if currentNode.Next == nil {
			break
		}

		if currentNode.Next.Value == valueToRemove {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}
}

func (l *LinkedList) GenerateLoopWithValueNode(loopValue int) {

	if l.Root == nil {
		return
	}

	currentNode := l.Root
	nodeOfValue := NewNode(0, nil)
	for {
		if currentNode.Next == nil {
			currentNode.Next = nodeOfValue
			break
		}

		if currentNode.Next.Value == loopValue {
			nodeOfValue = currentNode.Next
		}

		currentNode = currentNode.Next
	}

}
