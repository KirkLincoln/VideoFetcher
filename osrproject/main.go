package main

func validateInput() {
	/*  Validates number of cli arguments
	 *	and calls either fetch or search
	 */
}

func fetch() {
	/* Formats POST request for fetch and returns a string */


}

func search() {
	/* Calls regex and returns JSON */
}

func compareDirectory() {
	/* Check if directory exists, if not then provision one */

	/* Initiate ch */
	/* go fetch() {
	   	  go filterDirectoryContents(ch)
	   }
	 */
}

func filterDirectoryContents(ch chan <- string) {

}

func handler() {
	/*
	 *     validateInput() //entry
	 *      /            \
	 *  fetch()			search() // concurrent internal
	 *      \            /
	 *     compareDirectory()  // concurrent internal
	 *        \        /
	 *      return to main() // console return
	 *
	 */
}

func main() {

}
