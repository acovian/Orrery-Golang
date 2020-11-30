package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a solar system object:")
	for input.Scan() {

		if input.Text() == "sun" {
			fmt.Println("Sun")
		}
		if input.Text() == "mercury" {
			fmt.Println("Mercury")
		}
		if input.Text() == "venus" {
			fmt.Println(
				"Venus" +
					"Information:" +
					"The clouds of Venus are so toxic that they surpass the pH scale." +
					"Scientists detected Phosphine in the atmosphere, which is highly unusual on terrestrial planets. This chemical is typically considered to be a biosignature." +
					"Astrobiologists think it is possible that countless microbes are thriving in the Venusian atmosphere, subsisting off of the toxins and even UV radiation from the nearer Sun.")
		}
		fmt.Println("Enter a solar system object:")
	}
}
