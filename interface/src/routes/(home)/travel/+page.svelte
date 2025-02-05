<script>
    import Footer from "../Footer.svelte";
</script>

<div class="contentContainer">

    <h1>Where To Stay</h1>
    <h2>Kimpton Schofield Hotel</h2>
    <h3>2000 East Ninth Street, Cleveland, OH 44115</h3>
    <h3>
        (216) 357-3250
    </h3>

    <div style="display:grid; justify-content: center; max-width: 70vw; margin: auto; padding-bottom: 5vh">
        <img src="/img/hotel.png" alt="https://media-api.xogrp.com/images/ac80ba43-0a91-406c-b3b5-2b866ff529c8"/>
        <p style="max-width:500px; text-align:left">
            We've reserved a block of rooms at the Kimpton Schofield Hotel at the heart of downtown Cleveland. The hotel is very close to many popular bars and restaurants. We have options for both one night and two nights available in our block. Use the button below to book using our block discount!
        </p>
    </div>

    <div style="margin-bottom: 5vh">
        <a target="_blank" rel="noopener noreferrer" class="bookButton" href="https://www.ihg.com/kimptonhotels/hotels/us/en/find-hotels/select-roomrate?fromRedirect=true&qSrt=sBR&qIta=99801505&icdv=99801505&qSlH=CLESH&qCiD=30&qCiMy=042025&qCoD=01&qCoMy=052025&qGrpCd=M44&qAAR=6CBARC&qRtP=6CBARC&setPMCookies=true&qSHBrC=KI&qDest=2000%20East%20Ninth%20Street,%20Cleveland,%20OH,%20US&srb_u=1&qRmFltr=">Book Now!</a>
    </div>

    <Footer/>

</div>

<style>
    h1 {
        color: rgba(0, 0, 0, 0.4);
        font-size: clamp(20px, calc(1.25rem + ((1vw - 4px) * 0.9036)), 32px);
        font-weight: 400;
        text-transform: uppercase;
        letter-spacing: .42vw;
    }

    h2 {
        color: rgba(0, 0, 0, 0.4);
        font-size: clamp(16px, calc(1rem + ((1vw - 4px) * 0.9036)), 28px);
        font-weight: 400;
        text-transform: uppercase;
        letter-spacing: .42vw;
    }

    h3 {
        font-size: clamp(10px, calc(0.625rem + ((1vw - 4px) * 0.6024)), 18px);
        letter-spacing: .14vw;
    }

    img {
        width: 25vw;
        height: auto;
        margin: auto;
    }
    p {
        font-size: clamp(10px, calc(0.625rem + ((1vw - 4px) * 0.6024)), 18px);
    }

    .bookButton {
        width: 115px;
        height: 25px;
        background: #80ab7f;
        padding: 10px;
        text-align: center;
        border-radius: 5px;
        color: white;
        font-weight: bold;
        line-height: 25px;
        margin-bottom: 10vh;
    }
</style>