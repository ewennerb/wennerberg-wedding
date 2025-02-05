<script>
    import PersonCard from "./PersonCard.svelte";
    import Footer from "../Footer.svelte";

    const groomsmen = [
        {
            name: "Kyle Murphy",
            role: "Groomsman",
            relation: "Brother-in-Law of the Groom",
            photo: "/img/groomsmen/kyle.jpeg"
        },
        {
            name: "Anthony Mencini",
            role: "Groomsman",
            relation: "Friend of the Groom",
            photo: "/img/groomsmen/anthony.jpg"
        },
        {
            name: "Max Bostelman",
            role: "Groomsman",
            relation: "Brother of the Bride",
            photo: "/img/groomsmen/max.jpeg"
        },
        {
            name: "Matthew Dolinar",
            role: "Groomsman",
            relation: "Friend of the Groom",
            photo: "/img/groomsmen/matthew.jpg"
        },
        {
            name: "Jake Carrico",
            role: "Groomsman",
            relation: "Friend of the Groom",
            photo: "/img/groomsmen/placeholder.jpg"
        },
        {
            name: "James Lawless",
            role: "Groomsman",
            relation: "Friend of the Groom",
            photo: "/img/groomsmen/placeholder.jpg"
        }
    ];

    const bridesmaid = [
        {
            name: "Annie Murphy",
            role: "Bridesmaid",
            relation: "Sister of the Groom",
            photo: "/img/bridesmaids/Annie.PNG"
        },
        {
            name: "Grace Wennerberg",
            role: "Bridesmaid",
            relation: "Sister of the Groom",
            photo: "/img/bridesmaids/Grace.jpg"
        },
        {
            name: "Kelly Sullivan",
            role: "Bridesmaid",
            relation: "Friend of the Bride",
            photo: "/img/bridesmaids/Kelly.jpg"
        },
        {
            name: "Melanie Bostelman",
            role: "Bridesmaid",
            relation: "Sister of the Bride",
            photo: "/img/bridesmaids/Melanie.jpeg"
        },
        {
            name: "Molly Ryan",
            role: "Bridesmaid",
            relation: "Friend of the Bride",
            photo: "/img/bridesmaids/Molly.jpg"
        },
        {
            name: "Tavia Vitkauskas",
            role: "Bridesmaid",
            relation: "Friend of the Bride",
            photo: "/img/bridesmaids/Tavia.jpg"
        }
    ];

    const bestMan = {
        name: "Rodrigo Corrigan",
        role: "Best Man",
        relation: "Friend of the Groom",
        photo: "/img/groomsmen/placeholder.jpg"
    };

    const maidOfHonor = {
        name: "Madeline Tosanovic",
        role: "Maid of Honor",
        relation: "Friend of the Bride",
        photo: "/img/bridesmaids/Madeleine.jpg"
    }
</script>

<div class="contentContainer">
    <!-- Bridesmaids Section -->
    <section style="max-width:80vw; margin: auto">
        <h2>Bridesmaids</h2>
        <div class="honor">
            <PersonCard
                    photo={maidOfHonor.photo}
                    name={maidOfHonor.name}
                    role={maidOfHonor.role}
                    relation={maidOfHonor.relation}
            />
        </div>
        <div class="partyGrid">
            {#each bridesmaid as person}
                <PersonCard photo={person.photo} name={person.name} role={person.role} relation={person.relation}/>
            {/each}
        </div>
    </section>
    <!-- Groomsmen Section -->
    <section style="max-width:80vw; margin: auto">
        <h2 style="padding-top: 2vh">Groomsmen</h2>
        <div class="honor">
            <PersonCard
                    photo={bestMan.photo}
                    name={bestMan.name}
                    role={bestMan.role}
                    relation={bestMan.relation}
            />
        </div>
        <div class="partyGrid">
            {#each groomsmen as person}
                <PersonCard photo={person.photo} name={person.name} role={person.role} relation={person.relation}/>
            {/each}
        </div>
    </section>

    <Footer/>

</div>

<style>

    h2 {
        padding-bottom: 1vh;
        color: rgba(0, 0, 0, 0.4);
        font-size: clamp(20px, calc(1.25rem + ((1vw - 4px) * 0.9036)), 32px);
        font-weight: 400;
        text-transform: uppercase;
        letter-spacing: .42vw;
    }

    .partyGrid {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 1rem;
        justify-items: center;
    }

    .honor {
        margin: auto;
        display: flex;
        justify-content: center;
        max-width: 60%;
    }


</style>