<footer class="testFooter">
    <div class="footerBox">
        <div class="bottomText">M&E</div>
        <div class="footerDate">5.31.2025</div>
    </div>
    <div class="contactInfo">Problem with the site? Contact edwardwennerberg@gmail.com detailing your issue.</div>
</footer>


<style>
    .testFooter {
        align-items: center;
        display: flex;
        flex-direction: column;
        justify-content: flex-end;
        margin-top: auto;
        position: relative;
        width: 100%;
    }

    .footerBox {
        letter-spacing: 3.4px;
        line-height: 1;
        margin: 3vw auto;
        text-align: center;
        width: fit-content;
    }

    .bottomText {
        padding: 0 1vw 1vh;
        color: #404041;
        font-size: clamp(22px, calc(1.375rem + ((1vw - 4px) * 2.5602)), 56px);
        font-weight: normal;
        text-transform: uppercase;
        letter-spacing: 4px;
    }

    .footerDate {
        padding: 1vw 1vh 0;
        letter-spacing: 3px;
        color: #404041;
        font-size: clamp(14px, calc(0.875rem + ((1vw - 4px) * 0.6024)), 22px);;
        font-weight: 400;
        text-transform: uppercase;
        border-top: 1px solid #404041;
    }

    .contactInfo {
        max-width: 30vw;
        line-height: 1.75;
        letter-spacing: 0px;
        color: #404041;
        font-size: clamp(6px, calc(0.45rem + ((1vw - 4px)* 0.45)), 11px);
        font-weight: 400;
        margin: 12px 0;
        text-align: center;
    }
</style>