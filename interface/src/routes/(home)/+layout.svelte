<script>

    let { children } = $props();

    const today = new Date();
    const weddingDate = new Date("2025-05-31");
    const diffTime = Math.abs(weddingDate - today);
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));


</script>


<div>
    <div class="headerImage">
        <picture >
            <img src="/img/site-border-2.png" alt=""/>
        </picture>
    </div>
    <div class="mainContainer">
        <div class="homeHeader">
            <h1>Marina & Edward</h1>
            <h2 class="">May 31st • Cleveland, OH</h2>
            {#if diffDays > 0}
                <h2>{diffDays} {diffDays !== 1 ? "Days" : "Day"} Until Wedding!</h2>
            {/if}

            <h4>Hotel room block rates available until April 18th</h4>
            <h4>Click <a target="_blank" rel="noopener noreferrer" href="https://www.ihg.com/kimptonhotels/hotels/us/en/find-hotels/select-roomrate?fromRedirect=true&qSrt=sBR&qIta=99801505&icdv=99801505&qSlH=CLESH&qCiD=30&qCiMy=042025&qCoD=01&qCoMy=052025&qGrpCd=M44&qAAR=6CBARC&qRtP=6CBARC&setPMCookies=true&qSHBrC=KI&qDest=2000%20East%20Ninth%20Street,%20Cleveland,%20OH,%20US&srb_u=1&qRmFltr=">here</a> to book now!</h4>
        </div>
        <ul class="weddingNav">
            <li>
                <a href='/' data-item='Home'>Home</a>
            </li>
            <li>
                <a href='wedding-party' data-item="Wedding Party">Wedding Party</a>
            </li>
            <li>
                <a href='photos' data-item="Photos">Photos</a>
            </li>
            <li>
                <a href='rsvp' data-item='RSVP'>RSVP</a>
            </li>
            <li>
                <a href='faq' data-item='FAQ'>FAQ</a>
            </li>
            <li>
                <a href='travel' data-item='Travel'>Travel</a>
            </li>
            <li>
                <a href='things-to-do' data-item='Things To Do'>Things To Do</a>
            </li>
            <li>
                <a target="_blank" rel="noopener noreferrer" href='https://www.myregistry.com/wedding-registry/marina-bostelman-and-edward-wennerberg-cleveland-oh/4651928' data-item='Registry'>Registry</a>
            </li>
        </ul>
    </div>
</div>

<style>

    .homeHeader {
        text-align: center;
        width: 100%;
        z-index: 1;
    }

    .homeHeader h1 {
        color: rgba(0, 0, 0, 0.4);
        font-size: clamp(22px, calc(1.375rem + ((1vw - 4px) * 2.5602)), 56px);
        /*font-size: 3.9vw;*/
        font-weight: 400;
        text-transform: uppercase;
        letter-spacing: .42vw;
        padding-top: 25%;
        padding-bottom: 1.7vh;
    }

    .homeHeader h2 {
        line-height: 1;
        margin: 8px 0;
        letter-spacing: .13vw;
        color: rgba(0, 0, 0, 0.4);
        font-size: clamp(10px, calc(0.625rem + ((1vw - 4px) * 1.1295)), 25px);
        /*font-size: 1.3vw;*/
        font-weight: 400;
        text-transform: uppercase;
    }

    .homeHeader h4 {
        line-height: 1;
        letter-spacing: .13vw;
        color: rgb(168, 7, 7);
        font-size: clamp(6px, calc(0.625rem + ((1vw - 4px) * 1.1295)), 14px);
        font-weight: 400;
    }

    .mainContainer {
        align-items: center;
        justify-content: center;
        width: 80vw;
        text-align: center;
        margin-inline: auto;
    }


</style>

{@render children()}