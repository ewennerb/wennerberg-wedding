<script>
    export let src;
    export let name = "John Doe";
    export let bio = "Lorem ipsum fucker dolot";
</script>

<div style="margin-bottom:20px;">
    <div class="profileRow">
        <div class="circular-image">
            <img src={src} alt=""/>
        </div>
        <div class="profileInfo">
            <h4>{name}</h4>
            <p>{bio}</p>
        </div>
    </div>
</div>

<style>
    /* Todo: Alternate row-reverse */
    .profileRow {
        width: 90vw;
        /*max-width: 605px;*/
        /*max-width: 50vw;*/
        text-align: left;
        z-index: 2;
        display: flex;
        -webkit-box-pack: justify;
        justify-content: center;
        /*justify-content: space-between;*/
        -webkit-box-align: center;
        /*align-items: flex-start;*/
        align-items: center;
        /*todo: make conditional for flex-direction: ;*/
    }


    .profileInfo {
        padding-left: 5vw
    }

    .circular-image {
        width: 13vw;
        height: 13vw;
        border-radius: 50%;
        overflow: hidden;
    }

    /*Todo: fix padding and margins on h4 and p */
</style>