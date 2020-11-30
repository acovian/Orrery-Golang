package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a solar system object (or press control+C to quit):")
	for input.Scan() {

		if input.Text() == "Sun" || input.Text() == "sun" {
			fmt.Println("The Sun")
			fmt.Println("As of yet, no known form of life can exist on the Sun.")
		}
		if input.Text() == "Mercury" || input.Text() == "mercury" {
			fmt.Println("The Planet Mercury")
			fmt.Println("-----")
			fmt.Println("Information:")
			fmt.Println("-----")
			fmt.Println("Mercury orbits the sun at a distance of 0.4 astronomical units.")
			fmt.Println("-----")
			fmt.Println("Potential Biology:")
			fmt.Println("The surface temperature of Mercury alternates between searing heat during the day and frigidity at night.")
			fmt.Println("The sheer amount of solar radiation and lack of a suitable atmosphere make the presence of modern-day life on Mercury extremely unlikely. At least - on the surface.")
			fmt.Println("Astrobiologists think it is possible that at one point in the solar system's history, Mercury was a habitable world, complete with oceans and even microbial life at one point.")
		}
		if input.Text() == "Venus" || input.Text() == "venus" {
			fmt.Println("The Planet Venus")
			fmt.Println("-----")
			fmt.Println("Information:")
			fmt.Println("-----")
			fmt.Println("Venus orbits the sun at a distance of 0.7 astronomical units.")
			fmt.Println("-----")
			fmt.Println("Potential Biology:")
			fmt.Println("The clouds of Venus are so toxic that they surpass the pH scale.")
			fmt.Println("In 2020, scientists detected Phosphine in the atmosphere, which is a highly unusual phenomenon on abiotic terrestrial planets. This chemical is typically considered to be a biosignature.")
			fmt.Println("Astrobiologists think it is possible that countless microbes are thriving in the Venusian atmosphere, subsisting off of the toxins and even UV radiation from the nearer Sun.")
			fmt.Println("Any potential life that exists on Venus would be extremophilic.")
		}
		if input.Text() == "Mars" || input.Text() == "mars" {
			fmt.Println("The Planet Mars")
			fmt.Println("-----")
			fmt.Println("Information:")
			fmt.Println("-----")
			fmt.Println("Mars orbits the sun at a distance of 1.5 astronomical units.")
			fmt.Println("-----")
			fmt.Println("Potential Biology:")
			fmt.Println("The atmosphere of Mars is extremely thin. The planet receives a lethal amount of solar radiation due to the fact that its magnetosphere was shredded by merciless asteroid bombardments eons ago.")
			fmt.Println("However, scientists discovered the presence of methane, which is indicative of potential life. The methan is released seasonally, likely from extremophilic microorganisms living beneath the soil (only if it is certain that this is a biosignature).")
			fmt.Println("Inexplicably, life-detecting missions to Mars ceased in the 1970s, even though strange hints at the possibility of surface-level microbial life were present.")
			fmt.Println("If life is present on Mars, it is likely a remnant from its warmer, wetter past.")
		}
		if input.Text() == "Ceres" || input.Text() == "ceres" {
			fmt.Println("The Dwarf Planet Ceres")
			fmt.Println("-----")
			fmt.Println("Information:")
			fmt.Println("-----")
			fmt.Println("Ceres orbits the sun at a distance of 2.8 astronomical units.")
			fmt.Println("-----")
			fmt.Println("Potential Biology:")
			fmt.Println("Ceres is the largest celestial object between Mars and Jupiter. It lies in the famous asteroid belt. It is far too small to be a planet. Like Pluto, it has been relegated to the status of dwarf planet.")
			fmt.Println("Scientists have discovered that Ceres hosts large bodies of saltwater beneath its surface.")
			fmt.Println("There has been much speculation that Ceres could host aquatic life in its subterranean environments.")
			fmt.Println("Although what form this life would take - microbial or multicellular - is anyone's guess.")
		}
		if input.Text() == "Jupiter" || input.Text() == "jupiter" {
			fmt.Println("The Planet Jupiter")
			fmt.Println("-----")
			fmt.Println("Information:")
			fmt.Println("-----")
			fmt.Println("Jupiter, a gas giant and the largest planet in the solar system, orbits the sun at a distance of 5.2 astronomical units. The planet acts as a very effective shield against potentially destructive objects on their way towards the inner solar system.")
			fmt.Println("-----")
			fmt.Println("Potential Biology:")
			fmt.Println("Jupiter is the largest celestial object in the solar system besides the Sun. Most scientists have not seriously investigated the possibility of life existing in the clouds of Jupiter, whatever exotic form of biochemistry that life may implement.")
			fmt.Println("Scientists prefer instead to look at the planet's large moons for signs of life. Europa, the ice world with subterranean oceans, seems to bear the most potential for life.")
			fmt.Println("Callisto may also bear some form of subterranean aquatic life, and to a much lesser possibility, Ganymede - which is nearly the size of Mars, and bigger than Mercury.")
			fmt.Println("Although what form the life on these moons would take - microbial or multicellular - is anyone's guess.")
		}
		fmt.Println("Enter a solar system object (or press control+C to quit):")
	}
}
