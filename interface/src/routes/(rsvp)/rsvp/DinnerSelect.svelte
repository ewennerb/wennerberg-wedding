<script>

    export let guest;

    $: dinnerDisabled = guest.attending === "false" && guest.dinner_choice === "";

    const updateGuest = (field, value) => {
        guest[field] = value;
    }


</script>


<div class="knotRow">
    <div class="knotRow nameLabel">{guest.name}</div>
    <div class="knotRow inputStuff">
        <div class="buttonWrapper">
            <select class="guestDropdown" disabled={dinnerDisabled} bind:value={guest.dinner_choice} on:change={(e) => updateGuest('dinner_choice', e.target.value)}>
                <option value="" disabled selected>-- None --</option>
                <option value="steak">Filet Mignon</option>
                <option value="fish">Blackened Grouper</option>
                <option value="lasagna">Spinach Lasagna Roll</option>
            </select>
        </div>
    </div>
</div>


<style>
    /* Todo: Rename CSS Classes and update font info! */
    .knotRow {
        height: 80%;
        width: 480px;
        padding: 5px 20px 0px;
        text-align: left;
        z-index: 2;
        display: flex;
        -webkit-box-pack: justify;
        justify-content: space-between;
        -webkit-box-align: center;
        align-items: center;
        margin-bottom: 24px;
    }

    .knotRow .nameLabel {
        color: rgba(0, 0, 0, 0.4);
        font-weight: 400;
        display: flex;
        font-size: 20px;
        margin-bottom: 0px;
    }

    .knotRow .inputStuff {
        display: flex;
        width: 20rem;
        margin-bottom: 0px
    }

    .knotRow .inputStuff .buttonWrapper {
        height: 40px;
        font-size: 16px;
        border-radius: 5px;
        align-items: center;
        justify-content: center;
        display: flex;
        opacity: 1;
    }

    .acceptButton {
        height: 40px;
        width: 124px;
        font-size: 16px;
        border-radius: 2px;
        background-color: rgb(246, 233, 213);
        border: 2px solid rgba(0, 0, 0, 0.4);
        color: rgba(0, 0, 0, 0.4);
        opacity: 0.5;
        margin-left: 8px;
        margin-right: 12px;
        touch-action: manipulation;
    }

    .rsvpField {
        margin-right: 1rem;
        background-color: rgb(255, 255, 255);
        border-color: rgba(0, 0, 0, 0.4);
        border-radius: 3px;
        border-style: solid;
        border-width: 1px;
        color: rgba(0, 0, 0, 0.4);
        cursor: text;
        display: block;
        font-weight: 400;
        height: 90%;
        letter-spacing: 1px;
        line-height: 24px;
        transition-delay: 0s;
        transition-duration: 0.2s;
        transition-property: border-color;
        transition-timing-function: ease-out;
        width: 5rem;
        word-spacing: 0px;
        padding-block-start: 15px;
        padding-inline-end: 10px;
        padding-inline-start: 10px;
        padding-left: 10px;
        padding-right: 10px;
        padding-top: 0px;
        position: relative;
        text-align: start;
    }

    .guestDropdown {
        appearance: none;
        background-color: rgb(255, 255, 255);
        border-color: black;
        border-radius: 2px;
        border-style: solid;
        border-width: 1px;
        border-image-repeat: stretch;
        box-shadow: none;
        box-sizing: border-box;
        color: rgba(0, 0, 0, 0.4);
        cursor: text;
        display: block;
        font-weight: 400;
        height: 100%;
        letter-spacing: 1px;
        line-height: 24px;
        padding-block-start: 15px;
        padding-inline-end: 10px;
        padding-inline-start: 10px;
        padding-left: 10px;
        padding-right: 10px;
        padding-top: 0px;
        position: relative;
        text-align: start;
        text-indent: 0px;
        text-rendering: auto;
        text-size-adjust: 100%;
        touch-action: manipulation;
        transition-delay: 0s;
        transition-duration: 0.2s;
        transition-property: border-color;
        transition-timing-function: ease-out;
        width: 15rem;
        word-spacing: 0px;
    }

</style>