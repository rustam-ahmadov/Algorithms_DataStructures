package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Example() {

	// Compile the expression once, usually at init time.

	// Use raw strings to avoid having to quote the backslashes.

	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))

	fmt.Println(validID.MatchString("eve[7]"))

	fmt.Println(validID.MatchString("Job[48]"))

	fmt.Println(validID.MatchString("snakey"))

	// Output:

	// true

	// true

	// false

	// false

}

func ExampleMatch() {

	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))

	fmt.Println(matched, err)

	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))

	fmt.Println(matched, err)

	matched, err = regexp.Match(`a(b`, []byte(`seafood`))

	fmt.Println(matched, err)

	// Output:

	// true <nil>

	// false <nil>

	// false error parsing regexp: missing closing ): `a(b`

}

func ExampleMatchString() {

	matched, err := regexp.MatchString(`foo.*`, "seafood")

	fmt.Println(matched, err)

	matched, err = regexp.MatchString(`bar.*`, "seafood")

	fmt.Println(matched, err)

	matched, err = regexp.MatchString(`a(b`, "seafood")

	fmt.Println(matched, err)

	// Output:

	// true <nil>

	// false <nil>

	// false error parsing regexp: missing closing ): `a(b`

}

func ExampleQuoteMeta() {

	fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`))

	// Output:

	// Escaping symbols like: \.\+\*\?\(\)\|\[\]\{\}\^\$

}

func ExampleRegexp_Find() {

	re := regexp.MustCompile(`foo.?`)

	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))

	// Output:

	// "food"

}

func ExampleRegexp_FindAll() {

	re := regexp.MustCompile(`foo.?`)

	fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), -1))

	// Output:

	// ["food" "fool"]

}

func ExampleRegexp_FindAllSubmatch() {

	re := regexp.MustCompile(`foo(.?)`)

	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1))

	// Output:

	// [["food" "d"] ["fool" "l"]]

}

func ExampleRegexp_FindSubmatch() {

	re := regexp.MustCompile(`foo(.?)`)

	fmt.Printf("%q\n", re.FindSubmatch([]byte(`seafood fool`)))

	// Output:

	// ["food" "d"]

}

func ExampleRegexp_Match() {

	re := regexp.MustCompile(`foo.?`)

	fmt.Println(re.Match([]byte(`seafood fool`)))

	fmt.Println(re.Match([]byte(`something else`)))

	// Output:

	// true

	// false

}

func ExampleRegexp_FindString() {

	re := regexp.MustCompile(`foo.?`)

	fmt.Printf("%q\n", re.FindString("seafood fool"))

	fmt.Printf("%q\n", re.FindString("meat"))

	// Output:

	// "food"

	// ""

}

func ExampleRegexp_FindStringIndex() {

	re := regexp.MustCompile(`ab?`)

	fmt.Println(re.FindStringIndex("tablett"))

	fmt.Println(re.FindStringIndex("foo") == nil)

	// Output:

	// [1 3]

	// true

}

func ExampleRegexp_FindStringSubmatch() {

	re := regexp.MustCompile(`a(x*)b(y|z)c`)

	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))

	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))

	// Output:

	// ["axxxbyc" "xxx" "y"]

	// ["abzc" "" "z"]

}

func ExampleRegexp_FindAllString() {

	re := regexp.MustCompile(`a.`)

	fmt.Println(re.FindAllString("paranormal", -1))

	fmt.Println(re.FindAllString("paranormal", 2))

	fmt.Println(re.FindAllString("graal", -1))

	fmt.Println(re.FindAllString("none", -1))

	// Output:

	// [ar an al]

	// [ar an]

	// [aa]

	// []

}

func ExampleRegexp_FindAllStringSubmatch() {

	re := regexp.MustCompile(`a(x*)b`)

	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))

	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))

	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))

	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))

	// Output:

	// [["ab" ""]]

	// [["axxb" "xx"]]

	// [["ab" ""] ["axb" "x"]]

	// [["axxb" "xx"] ["ab" ""]]

}

func ExampleRegexp_FindAllStringSubmatchIndex() {

	re := regexp.MustCompile(`a(x*)b`)

	// Indices:

	//    01234567   012345678

	//    -ab-axb-   -axxb-ab-

	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))

	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))

	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))

	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))

	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))

	// Output:

	// [[1 3 2 2]]

	// [[1 5 2 4]]

	// [[1 3 2 2] [4 7 5 6]]

	// [[1 5 2 4] [6 8 7 7]]

	// []

}

func ExampleRegexp_FindSubmatchIndex() {

	re := regexp.MustCompile(`a(x*)b`)

	// Indices:

	//    01234567   012345678

	//    -ab-axb-   -axxb-ab-

	fmt.Println(re.FindSubmatchIndex([]byte("-ab-")))

	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-")))

	fmt.Println(re.FindSubmatchIndex([]byte("-ab-axb-")))

	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-ab-")))

	fmt.Println(re.FindSubmatchIndex([]byte("-foo-")))

	// Output:

	// [1 3 2 2]

	// [1 5 2 4]

	// [1 3 2 2]

	// [1 5 2 4]

	// []

}

func ExampleRegexp_Longest() {

	re := regexp.MustCompile(`a(|b)`)

	fmt.Println(re.FindString("ab"))

	re.Longest()

	fmt.Println(re.FindString("ab"))

	// Output:

	// a

	// ab

}

func ExampleRegexp_MatchString() {

	re := regexp.MustCompile(`(gopher){2}`)

	fmt.Println(re.MatchString("gopher"))

	fmt.Println(re.MatchString("gophergopher"))

	fmt.Println(re.MatchString("gophergophergopher"))

	// Output:

	// false

	// true

	// true

}

func ExampleRegexp_NumSubexp() {

	re0 := regexp.MustCompile(`a.`)

	fmt.Printf("%d\n", re0.NumSubexp())

	re := regexp.MustCompile(`(.*)((a)b)(.*)a`)

	fmt.Println(re.NumSubexp())

	// Output:

	// 0

	// 4

}

func ExampleRegexp_ReplaceAll() {

	re := regexp.MustCompile(`a(x*)b`)

	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("T")))

	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1")))

	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))

	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))

	re2 := regexp.MustCompile(`a(?P<1W>x*)b`)

	fmt.Printf("%s\n", re2.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))

	fmt.Printf("%s\n", re2.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))

	// Output:

	// -T-T-

	// --xx-

	// ---

	// -W-xxW-

	// --xx-

	// -W-xxW-

}

func ExampleRegexp_ReplaceAllLiteralString() {

	re := regexp.MustCompile(`a(x*)b`)

	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))

	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))

	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))

	// Output:

	// -T-T-

	// -$1-$1-

	// -${1}-${1}-

}

func ExampleRegexp_ReplaceAllString() {

	re := regexp.MustCompile(`a(x*)b`)

	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))

	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))

	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))

	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	re2 := regexp.MustCompile(`a(?P<1W>x*)b`)

	fmt.Printf("%s\n", re2.ReplaceAllString("-ab-axxb-", "$1W"))

	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	// Output:

	// -T-T-

	// --xx-

	// ---

	// -W-xxW-

	// --xx-

	// -W-xxW-

}

func ExampleRegexp_ReplaceAllStringFunc() {

	re := regexp.MustCompile(`[^aeiou]`)

	fmt.Println(re.ReplaceAllStringFunc("seafood fool", strings.ToUpper))

	// Output:

	// SeaFooD FooL

}

func ExampleRegexp_SubexpNames() {

	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)

	fmt.Println(re.MatchString("Alan Turing"))

	fmt.Printf("%q\n", re.SubexpNames())

	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])

	fmt.Println(reversed)

	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))

	// Output:

	// true

	// ["" "first" "last"]

	// ${last} ${first}

	// Turing Alan

}

func ExampleRegexp_SubexpIndex() {

	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)

	fmt.Println(re.MatchString("Alan Turing"))

	matches := re.FindStringSubmatch("Alan Turing")

	lastIndex := re.SubexpIndex("last")

	fmt.Printf("last => %d\n", lastIndex)

	fmt.Println(matches[lastIndex])

	// Output:

	// true

	// last => 2

	// Turing

}

func ExampleRegexp_Split() {

	a := regexp.MustCompile(`a`)

	fmt.Println(a.Split("banana", -1))

	fmt.Println(a.Split("banana", 0))

	fmt.Println(a.Split("banana", 1))

	fmt.Println(a.Split("banana", 2))

	zp := regexp.MustCompile(`z+`)

	fmt.Println(zp.Split("pizza", -1))

	fmt.Println(zp.Split("pizza", 0))

	fmt.Println(zp.Split("pizza", 1))

	fmt.Println(zp.Split("pizza", 2))

	// Output:

	// [b n n ]

	// []

	// [banana]

	// [b nana]

	// [pi a]

	// []

	// [pizza]

	// [pi a]

}

func ExampleRegexp_Expand() {

	content := []byte(`

	# comment line

	option1: value1

	option2: value2


	# another comment line

	option3: value3

`)

	// Regex pattern captures "key: value" pair from the content.

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by

	// referencing the values captured by the regex pattern.

	template := []byte("$key=$value\n")

	result := []byte{}

	// For each match of the regex in the content.

	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {

		// Apply the captured submatches to the template and append the output

		// to the result.

		result = pattern.Expand(result, template, content, submatches)

	}

	fmt.Println(string(result))

	// Output:

	// option1=value1

	// option2=value2

	// option3=value3

}

func ExampleRegexp_ExpandString() {

	content := `

	# comment line

	option1: value1

	option2: value2


	# another comment line

	option3: value3

`

	// Regex pattern captures "key: value" pair from the content.

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by

	// referencing the values captured by the regex pattern.

	template := "$key=$value\n"

	result := []byte{}

	// For each match of the regex in the content.

	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {

		// Apply the captured submatches to the template and append the output

		// to the result.

		result = pattern.ExpandString(result, template, content, submatches)

	}

	fmt.Println(string(result))

	// Output:

	// option1=value1

	// option2=value2

	// option3=value3

}

func ExampleRegexp_FindIndex() {

	content := []byte(`

	# comment line

	option1: value1

	option2: value2

`)

	// Regex pattern captures "key: value" pair from the content.

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	loc := pattern.FindIndex(content)

	fmt.Println(loc)

	fmt.Println(string(content[loc[0]:loc[1]]))

	// Output:

	// [18 33]

	// option1: value1

}

func ExampleRegexp_FindAllSubmatchIndex() {

	content := []byte(`

	# comment line

	option1: value1

	option2: value2

`)

	// Regex pattern captures "key: value" pair from the content.

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	allIndexes := pattern.FindAllSubmatchIndex(content, -1)

	for _, loc := range allIndexes {

		fmt.Println(loc)

		fmt.Println(string(content[loc[0]:loc[1]]))

		fmt.Println(string(content[loc[2]:loc[3]]))

		fmt.Println(string(content[loc[4]:loc[5]]))

	}

	// Output:

	// [18 33 18 25 27 33]

	// option1: value1

	// option1

	// value1

	// [35 50 35 42 44 50]

	// option2: value2

	// option2

	// value2

}

func ExampleRegexp_FindAllIndex() {

	content := []byte("London")

	re := regexp.MustCompile(`o.`)

	fmt.Println(re.FindAllIndex(content, 1))

	fmt.Println(re.FindAllIndex(content, -1))

	// Output:

	// [[1 3]]

	// [[1 3] [4 6]]

}

func main() {
	text := `<?xml version='1.0' encoding='UTF-8' standalone='yes' ?><hierarchy rotation="0"><node index="0" text="" resource-id="" class="android.widget.FrameLayout" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][1080,2400]"><node index="0" text="" resource-id="" class="android.widget.LinearLayout" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][1080,2400]"><node index="0" text="" resource-id="" class="android.widget.FrameLayout" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][1080,2400]"><node index="0" text="" resource-id="com.careem.acma:id/action_bar_root" class="android.widget.LinearLayout" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][1080,2400]"><node index="0" text="" resource-id="android:id/content" class="android.widget.FrameLayout" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][1080,2400]"><node index="0" text="" resource-id="" class="androidx.compose.ui.platform.ComposeView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][1080,2400]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][1080,2400]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][1080,73]" /><node index="1" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,73][1080,2400]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,155][1080,1147]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,155][1080,1147]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[23,403][1057,651]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="Go" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[68,448][226,606]" /><node index="1" text="Go" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[271,476][333,530]" /><node index="2" text="4" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[400,482][422,523]" /><node index="3" text="EGP 71.2" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[855,479][1012,526]" /><node index="4" text="4 mins away • 12:10" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[271,542][604,581]" /></node><node index="1" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[23,651][1057,899]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="GO Awfar" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[68,696][226,854]" /><node index="1" text="GO Awfar" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[271,724][481,778]" /><node index="2" text="4" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[548,730][570,771]" /><node index="3" text="EGP 46.48" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[816,727][1012,774]" /><node index="4" text="4 mins away • 12:08" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[271,790][614,829]" /></node><node index="2" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[23,899][1057,1147]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="White Taxi" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[68,944][226,1102]" /><node index="1" text="White Taxi" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[271,972][498,1026]" /><node index="2" text="4" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[565,978][587,1019]" /><node index="3" text="EGP 48.49" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[816,975][1012,1022]" /><node index="4" text="7 mins away • 12:10" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[271,1038][601,1077]" /></node><node index="3" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[23,155][1057,403]"><node index="0" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="Bike" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,155][260,437]" /><node index="1" text="Bike" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[271,228][363,282]" /><node index="2" text="1" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[430,234][446,275]" /><node index="3" text="EGP 30.78" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[821,231][1012,278]" /><node index="4" text="6 mins away • 12:12" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[271,294][602,333]" /></node></node></node><node index="2" text="" resource-id="" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[45,2062][1035,2220]"><node index="0" text="Yalla!" resource-id="" class="android.widget.TextView" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[484,2114][596,2168]" /><node index="1" text="" resource-id="" class="android.widget.Button" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[45,2062][1035,2220]" /></node></node></node></node></node></node></node></node><node index="1" text="" resource-id="android:id/navigationBarBackground" class="android.view.View" package="com.careem.acma" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,2265][1080,2400]" /></node></hierarchy>`
	re := regexp.MustCompile(`(<node[^>]*?clickable="false"[^>]*?focusable="false"[^>]*?>(<node[^>]*?clickable="true"[^>]*?>(<node[^>]*?/>){1,10}</node>){1,10}</node>)</node><node[^>]+?clickable="true".+?text="Yalla!"`)

	submatch := re.FindStringSubmatch(text)

	req := regexp.MustCompile(`(<node[^>]*?clickable="true"[^>]*?>(<node[^>]*?/>){1,10}</node>)`)

	//fmt.Println(submatch)

	res := req.FindAllString(submatch[0], -1)

	req3 := regexp.MustCompile(`<node[^>]*?clickable="false"[^>]*/>`)
	for _, tariff := range res {
		nodes := req3.FindAllString(tariff, -1)

		if len(nodes) != 5 {
			continue
		}

		textReg := regexp.MustCompile(`text="([^"]+)"`)

		awayReg := regexp.MustCompile(`(\d) mins? away • (\d{2}:\d{2})`)
		priceReg := regexp.MustCompile(`^[A-Z]{1,3} (\d+.\d+)`)

		nameText := textReg.FindStringSubmatch(nodes[1])
		seatText := textReg.FindStringSubmatch(nodes[2])
		priceText := textReg.FindStringSubmatch(nodes[3])
		awayText := textReg.FindStringSubmatch(nodes[4])

		fmt.Println("name:", nameText[1])

		seatCount, err := strconv.Atoi(seatText[1])
		if err != nil {
			// todo tecili
		}
		fmt.Println("seats:", seatCount)

		aways := awayReg.FindStringSubmatch(awayText[1])

		minutesInt, err := strconv.Atoi(aways[1]) // Atoi - это короткая запись для ParseInt
		waitTime := time.Duration(minutesInt) * time.Minute
		fmt.Println("wait time:", waitTime)

		now := time.Now()
		timeStr := aways[2]
		targetTime, err := time.Parse("15:04", timeStr)
		if err != nil {

		}
		// Создаём полное время на сегодня с часами и минутами из строки
		targetToday := time.Date(
			now.Year(), now.Month(), now.Day(),
			targetTime.Hour(), targetTime.Minute(), 0, 0,
			now.Location(),
		)

		// Если указанное время уже прошло сегодня, добавим сутки (значит, это завтра)
		if targetToday.Before(now) {
			targetToday = targetToday.Add(24 * time.Hour)
		}

		// Разница
		duration := targetToday.Sub(now)

		// Вывод в минутах
		fmt.Println("trip time:", int(duration.Minutes()))

		priceRes := priceReg.FindStringSubmatch(priceText[1])
		floatValue, err := strconv.ParseFloat(priceRes[1], 64)
		fmt.Println("price: ", floatValue)
	}

}
