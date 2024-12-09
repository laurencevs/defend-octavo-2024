package main

import (
	"aoc24/utils"
	"fmt"
	"os"
)

type fileData struct {
	id         int
	length     int
	spaceAfter int
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	if len(data)%2 == 1 {
		data = append(data, 0)
	}
	head := new(utils.LLNode[fileData]) // dummy
	tail := head
	for i := 0; i < len(data); i += 2 {
		tail = tail.InsertNext(fileData{
			id:         i / 2,
			length:     int(data[i] - '0'),
			spaceAfter: int(data[i+1] - '0'),
		})
	}
	head = head.Next
	tail.InsertNext(fileData{}) // so tail.Next.Prev exists

	fmt.Println(checksum(head, tail))
}

func checksum(head, tail *utils.LLNode[fileData]) int {
	for fileToMove := tail; fileToMove.Prev != nil; fileToMove = fileToMove.Prev {
		for candidateSpace := head; candidateSpace.Value.id != fileToMove.Value.id; candidateSpace = candidateSpace.Next {
			if candidateSpace.Value.spaceAfter >= fileToMove.Value.length {
				fileToMove.Pop()
				candidateSpace.InsertNext(fileData{
					id:         fileToMove.Value.id,
					length:     fileToMove.Value.length,
					spaceAfter: candidateSpace.Value.spaceAfter - fileToMove.Value.length,
				})
				candidateSpace.Value.spaceAfter = 0
				fileToMove.Next.Prev.Value.spaceAfter += fileToMove.Value.length + fileToMove.Value.spaceAfter
				break
			}
		}
	}

	total := 0
	position := 0
	for file := head; file != nil; file = file.Next {
		total += file.Value.id * (2*position + file.Value.length - 1) * file.Value.length / 2
		position += file.Value.length + file.Value.spaceAfter
	}

	return total
}
