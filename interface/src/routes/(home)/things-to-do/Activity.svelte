<script>
    export let name = "";
    export let address = "";
    export let desc = "";
    export let link = ""
</script>

<div style="padding-top: 1.5vh; padding-bottom: 1.5vh">
    <h2>{name}</h2>
    <h3>{address}</h3>
    <p>{desc}</p>
    <a target="_blank" rel="noopener noreferrer" class="siteButton" href={link}>Website</a>
</div>

<style>

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
    p {
        font-size: clamp(10px, calc(0.625rem + ((1vw - 4px) * 0.6024)), 18px);
    }
    .siteButton {
        width: 115px;
        height: 25px;
        background: #80ab7f;
        font-size: clamp(10px, calc(0.625rem + ((1vw - 4px) * 0.6024)), 18px);
        padding: clamp(6px, calc(0.375rem + ((1vw - 4px) * 0.3012)), 10px);
        text-align: center;
        border-radius: 5px;
        color: white;
        font-weight: bold;
        line-height: 25px;
        margin-bottom: 10vh;
    }
</style>