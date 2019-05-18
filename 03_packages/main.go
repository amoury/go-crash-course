package main

import (
	"fmt"
	"math"

	"github.com/amoury/go_crash_course_YT/03_packages/strutil"
)

func main() {
	fmt.Println(math.Max(32, 64))
	fmt.Println(math.Ceil(32.454))
	fmt.Println(strutil.Reverse("Ansar"))
}
