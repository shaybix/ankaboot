package ankaboot

// Greatly inspired by Rob Pike's talk on Lexical Analysis in Go
// and borrowed a lot of the code from slide, to get it to work.
// Will require to modify the code to be more suitable for my own
// use case.
//
type itemType int

// Define all of the 'items' the lexer will need to lex.
const (
	itemError     itemType = iota // Error
	itemEOF                       // End Of File
	itemLTag                      // Left Tag <
	itemRTag                      // Right Tag >
	itemRCTag                     //Right Close tag  />  [ <input />]
	itemLCTag                     // Left Close tag </   [ <a></a>]
	itemDocType                   // <!DOCTYPE html >
	itemHTML                      // <HTML>
	itemHead                      // <head>
	itemMeta                      // <meta >
	itemLink                      // <link />
	itemTitle                     // <title>
	itemStyle                     // <style>
	itemBody                      // <body>
	itemScript                    //  <script>
	itemFooter                    // <footer>
	itemTable                     // <table>
	itemTHead                     // <thead>
	itemTBody                     // <tbody>
	itemTh                        // <th>
	itemTr                        // <tr>
	itemTd                        // <td>
	itemDiv                       // <div>
	itemH1                        // <h1>
	itemH2                        // <h2>
	itemH3                        // <h3>
	itemH4                        // <h4>
	itemH5                        // <h5>
	itemH6                        // <h6>
	itemP                         // <p>
	itemSpan                      // <span>
	itemA                         // <a>
	itemImg                       // <img />
	itemUl                        // <ul>
	itemOl                        // <ol>
	itemLi                        // <li>
	itemBr                        // <br />
	itemHr                        // <hr />
	itemText                      // the values wrapped by the elements
	itemAttrId                    //
	itemAttrClass                 //
	itemAttrHref                  //
	itemAttrSrc                   //
	itemAttrType                  //
	itemAttrName                  //
	itemAttrValue                 //
	itemAttrCol                   //
	itemAttrRow                   //

)

// item struct that will be passed through lexer items channel to the
// parser.
type item struct {
	typ itemType
	val string
}

// lexer struct
type lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	items chan item
}

//stateFn represents a state of the scanner
// as a function that returns the next state (rob pike)
type stateFn func(*lexer) stateFn

// lex initialises the lexer
func lex(name, input string) (*lexer, chan item) {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}

	go l.run()
	return l, l.items
}

// run lexer
func (l *lexer) run() {
	// here will run
	close(l.items)
}

// emit tokens
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos

}

//next moves to the next item in the  slice
func (l *lexer) next() {
	// next
}

//peek looks ahead 1 in a slice and then backs up.
func (l *lexer) peek() {

}

//backup goes back 1 in a slice
func (l *lexer) backup() {

}

//close open channel
func close(items chan item) {
	// close channel
}

//lexInsideElement
func lexInsideElement() {
	// lex inside the element
}
