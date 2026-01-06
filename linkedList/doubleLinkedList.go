package linkedlist

type DoubleLinkedList struct {
	root *NodeDoubleLink
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		root: nil,
	}
}

func (d *DoubleLinkedList) Insert(newValue int) {
	newNode := NewNodeDoubleLink(newValue, nil, nil)
	if d.root == nil {
		d.root = newNode
		return
	}

	currentNode := d.root

	for {
		if currentNode.next == nil {
			newNode.previous = currentNode
			currentNode.next = newNode
			break
		}

		currentNode = currentNode.next
	}
}

func (d *DoubleLinkedList) Remove(valueToRemove int) {
	if d.root == nil {
		return
	}

	if d.root.value == valueToRemove {
		d.root = d.root.next
		if d.root != nil {
			d.root.previous = nil
		}
		return
	}

	currentNode := d.root

	for {
		if currentNode.next == nil {
			break
		}

		if currentNode.next.value == valueToRemove {
			tmpNode := currentNode
			currentNode.next = currentNode.next.next
			if currentNode.next != nil {
				currentNode.next.previous = tmpNode
			}
			break
		}
		currentNode = currentNode.next
	}
}

func (d *DoubleLinkedList) RemoveAll(valueToRemove int) {
	// if d.root == nil {
	// 	return
	// }  <<<--- no es necesario porque en el if del for hay un d.root != nil && ...
	//  y despues del for hay un d.root == nil { return

	for {
		if d.root != nil && d.root.value == valueToRemove {
			d.root = d.root.next
			if d.root != nil {
				d.root.previous = nil
			}
			continue
		}
		break
	}

	if d.root == nil {
		return
	}

	currentNode := d.root

	for {
		if currentNode.next == nil {
			break
		}
		if currentNode.next.value == valueToRemove {
			tmpNode := currentNode
			currentNode.next = currentNode.next.next
			if currentNode.next != nil {
				currentNode.next.previous = tmpNode
			}
		} else {
			currentNode = currentNode.next
		}
	}

}
