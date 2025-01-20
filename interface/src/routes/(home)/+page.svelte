<div class="contentContainer">
    <img src="/img/main.jpg" alt="" style="justify-content: center"/>

    <div class="eventBlock">
        <h2 class="eventHeader">Wedding Mass</h2>
        <div class="eventInfoLarge">
            <h2>May 31st, 2025</h2>
            <h2>1:00 PM–2:00 PM</h2>
        </div>
        <div class="eventInfoSmall">
            <h3>St. Stephen's Church</h3>
            <h3>1930 W 54th St, Cleveland, OH 44102</h3>
        </div>
    </div>
    <div class="eventBlock">
        <h2 class="eventHeader">Wedding Reception</h2>
        <div class="eventInfoLarge">
            <h2>May 31st, 2025</h2>
            <h2>5:00 PM–10:00 PM</h2>
        </div>
        <div class="eventInfoSmall">
            <h3 class="eventInfo">Windows on the River</h3>
            <h3 class="eventInfo">2000 Sycamore St, Cleveland, OH 44113</h3>
        </div>
    </div>

    <div class="lower">
        <img class="smallImg" src="/img/walk-away.jpg" alt=""/>
    </div>

    <footer>
        <h1>E&M</h1>
    </footer>


</div>

<style>
    .eventBlock {
        margin-bottom: 5rem;
    }

    .eventHeader {
        margin: 3rem 0rem 0.5rem;
        letter-spacing: 3px;
        color: rgb(64, 64, 65);
        font-size: clamp(20px, calc(1.25rem + ((1vw - 4px) * 0.9036)), 32px);
        text-transform: uppercase;
    }

    .eventInfoLarge {
        margin-top: 1rem;
    }

    .eventInfoLarge h2 {
        margin-top: 0.5rem;
        letter-spacing: 3px;
        color: rgb(64, 64, 65);
        font-size: clamp(14px, calc(0.875rem + ((1vw - 4px) * 0.6024)), 22px);
        font-weight: normal;
        text-transform: uppercase;
    }

    .eventInfoSmall {
        margin-top: 0.5rem;
    }

    .eventInfoSmall h3 {
        margin: 0.5rem;
        letter-spacing: 0px;
        color: rgb(64, 64, 65);
        font-size: clamp(12px, calc(0.75rem + ((1vw - 4px) * 0.4518)), 18px);
        font-weight: 400;
    }

    .lower {
        display: grid;
        justify-items: center;
    }


    .lower img {
        width: 35vw;
        /*max-width: 50%;*/
        /*max-height: 100vh;*/
        height: auto;
    }

    footer h1 {
        color: rgba(0, 0, 0, 0.4);
        font-size: 56px;
        font-weight: 400;
        text-transform: uppercase;
        letter-spacing: 6px;
        padding-top: 220px;
        padding-bottom: 15px;
    }

    /* Todo: get the contentContainer class to be only like 80% of the width */
    img {
        /*position: absolute;*/
        /*pointer-events: none;*/
        width: 100%;
        max-width: 70%;
        max-height: 100vh;
        height: 100%;
    }
</style>