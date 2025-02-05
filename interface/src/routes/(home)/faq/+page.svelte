<script>
    import Footer from "../Footer.svelte";
</script>

<div class="contentContainer">
    <div style="display:inline-block; max-width:500px; text-align:left; width:75vw">

        <h3>Q: What do I need to know about parking for ceremony and reception?</h3>
        <p>
            <b>Ceremony:</b> There is a small parking lot behind the church that is accessible from W 57th St.
            This parking lot will not be able to accommodate all wedding guests, but there is ample street parking
            available in the area surrounding the church. We strongly recommend allowing for a few extra minutes to find parking and walk
            to the church. All parking around te church is free.
        </p>
        <p>
            <b>Reception:</b> Windows on the River has a massive gated parking lot directly next to the venue.
            There is a $5 fee for this parking lot.
        </p>
        <p>
            <b>Hotel:</b> If you’re staying at the Kimpton Schofield, please contact the hotel for details on parking.
        </p>

        <h3>Q: Why is your wedding Mass going to be in Latin, and what can I expect?</h3>
        <p>A: We chose to have our wedding Mass said in Latin to honor the history of the Catholic Church and to
            emphasize the sacred nature of the Sacrament of Marriage and the vows we are making to each other. The
            structure and order of our wedding Mass will be the same as a normal wedding Mass, but most of the prayers
            of the Mass will be said in Latin rather than English. Our program will include a worship aid with English
            translations so you can follow along if you wish to. The Scripture readings, the homily (sermon), and our
            marriage vows will all be in English.
        </p>

        <h3>Q: What should we do between the ceremony and reception?</h3>
        <p>A: Whatever you want! We recommend you check out the "Things To Do" tab for some of our favorite spots around the city!</p>

        <h3>Q: What is the dress code?</h3>
        <p>A: Formal attire is required for the ceremony and reception.</p>

        <h3>Q: Can I bring a plus one?</h3>
        <p>A: Please refer to the RSVP tab to determine if you've been allowed a plus one. If you feel like we missed someone, please reach out.</p>


<!--        <h3>Q: Are children invited?</h3>-->
<!--        <p>A: Unless otherwise specified, we ask that you keep any children that aren't babies (2 yrs) home.</p>-->

<!--        <h3>Q: Can I bring a plus one?</h3>-->
<!--        <p>A: Please refer to the RSVP tab to determine if you've been allowed a plus one. If you feel like we missed someone, please reach out!</p>-->

<!--        <h3>Q: What's the parking situation like?</h3>-->
<!--        <p>A: There is plentiful street parking and a lot at the church. The reception venue has a massive parking lot that will be free for guests.</p>-->

<!--        <h3>Q: What about dietary restrictions?</h3>-->
<!--        <p>A: Please refer to the RSVP tab to indicate any allergies or dietary restrictions we should be aware of.</p>-->

<!--        <p style="color:red">Might roll these next 2 into a separate nav tab. </p>-->
<!--        <h3>What do we do between the mass and the reception?</h3>-->
<!--        <p>Check out the "Things to Do" tab for some of our favorite spots if you'd like to grab a snack or start pregaming!</p>-->

<!--        <h3>What is there to do in Cleveland?</h3>-->
<!--        <ul>-->
<!--            <li>Rock & Roll Hall of Fame</li>-->
<!--            <li>Cleveland Museum of Art</li>-->
<!--            <li>Cleveland Guardians Games on May 30th and June 1st against the Angels</li>-->
<!--        </ul>-->

    </div>

    <Footer/>

</div>

<style>
    h3 {
        font-family: "Times New Roman", serif;
        font-size: clamp(8px, calc(0.998rem + ((1vw - 4px) * 0.3012)), 20px);
        color: rgba(0, 0, 0, 0.4);
    }

    p {
        font-family: "Times New Roman", serif;
        font-size: clamp(8px, calc(0.875rem + ((1vw - 4px) * 0.3012)), 16px);
    }
</style>