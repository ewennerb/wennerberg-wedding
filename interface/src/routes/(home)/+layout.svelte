<script>

    let { children } = $props();

    const today = new Date();
    const weddingDate = new Date("05-31-2025");
    const diffTime = Math.abs(weddingDate - today);
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));


</script>


<!--<div class="testBG">-->
<!--    <img class="aggaR" src="/img/site-border-2.png" alt=""/>-->
<!--</div>-->

<!--<div style="justify-content: center; text-align: center">-->

<!--    <h1>Edward & Marina</h1>-->
<!--    <h2>May 31st • Cleveland, OH</h2>-->
<!--    <h4>{diffDays} Days Until Wedding!</h4>-->
<!--    <ul class="weddingNav">-->
<!--        <li>-->
<!--            <a href='/' data-item='Home'>Home</a>-->
<!--        </li>-->
<!--        <li>-->
<!--            <a href='travel' data-item='Travel'>Travel</a>-->
<!--        </li>-->
<!--        <li>-->
<!--            <a href='rsvp' data-item='RSVP'>RSVP</a>-->
<!--        </li>-->
<!--        <li>-->
<!--            <a href='faq' data-item='FAQ'>FAQ</a>-->
<!--        </li>-->
<!--        <li>-->
<!--            <a target="_blank" rel="noopener noreferrer" href='https://www.youtube.com/watch?v=dQw4w9WgXcQ' data-item='Registry'>Registry</a>-->
<!--        </li>-->
<!--    </ul>-->

<!--</div>-->


<div>
    <div class="headerImage">
        <picture >
            <img src="/img/site-border-2.png" alt=""/>
        </picture>
    </div>
    <div class="check">
        <div class="homeHeader">
            <h1>Edward & Marina</h1>
            <h2 class="">May 31st • Cleveland, OH</h2>
            <h2>{diffDays} Days Until Wedding!</h2>
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
                <a target="_blank" rel="noopener noreferrer" href='https://www.youtube.com/watch?v=dQw4w9WgXcQ' data-item='Registry'>Registry</a>
            </li>
        </ul>
    </div>
</div>

<!--    <img class="aggaR" src="/img/site-border-2.png" alt=""/>-->



<style>
    /*.headerImage {*/
    /*    width: 100%;*/
    /*    display: -webkit-box;*/
    /*    display: -webkit-flex;*/
    /*    display: -ms-flexbox;*/
    /*    display: flex;*/
    /*    -webkit-flex-direction: column;*/
    /*    -ms-flex-direction: column;*/
    /*    flex-direction: column;*/
    /*    -webkit-align-items: center;*/
    /*    -webkit-box-align: center;*/
    /*    -ms-flex-align: center;*/
    /*    align-items: center;*/
    /*}*/

    /*.headerImage picture {*/
    /*    display: block;*/
    /*    max-width: 100%;*/
    /*    max-height: 120px;*/
    /*    margin-bottom: 20px;*/
    /*    width: 100%;*/
    /*}*/

    /*.headerImage picture img {*/
    /*    position: absolute;*/
    /*    pointer-events: none;*/
    /*    width: 100%;*/
    /*    max-width: 100%;*/
    /*    max-height: 100vh;*/
    /*    height: 100%;*/
    /*    z-index: -1;*/
    /*}*/


    .homeHeader {
        text-align: center;
        width: 100%;
        z-index: 1;
    }

    .homeHeader h1 {
        /*margin: 16px auto;*/
        color: rgba(0, 0, 0, 0.4);
        font-size: 56px;
        font-weight: 400;
        text-transform: uppercase;
        letter-spacing: 6px;
        padding-top: 220px;
        padding-bottom: 15px;
    }

    .homeHeader h2 {
        line-height: 1;
        margin: 8px 0;
        letter-spacing: 2px;
        color: rgba(0, 0, 0, 0.4);
        font-size: 19px;
        font-weight: 400;
        text-transform: uppercase;
    }




    /*.testContainer {*/
    /*    position: relative;*/
    /*}*/

    .check {
        width: 100%;
        /*display: grid;*/
        /*float: left;*/
        /*position: absolute;*/

        justify-content: center;
        text-align: center;
    }


    .testBG {
        display: inline-block;
        overflow: hidden;
        position: relative;
        width: 100%;
    }

    .aggaR {
        pointer-events: none;
        width: 100%;
        max-width: 100%;
        max-height: 100vh;
        height: 100%;
        z-index: -1;
    }

    .imgbox {
        display: grid;
        height: 100%;
    }
    .center-fit {
        background-image: image("/img/site-border-2.png");
        width: 100%;
        max-width: 100%;
        max-height: 100vh;
    }


</style>

{@render children()}