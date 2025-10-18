// ps -o pid,ppid,comm | go run process-tree.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Process struct {
	PID      int
	PPID     int
	Command  string
	Children []*Process
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	processes := make(map[int]*Process)

	isHeader := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if isHeader && strings.HasPrefix(strings.ToUpper(line), "PID") {
			isHeader = false
			continue
		}

		fields := strings.Fields(line)
		fmt.Println("fields", fields)
		if len(fields) < 3 {
			continue
		}

		pid, err1 := strconv.Atoi(fields[0])
		ppid, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			continue
		}
		command := strings.Join(fields[2:], " ")

		processes[pid] = &Process{
			PID:     pid,
			PPID:    ppid,
			Command: command,
		}
	}

	// Build the tree
	for _, p := range processes {
		if parent, ok := processes[p.PPID]; ok {
			parent.Children = append(parent.Children, p)
		}
	}

	// Find roots (usually PID 1 or orphans)
	var roots []*Process
	for _, p := range processes {
		if _, ok := processes[p.PPID]; !ok {
			roots = append(roots, p)
		}
	}

	// Sort roots for consistent output
	sort.Slice(roots, func(i, j int) bool { return roots[i].PID < roots[j].PID })

	// Print tree recursively
	for _, root := range roots {
		printTree(root, "", true)
	}
}

func printTree(p *Process, prefix string, isLast bool) {
	fmt.Printf("%s%s─%s (%d)\n", prefix, branch(isLast), p.Command, p.PID)

	sort.Slice(p.Children, func(i, j int) bool { return p.Children[i].PID < p.Children[j].PID })

	for i, child := range p.Children {
		newPrefix := prefix
		if isLast {
			newPrefix += "   "
		} else {
			newPrefix += "│  "
		}
		printTree(child, newPrefix, i == len(p.Children)-1)
	}
}

func branch(isLast bool) string {
	if isLast {
		return "└"
	}
	return "├"
}
