<script>

    export let guest;

    $: dinnerDisabled = guest.attending === "false" && guest.dinner_choice === "";

    const updateGuest = (field, value) => {
        guest[field] = value;
    }


</script>


<div class="rsvpRow">
    <div class="rsvpRow nameLabel">{guest.name}</div>
    <div class="rsvpRow inputFlex">
        <div class="rsvpRow inputFlex buttonWrapper">
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

    .buttonWrapper {
        height: 40px;
        font-size: clamp(12px, calc(0.75rem + ((1vw - 4px) * 0.3012)), 16px);
        border-radius: 5px;
        align-items: center;
        justify-content: right;
        display: flex;
        opacity: 1;
    }

    @media (max-width: 400px) {
        .buttonWrapper {
            width: 50vw;
        }
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