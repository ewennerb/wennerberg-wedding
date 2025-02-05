<script>
    import Footer from "../Footer.svelte";
    import Activity from "./Activity.svelte";

    const foodAndDrink = [
        {
            "name": "Terrestrial Brewing Co.",
            "address": "7524 Father Frascati, Cleveland, OH 44102",
            "desc": "Excellent beer selection and a nice patio with views of Lake Erie.",
            "link": "https://terrestrialbrewing.com/"
        },
        {
            "name": "Collision Bend Brewing Co.",
            "address": "1250 Old River Rd, Cleveland, OH 44113",
            "desc": "Great food and craft beer located on the west bank of the Flats with a close-up view of the Cuyahoga River.",
            "link": "https://www.collisionbendbrewery.com/"
        },
        {
            "name": "Great Lakes Brewing Co.",
            "address": "2516 Market Ave, Cleveland, OH 44113",
            "desc": "Cleveland's most iconic brewery with amazing beer and food. Check out their brewery tour as well if you're looking for something to do on Friday or Sunday!",
            "link": "https://www.greatlakesbrewing.com/"
        },
        {
            "name": "Brewnuts",
            "address": "6501 Detroit Ave, Cleveland, OH 44102",
            "desc": "Novelty bar that sells a rotating menu of specialty donuts. The bar is for over 21s only, but the donuts can be enjoyed by all ages!",
            "link": "https://www.brewnutscleveland.com/"
        },
        {
            "name": "Super Electric Pinball Parlor",
            "address": "6500 Detroit Ave, Cleveland, OH 44102",
            "desc": "Pinball barcade with a lowkey and fun vibe. No quarters necessary!",
            "link": "https://www.superelectric.tv/"
        },
        {
            "name": "Tabletop Boardgame Cafe",
            "address": "1810 W 25th St, Cleveland, OH 44113",
            "desc": "Bright cafe offering a wide-range of board games to play. Serves sandwiches, snacks & coffee, plus beer & wine.",
            "link": "https://www.tabletopcleve.com/"
        },
    ]

    const activities = [
        {
            "name": "The Cleveland Museum of Art",
            "address": "11150 East Blvd, Cleveland, OH 44106",
            "desc": "Voted One of the largest and best art museums in the country (and admission is free!)",
            "link": "https://www.clevelandart.org/"
        },
        {
            "name": "Rock & Roll Hall of Fame",
            "address": "1100 E 9th St, Cleveland, OH 44114",
            "desc": "An iconic Cleveland museum dedicated to the history of rock and roll.",
            "link": "https://rockhall.com/ticketing/?gad_source=1"
        },
        {
            "name": "Greater Cleveland Aquarium",
            "address": "2000 Sycamore St. Cleveland, OH, 44113",
            "desc": "Cleveland's aquarium located inside an old power plant with all sorts of saltwater and freshwater critters to see. This building is also where the reception will be!",
            "link": "https://www.greaterclevelandaquarium.com/"
        },
        {
            "name": "Cuyahoga Valley National Park",
            "address": "Brecksville, OH",
            "desc": "33,000-acre national park with hiking trails, bike paths, a scenic railroad & vast array of flora & fauna",
            "link": "https://www.nps.gov/cuva/index.htm"
        },
        {
            "name": "Cleveland Guardians Game",
            "address": "2401 Ontario St. Cleveland, OH, 44115",
            "desc": "If you're feeling like catching a baseball game while in town, the Guards will be playing the Angels on May 30th and June 1st. Tickets are very affordable!",
            "link": "https://www.mlb.com/guardians/tickets"
        }
    ]

</script>

<div class="contentContainer">

    <div style="display: inline-block; max-width: 55vw;">
<!--        TODO: Better Spacing on the headers and between the text. It's ugly right now. -->
        <p>
            Whether you're visiting from out of town or looking to kill some time between the mass and the reception, there's lots of fun stuff to do in Cleveland! Here are some of our favorite restaurants and activities in and around downtown CLE.
        </p>

        <div>
            <h1>Bars & Restaurants</h1>
            {#each foodAndDrink as place}
                <Activity name={place.name} address={place.address} desc={place.desc} link={place.link}/>
            {/each}
        </div>


        <div>
            <h1>Activities</h1>
            {#each activities as place}
                <Activity name={place.name} address={place.address} desc={place.desc} link={place.link}/>
            {/each}
        </div>

    </div>

    <Footer/>

</div>

<style>

    .activityList {
        padding: 2vh;
    }

    h1 {
        color: rgba(0, 0, 0, 0.4);
        font-size: clamp(20px, calc(1.25rem + ((1vw - 4px) * 0.9036)), 32px);
        font-weight: 400;
        text-transform: uppercase;
        letter-spacing: .42vw;
    }

    p {
        font-size: clamp(10px, calc(0.625rem + ((1vw - 4px) * 0.6024)), 18px);
    }

</style>