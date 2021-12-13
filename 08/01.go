/*
--- Day 8: Seven Segment Search ---

You barely reach the safety of the cave when the whale smashes into the cave mouth, collapsing it. Sensors indicate another exit to this cave at a much greater depth, so you have no choice but to press on.

As your submarine slowly makes its way through the cave system, you notice that the four-digit seven-segment displays in your submarine are malfunctioning; they must have been damaged during the escape. You'll be in a lot of trouble without them, so you'd better figure out what's wrong.

Each digit of a seven-segment display is rendered by turning on or off any of seven segments named a through g:

  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
So, to render a 1, only segments c and f would be turned on; the rest would be off. To render a 7, only segments a, c, and f would be turned on.

The problem is that the signals which control the segments have been mixed up on each display. The submarine is still trying to display numbers by producing output on signal wires a through g, but those wires are connected to segments randomly. Worse, the wire/segment connections are mixed up separately for each four-digit display! (All of the digits within a display use the same connections, though.)

So, you might know that only signal wires b and g are turned on, but that doesn't mean segments b and g are turned on: the only digit that uses two segments is 1, so it must mean segments c and f are meant to be on. With just that information, you still can't tell which wire (b/g) goes to which segment (c/f). For that, you'll need to collect more information.

For each display, you watch the changing signals for a while, make a note of all ten unique signal patterns you see, and then write down a single four digit output value (your puzzle input). Using the signal patterns, you should be able to work out which pattern corresponds to which digit.

For example, here is what you might see in a single entry in your notes:

acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf
(The entry is wrapped here to two lines so it fits; in your notes, it will all be on a single line.)

Each entry consists of ten unique signal patterns, a | delimiter, and finally the four digit output value. Within an entry, the same wire/segment connections are used (but you don't know what the connections actually are). The unique signal patterns correspond to the ten different ways the submarine tries to render a digit using the current wire/segment connections. Because 7 is the only digit that uses three segments, dab in the above example means that to render a 7, signal lines d, a, and b are on. Because 4 is the only digit that uses four segments, eafb means that to render a 4, signal lines e, a, f, and b are on.

Using this information, you should be able to work out which combination of signal wires corresponds to each of the ten digits. Then, you can decode the four digit output value. Unfortunately, in the above example, all of the digits in the output value (cdfeb fcadb cdfeb cdbaf) use five segments and are more difficult to deduce.

For now, focus on the easy digits. Consider this larger example:

be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
fgae cfgab fg bagce
Because the digits 1, 4, 7, and 8 each use a unique number of segments, you should be able to tell which combinations of signals correspond to those digits. Counting only digits in the output values (the part after | on each line), in the above example, there are 26 instances of digits that use a unique number of segments (highlighted above).

In the output values, how many times do digits 1, 4, 7, or 8 appear?

*/

package main

import (
	"fmt"
	"os"
	// "math"
	"strings"
	"strconv"
)

type display [7]bool

func parseDisplay(str string) (display, error) {
	d := display{}
	m := []string{"a", "b", "c", "d", "e", "f", "g"}
	for idx, i := range m {
		if strings.Contains(str, i) {
			d[idx] = true
		}
	}
	return d, nil
}

type line struct {
	input []display
	output []display
}

func parseLine(str string) (line, error) {
	strs := strings.Split(str, " | ")
	inputStrs := strings.Split(strs[0], " ")
	inputDisplays := []display{}
	for _, i := range inputStrs {
		display, err := parseDisplay(i)
		if err != nil {
			return line{}, err
		}
		inputDisplays = append(inputDisplays, display)
	}
	
	outputStrs := strings.Split(strs[1], " ")
	outputDisplays := []display{}
	for _, i := range outputStrs {
		d, err := parseDisplay(i)
		if err != nil {
			return line{}, err
		}
		outputDisplays = append(outputDisplays, d)
	}
	return line{input: inputDisplays, output: outputDisplays}, nil
}

func main() {
	rlt, err := a()
	if err != nil {
		panic(err)
	}
	fmt.Println("Result A:", rlt)
	
	// rlt, err = b()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Result B:", rlt)
}

func a() (int, error) {
	// Read file
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	
	// Parse input
	strs := strings.Split(string(dat), "\n")
	lines := []line{}
	for _, i := range strs {
		l, err := parseLine(i)
		if err != nil {
			return 0, err
		}
		lines = append(lines, l)
	}
	
	// Calculate
	count := 0
	for _, i := range lines {
		for _, j := range i.output {
			segmentCount := 0
			for _, k := range j {
				if k {
					segmentCount += 1
				}
			}
			if segmentCount == 2 || segmentCount == 3 || segmentCount == 4 || segmentCount == 7 {
				count += 1
			}
		}
	}
 	return count, nil
}

func b() (int, error) {
	// Read file
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	
	// Parse input
	strs := strings.Split(string(dat), "\n")
	ints := []int{}
	for _, i := range strs {
		val, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}
		ints = append(ints, val)
	}
	
	// Calculate
	count := 0
	for i := 0; i < len(ints)-3; i++ {
		if ints[i] < ints[i+3] {
			count += 1
		}
	}
    return count, nil
}